// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package etrogpolygonzkevm

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

// EtrogpolygonzkevmMetaData contains all meta data concerning the Etrogpolygonzkevm contract.
var EtrogpolygonzkevmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecuteForcedBatchesNotPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitModeIsActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertModeIsActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertModeIsNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertNotExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TxWindowNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"noForcedBatchPending\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_lastForcedBatch\",\"type\":\"uint64\"}],\"name\":\"ActivateRevertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_prentBlockHash\",\"type\":\"bytes32\"}],\"name\":\"LogForcedBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData\",\"name\":\"forcedBatchData\",\"type\":\"tuple\"}],\"name\":\"enterRevertMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"executeAllForcedTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRevertModeActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRevertModeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_prevAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onRevertLastVerifiedBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrog.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801562000011575f80fd5b50604051620048823803806200488283398101604081905262000034916200006f565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d4565b6001600160a01b03811681146200006c575f80fd5b50565b5f805f806080858703121562000083575f80fd5b8451620000908162000057565b6020860151909450620000a38162000057565b6040860151909350620000b68162000057565b6060860151909250620000c98162000057565b939692955090935050565b60805160a05160c05160e051614670620002125f395f81816104fe01528181610a8701528181610c5e01528181610d5601528181610efd0152818161112f015281816114bc015281816115730152818161182d015281816118e80152818161197401528181611a5a01528181611d3f01528181611de7015281816121d0015281816126b70152818161271b0152818161273d0152818161287f01528181612a0401528181612aa30152818161342701528181613485015281816134a7015261352801525f818161066301528181611f4e0152818161200d0152818161238201528181612dfa01528181612e8b01526130fa01525f81816107170152818161104e015281816123f101528181612b9f015261316901525f818161076f0152818161081401528181611d7b015281816127cf0152612b7401526146705ff3fe608060405234801561000f575f80fd5b5060043610610332575f3560e01c80637a5460c5116101af578063c7fffd4b116100fe578063e46761c41161009e578063eaeb077b11610079578063eaeb077b146107b7578063eb5db56a146107ca578063f35dda47146107d3578063f851a440146107db575f80fd5b8063e46761c41461076a578063e57a0b4c14610791578063e7a7ed02146107a4575f80fd5b8063d02103ca116100d9578063d02103ca14610712578063d7bc90ff14610739578063db5b0ed714610744578063def57e5414610757575f80fd5b8063c7fffd4b146106e4578063c89e42df146106ec578063cfa8ed47146106ff575f80fd5b80639f26f84011610169578063ada8f91911610144578063ada8f91914610698578063b0afe154146106ab578063b151da2d146106b7578063c754c7ed146106ca575f80fd5b80639f26f8401461064b578063a3c573eb1461065e578063a652f26c14610685575f80fd5b80637a5460c5146105ce5780637cd76b8b146105ef5780638c3d73011461060257806391cafe321461060a57806394aa909f1461061d5780639e00187714610630575f80fd5b806340b5de6c11610285578063542028d5116102255780636e05d2cd116102005780636e05d2cd146105975780636ff512cc146105a057806371257022146105b3578063788aaa6a146105c6575f80fd5b8063542028d514610568578063676870d2146105705780636b8616ce14610578575f80fd5b806349b7b8021161026057806349b7b802146104f95780634c21fef3146105205780634e4877061461053457806352bdeb6d14610547575f80fd5b806340b5de6c1461049757806342308fab146104bf57806345605267146104c7575f80fd5b806326782247116102f05780632c111c06116102cb5780632c111c061461043257806332c2d153146104455780633c351e10146104585780633cbc795b1461046b575f80fd5b806326782247146103d757806328e18a89146104025780632acdc2b61461041f575f80fd5b8062d0295d14610336578063035089631461035157806303744a631461036c57806305835f3714610381578063107bf28c146103b557806311e892d4146103bd575b5f80fd5b61033e6107f3565b6040519081526020015b60405180910390f35b610359602081565b60405161ffff9091168152602001610348565b61037f61037a366004613989565b6108cc565b005b6103a860405180604001604052806008815260200167202021007270e02560c21b81525081565b6040516103489190613a07565b6103a8610b2e565b6103c560f981565b60405160ff9091168152602001610348565b6001546103ea906001600160a01b031681565b6040516001600160a01b039091168152602001610348565b600b5461040f9060ff1681565b6040519015158152602001610348565b61037f61042d366004613a2d565b610bba565b6008546103ea906001600160a01b031681565b61037f610453366004613a67565b610c5c565b6009546103ea906001600160a01b031681565b60095461048290600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610348565b6104a66001600160f81b031981565b6040516001600160f81b03199091168152602001610348565b61033e602481565b6007546104e190600160401b90046001600160401b031681565b6040516001600160401b039091168152602001610348565b6103ea7f000000000000000000000000000000000000000000000000000000000000000081565b603e5461040f90600160a01b900460ff1681565b61037f610542366004613aa6565b610cf7565b6103a860405180604001604052806002815260200161101760f31b81525081565b6103a8610e70565b610359601f81565b61033e610586366004613aa6565b60066020525f908152604090205481565b61033e60055481565b61037f6105ae366004613ac1565b610e7d565b61037f6105c1366004613aed565b610efb565b61040f6112ea565b6103a86040518060400160405280600281526020016180b960f01b81525081565b61037f6105fd366004613ac1565b61130c565b61037f61138a565b61037f610618366004613ac1565b611413565b61037f61062b366004613bda565b6114ba565b6103ea73a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b61037f610659366004613c44565b611a11565b6103ea7f000000000000000000000000000000000000000000000000000000000000000081565b6103a8610693366004613c82565b611eb1565b61037f6106a6366004613ac1565b612150565b61033e6405ca1ab1e081565b61037f6106c5366004613cf2565b6121ce565b6007546104e190600160801b90046001600160401b031681565b6103c560e481565b61037f6106fa366004613d09565b61221c565b6002546103ea906001600160a01b031681565b6103ea7f000000000000000000000000000000000000000000000000000000000000000081565b61033e635ca1ab1e81565b61037f610752366004613d77565b612288565b61037f610765366004613e3f565b612982565b6103ea7f000000000000000000000000000000000000000000000000000000000000000081565b603e546103ea906001600160a01b031681565b6007546104e1906001600160401b031681565b61037f6107c5366004613eb6565b6129c0565b61033e600a5481565b6103c5601b81565b5f546103ea906201000090046001600160a01b031681565b6040516370a0823160e01b81523060048201525f9081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610859573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061087d9190613efd565b6007549091505f906108a1906001600160401b03600160401b820481169116613f28565b6001600160401b03169050805f036108bb575f9250505090565b6108c58183613f4f565b9250505090565b4261012c600a546108dd9190613f6e565b11156108fc57604051632cd6ab1760e11b815260040160405180910390fd5b6008546001600160a01b0316801580159061092057506001600160a01b0381163314155b1561093e576040516324eff8c360e01b815260040160405180910390fd5b6007546001600160401b03600160401b8204811691168111156109745760405163c630a00d60e01b815260040160405180910390fd5b6007546001600160401b038083169116036109a25760405163115cad3f60e01b815260040160405180910390fd5b8251805160209182012081850151604080870151606088015191515f956109cc9594939101613f87565b60405160208183030381529060405280519060200120905060065f836109f190613faf565b9350836001600160401b03166001600160401b031681526020019081526020015f20548114610a335760405163671ebaaf60e11b815260040160405180910390fd5b4260b48560400151610a459190613fd4565b6001600160401b03161115610a6d57604051631637066960e11b815260040160405180910390fd5b42600a5560405163750d3eef60e11b8152600160048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ea1a7dde906024015f604051808303815f87803b158015610ad0575f80fd5b505af1158015610ae2573d5f803e3d5ffd5b5050604080513381526001600160401b03861660208201527fa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9935001905060405180910390a150505050565b60048054610b3b90613ff4565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6790613ff4565b8015610bb25780601f10610b8957610100808354040283529160200191610bb2565b820191905f5260205f20905b815481529060010190602001808311610b9557829003601f168201915b505050505081565b5f546201000090046001600160a01b03163314610bea57604051634755657960e01b815260040160405180910390fd5b603e54600160a01b900460ff16151581151503610c1a57604051632f873d5f60e11b815260040160405180910390fd5b603e805460ff60a01b1916600160a01b831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41905f90a150565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610ca557604051631736745960e31b815260040160405180910390fd5b806001600160a01b0316836001600160401b03167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610cea91815260200190565b60405180910390a3505050565b5f546201000090046001600160a01b03163314610d2757604051634755657960e01b815260040160405180910390fd5b62093a806001600160401b0382161115610d545760405163f5e37f2f60e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610db0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dd4919061402c565b610e0e576007546001600160401b03600160801b909104811690821610610e0e5760405163f5e37f2f60e01b815260040160405180910390fd5b6007805467ffffffffffffffff60801b1916600160801b6001600160401b038416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b60038054610b3b90613ff4565b5f546201000090046001600160a01b03163314610ead57604051634755657960e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610e65565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610f4457604051631736745960e31b815260040160405180910390fd5b5f54610100900460ff1615808015610f6257505f54600160ff909116105b80610f7b5750303b158015610f7b57505f5460ff166001145b610fe35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015611004575f805461ff0019166101001790555b5f61100e85612dca565b6009549091505f9061103a9088906001600160a01b03811690600160a01b900463ffffffff1685611eb1565b90505f818051906020012090505f4290505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110a8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110cc9190613efd565b90505f808483858f6110df600143614047565b406040516020016110f59695949392919061405a565b60408051808303601f190181529082905280516020909101206005819055639a908e7360e01b8252600160048301526024820181905291507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af115801561117d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111a1919061409e565b508c5f60026101000a8154816001600160a01b0302191690836001600160a01b031602179055508b60025f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555088600390816111fd9190614106565b50600461120a8982614106565b508c60085f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555062069780600760106101000a8154816001600160401b0302191690836001600160401b031602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e60405161128e939291906141c1565b60405180910390a150505050505080156112e1575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f4261012c600a546112fc9190613f6e565b1161130657505f90565b50600190565b5f546201000090046001600160a01b0316331461133c57604051634755657960e01b815260040160405180910390fd5b603e80546001600160a01b0319166001600160a01b0383169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a90602001610e65565b6001546001600160a01b031633146113b55760405163d1ec4b2360e01b815260040160405180910390fd5b6001545f805462010000600160b01b0319166001600160a01b039092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f546201000090046001600160a01b0316331461144357604051634755657960e01b815260040160405180910390fd5b6008546001600160a01b031661146c576040516319126e9b60e31b815260040160405180910390fd5b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610e65565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637fea08646040518163ffffffff1660e01b8152600401602060405180830381865afa158015611516573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061153a919061402c565b156115585760405163bf52841f60e01b815260040160405180910390fd5b4261012c600a546115699190613f6e565b1115806115f357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b40a15c36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115cd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115f1919061402c565b155b80156116025750600b5460ff16155b1561162057604051638ff06f2b60e01b815260040160405180910390fd5b835f8190036116425760405163cb591a5f60e01b815260040160405180910390fd5b6103e881111561166557604051635acfba9d60e11b815260040160405180910390fd5b6007546001600160401b0380821691611687918491600160401b900416613f6e565b11156116a65760405163c630a00d60e01b815260040160405180910390fd5b600754600554600160401b9091046001600160401b0316905f5b838110156117f1575f8989838181106116db576116db6141f3565b90506020028101906116ed9190614207565b6116f690614225565b90508361170281613faf565b825180516020918201208185015160408087015160608801519151959a509295505f94611733948794929101613f87565b60408051601f1981840301815291815281516020928301206001600160401b0389165f908152600690935291205490915081146117835760405163671ebaaf60e11b815260040160405180910390fd5b6001600160401b0386165f908152600660209081526040808320929092558481015185830151606087015193516117c3948a94889493923392910161405a565b60405160208183030381529060405280519060200120945050505080806117e990614230565b9150506116c0565b5060058190556007805467ffffffffffffffff60401b1916600160401b6001600160401b0385160217905560405163107b19b760e21b81525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906341ec66dc9061187490879086908c908c9033908d90600401614248565b6020604051808303815f875af1158015611890573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118b4919061409e565b6007549091506001600160401b03808216600160401b909204160361195a5760405163750d3eef60e11b81525f60048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ea1a7dde906024015f604051808303815f87803b158015611931575f80fd5b505af1158015611943573d5f803e3d5ffd5b5050600b805460ff191660011790555061195f9050565b42600a555b604051630a0291dd60e21b81525f60048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063280a4774906024015f604051808303815f87803b1580156119bd575f80fd5b505af11580156119cf573d5f803e3d5ffd5b50506040516001600160401b03841692507f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a491505f90a2505050505050505050565b6008546001600160a01b03168015801590611a3557506001600160a01b0381163314155b15611a53576040516324eff8c360e01b815260040160405180910390fd5b4262093a807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ab4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ad8919061409e565b611ae29190613fd4565b6001600160401b03161115611b0a57604051630f527b5360e21b815260040160405180910390fd5b815f819003611b2c5760405163cb591a5f60e01b815260040160405180910390fd5b6103e8811115611b4f57604051635acfba9d60e11b815260040160405180910390fd5b6007546001600160401b0380821691611b71918491600160401b900416613f6e565b1115611b905760405163c630a00d60e01b815260040160405180910390fd5b600754600554600160401b9091046001600160401b0316905f5b83811015611d39575f878783818110611bc557611bc56141f3565b9050602002810190611bd79190614207565b611be090614225565b905083611bec81613faf565b825180516020918201208185015160408087015160608801519151959a509295505f94611c1d948794929101613f87565b60408051601f1981840301815291815281516020928301206001600160401b0389165f90815260069093529120549091508114611c6d5760405163671ebaaf60e11b815260040160405180910390fd5b6001600160401b0386165f90815260066020526040812055611c90600188614047565b8403611ce45742600760109054906101000a90046001600160401b03168460400151611cbc9190613fd4565b6001600160401b03161115611ce45760405163c44a082160e01b815260040160405180910390fd5b848284602001518560400151338760600151604051602001611d0b9695949392919061405a565b6040516020818303038152906040528051906020012094505050508080611d3190614230565b915050611baa565b50611da27f000000000000000000000000000000000000000000000000000000000000000084611d676107f3565b611d719190614291565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190612f5a565b6005819055600780546001600160401b038416600160401b0267ffffffffffffffff60401b19909116179055604051639a908e7360e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390611e3290879086906004016001600160401b03929092168252602082015260400190565b6020604051808303815f875af1158015611e4e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e72919061409e565b6040519091506001600160401b038216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f87604051602401611ee3969594939291906142a8565b60408051601f198184030181529190526020810180516001600160e01b031663f811bff760e01b17905283519091506060905f03611fb75760f9601f8351611f2b91906142fc565b60405180604001604052806008815260200167202021007270e02560c21b8152507f000000000000000000000000000000000000000000000000000000000000000060405180604001604052806002815260200161101760f31b81525060e487604051602001611fa19796959493929190614317565b6040516020818303038152906040529050612072565b815161ffff1015611fdb57604051631245c7c160e11b815260040160405180910390fd5b815160f9611fea6020836142fc565b60405180604001604052806008815260200167202021007270e02560c21b8152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020016180b960f01b815250858860405160200161205f97969594939291906143aa565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa1580156120d0573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b0381166121045760405163668b0cb360e11b815260040160405180910390fd5b6040515f906121319084906405ca1ab1e090635ca1ab1e90601b906001600160f81b03199060200161443d565b60408051601f198184030181529190529450505050505b949350505050565b5f546201000090046001600160a01b0316331461218057604051634755657960e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610e65565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331461221757604051631736745960e31b815260040160405180910390fd5b600555565b5f546201000090046001600160a01b0316331461224c57604051634755657960e01b815260040160405180910390fd5b60036122588282614106565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610e659190613a07565b6002546001600160a01b031633146122b3576040516311e7be1560e01b815260040160405180910390fd5b4261012c600a546122c49190613f6e565b11156122e357604051632cd6ab1760e11b815260040160405180910390fd5b600b5460ff161561230757604051630c7e985960e01b815260040160405180910390fd5b855f8190036123295760405163cb591a5f60e01b815260040160405180910390fd5b6103e881111561234c57604051635acfba9d60e11b815260040160405180910390fd5b612357602442613f6e565b866001600160401b0316111561238057604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156123d8575f80fd5b505af11580156123ea573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa15801561244b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061246f9190613efd565b600754600554919250600160401b90046001600160401b031690815f805b86811015612646575f8e8e838181106124a8576124a86141f3565b9050608002018036038101906124be9190614480565b60408101519091506001600160401b0316156125c657856124de81613faf565b9650505f815f01518260200151836040015184606001516040516020016125089493929190613f87565b60408051601f1981840301815291815281516020928301206001600160401b038a165f908152600690935291205490915081146125585760405163671ebaaf60e11b815260040160405180910390fd5b85825f0151836020015184604001518f86606001516040516020016125829695949392919061405a565b60405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f905550612633565b80516040516125e2918591602001918252602082015260400190565b60405160208183030381529060405280519060200120925084815f0151888f8e5f801b60405160200161261a9695949392919061405a565b6040516020818303038152906040528051906020012094505b508061263e81614230565b91505061248d565b506007546001600160401b0390811690851611156126775760405163c630a00d60e01b815260040160405180910390fd5b6005839055856001600160401b038581169084161461270f575f61269b8487613f28565b90506126b06001600160401b03821683614047565b91506126e87f0000000000000000000000000000000000000000000000000000000000000000826001600160401b0316611d676107f3565b506007805467ffffffffffffffff60401b1916600160401b6001600160401b038816021790555b8015612858576127f7337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612797573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127bb9190613efd565b6127c59190614291565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016929190612fc2565b603e54604051633b51be4b60e01b81526001600160a01b0390911690633b51be4b9061282b9085908d908d906004016144f4565b5f6040518083038186803b158015612841575f80fd5b505afa158015612853573d5f803e3d5ffd5b505050505b604051639a908e7360e01b81526001600160401b0388166004820152602481018590525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af11580156128cd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128f1919061409e565b90506128fd8882613f28565b6001600160401b03168c6001600160401b03161461292e57604051630d0386cd60e11b815260040160405180910390fd5b806001600160401b03167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668860405161296991815260200190565b60405180910390a2505050505050505050505050505050565b603e54600160a01b900460ff166129ac576040516320864d6d60e21b815260040160405180910390fd5b6129b98585858585613000565b5050505050565b6008546001600160a01b031680158015906129e457506001600160a01b0381163314155b15612a02576040516324eff8c360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a5e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a82919061402c565b15612aa057604051630724b1a360e31b815260040160405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612afd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b219190613efd565b905082811115612b4457604051632354600f60e01b815260040160405180910390fd5b611388841115612b67576040516328a69b1f60e21b815260040160405180910390fd5b612b9c6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084612fc2565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bf9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c1d9190613efd565b600780549192506001600160401b03909116905f612c3a83613faf565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550508585604051612c6f929190614516565b6040519081900390208142612c85600143614047565b40604051602001612c999493929190613f87565b60408051601f1981840301815291815281516020928301206007546001600160401b03165f9081526006909352912055323303612d2a576007546040805183815233602082015260608183018190525f9082015290516001600160401b03909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2612d78565b6007546040516001600160401b03909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93190612d6f90849033908b908b90614525565b60405180910390a25b7f5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e998987937886868342612da8600143614047565b40604051612dba959493929190614559565b60405180910390a1505050505050565b60606001600160a01b03821615612f555760405163c00f14ab60e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa158015612e3e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612e659190810190614589565b60405163318aee3d60e01b81526001600160a01b0384811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015612ed1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ef591906145f1565b915091508163ffffffff165f14612f36576009805463ffffffff8416600160a01b026001600160c01b03199091166001600160a01b03841617179055612f52565b600980546001600160a01b0319166001600160a01b0386161790555b50505b919050565b6040516001600160a01b038316602482015260448101829052612fbd90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613628565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052612ffa9085906323b872dd60e01b90608401612f86565b50505050565b6002546001600160a01b0316331461302b576040516311e7be1560e01b815260040160405180910390fd5b600b5460ff161561304f57604051630c7e985960e01b815260040160405180910390fd5b4261012c600a546130609190613f6e565b111561307f57604051632cd6ab1760e11b815260040160405180910390fd5b835f8190036130a15760405163cb591a5f60e01b815260040160405180910390fd5b6103e88111156130c457604051635acfba9d60e11b815260040160405180910390fd5b6130cf602442613f6e565b846001600160401b031611156130f857604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613150575f80fd5b505af1158015613162573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156131c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131e79190613efd565b600754600554919250600160401b90046001600160401b031690815f5b858110156133b6575f8b8b8381811061321f5761321f6141f3565b90506020028101906132319190614207565b61323a90614225565b805180516020909101206040820151919250906001600160401b031615613347578561326581613faf565b9650505f8183602001518460400151856060015160405160200161328c9493929190613f87565b60408051601f1981840301815291815281516020928301206001600160401b038a165f908152600690935291205490915081146132dc5760405163671ebaaf60e11b815260040160405180910390fd5b8582846020015185604001518d87606001516040516020016133039695949392919061405a565b60405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f9055506133a1565b8151516201d4c0101561336d576040516328a69b1f60e21b815260040160405180910390fd5b60405161338890869083908a908f908e905f9060200161405a565b6040516020818303038152906040528051906020012094505b505080806133ae90614230565b915050613204565b506007546001600160401b0390811690841611156133e75760405163c630a00d60e01b815260040160405180910390fd5b6005829055846001600160401b038481169083161461347f575f61340b8386613f28565b90506134206001600160401b03821683614047565b91506134587f0000000000000000000000000000000000000000000000000000000000000000826001600160401b0316611d676107f3565b506007805467ffffffffffffffff60401b1916600160401b6001600160401b038716021790555b613501337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612797573d5f803e3d5ffd5b604051639a908e7360e01b81526001600160401b0387166004820152602481018490525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015613576573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061359a919061409e565b90506135a68782613f28565b6001600160401b0316896001600160401b0316146135d757604051630d0386cd60e11b815260040160405180910390fd5b806001600160401b03167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668760405161361291815260200190565b60405180910390a2505050505050505050505050565b5f61367c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166136f99092919063ffffffff16565b805190915015612fbd578080602001905181019061369a919061402c565b612fbd5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610fda565b606061214884845f85855f80866001600160a01b0316858760405161371e9190614629565b5f6040518083038185875af1925050503d805f8114613758576040519150601f19603f3d011682016040523d82523d5f602084013e61375d565b606091505b509150915061376e87838387613779565b979650505050505050565b606083156137e75782515f036137e0576001600160a01b0385163b6137e05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610fda565b5081612148565b61214883838151156137fc5781518083602001fd5b8060405162461bcd60e51b8152600401610fda9190613a07565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b038111828210171561384c5761384c613816565b60405290565b604051601f8201601f191681016001600160401b038111828210171561387a5761387a613816565b604052919050565b5f6001600160401b0382111561389a5761389a613816565b50601f01601f191660200190565b5f82601f8301126138b7575f80fd5b81356138ca6138c582613882565b613852565b8181528460208386010111156138de575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160401b038116811461390e575f80fd5b50565b8035612f55816138fa565b5f6080828403121561392c575f80fd5b61393461382a565b905081356001600160401b0381111561394b575f80fd5b613957848285016138a8565b825250602082013560208201526040820135613972816138fa565b806040830152506060820135606082015292915050565b5f60208284031215613999575f80fd5b81356001600160401b038111156139ae575f80fd5b6121488482850161391c565b5f5b838110156139d45781810151838201526020016139bc565b50505f910152565b5f81518084526139f38160208601602086016139ba565b601f01601f19169290920160200192915050565b602081525f613a1960208301846139dc565b9392505050565b801515811461390e575f80fd5b5f60208284031215613a3d575f80fd5b8135613a1981613a20565b6001600160a01b038116811461390e575f80fd5b8035612f5581613a48565b5f805f60608486031215613a79575f80fd5b8335613a84816138fa565b9250602084013591506040840135613a9b81613a48565b809150509250925092565b5f60208284031215613ab6575f80fd5b8135613a19816138fa565b5f60208284031215613ad1575f80fd5b8135613a1981613a48565b63ffffffff8116811461390e575f80fd5b5f805f805f8060c08789031215613b02575f80fd5b8635613b0d81613a48565b95506020870135613b1d81613a48565b94506040870135613b2d81613adc565b93506060870135613b3d81613a48565b925060808701356001600160401b0380821115613b58575f80fd5b613b648a838b016138a8565b935060a0890135915080821115613b79575f80fd5b50613b8689828a016138a8565b9150509295509295509295565b5f8083601f840112613ba3575f80fd5b5081356001600160401b03811115613bb9575f80fd5b6020830191508360208260051b8501011115613bd3575f80fd5b9250929050565b5f805f805f610360808789031215613bf0575f80fd5b86356001600160401b03811115613c05575f80fd5b613c1189828a01613b93565b9097509550506020870135935060408701359250808701881015613c33575f80fd5b506060860190509295509295909350565b5f8060208385031215613c55575f80fd5b82356001600160401b03811115613c6a575f80fd5b613c7685828601613b93565b90969095509350505050565b5f805f8060808587031215613c95575f80fd5b8435613ca081613adc565b93506020850135613cb081613a48565b92506040850135613cc081613adc565b915060608501356001600160401b03811115613cda575f80fd5b613ce6878288016138a8565b91505092959194509250565b5f60208284031215613d02575f80fd5b5035919050565b5f60208284031215613d19575f80fd5b81356001600160401b03811115613d2e575f80fd5b612148848285016138a8565b5f8083601f840112613d4a575f80fd5b5081356001600160401b03811115613d60575f80fd5b602083019150836020828501011115613bd3575f80fd5b5f805f805f805f60a0888a031215613d8d575f80fd5b87356001600160401b0380821115613da3575f80fd5b818a0191508a601f830112613db6575f80fd5b813581811115613dc4575f80fd5b8b60208260071b8501011115613dd8575f80fd5b60208301995080985050613dee60208b01613911565b9650613dfc60408b01613911565b9550613e0a60608b01613a5c565b945060808a0135915080821115613e1f575f80fd5b50613e2c8a828b01613d3a565b989b979a50959850939692959293505050565b5f805f805f60808688031215613e53575f80fd5b85356001600160401b03811115613e68575f80fd5b613e7488828901613b93565b9096509450506020860135613e88816138fa565b92506040860135613e98816138fa565b91506060860135613ea881613a48565b809150509295509295909350565b5f805f60408486031215613ec8575f80fd5b83356001600160401b03811115613edd575f80fd5b613ee986828701613d3a565b909790965060209590950135949350505050565b5f60208284031215613f0d575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03828116828216039080821115613f4857613f48613f14565b5092915050565b5f82613f6957634e487b7160e01b5f52601260045260245ffd5b500490565b80820180821115613f8157613f81613f14565b92915050565b938452602084019290925260c01b6001600160c01b0319166040830152604882015260680190565b5f6001600160401b03808316818103613fca57613fca613f14565b6001019392505050565b6001600160401b03818116838216019080821115613f4857613f48613f14565b600181811c9082168061400857607f821691505b60208210810361402657634e487b7160e01b5f52602260045260245ffd5b50919050565b5f6020828403121561403c575f80fd5b8151613a1981613a20565b81810381811115613f8157613f81613f14565b9586526020860194909452604085019290925260c01b6001600160c01b0319166060808501919091521b6001600160601b0319166068830152607c820152609c0190565b5f602082840312156140ae575f80fd5b8151613a19816138fa565b601f821115612fbd575f81815260208120601f850160051c810160208610156140df5750805b601f850160051c820191505b818110156140fe578281556001016140eb565b505050505050565b81516001600160401b0381111561411f5761411f613816565b6141338161412d8454613ff4565b846140b9565b602080601f831160018114614166575f841561414f5750858301515b5f19600386901b1c1916600185901b1785556140fe565b5f85815260208120601f198616915b8281101561419457888601518255948401946001909101908401614175565b50858210156141b157878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f6141d360608301866139dc565b6020830194909452506001600160a01b0391909116604090910152919050565b634e487b7160e01b5f52603260045260245ffd5b5f8235607e1983360301811261421b575f80fd5b9190910192915050565b5f613f81368361391c565b5f6001820161424157614241613f14565b5060010190565b6001600160401b03871681526020810186905260408101859052606081018490526001600160a01b03831660808201526103a081016103008360a0840137979650505050505050565b8082028115828204841417613f8157613f81613f14565b63ffffffff87811682526001600160a01b038781166020840152908616604083015284811660608301528316608082015260c060a082018190525f906142f0908301846139dc565b98975050505050505050565b61ffff818116838216019080821115613f4857613f48613f14565b5f60ff60f81b808a60f81b16835261ffff60f01b8960f01b1660018401528751614348816003860160208c016139ba565b606088901b6001600160601b0319166003918501918201528651614373816017840160208b016139ba565b808201915050818660f81b166017820152845191506143998260188301602088016139ba565b016018019998505050505050505050565b60ff60f81b8860f81b1681525f61ffff60f01b808960f01b16600184015287516143db816003860160208c016139ba565b606088901b6001600160601b0319166003918501918201528651614406816017840160208b016139ba565b808201915050818660f01b1660178201528451915061442c8260198301602088016139ba565b016019019998505050505050505050565b5f865161444e818460208b016139ba565b9190910194855250602084019290925260f81b6001600160f81b03199081166040840152166041820152604201919050565b5f60808284031215614490575f80fd5b61449861382a565b823581526020830135602082015260408301356144b4816138fa565b60408201526060928301359281019290925250919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b838152604060208201525f61450d6040830184866144cc565b95945050505050565b818382375f9101908152919050565b8481526001600160a01b03841660208201526060604082018190525f9061454f90830184866144cc565b9695505050505050565b608081525f61456c6080830187896144cc565b602083019590955250604081019290925260609091015292915050565b5f60208284031215614599575f80fd5b81516001600160401b038111156145ae575f80fd5b8201601f810184136145be575f80fd5b80516145cc6138c582613882565b8181528560208385010111156145e0575f80fd5b61450d8260208301602086016139ba565b5f8060408385031215614602575f80fd5b825161460d81613adc565b602084015190925061461e81613a48565b809150509250929050565b5f825161421b8184602087016139ba56fea26469706673582212209ddbe5c07a8895c3828e81eae80ec2c11d03b2e850f0dc4c84a869b3f883316664736f6c63430008140033",
}

// EtrogpolygonzkevmABI is the input ABI used to generate the binding from.
// Deprecated: Use EtrogpolygonzkevmMetaData.ABI instead.
var EtrogpolygonzkevmABI = EtrogpolygonzkevmMetaData.ABI

// EtrogpolygonzkevmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtrogpolygonzkevmMetaData.Bin instead.
var EtrogpolygonzkevmBin = EtrogpolygonzkevmMetaData.Bin

// DeployEtrogpolygonzkevm deploys a new Ethereum contract, binding an instance of Etrogpolygonzkevm to it.
func DeployEtrogpolygonzkevm(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Etrogpolygonzkevm, error) {
	parsed, err := EtrogpolygonzkevmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtrogpolygonzkevmBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Etrogpolygonzkevm{EtrogpolygonzkevmCaller: EtrogpolygonzkevmCaller{contract: contract}, EtrogpolygonzkevmTransactor: EtrogpolygonzkevmTransactor{contract: contract}, EtrogpolygonzkevmFilterer: EtrogpolygonzkevmFilterer{contract: contract}}, nil
}

// Etrogpolygonzkevm is an auto generated Go binding around an Ethereum contract.
type Etrogpolygonzkevm struct {
	EtrogpolygonzkevmCaller     // Read-only binding to the contract
	EtrogpolygonzkevmTransactor // Write-only binding to the contract
	EtrogpolygonzkevmFilterer   // Log filterer for contract events
}

// EtrogpolygonzkevmCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtrogpolygonzkevmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtrogpolygonzkevmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtrogpolygonzkevmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtrogpolygonzkevmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtrogpolygonzkevmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtrogpolygonzkevmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtrogpolygonzkevmSession struct {
	Contract     *Etrogpolygonzkevm // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EtrogpolygonzkevmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtrogpolygonzkevmCallerSession struct {
	Contract *EtrogpolygonzkevmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EtrogpolygonzkevmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtrogpolygonzkevmTransactorSession struct {
	Contract     *EtrogpolygonzkevmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EtrogpolygonzkevmRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtrogpolygonzkevmRaw struct {
	Contract *Etrogpolygonzkevm // Generic contract binding to access the raw methods on
}

// EtrogpolygonzkevmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtrogpolygonzkevmCallerRaw struct {
	Contract *EtrogpolygonzkevmCaller // Generic read-only contract binding to access the raw methods on
}

// EtrogpolygonzkevmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtrogpolygonzkevmTransactorRaw struct {
	Contract *EtrogpolygonzkevmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtrogpolygonzkevm creates a new instance of Etrogpolygonzkevm, bound to a specific deployed contract.
func NewEtrogpolygonzkevm(address common.Address, backend bind.ContractBackend) (*Etrogpolygonzkevm, error) {
	contract, err := bindEtrogpolygonzkevm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Etrogpolygonzkevm{EtrogpolygonzkevmCaller: EtrogpolygonzkevmCaller{contract: contract}, EtrogpolygonzkevmTransactor: EtrogpolygonzkevmTransactor{contract: contract}, EtrogpolygonzkevmFilterer: EtrogpolygonzkevmFilterer{contract: contract}}, nil
}

// NewEtrogpolygonzkevmCaller creates a new read-only instance of Etrogpolygonzkevm, bound to a specific deployed contract.
func NewEtrogpolygonzkevmCaller(address common.Address, caller bind.ContractCaller) (*EtrogpolygonzkevmCaller, error) {
	contract, err := bindEtrogpolygonzkevm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmCaller{contract: contract}, nil
}

// NewEtrogpolygonzkevmTransactor creates a new write-only instance of Etrogpolygonzkevm, bound to a specific deployed contract.
func NewEtrogpolygonzkevmTransactor(address common.Address, transactor bind.ContractTransactor) (*EtrogpolygonzkevmTransactor, error) {
	contract, err := bindEtrogpolygonzkevm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmTransactor{contract: contract}, nil
}

// NewEtrogpolygonzkevmFilterer creates a new log filterer instance of Etrogpolygonzkevm, bound to a specific deployed contract.
func NewEtrogpolygonzkevmFilterer(address common.Address, filterer bind.ContractFilterer) (*EtrogpolygonzkevmFilterer, error) {
	contract, err := bindEtrogpolygonzkevm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmFilterer{contract: contract}, nil
}

// bindEtrogpolygonzkevm binds a generic wrapper to an already deployed contract.
func bindEtrogpolygonzkevm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtrogpolygonzkevmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etrogpolygonzkevm *EtrogpolygonzkevmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etrogpolygonzkevm.Contract.EtrogpolygonzkevmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etrogpolygonzkevm *EtrogpolygonzkevmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.EtrogpolygonzkevmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etrogpolygonzkevm *EtrogpolygonzkevmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.EtrogpolygonzkevmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etrogpolygonzkevm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.GLOBALEXITROOTMANAGERL2(&_Etrogpolygonzkevm.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.GLOBALEXITROOTMANAGERL2(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGEPARAMS(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGEPARAMS(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXCONSTANTBYTES(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXCONSTANTBYTES(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Etrogpolygonzkevm.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Etrogpolygonzkevm.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Etrogpolygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.SIGNATUREINITIALIZETXR(&_Etrogpolygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.SIGNATUREINITIALIZETXR(&_Etrogpolygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.SIGNATUREINITIALIZETXS(&_Etrogpolygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.SIGNATUREINITIALIZETXS(&_Etrogpolygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Etrogpolygonzkevm.Contract.SIGNATUREINITIALIZETXV(&_Etrogpolygonzkevm.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Etrogpolygonzkevm.Contract.SIGNATUREINITIALIZETXV(&_Etrogpolygonzkevm.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Etrogpolygonzkevm.Contract.TIMESTAMPRANGE(&_Etrogpolygonzkevm.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Etrogpolygonzkevm.Contract.TIMESTAMPRANGE(&_Etrogpolygonzkevm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) Admin() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.Admin(&_Etrogpolygonzkevm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) Admin() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.Admin(&_Etrogpolygonzkevm.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) BridgeAddress() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.BridgeAddress(&_Etrogpolygonzkevm.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) BridgeAddress() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.BridgeAddress(&_Etrogpolygonzkevm.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Etrogpolygonzkevm.Contract.CalculatePolPerForceBatch(&_Etrogpolygonzkevm.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Etrogpolygonzkevm.Contract.CalculatePolPerForceBatch(&_Etrogpolygonzkevm.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.DataAvailabilityProtocol(&_Etrogpolygonzkevm.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.DataAvailabilityProtocol(&_Etrogpolygonzkevm.CallOpts)
}

// ExitMode is a free data retrieval call binding the contract method 0x28e18a89.
//
// Solidity: function exitMode() view returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) ExitMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "exitMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExitMode is a free data retrieval call binding the contract method 0x28e18a89.
//
// Solidity: function exitMode() view returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) ExitMode() (bool, error) {
	return _Etrogpolygonzkevm.Contract.ExitMode(&_Etrogpolygonzkevm.CallOpts)
}

// ExitMode is a free data retrieval call binding the contract method 0x28e18a89.
//
// Solidity: function exitMode() view returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) ExitMode() (bool, error) {
	return _Etrogpolygonzkevm.Contract.ExitMode(&_Etrogpolygonzkevm.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) ForceBatchAddress() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.ForceBatchAddress(&_Etrogpolygonzkevm.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.ForceBatchAddress(&_Etrogpolygonzkevm.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) ForceBatchTimeout() (uint64, error) {
	return _Etrogpolygonzkevm.Contract.ForceBatchTimeout(&_Etrogpolygonzkevm.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Etrogpolygonzkevm.Contract.ForceBatchTimeout(&_Etrogpolygonzkevm.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.ForcedBatches(&_Etrogpolygonzkevm.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.ForcedBatches(&_Etrogpolygonzkevm.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) GasTokenAddress() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.GasTokenAddress(&_Etrogpolygonzkevm.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) GasTokenAddress() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.GasTokenAddress(&_Etrogpolygonzkevm.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) GasTokenNetwork() (uint32, error) {
	return _Etrogpolygonzkevm.Contract.GasTokenNetwork(&_Etrogpolygonzkevm.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) GasTokenNetwork() (uint32, error) {
	return _Etrogpolygonzkevm.Contract.GasTokenNetwork(&_Etrogpolygonzkevm.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.GenerateInitializeTransaction(&_Etrogpolygonzkevm.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Etrogpolygonzkevm.Contract.GenerateInitializeTransaction(&_Etrogpolygonzkevm.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) GlobalExitRootManager() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.GlobalExitRootManager(&_Etrogpolygonzkevm.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.GlobalExitRootManager(&_Etrogpolygonzkevm.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Etrogpolygonzkevm.Contract.IsSequenceWithDataAvailabilityAllowed(&_Etrogpolygonzkevm.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Etrogpolygonzkevm.Contract.IsSequenceWithDataAvailabilityAllowed(&_Etrogpolygonzkevm.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) LastAccInputHash() ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.LastAccInputHash(&_Etrogpolygonzkevm.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Etrogpolygonzkevm.Contract.LastAccInputHash(&_Etrogpolygonzkevm.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) LastForceBatch() (uint64, error) {
	return _Etrogpolygonzkevm.Contract.LastForceBatch(&_Etrogpolygonzkevm.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) LastForceBatch() (uint64, error) {
	return _Etrogpolygonzkevm.Contract.LastForceBatch(&_Etrogpolygonzkevm.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) LastForceBatchSequenced() (uint64, error) {
	return _Etrogpolygonzkevm.Contract.LastForceBatchSequenced(&_Etrogpolygonzkevm.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Etrogpolygonzkevm.Contract.LastForceBatchSequenced(&_Etrogpolygonzkevm.CallOpts)
}

// LastRevertModeTimestamp is a free data retrieval call binding the contract method 0xeb5db56a.
//
// Solidity: function lastRevertModeTimestamp() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) LastRevertModeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "lastRevertModeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRevertModeTimestamp is a free data retrieval call binding the contract method 0xeb5db56a.
//
// Solidity: function lastRevertModeTimestamp() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) LastRevertModeTimestamp() (*big.Int, error) {
	return _Etrogpolygonzkevm.Contract.LastRevertModeTimestamp(&_Etrogpolygonzkevm.CallOpts)
}

// LastRevertModeTimestamp is a free data retrieval call binding the contract method 0xeb5db56a.
//
// Solidity: function lastRevertModeTimestamp() view returns(uint256)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) LastRevertModeTimestamp() (*big.Int, error) {
	return _Etrogpolygonzkevm.Contract.LastRevertModeTimestamp(&_Etrogpolygonzkevm.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) NetworkName() (string, error) {
	return _Etrogpolygonzkevm.Contract.NetworkName(&_Etrogpolygonzkevm.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) NetworkName() (string, error) {
	return _Etrogpolygonzkevm.Contract.NetworkName(&_Etrogpolygonzkevm.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) PendingAdmin() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.PendingAdmin(&_Etrogpolygonzkevm.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) PendingAdmin() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.PendingAdmin(&_Etrogpolygonzkevm.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) Pol() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.Pol(&_Etrogpolygonzkevm.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) Pol() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.Pol(&_Etrogpolygonzkevm.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) RollupManager() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.RollupManager(&_Etrogpolygonzkevm.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) RollupManager() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.RollupManager(&_Etrogpolygonzkevm.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) TrustedSequencer() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.TrustedSequencer(&_Etrogpolygonzkevm.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) TrustedSequencer() (common.Address, error) {
	return _Etrogpolygonzkevm.Contract.TrustedSequencer(&_Etrogpolygonzkevm.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Etrogpolygonzkevm.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) TrustedSequencerURL() (string, error) {
	return _Etrogpolygonzkevm.Contract.TrustedSequencerURL(&_Etrogpolygonzkevm.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmCallerSession) TrustedSequencerURL() (string, error) {
	return _Etrogpolygonzkevm.Contract.TrustedSequencerURL(&_Etrogpolygonzkevm.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.AcceptAdminRole(&_Etrogpolygonzkevm.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.AcceptAdminRole(&_Etrogpolygonzkevm.TransactOpts)
}

// EnterRevertMode is a paid mutator transaction binding the contract method 0x03744a63.
//
// Solidity: function enterRevertMode((bytes,bytes32,uint64,bytes32) forcedBatchData) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) EnterRevertMode(opts *bind.TransactOpts, forcedBatchData PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "enterRevertMode", forcedBatchData)
}

// EnterRevertMode is a paid mutator transaction binding the contract method 0x03744a63.
//
// Solidity: function enterRevertMode((bytes,bytes32,uint64,bytes32) forcedBatchData) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) EnterRevertMode(forcedBatchData PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.EnterRevertMode(&_Etrogpolygonzkevm.TransactOpts, forcedBatchData)
}

// EnterRevertMode is a paid mutator transaction binding the contract method 0x03744a63.
//
// Solidity: function enterRevertMode((bytes,bytes32,uint64,bytes32) forcedBatchData) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) EnterRevertMode(forcedBatchData PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.EnterRevertMode(&_Etrogpolygonzkevm.TransactOpts, forcedBatchData)
}

// ExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x94aa909f.
//
// Solidity: function executeAllForcedTransactions((bytes,bytes32,uint64,bytes32)[] batches, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) ExecuteAllForcedTransactions(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "executeAllForcedTransactions", batches, newLocalExitRoot, newStateRoot, proof)
}

// ExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x94aa909f.
//
// Solidity: function executeAllForcedTransactions((bytes,bytes32,uint64,bytes32)[] batches, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) ExecuteAllForcedTransactions(batches []PolygonRollupBaseEtrogBatchData, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.ExecuteAllForcedTransactions(&_Etrogpolygonzkevm.TransactOpts, batches, newLocalExitRoot, newStateRoot, proof)
}

// ExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x94aa909f.
//
// Solidity: function executeAllForcedTransactions((bytes,bytes32,uint64,bytes32)[] batches, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) ExecuteAllForcedTransactions(batches []PolygonRollupBaseEtrogBatchData, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.ExecuteAllForcedTransactions(&_Etrogpolygonzkevm.TransactOpts, batches, newLocalExitRoot, newStateRoot, proof)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.ForceBatch(&_Etrogpolygonzkevm.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.ForceBatch(&_Etrogpolygonzkevm.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.Initialize(&_Etrogpolygonzkevm.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.Initialize(&_Etrogpolygonzkevm.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// IsRevertModeActive is a paid mutator transaction binding the contract method 0x788aaa6a.
//
// Solidity: function isRevertModeActive() returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) IsRevertModeActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "isRevertModeActive")
}

// IsRevertModeActive is a paid mutator transaction binding the contract method 0x788aaa6a.
//
// Solidity: function isRevertModeActive() returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) IsRevertModeActive() (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.IsRevertModeActive(&_Etrogpolygonzkevm.TransactOpts)
}

// IsRevertModeActive is a paid mutator transaction binding the contract method 0x788aaa6a.
//
// Solidity: function isRevertModeActive() returns(bool)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) IsRevertModeActive() (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.IsRevertModeActive(&_Etrogpolygonzkevm.TransactOpts)
}

// OnRevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xb151da2d.
//
// Solidity: function onRevertLastVerifiedBatch(bytes32 _prevAccInputHash) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) OnRevertLastVerifiedBatch(opts *bind.TransactOpts, _prevAccInputHash [32]byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "onRevertLastVerifiedBatch", _prevAccInputHash)
}

// OnRevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xb151da2d.
//
// Solidity: function onRevertLastVerifiedBatch(bytes32 _prevAccInputHash) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) OnRevertLastVerifiedBatch(_prevAccInputHash [32]byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.OnRevertLastVerifiedBatch(&_Etrogpolygonzkevm.TransactOpts, _prevAccInputHash)
}

// OnRevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xb151da2d.
//
// Solidity: function onRevertLastVerifiedBatch(bytes32 _prevAccInputHash) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) OnRevertLastVerifiedBatch(_prevAccInputHash [32]byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.OnRevertLastVerifiedBatch(&_Etrogpolygonzkevm.TransactOpts, _prevAccInputHash)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.OnVerifyBatches(&_Etrogpolygonzkevm.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.OnVerifyBatches(&_Etrogpolygonzkevm.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "sequenceBatches", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SequenceBatches(&_Etrogpolygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SequenceBatches(&_Etrogpolygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "sequenceBatchesValidium", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SequenceBatchesValidium(&_Etrogpolygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SequenceBatchesValidium(&_Etrogpolygonzkevm.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SequenceForceBatches(&_Etrogpolygonzkevm.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SequenceForceBatches(&_Etrogpolygonzkevm.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetDataAvailabilityProtocol(&_Etrogpolygonzkevm.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetDataAvailabilityProtocol(&_Etrogpolygonzkevm.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetForceBatchAddress(&_Etrogpolygonzkevm.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetForceBatchAddress(&_Etrogpolygonzkevm.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetForceBatchTimeout(&_Etrogpolygonzkevm.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetForceBatchTimeout(&_Etrogpolygonzkevm.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetTrustedSequencer(&_Etrogpolygonzkevm.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetTrustedSequencer(&_Etrogpolygonzkevm.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetTrustedSequencerURL(&_Etrogpolygonzkevm.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SetTrustedSequencerURL(&_Etrogpolygonzkevm.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SwitchSequenceWithDataAvailability(&_Etrogpolygonzkevm.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.SwitchSequenceWithDataAvailability(&_Etrogpolygonzkevm.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.TransferAdminRole(&_Etrogpolygonzkevm.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Etrogpolygonzkevm.Contract.TransferAdminRole(&_Etrogpolygonzkevm.TransactOpts, newPendingAdmin)
}

// EtrogpolygonzkevmAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmAcceptAdminRoleIterator struct {
	Event *EtrogpolygonzkevmAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmAcceptAdminRole)
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
		it.Event = new(EtrogpolygonzkevmAcceptAdminRole)
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
func (it *EtrogpolygonzkevmAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmAcceptAdminRole represents a AcceptAdminRole event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*EtrogpolygonzkevmAcceptAdminRoleIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmAcceptAdminRoleIterator{contract: _Etrogpolygonzkevm.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmAcceptAdminRole)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseAcceptAdminRole(log types.Log) (*EtrogpolygonzkevmAcceptAdminRole, error) {
	event := new(EtrogpolygonzkevmAcceptAdminRole)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmActivateRevertModeIterator is returned from FilterActivateRevertMode and is used to iterate over the raw logs and unpacked data for ActivateRevertMode events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmActivateRevertModeIterator struct {
	Event *EtrogpolygonzkevmActivateRevertMode // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmActivateRevertModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmActivateRevertMode)
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
		it.Event = new(EtrogpolygonzkevmActivateRevertMode)
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
func (it *EtrogpolygonzkevmActivateRevertModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmActivateRevertModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmActivateRevertMode represents a ActivateRevertMode event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmActivateRevertMode struct {
	Sender          common.Address
	LastForcedBatch uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterActivateRevertMode is a free log retrieval operation binding the contract event 0xa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9.
//
// Solidity: event ActivateRevertMode(address _sender, uint64 _lastForcedBatch)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterActivateRevertMode(opts *bind.FilterOpts) (*EtrogpolygonzkevmActivateRevertModeIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "ActivateRevertMode")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmActivateRevertModeIterator{contract: _Etrogpolygonzkevm.contract, event: "ActivateRevertMode", logs: logs, sub: sub}, nil
}

// WatchActivateRevertMode is a free log subscription operation binding the contract event 0xa4331640971235cc6212bce0ce97a6a53e440af92617e4121cb6ffbc874ec6c9.
//
// Solidity: event ActivateRevertMode(address _sender, uint64 _lastForcedBatch)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchActivateRevertMode(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmActivateRevertMode) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "ActivateRevertMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmActivateRevertMode)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "ActivateRevertMode", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseActivateRevertMode(log types.Log) (*EtrogpolygonzkevmActivateRevertMode, error) {
	event := new(EtrogpolygonzkevmActivateRevertMode)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "ActivateRevertMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmForceBatchIterator struct {
	Event *EtrogpolygonzkevmForceBatch // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmForceBatch)
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
		it.Event = new(EtrogpolygonzkevmForceBatch)
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
func (it *EtrogpolygonzkevmForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmForceBatch represents a ForceBatch event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*EtrogpolygonzkevmForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmForceBatchIterator{contract: _Etrogpolygonzkevm.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmForceBatch)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseForceBatch(log types.Log) (*EtrogpolygonzkevmForceBatch, error) {
	event := new(EtrogpolygonzkevmForceBatch)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmInitialSequenceBatchesIterator struct {
	Event *EtrogpolygonzkevmInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmInitialSequenceBatches)
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
		it.Event = new(EtrogpolygonzkevmInitialSequenceBatches)
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
func (it *EtrogpolygonzkevmInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmInitialSequenceBatches represents a InitialSequenceBatches event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*EtrogpolygonzkevmInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmInitialSequenceBatchesIterator{contract: _Etrogpolygonzkevm.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmInitialSequenceBatches)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseInitialSequenceBatches(log types.Log) (*EtrogpolygonzkevmInitialSequenceBatches, error) {
	event := new(EtrogpolygonzkevmInitialSequenceBatches)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmInitializedIterator struct {
	Event *EtrogpolygonzkevmInitialized // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmInitialized)
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
		it.Event = new(EtrogpolygonzkevmInitialized)
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
func (it *EtrogpolygonzkevmInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmInitialized represents a Initialized event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterInitialized(opts *bind.FilterOpts) (*EtrogpolygonzkevmInitializedIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmInitializedIterator{contract: _Etrogpolygonzkevm.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmInitialized) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmInitialized)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseInitialized(log types.Log) (*EtrogpolygonzkevmInitialized, error) {
	event := new(EtrogpolygonzkevmInitialized)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmLogForcedBatchDataIterator is returned from FilterLogForcedBatchData and is used to iterate over the raw logs and unpacked data for LogForcedBatchData events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmLogForcedBatchDataIterator struct {
	Event *EtrogpolygonzkevmLogForcedBatchData // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmLogForcedBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmLogForcedBatchData)
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
		it.Event = new(EtrogpolygonzkevmLogForcedBatchData)
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
func (it *EtrogpolygonzkevmLogForcedBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmLogForcedBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmLogForcedBatchData represents a LogForcedBatchData event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmLogForcedBatchData struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Timestamp          *big.Int
	PrentBlockHash     [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogForcedBatchData is a free log retrieval operation binding the contract event 0x5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e9989879378.
//
// Solidity: event LogForcedBatchData(bytes _transactions, bytes32 _lastGlobalExitRoot, uint256 _timestamp, bytes32 _prentBlockHash)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterLogForcedBatchData(opts *bind.FilterOpts) (*EtrogpolygonzkevmLogForcedBatchDataIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "LogForcedBatchData")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmLogForcedBatchDataIterator{contract: _Etrogpolygonzkevm.contract, event: "LogForcedBatchData", logs: logs, sub: sub}, nil
}

// WatchLogForcedBatchData is a free log subscription operation binding the contract event 0x5e644fcc5facab2dccd802043da91b15d76aeec7a6c4886a48e66e9989879378.
//
// Solidity: event LogForcedBatchData(bytes _transactions, bytes32 _lastGlobalExitRoot, uint256 _timestamp, bytes32 _prentBlockHash)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchLogForcedBatchData(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmLogForcedBatchData) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "LogForcedBatchData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmLogForcedBatchData)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "LogForcedBatchData", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseLogForcedBatchData(log types.Log) (*EtrogpolygonzkevmLogForcedBatchData, error) {
	event := new(EtrogpolygonzkevmLogForcedBatchData)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "LogForcedBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSequenceBatchesIterator struct {
	Event *EtrogpolygonzkevmSequenceBatches // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSequenceBatches)
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
		it.Event = new(EtrogpolygonzkevmSequenceBatches)
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
func (it *EtrogpolygonzkevmSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSequenceBatches represents a SequenceBatches event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*EtrogpolygonzkevmSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSequenceBatchesIterator{contract: _Etrogpolygonzkevm.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSequenceBatches)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSequenceBatches(log types.Log) (*EtrogpolygonzkevmSequenceBatches, error) {
	event := new(EtrogpolygonzkevmSequenceBatches)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSequenceForceBatchesIterator struct {
	Event *EtrogpolygonzkevmSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSequenceForceBatches)
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
		it.Event = new(EtrogpolygonzkevmSequenceForceBatches)
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
func (it *EtrogpolygonzkevmSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSequenceForceBatches represents a SequenceForceBatches event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*EtrogpolygonzkevmSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSequenceForceBatchesIterator{contract: _Etrogpolygonzkevm.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSequenceForceBatches)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSequenceForceBatches(log types.Log) (*EtrogpolygonzkevmSequenceForceBatches, error) {
	event := new(EtrogpolygonzkevmSequenceForceBatches)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetDataAvailabilityProtocolIterator struct {
	Event *EtrogpolygonzkevmSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSetDataAvailabilityProtocol)
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
		it.Event = new(EtrogpolygonzkevmSetDataAvailabilityProtocol)
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
func (it *EtrogpolygonzkevmSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*EtrogpolygonzkevmSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSetDataAvailabilityProtocolIterator{contract: _Etrogpolygonzkevm.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSetDataAvailabilityProtocol)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*EtrogpolygonzkevmSetDataAvailabilityProtocol, error) {
	event := new(EtrogpolygonzkevmSetDataAvailabilityProtocol)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetForceBatchAddressIterator struct {
	Event *EtrogpolygonzkevmSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSetForceBatchAddress)
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
		it.Event = new(EtrogpolygonzkevmSetForceBatchAddress)
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
func (it *EtrogpolygonzkevmSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSetForceBatchAddress represents a SetForceBatchAddress event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*EtrogpolygonzkevmSetForceBatchAddressIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSetForceBatchAddressIterator{contract: _Etrogpolygonzkevm.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSetForceBatchAddress)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSetForceBatchAddress(log types.Log) (*EtrogpolygonzkevmSetForceBatchAddress, error) {
	event := new(EtrogpolygonzkevmSetForceBatchAddress)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetForceBatchTimeoutIterator struct {
	Event *EtrogpolygonzkevmSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSetForceBatchTimeout)
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
		it.Event = new(EtrogpolygonzkevmSetForceBatchTimeout)
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
func (it *EtrogpolygonzkevmSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*EtrogpolygonzkevmSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSetForceBatchTimeoutIterator{contract: _Etrogpolygonzkevm.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSetForceBatchTimeout)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSetForceBatchTimeout(log types.Log) (*EtrogpolygonzkevmSetForceBatchTimeout, error) {
	event := new(EtrogpolygonzkevmSetForceBatchTimeout)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetTrustedSequencerIterator struct {
	Event *EtrogpolygonzkevmSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSetTrustedSequencer)
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
		it.Event = new(EtrogpolygonzkevmSetTrustedSequencer)
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
func (it *EtrogpolygonzkevmSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSetTrustedSequencer represents a SetTrustedSequencer event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*EtrogpolygonzkevmSetTrustedSequencerIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSetTrustedSequencerIterator{contract: _Etrogpolygonzkevm.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSetTrustedSequencer)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSetTrustedSequencer(log types.Log) (*EtrogpolygonzkevmSetTrustedSequencer, error) {
	event := new(EtrogpolygonzkevmSetTrustedSequencer)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetTrustedSequencerURLIterator struct {
	Event *EtrogpolygonzkevmSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSetTrustedSequencerURL)
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
		it.Event = new(EtrogpolygonzkevmSetTrustedSequencerURL)
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
func (it *EtrogpolygonzkevmSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*EtrogpolygonzkevmSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSetTrustedSequencerURLIterator{contract: _Etrogpolygonzkevm.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSetTrustedSequencerURL)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSetTrustedSequencerURL(log types.Log) (*EtrogpolygonzkevmSetTrustedSequencerURL, error) {
	event := new(EtrogpolygonzkevmSetTrustedSequencerURL)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator struct {
	Event *EtrogpolygonzkevmSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmSwitchSequenceWithDataAvailability)
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
		it.Event = new(EtrogpolygonzkevmSwitchSequenceWithDataAvailability)
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
func (it *EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmSwitchSequenceWithDataAvailabilityIterator{contract: _Etrogpolygonzkevm.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmSwitchSequenceWithDataAvailability)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*EtrogpolygonzkevmSwitchSequenceWithDataAvailability, error) {
	event := new(EtrogpolygonzkevmSwitchSequenceWithDataAvailability)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmTransferAdminRoleIterator struct {
	Event *EtrogpolygonzkevmTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmTransferAdminRole)
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
		it.Event = new(EtrogpolygonzkevmTransferAdminRole)
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
func (it *EtrogpolygonzkevmTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmTransferAdminRole represents a TransferAdminRole event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*EtrogpolygonzkevmTransferAdminRoleIterator, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmTransferAdminRoleIterator{contract: _Etrogpolygonzkevm.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmTransferAdminRole)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseTransferAdminRole(log types.Log) (*EtrogpolygonzkevmTransferAdminRole, error) {
	event := new(EtrogpolygonzkevmTransferAdminRole)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtrogpolygonzkevmVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmVerifyBatchesIterator struct {
	Event *EtrogpolygonzkevmVerifyBatches // Event containing the contract specifics and raw log

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
func (it *EtrogpolygonzkevmVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtrogpolygonzkevmVerifyBatches)
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
		it.Event = new(EtrogpolygonzkevmVerifyBatches)
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
func (it *EtrogpolygonzkevmVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtrogpolygonzkevmVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtrogpolygonzkevmVerifyBatches represents a VerifyBatches event raised by the Etrogpolygonzkevm contract.
type EtrogpolygonzkevmVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*EtrogpolygonzkevmVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &EtrogpolygonzkevmVerifyBatchesIterator{contract: _Etrogpolygonzkevm.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *EtrogpolygonzkevmVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Etrogpolygonzkevm.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtrogpolygonzkevmVerifyBatches)
				if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Etrogpolygonzkevm *EtrogpolygonzkevmFilterer) ParseVerifyBatches(log types.Log) (*EtrogpolygonzkevmVerifyBatches, error) {
	event := new(EtrogpolygonzkevmVerifyBatches)
	if err := _Etrogpolygonzkevm.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
