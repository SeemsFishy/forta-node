package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"

	"forta-network/forta-node/clients"
	"forta-network/forta-node/clients/messaging"
	"forta-network/forta-node/config"
	"forta-network/forta-node/contracts"
	"forta-network/forta-node/ethereum"
	"forta-network/forta-node/services/registry/regtypes"
	"forta-network/forta-node/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// RegistryService listens to the agent scanner list changes so the node can stay in sync.
type RegistryService struct {
	cfg            config.Config
	scannerAddress common.Address
	msgClient      clients.MessageClient

	rpcClient  *rpc.Client
	agentReg   AgentRegistryCaller
	scannerReg ScannerRegistryCaller
	ipfsClient IPFSClient
	ethClient  EthClient

	agentsConfigs []*config.AgentConfig
	done          chan struct{}
	version       string
	sem           *semaphore.Weighted
}

// AgentRegistryCaller calls the agent registry contract.
type AgentRegistryCaller interface {
	AgentReference(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (string, error)
	AgentLatestVersion(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error)
}

// ScannerRegistryCaller calls the scanner registry contract.
type ScannerRegistryCaller interface {
	AgentLength(opts *bind.CallOpts, scanner common.Address) (*big.Int, error)
	AgentAt(opts *bind.CallOpts, scanner common.Address, index *big.Int) ([32]byte, *big.Int, bool, string, bool, error)
	GetAgentListHash(opts *bind.CallOpts, scanner common.Address) ([32]byte, error)
}

// IPFSClient interacts with an IPFS Gateway.
type IPFSClient interface {
	GetAgentFile(cid string) (*regtypes.AgentFile, error)
}

// EthClient interacts with the Ethereum API.
type EthClient interface {
	ethereum.Client
}

// New creates a new service.
func New(cfg config.Config, scannerAddress common.Address, msgClient clients.MessageClient) *RegistryService {
	var ipfsURL string
	if cfg.Registry.IPFSGateway != nil {
		ipfsURL = *cfg.Registry.IPFSGateway
	} else {
		ipfsURL = config.DefaultIPFSGateway
	}

	return &RegistryService{
		cfg:            cfg,
		scannerAddress: scannerAddress,
		msgClient:      msgClient,
		ipfsClient:     &ipfsClient{ipfsURL},
		done:           make(chan struct{}),
	}
}

// Init only initializes the service.
func (rs *RegistryService) Init() (err error) {
	rs.rpcClient, err = ethereum.NewRpcClient(rs.cfg.Registry.Ethereum.JsonRpcUrl)
	if err != nil {
		return
	}
	agentRegAddress := config.GetEnvDefaults(rs.cfg.Production).DefaultAgentRegistryContractAddress
	client := ethclient.NewClient(rs.rpcClient)
	rs.agentReg, err = contracts.NewAgentRegistryCaller(common.HexToAddress(agentRegAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create the agent registry caller: %v", err)
	}
	rs.scannerReg, err = contracts.NewScannerRegistryCaller(common.HexToAddress(rs.cfg.Registry.ContractAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create the scanner registry caller: %v", err)
	}
	// used for getting the latest block number so that we can query consistent state
	log.Infof("Creating Caller: %s", rs.Name())
	ethClient, err := ethereum.NewStreamEthClient(context.Background(), rs.cfg.Registry.Ethereum.JsonRpcUrl)
	if err != nil {
		return err
	}
	rs.ethClient = ethClient
	return nil
}

// Start initializes and starts the registry service.
func (rs *RegistryService) Start() error {
	log.Infof("Starting %s", rs.Name())
	if err := rs.Init(); err != nil {
		return err
	}
	rs.sem = semaphore.NewWeighted(1)
	return rs.start()
}

// FindAgentGlobally prepares the config for an agent, optionally by using a specific version.
// It uses the agent registry directly and disregards the scanner registry.
func (rs *RegistryService) FindAgentGlobally(agentID string, version uint64) (config.AgentConfig, error) {
	opts, err := rs.optsWithLatestBlock()
	if err != nil {
		return config.AgentConfig{}, err
	}
	agentIDBytes := ([32]byte)(common.HexToHash(agentID))

	if version == 0 {
		latestVersion, err := rs.agentReg.AgentLatestVersion(opts, agentIDBytes)
		if err != nil {
			return config.AgentConfig{}, fmt.Errorf("failed to get the latest version of the agent: %v", err)
		}
		version = latestVersion.Uint64()
	}

	agentRef, err := rs.agentReg.AgentReference(opts, agentIDBytes, big.NewInt(0).SetUint64(version))
	if err != nil {
		return config.AgentConfig{}, fmt.Errorf("failed to get the latest ref: %v", err)
	}

	return rs.makeAgentConfig(agentIDBytes, agentRef)
}

// ReadLocalAgents tries to read the local agents and silently returns an
// empty array if the file is not readable or not found.
func (rs *RegistryService) ReadLocalAgents() ([]*config.AgentConfig, error) {
	var agents []*config.AgentConfig
	b, err := ioutil.ReadFile(rs.cfg.LocalAgentsPath)
	if err == nil {
		if err := json.Unmarshal(b, &agents); err != nil {
			return nil, fmt.Errorf("failed to unmarshal the local agents file: %v", err)
		}
	}
	return agents, nil
}

// WriteLocalAgents writes the agents to the local list.
func (rs *RegistryService) WriteLocalAgents(agents []*config.AgentConfig) error {
	if len(agents) == 0 {
		return nil
	}
	b, _ := json.MarshalIndent(agents, "", "  ")
	return ioutil.WriteFile(rs.cfg.LocalAgentsPath, b, 0644)
}

func (rs *RegistryService) start() error {
	go func() {
		//TODO: possibly make this configurable, but 15s per block is normal
		ticker := time.NewTicker(15 * time.Second)
		for {
			if err := rs.publishLatestAgents(); err != nil {
				log.Errorf("failed to publish the latest agents: %v", err)
			}
			<-ticker.C
		}
	}()

	return nil
}

func (rs *RegistryService) publishLatestAgents() error {
	// only allow one executor at a time, even if slow
	if rs.sem.TryAcquire(1) {
		defer rs.sem.Release(1)
		// opts is nil so we get the latest scanner list version
		version, err := rs.scannerReg.GetAgentListHash(nil, rs.scannerAddress)
		if err != nil {
			return fmt.Errorf("failed to get the scanner list agents version: %v", err)
		}
		versionStr := string(version[:])
		// if versions change, then get the full list of agents
		if rs.version == "" || rs.version != versionStr {
			log.Infof("registry: agent version changed %s->%s", rs.version, versionStr)
			rs.version = versionStr
			rs.agentsConfigs, err = rs.getLatestAgents()
			if err != nil {
				return fmt.Errorf("failed to get latest agents: %v", err)
			}
			rs.msgClient.Publish(messaging.SubjectAgentsVersionsLatest, rs.agentsConfigs)
		} else {
			log.Info("registry: no agent changes detected")
		}
	}
	return nil
}

func (rs *RegistryService) optsWithLatestBlock() (*bind.CallOpts, error) {
	blk, err := rs.ethClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get the block for agents: %v", err)
	}
	num, err := utils.HexToBigInt(blk.Number)
	if err != nil {
		return nil, err
	}
	return &bind.CallOpts{
		BlockNumber: num,
	}, nil
}

func (rs *RegistryService) getLatestAgents() ([]*config.AgentConfig, error) {
	var agentConfigs []*config.AgentConfig

	opts, err := rs.optsWithLatestBlock()
	if err != nil {
		return nil, err
	}

	lengthBig, err := rs.scannerReg.AgentLength(opts, rs.scannerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get the scanner list agents length: %v", err)
	}
	log.Infof("registry: getLatestAgents(%s) = %s", rs.scannerAddress.Hex(), lengthBig.Text(10))
	// TODO: If we are going to get 100s of agents, we probably need to batch the calls here.
	length := int(lengthBig.Int64())
	for i := 0; i < length; i++ {
		agentID, _, _, agentRef, disabled, err := rs.scannerReg.AgentAt(opts, rs.scannerAddress, big.NewInt(int64(i)))
		if err != nil {
			return nil, fmt.Errorf("failed to get agent at index '%d' in scanner list '%s': %v", i, rs.scannerAddress.String(), err)
		}
		// if agent dev disables agent, this will prevent it from running
		if !disabled {
			agentCfg, err := rs.makeAgentConfig(agentID, agentRef)
			if err != nil {
				return nil, fmt.Errorf("failed to make agent config: %v", err)
			}
			agentConfigs = append(agentConfigs, &agentCfg)
		}
	}

	// Also include local agents if any.
	localAgents, _ := rs.ReadLocalAgents()

	return append(agentConfigs, localAgents...), nil
}

func (rs *RegistryService) makeAgentConfig(agentID [32]byte, ref string) (agentCfg config.AgentConfig, err error) {
	agentCfg.ID = (common.Hash)(agentID).String()
	if len(ref) == 0 {
		return
	}

	var (
		agentData *regtypes.AgentFile
	)
	for i := 0; i < 10; i++ {
		agentData, err = rs.ipfsClient.GetAgentFile(ref)
		if err == nil {
			break
		}
	}
	if err != nil {
		err = fmt.Errorf("failed to load the agent file using ipfs ref: %v", err)
		return
	}

	var ok bool
	agentCfg.Image, ok = utils.ValidateImageRef(rs.cfg.Registry.ContainerRegistry, agentData.Manifest.ImageReference)
	if !ok {
		log.Warnf("invalid agent reference - skipping: %s", agentCfg.Image)
	}

	return
}

// Stop stops the registry service.
func (rs *RegistryService) Stop() error {
	return nil
}

// Name returns the name of the service.
func (rs *RegistryService) Name() string {
	return "RegistryService"
}
