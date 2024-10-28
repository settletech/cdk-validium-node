// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevm

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PolygonRollupBaseEtrogBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonValidiumEtrogValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonValidiumEtrogValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonzkevmMetaData contains all meta data concerning the Polygonzkevm contract.
var PolygonzkevmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecuteForcedBatchesNotPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitModeIsActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertModeIsActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertModeIsNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertNotExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TxWindowNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerifyOnRevertNotPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"noForcedBatchPending\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_lastForcedBatch\",\"type\":\"uint64\"}],\"name\":\"ActivateRevertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_prentBlockHash\",\"type\":\"bytes32\"}],\"name\":\"LogForcedBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEST_VAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData\",\"name\":\"forcedBatchData\",\"type\":\"tuple\"}],\"name\":\"enterRevertMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrog.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"executeAllForcedTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRevertModeActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRevertModeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_prevAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onRevertLastVerifiedBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrog.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040526005603e5534801562000016575f80fd5b5060405162004b0838038062004b08833981016040819052620000399162000074565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d9565b6001600160a01b038116811462000071575f80fd5b50565b5f805f806080858703121562000088575f80fd5b845162000095816200005c565b6020860151909450620000a8816200005c565b6040860151909350620000bb816200005c565b6060860151909250620000ce816200005c565b939692955090935050565b60805160a05160c05160e0516148ce6200023a5f395f818161052101528181610aac01528181610b0d01528181610cf101528181610de901528181610f90015281816111c2015281816113a101528181611458015281816118b5015281816119540152818161197601528181611ab801528181611b9801528181611e5a0152818161213f015281816121ac0152818161259501528181612a7c01528181612ae001528181612b0201528181612be401528181612d1501528181612db40152818161373801528181613796015281816137b8015261383901525f81816106860152818161158001528181612313015281816123d2015281816127470152818161310b0152818161319c015261340b01525f818161073a015281816110e1015281816115ef015281816127b601528181612eb0015261347a01525f818161079201528181610837015281816118fa01528181611a080152612e8501526148ce5ff3fe608060405234801561000f575f80fd5b506004361061034c575f3560e01c8063788aaa6a116101c9578063c7fffd4b116100fe578063e46761c41161009e578063eaeb077b11610079578063eaeb077b146107da578063eb5db56a146107ed578063f35dda47146107f6578063f851a440146107fe575f80fd5b8063e46761c41461078d578063e57a0b4c146107b4578063e7a7ed02146107c7575f80fd5b8063d02103ca116100d9578063d02103ca14610735578063d7bc90ff1461075c578063db5b0ed714610767578063def57e541461077a575f80fd5b8063c7fffd4b14610707578063c89e42df1461070f578063cfa8ed4714610722575f80fd5b80639f26f84011610169578063ada8f91911610144578063ada8f919146106bb578063b0afe154146106ce578063b151da2d146106da578063c754c7ed146106ed575f80fd5b80639f26f8401461066e578063a3c573eb14610681578063a652f26c146106a8575f80fd5b80637cd76b8b116101a45780637cd76b8b146106255780638c3d73011461063857806391cafe32146106405780639e00187714610653575f80fd5b8063788aaa6a146105e95780637a5460c5146105f15780637cbe817f14610612575f80fd5b80633cbc795b1161029f57806352bdeb6d1161023f5780636b8616ce1161021a5780636b8616ce1461059b5780636e05d2cd146105ba5780636ff512cc146105c357806371257022146105d6575f80fd5b806352bdeb6d1461056a578063542028d51461058b578063676870d214610593575f80fd5b8063456052671161027a57806345605267146104ea57806349b7b8021461051c5780634c21fef3146105435780634e48770614610557575f80fd5b80633cbc795b1461048e57806340b5de6c146104ba57806342308fab146104e2575f80fd5b8063236ae0c01161030a5780632acdc2b6116102e55780632acdc2b6146104425780632c111c061461045557806332c2d153146104685780633c351e101461047b575f80fd5b8063236ae0c0146103f157806326782247146103fa57806328e18a8914610425575f80fd5b8062d0295d14610350578063035089631461036b57806303744a631461038657806305835f371461039b578063107bf28c146103cf57806311e892d4146103d7575b5f80fd5b610358610816565b6040519081526020015b60405180910390f35b610373602081565b60405161ffff9091168152602001610362565b610399610394366004613c9a565b6108ef565b005b6103c260405180604001604052806008815260200167202021007270e02560c21b81525081565b6040516103629190613d18565b6103c2610bc1565b6103df60f981565b60405160ff9091168152602001610362565b610358603e5481565b60015461040d906001600160a01b031681565b6040516001600160a01b039091168152602001610362565b600b546104329060ff1681565b6040519015158152602001610362565b610399610450366004613d3e565b610c4d565b60085461040d906001600160a01b031681565b610399610476366004613d78565b610cef565b60095461040d906001600160a01b031681565b6009546104a590600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610362565b6104c96001600160f81b031981565b6040516001600160f81b03199091168152602001610362565b610358602481565b60075461050490600160401b90046001600160401b031681565b6040516001600160401b039091168152602001610362565b61040d7f000000000000000000000000000000000000000000000000000000000000000081565b603f5461043290600160a01b900460ff1681565b610399610565366004613db7565b610d8a565b6103c260405180604001604052806002815260200161101760f31b81525081565b6103c2610f03565b610373601f81565b6103586105a9366004613db7565b60066020525f908152604090205481565b61035860055481565b6103996105d1366004613dd2565b610f10565b6103996105e4366004613dfe565b610f8e565b61043261137d565b6103c26040518060400160405280600281526020016180b960f01b81525081565b610399610620366004613ee8565b61139f565b610399610633366004613dd2565b611c63565b610399611ce1565b61039961064e366004613dd2565b611d6a565b61040d73a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b61039961067c366004613ff0565b611e11565b61040d7f000000000000000000000000000000000000000000000000000000000000000081565b6103c26106b636600461402e565b612276565b6103996106c9366004613dd2565b612515565b6103586405ca1ab1e081565b6103996106e836600461409e565b612593565b60075461050490600160801b90046001600160401b031681565b6103df60e481565b61039961071d3660046140b5565b6125e1565b60025461040d906001600160a01b031681565b61040d7f000000000000000000000000000000000000000000000000000000000000000081565b610358635ca1ab1e81565b610399610775366004613ee8565b61264d565b6103996107883660046140e6565b612c93565b61040d7f000000000000000000000000000000000000000000000000000000000000000081565b603f5461040d906001600160a01b031681565b600754610504906001600160401b031681565b6103996107e836600461415d565b612cd1565b610358600a5481565b6103df601b81565b5f5461040d906201000090046001600160a01b031681565b6040516370a0823160e01b81523060048201525f9081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561087c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108a091906141a4565b6007549091505f906108c4906001600160401b03600160401b8204811691166141cf565b6001600160401b03169050805f036108de575f9250505090565b6108e881836141f6565b9250505090565b42610e10600a546109009190614215565b111561091f57604051632cd6ab1760e11b815260040160405180910390fd5b6008546001600160a01b0316801580159061094357506001600160a01b0381163314155b15610961576040516324eff8c360e01b815260040160405180910390fd5b6007546001600160401b03600160401b8204811691168111156109975760405163c630a00d60e01b815260040160405180910390fd5b6007546001600160401b038083169116036109c55760405163115cad3f60e01b815260040160405180910390fd5b8251805160209182012081850151604080870151606088015191515f956109ef959493910161422e565b60405160208183030381529060405280519060200120905060065f83610a1490614256565b9350836001600160401b03166001600160401b031681526020019081526020015f20548114610a565760405163671ebaaf60e11b815260040160405180910390fd5b42620697808560400151610a6a919061427b565b6001600160401b03161115610a9257604051631637066960e11b815260040160405180910390fd5b42600a5560405163750d3eef60e11b8152600160048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ea1a7dde906024015f604051808303815f87803b158015610af5575f80fd5b505af1158015610b07573d5f803e3d5ffd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8d2b0736040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610b63575f80fd5b505af1158015610b75573d5f803e3d5ffd5b5050604080513381526001600160401b03861660208201527fa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9935001905060405180910390a150505050565b60048054610bce9061429b565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfa9061429b565b8015610c455780601f10610c1c57610100808354040283529160200191610c45565b820191905f5260205f20905b815481529060010190602001808311610c2857829003601f168201915b505050505081565b5f546201000090046001600160a01b03163314610c7d57604051634755657960e01b815260040160405180910390fd5b603f54600160a01b900460ff16151581151503610cad57604051632f873d5f60e11b815260040160405180910390fd5b603f805460ff60a01b1916600160a01b831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41905f90a150565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610d3857604051631736745960e31b815260040160405180910390fd5b806001600160a01b0316836001600160401b03167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610d7d91815260200190565b60405180910390a3505050565b5f546201000090046001600160a01b03163314610dba57604051634755657960e01b815260040160405180910390fd5b62093a806001600160401b0382161115610de75760405163f5e37f2f60e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e6791906142d3565b610ea1576007546001600160401b03600160801b909104811690821610610ea15760405163f5e37f2f60e01b815260040160405180910390fd5b6007805467ffffffffffffffff60801b1916600160801b6001600160401b038416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b60038054610bce9061429b565b5f546201000090046001600160a01b03163314610f4057604051634755657960e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610ef8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610fd757604051631736745960e31b815260040160405180910390fd5b5f54610100900460ff1615808015610ff557505f54600160ff909116105b8061100e5750303b15801561100e57505f5460ff166001145b6110765760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015611097575f805461ff0019166101001790555b5f6110a1856130db565b6009549091505f906110cd9088906001600160a01b03811690600160a01b900463ffffffff1685612276565b90505f818051906020012090505f4290505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561113b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061115f91906141a4565b90505f808483858f6111726001436142ee565b4060405160200161118896959493929190614301565b60408051808303601f190181529082905280516020909101206005819055639a908e7360e01b8252600160048301526024820181905291507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015611210573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112349190614345565b508c5f60026101000a8154816001600160a01b0302191690836001600160a01b031602179055508b60025f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550886003908161129091906143ad565b50600461129d89826143ad565b508c60085f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555062069780600760106101000a8154816001600160401b0302191690836001600160401b031602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e60405161132193929190614468565b60405180910390a15050505050508015611374575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f42610e10600a5461138f9190614215565b1161139957505f90565b50600190565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637fea08646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141f91906142d3565b1561143d5760405163bf52841f60e01b815260040160405180910390fd5b42610e10600a5461144e9190614215565b1115806114d857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b40a15c36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114b2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114d691906142d3565b155b80156114e75750600b5460ff16155b1561150557604051638ff06f2b60e01b815260040160405180910390fd5b855f8190036115275760405163cb591a5f60e01b815260040160405180910390fd5b6103e881111561154a57604051635acfba9d60e11b815260040160405180910390fd5b611555602442614215565b866001600160401b0316111561157e57604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156115d6575f80fd5b505af11580156115e8573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015611649573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061166d91906141a4565b600754600554919250600160401b90046001600160401b031690815f805b86811015611844575f8e8e838181106116a6576116a661449a565b9050608002018036038101906116bc91906144ae565b60408101519091506001600160401b0316156117c457856116dc81614256565b9650505f815f0151826020015183604001518460600151604051602001611706949392919061422e565b60408051601f1981840301815291815281516020928301206001600160401b038a165f908152600690935291205490915081146117565760405163671ebaaf60e11b815260040160405180910390fd5b85825f01518360200151846040015133866060015160405160200161178096959493929190614301565b60405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f905550611831565b80516040516117e0918591602001918252602082015260400190565b60405160208183030381529060405280519060200120925084815f0151888f8e5f801b60405160200161181896959493929190614301565b6040516020818303038152906040528051906020012094505b508061183c816144fa565b91505061168b565b506007546001600160401b0390811690851611156118755760405163c630a00d60e01b815260040160405180910390fd5b6005839055856001600160401b0385811690841614611948575f61189984876141cf565b90506118ae6001600160401b038216836142ee565b91506119217f0000000000000000000000000000000000000000000000000000000000000000826001600160401b03166118e6610816565b6118f09190614512565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016919061326b565b506007805467ffffffffffffffff60401b1916600160401b6001600160401b038816021790555b8015611a9157611a30337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119d0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119f491906141a4565b6119fe9190614512565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169291906132d3565b603f54604051633b51be4b60e01b81526001600160a01b0390911690633b51be4b90611a649085908d908d90600401614551565b5f6040518083038186803b158015611a7a575f80fd5b505afa158015611a8c573d5f803e3d5ffd5b505050505b60405163287d2d0b60e11b81526001600160401b0388166004820152602481018590525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906350fa5a16906044016020604051808303815f875af1158015611b06573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b2a9190614345565b9050611b3688826141cf565b6001600160401b03168c6001600160401b031614611b6757604051630d0386cd60e11b815260040160405180910390fd5b6007546001600160401b03808216600160401b9092041603611c0a5760405163750d3eef60e11b81525f60048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ea1a7dde906024015f604051808303815f87803b158015611be1575f80fd5b505af1158015611bf3573d5f803e3d5ffd5b5050600b805460ff1916600117905550611c0f9050565b42600a555b806001600160401b03167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76688604051611c4a91815260200190565b60405180910390a2505050505050505050505050505050565b5f546201000090046001600160a01b03163314611c9357604051634755657960e01b815260040160405180910390fd5b603f80546001600160a01b0319166001600160a01b0383169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a90602001610ef8565b6001546001600160a01b03163314611d0c5760405163d1ec4b2360e01b815260040160405180910390fd5b6001545f805462010000600160b01b0319166001600160a01b039092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f546201000090046001600160a01b03163314611d9a57604051634755657960e01b815260040160405180910390fd5b6008546001600160a01b0316611dc3576040516319126e9b60e31b815260040160405180910390fd5b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610ef8565b6008546001600160a01b03168015801590611e3557506001600160a01b0381163314155b15611e53576040516324eff8c360e01b815260040160405180910390fd5b4262093a807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611eb4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ed89190614345565b611ee2919061427b565b6001600160401b03161115611f0a57604051630f527b5360e21b815260040160405180910390fd5b815f819003611f2c5760405163cb591a5f60e01b815260040160405180910390fd5b6103e8811115611f4f57604051635acfba9d60e11b815260040160405180910390fd5b6007546001600160401b0380821691611f71918491600160401b900416614215565b1115611f905760405163c630a00d60e01b815260040160405180910390fd5b600754600554600160401b9091046001600160401b0316905f5b83811015612139575f878783818110611fc557611fc561449a565b9050602002810190611fd79190614573565b611fe090614591565b905083611fec81614256565b825180516020918201208185015160408087015160608801519151959a509295505f9461201d94879492910161422e565b60408051601f1981840301815291815281516020928301206001600160401b0389165f9081526006909352912054909150811461206d5760405163671ebaaf60e11b815260040160405180910390fd5b6001600160401b0386165f908152600660205260408120556120906001886142ee565b84036120e45742600760109054906101000a90046001600160401b031684604001516120bc919061427b565b6001600160401b031611156120e45760405163c44a082160e01b815260040160405180910390fd5b84828460200151856040015133876060015160405160200161210b96959493929190614301565b6040516020818303038152906040528051906020012094505050508080612131906144fa565b915050611faa565b506121677f0000000000000000000000000000000000000000000000000000000000000000846118e6610816565b6005819055600780546001600160401b038416600160401b0267ffffffffffffffff60401b19909116179055604051639a908e7360e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639a908e73906121f790879086906004016001600160401b03929092168252602082015260400190565b6020604051808303815f875af1158015612213573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122379190614345565b6040519091506001600160401b038216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f876040516024016122a89695949392919061459c565b60408051601f198184030181529190526020810180516001600160e01b031663f811bff760e01b17905283519091506060905f0361237c5760f9601f83516122f091906145f0565b60405180604001604052806008815260200167202021007270e02560c21b8152507f000000000000000000000000000000000000000000000000000000000000000060405180604001604052806002815260200161101760f31b81525060e487604051602001612366979695949392919061460b565b6040516020818303038152906040529050612437565b815161ffff10156123a057604051631245c7c160e11b815260040160405180910390fd5b815160f96123af6020836145f0565b60405180604001604052806008815260200167202021007270e02560c21b8152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020016180b960f01b8152508588604051602001612424979695949392919061469e565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612495573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b0381166124c95760405163668b0cb360e11b815260040160405180910390fd5b6040515f906124f69084906405ca1ab1e090635ca1ab1e90601b906001600160f81b031990602001614731565b60408051601f198184030181529190529450505050505b949350505050565b5f546201000090046001600160a01b0316331461254557604051634755657960e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610ef8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633146125dc57604051631736745960e31b815260040160405180910390fd5b600555565b5f546201000090046001600160a01b0316331461261157604051634755657960e01b815260040160405180910390fd5b600361261d82826143ad565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610ef89190613d18565b6002546001600160a01b03163314612678576040516311e7be1560e01b815260040160405180910390fd5b42610e10600a546126899190614215565b11156126a857604051632cd6ab1760e11b815260040160405180910390fd5b600b5460ff16156126cc57604051630c7e985960e01b815260040160405180910390fd5b855f8190036126ee5760405163cb591a5f60e01b815260040160405180910390fd5b6103e881111561271157604051635acfba9d60e11b815260040160405180910390fd5b61271c602442614215565b866001600160401b0316111561274557604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561279d575f80fd5b505af11580156127af573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015612810573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061283491906141a4565b600754600554919250600160401b90046001600160401b031690815f805b86811015612a0b575f8e8e8381811061286d5761286d61449a565b90506080020180360381019061288391906144ae565b60408101519091506001600160401b03161561298b57856128a381614256565b9650505f815f01518260200151836040015184606001516040516020016128cd949392919061422e565b60408051601f1981840301815291815281516020928301206001600160401b038a165f9081526006909352912054909150811461291d5760405163671ebaaf60e11b815260040160405180910390fd5b85825f0151836020015184604001518f866060015160405160200161294796959493929190614301565b60405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f9055506129f8565b80516040516129a7918591602001918252602082015260400190565b60405160208183030381529060405280519060200120925084815f0151888f8e5f801b6040516020016129df96959493929190614301565b6040516020818303038152906040528051906020012094505b5080612a03816144fa565b915050612852565b506007546001600160401b039081169085161115612a3c5760405163c630a00d60e01b815260040160405180910390fd5b6005839055856001600160401b0385811690841614612ad4575f612a6084876141cf565b9050612a756001600160401b038216836142ee565b9150612aad7f0000000000000000000000000000000000000000000000000000000000000000826001600160401b03166118e6610816565b506007805467ffffffffffffffff60401b1916600160401b6001600160401b038816021790555b8015612bbd57612b5c337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119d0573d5f803e3d5ffd5b603f54604051633b51be4b60e01b81526001600160a01b0390911690633b51be4b90612b909085908d908d90600401614551565b5f6040518083038186803b158015612ba6575f80fd5b505afa158015612bb8573d5f803e3d5ffd5b505050505b604051639a908e7360e01b81526001600160401b0388166004820152602481018590525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015612c32573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c569190614345565b9050612c6288826141cf565b6001600160401b03168c6001600160401b031614611c0f57604051630d0386cd60e11b815260040160405180910390fd5b603f54600160a01b900460ff16612cbd576040516320864d6d60e21b815260040160405180910390fd5b612cca8585858585613311565b5050505050565b6008546001600160a01b03168015801590612cf557506001600160a01b0381163314155b15612d13576040516324eff8c360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d6f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d9391906142d3565b15612db157604051630724b1a360e31b815260040160405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e0e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e3291906141a4565b905082811115612e5557604051632354600f60e01b815260040160405180910390fd5b611388841115612e78576040516328a69b1f60e21b815260040160405180910390fd5b612ead6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163330846132d3565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f0a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f2e91906141a4565b600780549192506001600160401b03909116905f612f4b83614256565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550508585604051612f80929190614774565b6040519081900390208142612f966001436142ee565b40604051602001612faa949392919061422e565b60408051601f1981840301815291815281516020928301206007546001600160401b03165f908152600690935291205532330361303b576007546040805183815233602082015260608183018190525f9082015290516001600160401b03909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2613089565b6007546040516001600160401b03909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061308090849033908b908b90614783565b60405180910390a25b7f5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e9989879378868683426130b96001436142ee565b406040516130cb9594939291906147b7565b60405180910390a1505050505050565b60606001600160a01b038216156132665760405163c00f14ab60e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa15801561314f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261317691908101906147e7565b60405163318aee3d60e01b81526001600160a01b0384811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa1580156131e2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613206919061484f565b915091508163ffffffff165f14613247576009805463ffffffff8416600160a01b026001600160c01b03199091166001600160a01b03841617179055613263565b600980546001600160a01b0319166001600160a01b0386161790555b50505b919050565b6040516001600160a01b0383166024820152604481018290526132ce90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613939565b505050565b6040516001600160a01b038085166024830152831660448201526064810182905261330b9085906323b872dd60e01b90608401613297565b50505050565b6002546001600160a01b0316331461333c576040516311e7be1560e01b815260040160405180910390fd5b600b5460ff161561336057604051630c7e985960e01b815260040160405180910390fd5b42610e10600a546133719190614215565b111561339057604051632cd6ab1760e11b815260040160405180910390fd5b835f8190036133b25760405163cb591a5f60e01b815260040160405180910390fd5b6103e88111156133d557604051635acfba9d60e11b815260040160405180910390fd5b6133e0602442614215565b846001600160401b0316111561340957604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613461575f80fd5b505af1158015613473573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134d4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134f891906141a4565b600754600554919250600160401b90046001600160401b031690815f5b858110156136c7575f8b8b838181106135305761353061449a565b90506020028101906135429190614573565b61354b90614591565b805180516020909101206040820151919250906001600160401b031615613658578561357681614256565b9650505f8183602001518460400151856060015160405160200161359d949392919061422e565b60408051601f1981840301815291815281516020928301206001600160401b038a165f908152600690935291205490915081146135ed5760405163671ebaaf60e11b815260040160405180910390fd5b8582846020015185604001518d876060015160405160200161361496959493929190614301565b60405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f9055506136b2565b8151516201d4c0101561367e576040516328a69b1f60e21b815260040160405180910390fd5b60405161369990869083908a908f908e905f90602001614301565b6040516020818303038152906040528051906020012094505b505080806136bf906144fa565b915050613515565b506007546001600160401b0390811690841611156136f85760405163c630a00d60e01b815260040160405180910390fd5b6005829055846001600160401b0384811690831614613790575f61371c83866141cf565b90506137316001600160401b038216836142ee565b91506137697f0000000000000000000000000000000000000000000000000000000000000000826001600160401b03166118e6610816565b506007805467ffffffffffffffff60401b1916600160401b6001600160401b038716021790555b613812337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119d0573d5f803e3d5ffd5b604051639a908e7360e01b81526001600160401b0387166004820152602481018490525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015613887573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138ab9190614345565b90506138b787826141cf565b6001600160401b0316896001600160401b0316146138e857604051630d0386cd60e11b815260040160405180910390fd5b806001600160401b03167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668760405161392391815260200190565b60405180910390a2505050505050505050505050565b5f61398d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613a0a9092919063ffffffff16565b8051909150156132ce57808060200190518101906139ab91906142d3565b6132ce5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161106d565b606061250d84845f85855f80866001600160a01b03168587604051613a2f9190614887565b5f6040518083038185875af1925050503d805f8114613a69576040519150601f19603f3d011682016040523d82523d5f602084013e613a6e565b606091505b5091509150613a7f87838387613a8a565b979650505050505050565b60608315613af85782515f03613af1576001600160a01b0385163b613af15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161106d565b508161250d565b61250d8383815115613b0d5781518083602001fd5b8060405162461bcd60e51b815260040161106d9190613d18565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b0381118282101715613b5d57613b5d613b27565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613b8b57613b8b613b27565b604052919050565b5f6001600160401b03821115613bab57613bab613b27565b50601f01601f191660200190565b5f82601f830112613bc8575f80fd5b8135613bdb613bd682613b93565b613b63565b818152846020838601011115613bef575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160401b0381168114613c1f575f80fd5b50565b803561326681613c0b565b5f60808284031215613c3d575f80fd5b613c45613b3b565b905081356001600160401b03811115613c5c575f80fd5b613c6884828501613bb9565b825250602082013560208201526040820135613c8381613c0b565b806040830152506060820135606082015292915050565b5f60208284031215613caa575f80fd5b81356001600160401b03811115613cbf575f80fd5b61250d84828501613c2d565b5f5b83811015613ce5578181015183820152602001613ccd565b50505f910152565b5f8151808452613d04816020860160208601613ccb565b601f01601f19169290920160200192915050565b602081525f613d2a6020830184613ced565b9392505050565b8015158114613c1f575f80fd5b5f60208284031215613d4e575f80fd5b8135613d2a81613d31565b6001600160a01b0381168114613c1f575f80fd5b803561326681613d59565b5f805f60608486031215613d8a575f80fd5b8335613d9581613c0b565b9250602084013591506040840135613dac81613d59565b809150509250925092565b5f60208284031215613dc7575f80fd5b8135613d2a81613c0b565b5f60208284031215613de2575f80fd5b8135613d2a81613d59565b63ffffffff81168114613c1f575f80fd5b5f805f805f8060c08789031215613e13575f80fd5b8635613e1e81613d59565b95506020870135613e2e81613d59565b94506040870135613e3e81613ded565b93506060870135613e4e81613d59565b925060808701356001600160401b0380821115613e69575f80fd5b613e758a838b01613bb9565b935060a0890135915080821115613e8a575f80fd5b50613e9789828a01613bb9565b9150509295509295509295565b5f8083601f840112613eb4575f80fd5b5081356001600160401b03811115613eca575f80fd5b602083019150836020828501011115613ee1575f80fd5b9250929050565b5f805f805f805f60a0888a031215613efe575f80fd5b87356001600160401b0380821115613f14575f80fd5b818a0191508a601f830112613f27575f80fd5b813581811115613f35575f80fd5b8b60208260071b8501011115613f49575f80fd5b60208301995080985050613f5f60208b01613c22565b9650613f6d60408b01613c22565b9550613f7b60608b01613d6d565b945060808a0135915080821115613f90575f80fd5b50613f9d8a828b01613ea4565b989b979a50959850939692959293505050565b5f8083601f840112613fc0575f80fd5b5081356001600160401b03811115613fd6575f80fd5b6020830191508360208260051b8501011115613ee1575f80fd5b5f8060208385031215614001575f80fd5b82356001600160401b03811115614016575f80fd5b61402285828601613fb0565b90969095509350505050565b5f805f8060808587031215614041575f80fd5b843561404c81613ded565b9350602085013561405c81613d59565b9250604085013561406c81613ded565b915060608501356001600160401b03811115614086575f80fd5b61409287828801613bb9565b91505092959194509250565b5f602082840312156140ae575f80fd5b5035919050565b5f602082840312156140c5575f80fd5b81356001600160401b038111156140da575f80fd5b61250d84828501613bb9565b5f805f805f608086880312156140fa575f80fd5b85356001600160401b0381111561410f575f80fd5b61411b88828901613fb0565b909650945050602086013561412f81613c0b565b9250604086013561413f81613c0b565b9150606086013561414f81613d59565b809150509295509295909350565b5f805f6040848603121561416f575f80fd5b83356001600160401b03811115614184575f80fd5b61419086828701613ea4565b909790965060209590950135949350505050565b5f602082840312156141b4575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038281168282160390808211156141ef576141ef6141bb565b5092915050565b5f8261421057634e487b7160e01b5f52601260045260245ffd5b500490565b80820180821115614228576142286141bb565b92915050565b938452602084019290925260c01b6001600160c01b0319166040830152604882015260680190565b5f6001600160401b03808316818103614271576142716141bb565b6001019392505050565b6001600160401b038181168382160190808211156141ef576141ef6141bb565b600181811c908216806142af57607f821691505b6020821081036142cd57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f602082840312156142e3575f80fd5b8151613d2a81613d31565b81810381811115614228576142286141bb565b9586526020860194909452604085019290925260c01b6001600160c01b0319166060808501919091521b6001600160601b0319166068830152607c820152609c0190565b5f60208284031215614355575f80fd5b8151613d2a81613c0b565b601f8211156132ce575f81815260208120601f850160051c810160208610156143865750805b601f850160051c820191505b818110156143a557828155600101614392565b505050505050565b81516001600160401b038111156143c6576143c6613b27565b6143da816143d4845461429b565b84614360565b602080601f83116001811461440d575f84156143f65750858301515b5f19600386901b1c1916600185901b1785556143a5565b5f85815260208120601f198616915b8281101561443b5788860151825594840194600190910190840161441c565b508582101561445857878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f61447a6060830186613ced565b6020830194909452506001600160a01b0391909116604090910152919050565b634e487b7160e01b5f52603260045260245ffd5b5f608082840312156144be575f80fd5b6144c6613b3b565b823581526020830135602082015260408301356144e281613c0b565b60408201526060928301359281019290925250919050565b5f6001820161450b5761450b6141bb565b5060010190565b8082028115828204841417614228576142286141bb565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b838152604060208201525f61456a604083018486614529565b95945050505050565b5f8235607e19833603018112614587575f80fd5b9190910192915050565b5f6142283683613c2d565b63ffffffff87811682526001600160a01b038781166020840152908616604083015284811660608301528316608082015260c060a082018190525f906145e490830184613ced565b98975050505050505050565b61ffff8181168382160190808211156141ef576141ef6141bb565b5f60ff60f81b808a60f81b16835261ffff60f01b8960f01b166001840152875161463c816003860160208c01613ccb565b606088901b6001600160601b0319166003918501918201528651614667816017840160208b01613ccb565b808201915050818660f81b1660178201528451915061468d826018830160208801613ccb565b016018019998505050505050505050565b60ff60f81b8860f81b1681525f61ffff60f01b808960f01b16600184015287516146cf816003860160208c01613ccb565b606088901b6001600160601b03191660039185019182015286516146fa816017840160208b01613ccb565b808201915050818660f01b16601782015284519150614720826019830160208801613ccb565b016019019998505050505050505050565b5f8651614742818460208b01613ccb565b9190910194855250602084019290925260f81b6001600160f81b03199081166040840152166041820152604201919050565b818382375f9101908152919050565b8481526001600160a01b03841660208201526060604082018190525f906147ad9083018486614529565b9695505050505050565b608081525f6147ca608083018789614529565b602083019590955250604081019290925260609091015292915050565b5f602082840312156147f7575f80fd5b81516001600160401b0381111561480c575f80fd5b8201601f8101841361481c575f80fd5b805161482a613bd682613b93565b81815285602083850101111561483e575f80fd5b61456a826020830160208601613ccb565b5f8060408385031215614860575f80fd5b825161486b81613ded565b602084015190925061487c81613d59565b809150509250929050565b5f8251614587818460208701613ccb56fea264697066735822122066ba0caa282cc227895ef0f8158a58a4adf45f9e3e0fa57b2538ca0219a8f0b064736f6c63430008140033",
}

// PolygonzkevmABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmMetaData.ABI instead.
var PolygonzkevmABI = PolygonzkevmMetaData.ABI

// PolygonzkevmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmMetaData.Bin instead.
var PolygonzkevmBin = PolygonzkevmMetaData.Bin

// DeployPolygonzkevm deploys a new Ethereum contract, binding an instance of Polygonzkevm to it.
func DeployPolygonzkevm(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonzkevm, error) {
	parsed, err := PolygonzkevmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevm{PolygonzkevmCaller: PolygonzkevmCaller{contract: contract}, PolygonzkevmTransactor: PolygonzkevmTransactor{contract: contract}, PolygonzkevmFilterer: PolygonzkevmFilterer{contract: contract}}, nil
}

// Polygonzkevm is an auto generated Go binding around an Ethereum contract.
type Polygonzkevm struct {
	PolygonzkevmCaller     // Read-only binding to the contract
	PolygonzkevmTransactor // Write-only binding to the contract
	PolygonzkevmFilterer   // Log filterer for contract events
}

// PolygonzkevmCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmSession struct {
	Contract     *Polygonzkevm     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolygonzkevmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmCallerSession struct {
	Contract *PolygonzkevmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PolygonzkevmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmTransactorSession struct {
	Contract     *PolygonzkevmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PolygonzkevmRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmRaw struct {
	Contract *Polygonzkevm // Generic contract binding to access the raw methods on
}

// PolygonzkevmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmCallerRaw struct {
	Contract *PolygonzkevmCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmTransactorRaw struct {
	Contract *PolygonzkevmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevm creates a new instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevm(address common.Address, backend bind.ContractBackend) (*Polygonzkevm, error) {
	contract, err := bindPolygonzkevm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevm{PolygonzkevmCaller: PolygonzkevmCaller{contract: contract}, PolygonzkevmTransactor: PolygonzkevmTransactor{contract: contract}, PolygonzkevmFilterer: PolygonzkevmFilterer{contract: contract}}, nil
}

// NewPolygonzkevmCaller creates a new read-only instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevmCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmCaller, error) {
	contract, err := bindPolygonzkevm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmCaller{contract: contract}, nil
}

// NewPolygonzkevmTransactor creates a new write-only instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevmTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmTransactor, error) {
	contract, err := bindPolygonzkevm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmTransactor{contract: contract}, nil
}

// NewPolygonzkevmFilterer creates a new log filterer instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevmFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmFilterer, error) {
	contract, err := bindPolygonzkevm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmFilterer{contract: contract}, nil
}

// bindPolygonzkevm binds a generic wrapper to an already deployed contract.
func bindPolygonzkevm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevm *PolygonzkevmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevm.Contract.PolygonzkevmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevm *PolygonzkevmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.PolygonzkevmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevm *PolygonzkevmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.PolygonzkevmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevm *PolygonzkevmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevm *PolygonzkevmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevm *PolygonzkevmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevm.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevm.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonzkevm.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevm.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonzkevm.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevm.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonzkevm.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevm.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonzkevm.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevm *PolygonzkevmCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevm *PolygonzkevmSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevm.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonzkevm *PolygonzkevmCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonzkevm.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevm.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonzkevm.Contract.SIGNATUREINITIALIZETXR(&_Polygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevm.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonzkevm.Contract.SIGNATUREINITIALIZETXS(&_Polygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevm.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonzkevm *PolygonzkevmCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonzkevm.Contract.SIGNATUREINITIALIZETXV(&_Polygonzkevm.CallOpts)
}

// TESTVAR is a free data retrieval call binding the contract method 0x236ae0c0.
//
// Solidity: function TEST_VAR() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) TESTVAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "TEST_VAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TESTVAR is a free data retrieval call binding the contract method 0x236ae0c0.
//
// Solidity: function TEST_VAR() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) TESTVAR() (*big.Int, error) {
	return _Polygonzkevm.Contract.TESTVAR(&_Polygonzkevm.CallOpts)
}

// TESTVAR is a free data retrieval call binding the contract method 0x236ae0c0.
//
// Solidity: function TEST_VAR() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) TESTVAR() (*big.Int, error) {
	return _Polygonzkevm.Contract.TESTVAR(&_Polygonzkevm.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevm.Contract.TIMESTAMPRANGE(&_Polygonzkevm.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonzkevm.Contract.TIMESTAMPRANGE(&_Polygonzkevm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) Admin() (common.Address, error) {
	return _Polygonzkevm.Contract.Admin(&_Polygonzkevm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevm.Contract.Admin(&_Polygonzkevm.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.BridgeAddress(&_Polygonzkevm.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.BridgeAddress(&_Polygonzkevm.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevm.Contract.CalculatePolPerForceBatch(&_Polygonzkevm.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonzkevm.Contract.CalculatePolPerForceBatch(&_Polygonzkevm.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonzkevm.Contract.DataAvailabilityProtocol(&_Polygonzkevm.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonzkevm.Contract.DataAvailabilityProtocol(&_Polygonzkevm.CallOpts)
}

// ExitMode is a free data retrieval call binding the contract method 0x28e18a89.
//
// Solidity: function exitMode() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) ExitMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "exitMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExitMode is a free data retrieval call binding the contract method 0x28e18a89.
//
// Solidity: function exitMode() view returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) ExitMode() (bool, error) {
	return _Polygonzkevm.Contract.ExitMode(&_Polygonzkevm.CallOpts)
}

// ExitMode is a free data retrieval call binding the contract method 0x28e18a89.
//
// Solidity: function exitMode() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) ExitMode() (bool, error) {
	return _Polygonzkevm.Contract.ExitMode(&_Polygonzkevm.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.ForceBatchAddress(&_Polygonzkevm.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.ForceBatchAddress(&_Polygonzkevm.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.ForceBatchTimeout(&_Polygonzkevm.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.ForceBatchTimeout(&_Polygonzkevm.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevm.Contract.ForcedBatches(&_Polygonzkevm.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevm.Contract.ForcedBatches(&_Polygonzkevm.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.GasTokenAddress(&_Polygonzkevm.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.GasTokenAddress(&_Polygonzkevm.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevm *PolygonzkevmCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevm *PolygonzkevmSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevm.Contract.GasTokenNetwork(&_Polygonzkevm.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevm *PolygonzkevmCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevm.Contract.GasTokenNetwork(&_Polygonzkevm.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevm *PolygonzkevmSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevm.Contract.GenerateInitializeTransaction(&_Polygonzkevm.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonzkevm.Contract.GenerateInitializeTransaction(&_Polygonzkevm.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevm.Contract.GlobalExitRootManager(&_Polygonzkevm.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevm.Contract.GlobalExitRootManager(&_Polygonzkevm.CallOpts)
}

// IsRevertModeActive is a free data retrieval call binding the contract method 0x788aaa6a.
//
// Solidity: function isRevertModeActive() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) IsRevertModeActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "isRevertModeActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevertModeActive is a free data retrieval call binding the contract method 0x788aaa6a.
//
// Solidity: function isRevertModeActive() view returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) IsRevertModeActive() (bool, error) {
	return _Polygonzkevm.Contract.IsRevertModeActive(&_Polygonzkevm.CallOpts)
}

// IsRevertModeActive is a free data retrieval call binding the contract method 0x788aaa6a.
//
// Solidity: function isRevertModeActive() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) IsRevertModeActive() (bool, error) {
	return _Polygonzkevm.Contract.IsRevertModeActive(&_Polygonzkevm.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonzkevm.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonzkevm.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonzkevm.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonzkevm.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevm.Contract.LastAccInputHash(&_Polygonzkevm.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonzkevm.Contract.LastAccInputHash(&_Polygonzkevm.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatch(&_Polygonzkevm.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatch(&_Polygonzkevm.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatchSequenced(&_Polygonzkevm.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatchSequenced(&_Polygonzkevm.CallOpts)
}

// LastRevertModeTimestamp is a free data retrieval call binding the contract method 0xeb5db56a.
//
// Solidity: function lastRevertModeTimestamp() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) LastRevertModeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastRevertModeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRevertModeTimestamp is a free data retrieval call binding the contract method 0xeb5db56a.
//
// Solidity: function lastRevertModeTimestamp() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) LastRevertModeTimestamp() (*big.Int, error) {
	return _Polygonzkevm.Contract.LastRevertModeTimestamp(&_Polygonzkevm.CallOpts)
}

// LastRevertModeTimestamp is a free data retrieval call binding the contract method 0xeb5db56a.
//
// Solidity: function lastRevertModeTimestamp() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastRevertModeTimestamp() (*big.Int, error) {
	return _Polygonzkevm.Contract.LastRevertModeTimestamp(&_Polygonzkevm.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevm *PolygonzkevmCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevm *PolygonzkevmSession) NetworkName() (string, error) {
	return _Polygonzkevm.Contract.NetworkName(&_Polygonzkevm.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevm *PolygonzkevmCallerSession) NetworkName() (string, error) {
	return _Polygonzkevm.Contract.NetworkName(&_Polygonzkevm.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevm.Contract.PendingAdmin(&_Polygonzkevm.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevm.Contract.PendingAdmin(&_Polygonzkevm.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) Pol() (common.Address, error) {
	return _Polygonzkevm.Contract.Pol(&_Polygonzkevm.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) Pol() (common.Address, error) {
	return _Polygonzkevm.Contract.Pol(&_Polygonzkevm.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) RollupManager() (common.Address, error) {
	return _Polygonzkevm.Contract.RollupManager(&_Polygonzkevm.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) RollupManager() (common.Address, error) {
	return _Polygonzkevm.Contract.RollupManager(&_Polygonzkevm.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevm.Contract.TrustedSequencer(&_Polygonzkevm.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevm.Contract.TrustedSequencer(&_Polygonzkevm.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevm *PolygonzkevmCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevm *PolygonzkevmSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevm.Contract.TrustedSequencerURL(&_Polygonzkevm.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevm *PolygonzkevmCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevm.Contract.TrustedSequencerURL(&_Polygonzkevm.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevm *PolygonzkevmTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevm *PolygonzkevmSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.AcceptAdminRole(&_Polygonzkevm.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.AcceptAdminRole(&_Polygonzkevm.TransactOpts)
}

// EnterRevertMode is a paid mutator transaction binding the contract method 0x03744a63.
//
// Solidity: function enterRevertMode((bytes,bytes32,uint64,bytes32) forcedBatchData) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) EnterRevertMode(opts *bind.TransactOpts, forcedBatchData PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "enterRevertMode", forcedBatchData)
}

// EnterRevertMode is a paid mutator transaction binding the contract method 0x03744a63.
//
// Solidity: function enterRevertMode((bytes,bytes32,uint64,bytes32) forcedBatchData) returns()
func (_Polygonzkevm *PolygonzkevmSession) EnterRevertMode(forcedBatchData PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.EnterRevertMode(&_Polygonzkevm.TransactOpts, forcedBatchData)
}

// EnterRevertMode is a paid mutator transaction binding the contract method 0x03744a63.
//
// Solidity: function enterRevertMode((bytes,bytes32,uint64,bytes32) forcedBatchData) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) EnterRevertMode(forcedBatchData PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.EnterRevertMode(&_Polygonzkevm.TransactOpts, forcedBatchData)
}

// ExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x7cbe817f.
//
// Solidity: function executeAllForcedTransactions((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ExecuteAllForcedTransactions(opts *bind.TransactOpts, batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "executeAllForcedTransactions", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// ExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x7cbe817f.
//
// Solidity: function executeAllForcedTransactions((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonzkevm *PolygonzkevmSession) ExecuteAllForcedTransactions(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ExecuteAllForcedTransactions(&_Polygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// ExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x7cbe817f.
//
// Solidity: function executeAllForcedTransactions((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ExecuteAllForcedTransactions(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ExecuteAllForcedTransactions(&_Polygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevm *PolygonzkevmSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ForceBatch(&_Polygonzkevm.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ForceBatch(&_Polygonzkevm.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevm *PolygonzkevmSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.Initialize(&_Polygonzkevm.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.Initialize(&_Polygonzkevm.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnRevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xb151da2d.
//
// Solidity: function onRevertLastVerifiedBatch(bytes32 _prevAccInputHash) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) OnRevertLastVerifiedBatch(opts *bind.TransactOpts, _prevAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "onRevertLastVerifiedBatch", _prevAccInputHash)
}

// OnRevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xb151da2d.
//
// Solidity: function onRevertLastVerifiedBatch(bytes32 _prevAccInputHash) returns()
func (_Polygonzkevm *PolygonzkevmSession) OnRevertLastVerifiedBatch(_prevAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.OnRevertLastVerifiedBatch(&_Polygonzkevm.TransactOpts, _prevAccInputHash)
}

// OnRevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xb151da2d.
//
// Solidity: function onRevertLastVerifiedBatch(bytes32 _prevAccInputHash) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) OnRevertLastVerifiedBatch(_prevAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.OnRevertLastVerifiedBatch(&_Polygonzkevm.TransactOpts, _prevAccInputHash)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevm *PolygonzkevmSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.OnVerifyBatches(&_Polygonzkevm.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.OnVerifyBatches(&_Polygonzkevm.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "sequenceBatches", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonzkevm *PolygonzkevmSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceBatches(&_Polygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceBatches(&_Polygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "sequenceBatchesValidium", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonzkevm *PolygonzkevmSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceBatchesValidium(&_Polygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceBatchesValidium(&_Polygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevm *PolygonzkevmSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceForceBatches(&_Polygonzkevm.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceForceBatches(&_Polygonzkevm.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetDataAvailabilityProtocol(&_Polygonzkevm.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetDataAvailabilityProtocol(&_Polygonzkevm.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetForceBatchAddress(&_Polygonzkevm.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetForceBatchAddress(&_Polygonzkevm.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetForceBatchTimeout(&_Polygonzkevm.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetForceBatchTimeout(&_Polygonzkevm.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencer(&_Polygonzkevm.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencer(&_Polygonzkevm.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencerURL(&_Polygonzkevm.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencerURL(&_Polygonzkevm.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonzkevm *PolygonzkevmSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SwitchSequenceWithDataAvailability(&_Polygonzkevm.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SwitchSequenceWithDataAvailability(&_Polygonzkevm.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevm *PolygonzkevmSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.TransferAdminRole(&_Polygonzkevm.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.TransferAdminRole(&_Polygonzkevm.TransactOpts, newPendingAdmin)
}

// PolygonzkevmAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevm contract.
type PolygonzkevmAcceptAdminRoleIterator struct {
	Event *PolygonzkevmAcceptAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmAcceptAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmAcceptAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevm contract.
type PolygonzkevmAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmAcceptAdminRoleIterator{contract: _Polygonzkevm.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmAcceptAdminRole)
				if err := _Polygonzkevm.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmAcceptAdminRole, error) {
	event := new(PolygonzkevmAcceptAdminRole)
	if err := _Polygonzkevm.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmActivateRevertModeIterator is returned from FilterActivateRevertMode and is used to iterate over the raw logs and unpacked data for ActivateRevertMode events raised by the Polygonzkevm contract.
type PolygonzkevmActivateRevertModeIterator struct {
	Event *PolygonzkevmActivateRevertMode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmActivateRevertModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmActivateRevertMode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmActivateRevertMode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmActivateRevertModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmActivateRevertModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmActivateRevertMode represents a ActivateRevertMode event raised by the Polygonzkevm contract.
type PolygonzkevmActivateRevertMode struct {
	Sender          common.Address
	LastForcedBatch uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterActivateRevertMode is a free log retrieval operation binding the contract event 0xa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9.
//
// Solidity: event ActivateRevertMode(address _sender, uint64 _lastForcedBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterActivateRevertMode(opts *bind.FilterOpts) (*PolygonzkevmActivateRevertModeIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "ActivateRevertMode")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmActivateRevertModeIterator{contract: _Polygonzkevm.contract, event: "ActivateRevertMode", logs: logs, sub: sub}, nil
}

// WatchActivateRevertMode is a free log subscription operation binding the contract event 0xa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9.
//
// Solidity: event ActivateRevertMode(address _sender, uint64 _lastForcedBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchActivateRevertMode(opts *bind.WatchOpts, sink chan<- *PolygonzkevmActivateRevertMode) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "ActivateRevertMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmActivateRevertMode)
				if err := _Polygonzkevm.contract.UnpackLog(event, "ActivateRevertMode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActivateRevertMode is a log parse operation binding the contract event 0xa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9.
//
// Solidity: event ActivateRevertMode(address _sender, uint64 _lastForcedBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseActivateRevertMode(log types.Log) (*PolygonzkevmActivateRevertMode, error) {
	event := new(PolygonzkevmActivateRevertMode)
	if err := _Polygonzkevm.contract.UnpackLog(event, "ActivateRevertMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevm contract.
type PolygonzkevmForceBatchIterator struct {
	Event *PolygonzkevmForceBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmForceBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmForceBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmForceBatch represents a ForceBatch event raised by the Polygonzkevm contract.
type PolygonzkevmForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmForceBatchIterator{contract: _Polygonzkevm.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmForceBatch)
				if err := _Polygonzkevm.contract.UnpackLog(event, "ForceBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmForceBatch, error) {
	event := new(PolygonzkevmForceBatch)
	if err := _Polygonzkevm.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonzkevm contract.
type PolygonzkevmInitialSequenceBatchesIterator struct {
	Event *PolygonzkevmInitialSequenceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmInitialSequenceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmInitialSequenceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonzkevm contract.
type PolygonzkevmInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonzkevmInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmInitialSequenceBatchesIterator{contract: _Polygonzkevm.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmInitialSequenceBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialSequenceBatches is a log parse operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonzkevmInitialSequenceBatches, error) {
	event := new(PolygonzkevmInitialSequenceBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevm contract.
type PolygonzkevmInitializedIterator struct {
	Event *PolygonzkevmInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmInitialized represents a Initialized event raised by the Polygonzkevm contract.
type PolygonzkevmInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmInitializedIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmInitializedIterator{contract: _Polygonzkevm.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmInitialized)
				if err := _Polygonzkevm.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseInitialized(log types.Log) (*PolygonzkevmInitialized, error) {
	event := new(PolygonzkevmInitialized)
	if err := _Polygonzkevm.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmLogForcedBatchDataIterator is returned from FilterLogForcedBatchData and is used to iterate over the raw logs and unpacked data for LogForcedBatchData events raised by the Polygonzkevm contract.
type PolygonzkevmLogForcedBatchDataIterator struct {
	Event *PolygonzkevmLogForcedBatchData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmLogForcedBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmLogForcedBatchData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmLogForcedBatchData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmLogForcedBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmLogForcedBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmLogForcedBatchData represents a LogForcedBatchData event raised by the Polygonzkevm contract.
type PolygonzkevmLogForcedBatchData struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Timestamp          *big.Int
	PrentBlockHash     [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogForcedBatchData is a free log retrieval operation binding the contract event 0x5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e9989879378.
//
// Solidity: event LogForcedBatchData(bytes _transactions, bytes32 _lastGlobalExitRoot, uint256 _timestamp, bytes32 _prentBlockHash)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterLogForcedBatchData(opts *bind.FilterOpts) (*PolygonzkevmLogForcedBatchDataIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "LogForcedBatchData")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmLogForcedBatchDataIterator{contract: _Polygonzkevm.contract, event: "LogForcedBatchData", logs: logs, sub: sub}, nil
}

// WatchLogForcedBatchData is a free log subscription operation binding the contract event 0x5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e9989879378.
//
// Solidity: event LogForcedBatchData(bytes _transactions, bytes32 _lastGlobalExitRoot, uint256 _timestamp, bytes32 _prentBlockHash)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchLogForcedBatchData(opts *bind.WatchOpts, sink chan<- *PolygonzkevmLogForcedBatchData) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "LogForcedBatchData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmLogForcedBatchData)
				if err := _Polygonzkevm.contract.UnpackLog(event, "LogForcedBatchData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogForcedBatchData is a log parse operation binding the contract event 0x5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e9989879378.
//
// Solidity: event LogForcedBatchData(bytes _transactions, bytes32 _lastGlobalExitRoot, uint256 _timestamp, bytes32 _prentBlockHash)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseLogForcedBatchData(log types.Log) (*PolygonzkevmLogForcedBatchData, error) {
	event := new(PolygonzkevmLogForcedBatchData)
	if err := _Polygonzkevm.contract.UnpackLog(event, "LogForcedBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevm contract.
type PolygonzkevmSequenceBatchesIterator struct {
	Event *PolygonzkevmSequenceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSequenceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSequenceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSequenceBatches represents a SequenceBatches event raised by the Polygonzkevm contract.
type PolygonzkevmSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSequenceBatchesIterator{contract: _Polygonzkevm.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSequenceBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequenceBatches is a log parse operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmSequenceBatches, error) {
	event := new(PolygonzkevmSequenceBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevm contract.
type PolygonzkevmSequenceForceBatchesIterator struct {
	Event *PolygonzkevmSequenceForceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSequenceForceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSequenceForceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevm contract.
type PolygonzkevmSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSequenceForceBatchesIterator{contract: _Polygonzkevm.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSequenceForceBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmSequenceForceBatches, error) {
	event := new(PolygonzkevmSequenceForceBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Polygonzkevm contract.
type PolygonzkevmSetDataAvailabilityProtocolIterator struct {
	Event *PolygonzkevmSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetDataAvailabilityProtocol)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSetDataAvailabilityProtocol)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Polygonzkevm contract.
type PolygonzkevmSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*PolygonzkevmSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetDataAvailabilityProtocolIterator{contract: _Polygonzkevm.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetDataAvailabilityProtocol)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDataAvailabilityProtocol is a log parse operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*PolygonzkevmSetDataAvailabilityProtocol, error) {
	event := new(PolygonzkevmSetDataAvailabilityProtocol)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonzkevm contract.
type PolygonzkevmSetForceBatchAddressIterator struct {
	Event *PolygonzkevmSetForceBatchAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetForceBatchAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSetForceBatchAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonzkevm contract.
type PolygonzkevmSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonzkevmSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetForceBatchAddressIterator{contract: _Polygonzkevm.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetForceBatchAddress)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForceBatchAddress is a log parse operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonzkevmSetForceBatchAddress, error) {
	event := new(PolygonzkevmSetForceBatchAddress)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevm contract.
type PolygonzkevmSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmSetForceBatchTimeout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetForceBatchTimeout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSetForceBatchTimeout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevm contract.
type PolygonzkevmSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetForceBatchTimeoutIterator{contract: _Polygonzkevm.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetForceBatchTimeout)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmSetForceBatchTimeout, error) {
	event := new(PolygonzkevmSetForceBatchTimeout)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencerIterator struct {
	Event *PolygonzkevmSetTrustedSequencer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetTrustedSequencer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSetTrustedSequencer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetTrustedSequencerIterator{contract: _Polygonzkevm.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetTrustedSequencer)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmSetTrustedSequencer, error) {
	event := new(PolygonzkevmSetTrustedSequencer)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmSetTrustedSequencerURL // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetTrustedSequencerURL)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSetTrustedSequencerURL)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetTrustedSequencerURLIterator{contract: _Polygonzkevm.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetTrustedSequencerURL)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmSetTrustedSequencerURL)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Polygonzkevm contract.
type PolygonzkevmSwitchSequenceWithDataAvailabilityIterator struct {
	Event *PolygonzkevmSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSwitchSequenceWithDataAvailability)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmSwitchSequenceWithDataAvailability)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Polygonzkevm contract.
type PolygonzkevmSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*PolygonzkevmSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSwitchSequenceWithDataAvailabilityIterator{contract: _Polygonzkevm.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSwitchSequenceWithDataAvailability)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwitchSequenceWithDataAvailability is a log parse operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*PolygonzkevmSwitchSequenceWithDataAvailability, error) {
	event := new(PolygonzkevmSwitchSequenceWithDataAvailability)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevm contract.
type PolygonzkevmTransferAdminRoleIterator struct {
	Event *PolygonzkevmTransferAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmTransferAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmTransferAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevm contract.
type PolygonzkevmTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmTransferAdminRoleIterator{contract: _Polygonzkevm.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmTransferAdminRole)
				if err := _Polygonzkevm.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmTransferAdminRole, error) {
	event := new(PolygonzkevmTransferAdminRole)
	if err := _Polygonzkevm.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevm contract.
type PolygonzkevmVerifyBatchesIterator struct {
	Event *PolygonzkevmVerifyBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonzkevmVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmVerifyBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonzkevmVerifyBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonzkevmVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmVerifyBatches represents a VerifyBatches event raised by the Polygonzkevm contract.
type PolygonzkevmVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmVerifyBatchesIterator{contract: _Polygonzkevm.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmVerifyBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmVerifyBatches, error) {
	event := new(PolygonzkevmVerifyBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
