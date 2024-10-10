// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupmanager

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

// LegacyZKEVMStateVariablesPendingState is an auto generated low-level Go binding around an user-defined struct.
type LegacyZKEVMStateVariablesPendingState struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}

// LegacyZKEVMStateVariablesSequencedBatchData is an auto generated low-level Go binding around an user-defined struct.
type LegacyZKEVMStateVariablesSequencedBatchData struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}

// PolygonRollupManagerStateRootData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupManagerStateRootData struct {
	Timestamp         uint64
	InitVerifiedBatch uint64
	LastLocalExitRoot [32]byte
	OldStateRoot      [32]byte
}

// PolygonrollupmanagerMetaData contains all meta data concerning the Polygonrollupmanager contract.
var PolygonrollupmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingLocalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRevertBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameRevertStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerifiedBeyondFinalityPeriod\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"AllowToRevertBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_lastLocalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"LastLocalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_prevlastVerifiedBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"}],\"name\":\"RevertLastVerifiedBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FINALITY_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBaseRollback\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupPendingStateTransitions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structLegacyZKEVMStateVariables.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"verifiedBatch\",\"type\":\"uint64\"}],\"name\":\"getRollupVerifiedBatchData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupManager.StateRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustedAggregatorTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"},{\"internalType\":\"contractPolygonZkEVMExistentEtrog\",\"name\":\"polygonZkEVM\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"zkEVMVerifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMForkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMChainID\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAllowedToRevertBatches\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRevertBatchesExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"onExecuteAllForcedTransactions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertLastVerifiedBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"internalType\":\"contractIPolygonRollupBaseRollback\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIdToStateRootData\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nextLocalExitRootBatch\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setAllowedToRevertBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_rollupID\",\"type\":\"uint32\"}],\"name\":\"setLastLocalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setRevertBatchesExecuted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e0346200017e57601f6200617c38819003918201601f19168301916001600160401b0383118484101762000182578084926060946040528339810103126200017e5780516001600160a01b03919082811681036200017e5760208201519183831683036200017e576040015192831683036200017e5760805260c05260a0525f5460ff8160081c16620001295760ff80821610620000ee575b604051615fe59081620001978239608051818181610fd2015261569c015260a051818181610dc50152818161144d0152615028015260c051818181610a650152818161273a0152818161487701526154080152f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f62000099565b60405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080604052600436101562000012575f80fd5b5f803560e01c80630645af091462003382578063066ec0121462003359578063080b311114620033095780630a0d9fbe14620032dd57806311f6b287146200329257806312b86e1914620031b75780631489ed1014620030dc57806315064c9614620030b75780631608859c14620030145780631796a1ae1462002feb5780631816b7e51462002f3e5780632072f6c51462002e6c578063248a9ca31462002e3d578063252801691462002d89578063280a47741462002cf95780632f2ff15d1462002c5457806330490fdd1462002bf057806330c27dde1462002bc757806336568abe1462002b78578063394218e91462002ac357806341ec66dc14620023e3578063477fa27014620023c357806355a71ee0146200236857806360469169146200231e57806365c0504d14620022975780637222020f14620021be578063727885e91462001d935780637975fcfe1462001d1a5780637c52fc1d1462001cfb5780637fb6e76a1462001cb95780637fea08641462001c94578063841b24d71462001c7157806387c20c0114620018635780638bd4f07114620017ca57806391d14854146200177b57806399f5634e146200175c5780639a908e7314620016055780639c9f3dfe146200155f578063a066215c14620014b6578063a217fddf1462001498578063a2967d991462001471578063a3c573eb146200142b578063afd23cbe1462001404578063b40a15c314620013dc578063b99d0ad7146200133d578063c1acbc341462001311578063c4c928c21462001038578063ceee281d1462000ff6578063d02103ca1462000fb0578063d5073f6f1462000ef6578063d547741f1462000eb0578063d939b3151462000e87578063dbc169761462000d55578063dbebee191462000d10578063dde0ff771462000ce4578063df6058c01462000cb9578063e0bfd3d21462000a89578063e46761c41462000a43578063e8d2b07314620006fb578063ea1a7dde1462000659578063f34eb8eb1462000407578063f4e9267514620003e15763f9c4c2ae146200030f575f80fd5b34620003de576020366003190112620003de5760406101809163ffffffff6200033762004311565b16815260826020522060ff81549160018101546005820154916007600682015491015492604051956001600160a01b039384821688526001600160401b039485809360a01c1660208a01528116604089015260a01c166060870152608086015281811660a0860152818160401c1660c0860152818160801c1660e086015260c01c6101008501528082166101208501528160401c1661014084015260801c16610160820152f35b80fd5b5034620003de5780600319360112620003de57602063ffffffff60815416604051908152f35b5034620003de5760c0366003190112620003de576200042562004270565b6200042f620042b5565b906200043a620042e3565b906064359160ff8316809303620006555760843560a435936001600160401b0394858111620006515762000473903690600401620044db565b335f9081527fb3124b64d0942b276d75ea451f4dd8b83bf9442b427f9eb873ac771c92e65f02602052604090205490939060ff16156200063f57607f549663ffffffff9765ffffffff0000620004ce8a8360101c1662004599565b998a169960101b169065ffffffff0000191617607f556040519060c0820190828210898311176200062b5789916040528a6001600160a01b0380991698898552806020860193169a8b84528060408701971696878152606087019289845280608089019580875260a08a01988d8a528152608060205260409020985116906bffffffffffffffffffffffff60a01b91828a541617895560018901965116908654161785555116620005a19084908154906001600160401b0360a01b9060a01b16906001600160401b0360a01b1916179055565b519082549060ff60e81b9051151560e81b169160ff60e01b9060e01b169061ffff60e01b1916171790555190600201556040519586958652602086015260408501526060840152608083015260a0820160c0905260c0820162000604916200454a565b037fa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b5291a280f35b634e487b7160e01b5f52604160045260245ffd5b604051637615be1f60e11b8152600490fd5b8780fd5b5f80fd5b5034620003de576020366003190112620003de57600435801515808203620006f757338352608360205263ffffffff60408420541615620006e557607f805460ff191660ff929092169190911790556040805191151582523360208301527f4bd13f042de7fee85a224d6bbb9efaab00a2e510a67c51f907c48701821927019190819081015b0390a180f35b6040516371653c1560e01b8152600490fd5b8280fd5b5034620003de5780600319360112620003de573381526020906083825263ffffffff6040822054169182825260828152604082209060ff607f5416158015620009a8575b62000996576006820180546001600160401b039586808360401c16921690808752607e855260408720835f5285526200077b60405f20620055fb565b61025889825116018981116200098257894291161062000970578501516003870196949089169290815b8a81169085821115620007d257505f5287875260405f208960018201818d825460401c16935555620007a5565b8a9150968b8095979a96999816956001600160401b03199680888654161785555f52885260405f20546001600160a01b0382541690813b156200096c57839160248392604051948593849263b151da2d60e01b845260048401525af18015620009615762000932575b50620008fe620008cf8b98958760409c9b99966200091799967fc2ac6a39960d89d285c7e09e8ba997a82779fa1d2d673b6ca813cdbf35309ddc9f876002916200089167ffffffffffffffff60401b9a62004643565b5f52018c52808f5f2055898152607e8c528e81208d5f528c5260028f5f208281558260018201550155620008c68d8262004571565b541690620045e6565b9383620008e260855496828816620045e6565b169384620008f18c8c620045e6565b928716178c1c16620045e6565b891b16916001600160801b03191617176085556200563a565b61010061ffff19607f541617607f558351928352820152a180f35b9162000947819a9997949b9895969362004419565b6200095d578b999295979891949396996200083b565b8880fd5b6040513d85823e3d90fd5b8380fd5b604051632d5f6a9560e01b8152600490fd5b634e487b7160e01b5f52601160045260245ffd5b604051631f50ffb960e31b8152600490fd5b5082816001600160a01b038454166004604051809481936375aedab560e11b83525af190811562000a3857849162000a05575b5061012c8101809111620009f15742116200073f565b634e487b7160e01b84526011600452602484fd5b90508181813d831162000a30575b62000a1f81836200449d565b810103126200065557515f620009db565b503d62000a13565b6040513d86823e3d90fd5b5034620003de5780600319360112620003de5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034620003de5760c0366003190112620003de576004356001600160a01b03808216809203620006f75762000abd620042b5565b62000ac7620042e3565b606435936001600160401b039081861690818703620006555760a4359260ff8416938481036200065557335f9081527ff444741efd45fa05f38650e87f3d8ca01816575bed45a01e78186ae5f3bc8355602090815260409091205490999060ff16156200063f57848b5260848a5263ffffffff978860408d20541662000ca757848c5260838b528860408d20541662000c95576081549862000b6b818b1662004599565b63ffffffff199a8b1691169081176081555f87815260848d52604080822080548d168417905587825260838e528082208054909c168317909b5581815260828d529990992080546001600160a01b03199081168717825560018201805467ffffffffffffffff60a01b191660a08c901b67ffffffffffffffff60a01b16178155919b919290919083549092169116179055885467ffffffffffffffff60a01b191660a09190911b67ffffffffffffffff60a01b1617885560078801805460ff60801b191660809290921b60ff60801b16919091179055604051941684528684015260408301526060820152608081015f905260a07fadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f85091a28280526002019052608435604082205580f35b604051630d409b9360e41b8152600490fd5b6040516337c8fe0960e11b8152600490fd5b5034620003de576020366003190112620003de5762000ce162000cdb62004311565b6200563a565b80f35b5034620003de5780600319360112620003de5760206001600160401b0360855460401c16604051908152f35b5034620003de576020366003190112620003de576001600160401b036002604060209363ffffffff62000d4262004311565b168152607e855220015416604051908152f35b5034620003de5780600319360112620003de57335f9081527f307dd781c86f141a460373e0117a8033afea21c02f6121aff0f5b5553df24f4a602052604090205460ff16156200063f576001600160401b0342166001600160401b03196088541617608855806001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1562000e8457818091600460405180948193636de0b4bb60e11b83525af1801562000e795762000e61575b5050606f5460ff81161562000e4f5760ff1916606f557f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b38180a180f35b604051635386698160e01b8152600490fd5b62000e6c9062004419565b620003de57805f62000e12565b6040513d84823e3d90fd5b50fd5b5034620003de5780600319360112620003de5760206001600160401b0360865416604051908152f35b5034620003de576040366003190112620003de5762000ce160043562000ed5620042b5565b90808452603460205262000ef0600160408620015462005855565b62005876565b5034620003de576020366003190112620003de57335f9081527fe0205af17734f857fb538b40523c7cc189a1ab8f1ec33be516761c6b465766b660205260409020546004359060ff16156200063f57683635c9adc5dea000008111801562000fa2575b62000f90576020817ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b292608755604051908152a180f35b604051638586952560e01b8152600490fd5b50633b9aca00811062000f59565b5034620003de5780600319360112620003de5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034620003de576020366003190112620003de5763ffffffff60406020926001600160a01b036200102662004270565b16815260838452205416604051908152f35b5034620003de576060366003190112620003de576004356001600160a01b03808216809203620006f75760243563ffffffff80821680920362000655576001600160401b03926044358481116200130d57366023820112156200130d57806004013590858211620006515736602483830101116200065157335f9081527ff14f5a8ad59d90593602e905b358229bff5ceea677d5bf0f5a1701793550a9a6602090815260409091205490949060ff16156200063f5785158015620012fc575b620012ea5787895260838552604089205416968715620012d857878952608285526040892093600785019788549088818360401c1614620012c657888c526080885260408c20916001830190815490600160ff8360e81c16151514620012b45760ff808360e01c169160801c1603620012a25782620011ba928f9a600188910193166bffffffffffffffffffffffff60a01b8454161783555460a01c168154906001600160401b0360a01b9060a01b16906001600160401b0360a01b1916179055565b620011c6898b62004571565b8a875260828852620011db604088206200537f565b1698896001600160401b0319825416179055541690803b156200129e578492836064602494876040519889978896879463278f794360e11b8652600486015260408286015282604486015201848401378181018301849052601f01601f191681010301925af1801562000e795762001282575b5050907ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d926040928351928352820152a280f35b620012909093929362004419565b6200129e5790845f6200124e565b8480fd5b604051635aa0d5f160e11b8152600490fd5b604051633b8d3d9960e01b8152600490fd5b604051634f61d51960e01b8152600490fd5b6040516374a086a360e01b8152600490fd5b604051637512e5cb60e01b8152600490fd5b5080607f5460101c168611620010f7565b8680fd5b5034620003de5780600319360112620003de5760206001600160401b0360855460801c16604051908152f35b5034620003de576040366003190112620003de57620013a360406080926200136462004311565b63ffffffff62001373620042cc565b916200137e620055d5565b5016825260826020526001600160401b036004848420019116825260205220620055fb565b620013da6040518092606080916001600160401b038082511685526020820151166020850152604081015160408501520151910152565bf35b5034620003de5780600319360112620003de57602060ff607f5460081c166040519015158152f35b5034620003de5780600319360112620003de57602061ffff60865460801c16604051908152f35b5034620003de5780600319360112620003de5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034620003de5780600319360112620003de576020620014906200514e565b604051908152f35b5034620003de5780600319360112620003de57602090604051908152f35b5034620003de576020366003190112620003de57620014d4620042fa565b620014de6200581c565b6001600160401b038116906201518082116200154d57608680546fffffffffffffffff00000000000000001916604092831b67ffffffffffffffff60401b16179055519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602090a180f35b604051631c0cfbfd60e31b8152600490fd5b5034620003de576020366003190112620003de576200157d620042fa565b620015876200581c565b60ff606f541615620015dd575b60206001600160401b037fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c759216806001600160401b03196086541617608655604051908152a180f35b6001600160401b03806086541690821610620015945760405163048a05a960e41b8152600490fd5b5034620003de576040366003190112620003de5762001623620042fa565b60ff606f54166200174a573382526020916083835263ffffffff604082205416908115620006e5576001600160401b03808416156200173857620017016040838562001707955260828852209160855490806200168388828516620045ca565b16966001600160401b031997888094161760855560068501805483620016ad8183169485620045ca565b16998a911617905560405192620016c4846200442d565b60243584528260018b8601958242168752604081019485528b5f52600389018d5260405f2090518155019451169084541617835551169062004571565b62004bfa565b7f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a2583604051848152a2604051908152f35b604051632590ccf960e01b8152600490fd5b604051630bc011ff60e21b8152600490fd5b5034620003de5780600319360112620003de57602062001490620053ed565b5034620003de576040366003190112620003de576001600160a01b036040620017a3620042b5565b92600435815260346020522091165f52602052602060ff60405f2054166040519015158152f35b5034620003de57620017dc3662004325565b969592919360ff606f96929654166200174a577f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a701097866004966001600160401b03956200183f948660408f63ffffffff829f168152608260205220998a62004db4565b16865201602052600282852001549082519182526020820152a162000ce16200501e565b5034620003de5762001875366200439b565b9060ff606f989798969394959654166200174a5763ffffffff87168952608260205260408920976001600160401b0387165f5260038901602052620018d16001600160401b03600160405f2001541660855460c01c90620045ca565b6001600160401b034291161162001c5f576103e86001600160401b03620018f9868a620045e6565b161162001c4d578587620019109587938c620046d4565b6200191b856200537f565b94839587906001600160401b03620019348288620045e6565b169060865492620019526001600160401b038560401c164262004ff1565b995b6001600160401b0381166001600160401b038416811462001c36575f5260038601602052600160405f200154906001600160401b0382168c105f14620019a7575060401c6001600160401b031662001954565b620019d49495969798999a9b50620019cb92506001600160401b03939150620045e6565b16809262004ff1565b918183101562001bdb575003600c81111562001bd0575062001a0e600c5b6087548161ffff60865460801c160a02906103e80a9062004fff565b6087555b608754683635c9adc5dea00000908181111562001baf57506087555b6086546001600160401b031662001acd579062001a958262001a768660067faac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b497960162004571565b6001600160401b0386165f52600281016020528260405f205562004643565b62001aa0856200563a565b604080516001600160401b039590951685526020850191909152830152339263ffffffff1691606090a380f35b62001ad88162004bfa565b600681019081546001600160401b038160801c166001600160401b03811462001b9b5767ffffffffffffffff60801b19909116600190910160801b67ffffffffffffffff60801b1617918290556040517faac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b494939262001b95926004916001600160401b039062001b688562004449565b8142168552818a16602086015287604086015286606086015260801c168a52016020526040882062004600565b62001aa0565b634e487b7160e01b89526011600452602489fd5b9050633b9aca0080911062001bc6575b5062001a2e565b6087555f62001bbf565b62001a0e90620019f2565b9062001c219203600c81115f1462001c2a575062001c16600c915b670de0b6b3a76400009261ffff816103e80a9260801c160a830262004fff565b906087540262004fff565b60875562001a12565b62001c16909162001bf6565b505092939495969798505081620019d49162004ff1565b604051635acfba9d60e11b8152600490fd5b604051638a0704d360e01b8152600490fd5b5034620003de5780600319360112620003de57602060855460c01c604051908152f35b5034620003de5780600319360112620003de57602060ff607f54166040519015158152f35b5034620003de576020366003190112620003de5763ffffffff60406020926001600160401b0362001ce9620042fa565b16815260848452205416604051908152f35b5034620003de5780600319360112620003de5760206040516102588152f35b5034620003de5760c0366003190112620003de5762001d7a62001d8f9162001d4162004311565b9062001d4c620042cc565b9063ffffffff62001d5c620042e3565b93168152608260205260a435926084359260406064359320620054af565b6040519182916020835260208301906200454a565b0390f35b5034620003de5760e0366003190112620003de5762001db162004311565b9062001dbc620042cc565b91604435906001600160a01b0382168203620006555762001ddc62004287565b62001de66200429e565b9460a4356001600160401b038111620021ba5762001e09903690600401620044db565b9060c4356001600160401b0381116200130d5762001e2c903690600401620044db565b335f9081527f41902c3bd358ff4b85ee46864bd81ee453a1b3cfd703a58950be8d28ae030661602052604090205490919060ff16156200063f5763ffffffff85161580156200219f575b620012ea5763ffffffff8516875260806020526040872092600160ff8186015460e81c16151514620012b4576001600160401b0382168852608460205263ffffffff60408920541662000ca7576081549163ffffffff62001ed981851662004599565b16809363ffffffff1916176081556001600160a01b03855416996040519a62001f028c62004465565b8a8c5260405190816106128101106001600160401b03610612840111176200218b5762001f548c9d83926106126200595e8539610612840190815230602082015260606040820181905201906200454a565b03908bf095861562002180576001600160a01b037f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641926001600160401b0385620020bc8f60ff8f916080988f63ffffffff8f9d62002099928f956001968c6007941682526084602052604082208186198254161790558e831682526083602052604082208186198254161790558152608260205260408120918e6bffffffffffffffffffffffff60a01b9116818454161783558c8888015460a01c1690620020418985019283908154906001600160401b0360a01b9060a01b16906001600160401b0360a01b1916179055565b8f8989015416908254161790556200207b8b83908154906001600160401b0360a01b9060a01b16906001600160401b0360a01b1916179055565b60406002870154918080526002840160205220550195168562004571565b0154825460ff60801b191660e09190911c9190911660801b60ff60801b16179055565b63ffffffff6040519d168d528c6020858d169101521660408c01521698896060820152a26001600160a01b0384163b1562000651576001600160a01b03966200215988946200214686948c9a8b976040519d8e9c8d9b8c99633892b81160e11b8b521660048a01521660248801526044870152606486015260c0608486015260c48501906200454a565b8381036003190160a4850152906200454a565b0393165af1801562000e79576200216d5750f35b620021789062004419565b620003de5780f35b6040513d8c823e3d90fd5b634e487b7160e01b8c52604160045260248cfd5b5063ffffffff607f5460101c1663ffffffff86161162001e76565b8580fd5b5034620003de576020366003190112620003de57620021dc62004311565b335f9081527fd26c1f3d72948c9f5ca2b28d010847694e18f81d42613c046515d0839b64354c602052604090205460ff16156200063f5763ffffffff80911690811590811562002285575b50620012ea578082526080602052600160408320018054600160ff8260e81c16151514620012b45760ff60e81b1916600160e81b1790557f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e448280a280f35b9050607f5460101c1681115f62002227565b5034620003de576020366003190112620003de57604060c09163ffffffff620022bf62004311565b1681526080602052206001600160a01b0360ff81835416926002600182015491015492604051948552811660208501526001600160401b038160a01c166040850152818160e01c16606085015260e81c161515608083015260a0820152f35b5034620003de5780600319360112620003de57608754906064820291808304606414901517156200235457602082604051908152f35b634e487b7160e01b81526011600452602490fd5b5034620003de576040366003190112620003de576001600160401b03600260406200239262004311565b9363ffffffff620023a2620042cc565b951681526082602052200191165f52602052602060405f2054604051908152f35b5034620003de5780600319360112620003de576020608754604051908152f35b5034620003de576103a0366003190112620003de5762002402620042fa565b6200240c6200429e565b90366103a411620006f75760ff606f54166200174a57338352608360205263ffffffff604084205416928315620006e5576001600160401b038216156200173857808460409252608260205220926085546001600160401b036200247384828416620045ca565b16906001600160401b03191617608555600684015492620025106001600160401b03620024a385828816620045ca565b1694856001600160401b031982161760068801556001600160401b0360405191620024ce836200442d565b6024358352816001602085019482421686528260408201941684528a5f5260038c0160205260405f209051815501935116821984541617835551169062004571565b817f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a256020604051878152a2620025468562004643565b6200255d62002555866200537f565b9384620045ca565b926103e86001600160401b03620025758387620045e6565b161162001c4d5762002587866200537f565b6001600160401b036007880154166001600160401b0383161062002ab1576001600160401b0382165f526002870160205260405f205490811562002a9f576001600160401b0381166001600160401b0384161162002a8d576001600160401b0381166001600160401b038716111562002a7b5760205f6200262a62002616606435868d8c8a60443592620054af565b836040519282848094519384920162004527565b8101039060025afa156200298b577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000015f51066001600160a01b0360018a01541660405191620026798362004465565b8252604051634890ed4560e11b815291829061030060a460048401375f61030483015b6001821062002a6057505050610324816020935afa9081156200298b575f9162002a2a575b501562002a1857620026d8620027cd9187620045e6565b93620026f8620026e7620053ed565b6001600160401b03871690620046c0565b90604051916001600160a01b03602084019263a9059cbb60e01b84521660248401526044830152604482526200272e8262004449565b5f806001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692604051946200276a8662004481565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1562002a0f573d91620027ae83620044bf565b92620027be60405194856200449d565b83523d5f602085013e62004b57565b8051908162002996575b50506001600160a01b038754165f5260836020526200285763ffffffff60405f20541691604051906200280a8262004449565b6001600160401b03421682526001600160401b038516602083015260443560408301526060820152825f52607e60205260405f206001600160401b0388165f5260205260405f2062004600565b5f52607e6020526001600160401b03600160405f200191165f5260205260405f206001600160401b0384166001600160401b031982541617905567ffffffffffffffff60401b620028b8608554926001600160401b038460401c16620045ca565b60401b166001600160401b0360801b4260801b169177ffffffffffffffffffffffffffffffff0000000000000000191617176085556001600160a01b03845416803b1562000655575f80916064604051809481936332c2d15360e01b83526001600160401b0389166004840152833560248401523360448401525af19485156200298b576001600160401b03620014909460029360209862002979575b5062002965816006850162004571565b165f5201845260643560405f20556200563a565b620029849062004419565b5f62002955565b6040513d5f823e3d90fd5b602080620029a99383010191016200466b565b15620029b7575f80620027d7565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b60609162004b57565b6040516309bde33960e01b8152600490fd5b62002a51915060203d60201162002a58575b62002a4881836200449d565b8101906200466b565b5f620026c1565b503d62002a3c565b8293506020809160019394518152019301910184926200269c565b60405163b9b18f5760e01b8152600490fd5b604051630f2b74f160e11b8152600490fd5b6040516324cbdcc360e11b8152600490fd5b60405163ead1340b60e01b8152600490fd5b5034620003de576020366003190112620003de5762002ae1620042fa565b62002aeb6200581c565b60ff606f54161562002b50575b60207f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1916085546001600160c01b036001600160401b0360c01b8360c01b169116176085556001600160401b0360405191168152a180f35b60855460c01c6001600160401b0382161062002af85760405163401636df60e01b8152600490fd5b5034620003de576040366003190112620003de5762002b96620042b5565b336001600160a01b0382160362002bb55762000ce19060043562005876565b604051630b4ad1cd60e31b8152600490fd5b5034620003de5780600319360112620003de5760206001600160401b0360885416604051908152f35b5034620003de576040366003190112620003de576001600160401b03604062002c1862004311565b9263ffffffff62002c28620042cc565b9462002c33620055d5565b50168152607e6020522091165f526020526080620013a360405f20620055fb565b5034620003de576040366003190112620003de5760043562002c75620042b5565b818352603460205262002c8f600160408520015462005855565b81835260346020526001600160a01b0360408420911690815f5260205260ff60405f2054161562002cbe578280f35b818352603460205260408320815f5260205260405f20600160ff1982541617905533915f8051602062005f708339815191528480a45f808280f35b5034620003de576020366003190112620003de57600435801515808203620006f757338352608360205263ffffffff60408420541615620006e557607f805461ff00191660089290921b61ff00169190911790556040805191151582523360208301527f4bd13f042de7fee85a224d6bbb9efaab00a2e510a67c51f907c4870182192701919081908101620006df565b5034620003de576040366003190112620003de576060906003604062002dae62004311565b9263ffffffff62002dbe620042cc565b948284805162002dce816200442d565b8281528260208201520152168152608260205220016001600160401b038092165f5260205260405f209060405162002e06816200442d565b81600184549485845201549181604060208301928286168452019360401c1683526040519485525116602084015251166040820152f35b5034620003de576020366003190112620003de5760016040602092600435815260348452200154604051908152f35b5034620003de5780600319360112620003de577f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e8152603460205260408120335f5260205260ff60405f2054161562002eca575b62000ce16200501e565b6001600160401b038060855460801c1690811591821562002f24575b50811562002f08575b501562002ec05760405163692baaad60e11b8152600490fd5b905062002f198160885416620045af565b429116115f62002eef565b62002f31919250620045af565b8142911611905f62002ee6565b5034620003de576020366003190112620003de5760043561ffff811690818103620006f75762002f6d6200581c565b6103e88210801562002fdf575b62002fcd576086805461ffff60801b191660809290921b61ffff60801b169190911790556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c590602090a180f35b604051630984a67960e31b8152600490fd5b506103ff821162002f7a565b5034620003de5780600319360112620003de57602063ffffffff607f5460101c16604051908152f35b5034620003de576040366003190112620003de576200303262004311565b63ffffffff62003041620042cc565b91168252608260205260408220905f8051602062005f908339815191528352603460205260408320335f5260205260ff60405f2054161562003089575b62000ce19162004ca4565b60ff606f54166200174a57620030a08183620053ba565b6200307e57604051630674f25160e11b8152600490fd5b5034620003de5780600319360112620003de57602060ff606f54166040519015158152f35b5034620003de57620030ee366200439b565b62003101979394959297969196620057e3565b63ffffffff841696878a52608260205288868860408d2096620031259688620046d4565b62003134846006830162004571565b6001600160401b0384165f52600281016020528560405f2055620031589062004643565b62003163906200563a565b60405191829133956200319092846040919493926001600160401b03606083019616825260208201520152565b037fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d391a380f35b5034620003de57620031c93662004325565b620031db9793929596919497620057e3565b84868463ffffffff87169a8b8d52608260205260408d209a620031ff968c62004db4565b6200320e816006870162004571565b6001600160401b031693845f52600281016020528260405f2055620032339062004643565b6200323e906200563a565b61127560c71b6085546001600160c01b0316176085556040519283526020830152604082015233606082015260807f3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e91a280f35b5034620003de576020366003190112620003de57620032cc604060209263ffffffff620032be62004311565b16815260828452206200537f565b6001600160401b0360405191168152f35b5034620003de5780600319360112620003de5760206001600160401b0360865460401c16604051908152f35b5034620003de576040366003190112620003de576200334f60209160406200333062004311565b9163ffffffff62003340620042cc565b931681526082855220620053ba565b6040519015158152f35b5034620003de5780600319360112620003de5760206001600160401b0360855416604051908152f35b5034620003de57610140366003190112620003de57620033a162004270565b620033ab620042cc565b620033b5620042e3565b91620033c062004287565b620033ca6200429e565b60a435926001600160a01b03841684036200130d5760c435956001600160a01b0387168703620006515760e435956001600160a01b03871687036200095d576001600160401b03610104351661010435036200095d576001600160401b03610124351661012435036200095d5788549160ff8360081c16158062004262575b156200420e5761010260ff9361ffff19161791828b55608654916001600160c01b03608554916001600160401b0360c01b9060c01b1691161760855567016345785d8a000060875569070800000000000000006101f560811b926001600160401b0361ffff60801b199316906001600160801b0319161717161760865560081c1615620041b5575f8051602062005f90833981519152875260346020526001600160a01b036040882091169081885260205260ff604088205416156200415c575b5085805260346020526001600160a01b036040872091169081875260205260ff6040872054161562004120575b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59080875260346020526040872082885260205260ff60408820541615620040e3575b507f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e80875260346020526040872082885260205260ff60408820541615620040a6575b507f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac9081875260346020526040872081885260205260ff6040882054161562004069575b50507fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd8086526034602052604086206001600160a01b038316875260205260ff6040872054161562004023575b507fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd088086526034602052604086206001600160a01b038316875260205260ff6040872054161562003fdd575b507f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f48086526034602052604086206001600160a01b038316875260205260ff6040872054161562003f97575b507fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db18086526034602052604086206001600160a01b038316875260205260ff6040872054161562003f51575b507f736419170ed02bfd51e6789f87c97c1528041947ae9835536f00b50fe51a982e8086526034602052604086206001600160a01b038316875260205260ff6040872054161562003f0b575b505f8051602062005f908339815191528552603460205260016040862001908154917f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0809155807fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff935f8051602062005f90833981519152858a80a4808752604087206001600160a01b038316885260205260ff6040882054161562003ec5575b507f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb8087526034602052604087206001600160a01b038316885260205260ff6040882054161562003e75575b50507f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e808652603460205260408620600181018054837f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff28595195868094558a80a46001600160a01b038416875260205260ff6040872054161562003e2f575b508085526034602052604085206001600160a01b038316865260205260ff6040862054161562003ddf575b50506001600160401b0360735460401c166001600160401b03607454169081810362003dcd5762003aef8260066081549563ffffffff6200395581891662004599565b63ffffffff199889169116908117608155610124356001600160401b03168a52608460209081526040808c2080548b16841790556001600160a01b038b8116808e5260838452828e208054909c168517909b55838d5260829092528b2080546001600160a01b0319908116909a17815560018101805467ffffffffffffffff60a01b19166101043560a01b67ffffffffffffffff60a01b161781559299909390919083549092169116179055805467ffffffffffffffff60a01b19166101243560a01b67ffffffffffffffff60a01b16178155600781019660ff60801b1988541688557fadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f85060a06040516001600160401b03610104351681526001600160a01b038c1660208201526001600160401b03610124351660408201528c6060820152866080820152a2828952607560205260408920546002820160205260408a20558489526072602052604089206003820160205260408a209080820362003d96575b5050016001600160401b0319938482541617815562004571565b8254161790556001600160a01b03607a5416906001600160a01b03606f5460581c1660726020526040842054906001600160a01b0383163b156200129e5760405193635d6717a560e01b85526004850152602484015260a0604484015283906076548060011c9060018116801562003d8b575b60208310811462003d77578260a488015290815f1462003d58575060011462003cfe575b508382036003190160648501526077548592600182811c92908116801562003cf3575b60208410811462003cdf57918391888a80999795829997855292835f1462003cb15750505060011462003c47575b50506001600160a01b03848093879360848301520393165af1801562000e795762003c35575b5061ff001981541681557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160028152a180f35b62003c409062004419565b5f62003bfd565b607787529192508591907f7901cb5addcae2d210a531c604a76a660d77039093bac314de0816a16392aff15b81841062003c9157505001602001906001600160a01b038462003bd7565b805460208486018101919091528a98508997509093019260010162003c73565b60ff191660208581019190915294151560051b90930190930194506001600160a01b0392915062003bd79050565b634e487b7160e01b89526022600452602489fd5b92607f169262003ba9565b60768652909150847fb5732705f5241370a28908c2fe1303cb223f03b90d857fd0573f003f79fefed45b82821062003d40575060c49150840101905f62003b86565b6001816020925460c4858a0101520191019062003d28565b905060c492935060ff191682860152151560051b840101905f62003b86565b634e487b7160e01b88526022600452602488fd5b91607f169162003b62565b6001600160401b0360018062003dc594845481550192018181541682198454161783555460401c169062004571565b5f8062003ad5565b604051632e4cc54360e11b8152600490fd5b8085526034602052604085206001600160a01b038316865260205260408520600160ff198254161790556001600160a01b03339216905f8051602062005f708339815191528680a45f8062003912565b8086526034602090815260408088206001600160a01b038616808a5292528720805460ff1916600117905533915f8051602062005f708339815191528880a45f620038e7565b8087526034602052604087206001600160a01b038316885260205260408720600160ff198254161790556001600160a01b03339216905f8051602062005f708339815191528880a45f806200386a565b8087526034602090815260408089206001600160a01b038516808b5292528820805460ff1916600117905533915f8051602062005f708339815191528980a45f6200381e565b8086526034602090815260408088206001600160a01b038516808a5292528720805460ff1916600117905533915f8051602062005f708339815191528880a45f6200377d565b8086526034602090815260408088206001600160a01b038516808a5292528720805460ff1916600117905533915f8051602062005f708339815191528880a45f62003731565b8086526034602090815260408088206001600160a01b038516808a5292528720805460ff1916600117905533915f8051602062005f708339815191528880a45f620036e5565b8086526034602090815260408088206001600160a01b038516808a5292528720805460ff1916600117905533915f8051602062005f708339815191528880a45f62003699565b8086526034602090815260408088206001600160a01b038516808a5292528720805460ff1916600117905533915f8051602062005f708339815191528880a45f6200364d565b81875260346020526040872081885260205260408720600160ff1982541617905533915f8051602062005f708339815191528880a45f8062003600565b80875260346020526040872082885260205260408720600160ff198254161790558133915f8051602062005f708339815191528980a45f620035bc565b80875260346020526040872082885260205260408720600160ff198254161790558133915f8051602062005f708339815191528980a45f62003579565b85805260346020526040862081875260205260408620600160ff198254161790553381875f8051602062005f708339815191528180a462003537565b5f8051602062005f90833981519152875260346020526040872081885260205260408720600160ff1982541617905533905f8051602062005f908339815191525f8051602062005f708339815191528980a45f6200350a565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b62461bcd60e51b608090815260206084908152602e60a4527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160c4526d191e481a5b9a5d1a585b1a5e995960921b60e45290fd5b50600260ff84161062003449565b600435906001600160a01b03821682036200065557565b606435906001600160a01b03821682036200065557565b608435906001600160a01b03821682036200065557565b602435906001600160a01b03821682036200065557565b602435906001600160401b03821682036200065557565b604435906001600160401b03821682036200065557565b600435906001600160401b03821682036200065557565b6004359063ffffffff821682036200065557565b6103e0600319820112620006555760043563ffffffff811681036200065557916001600160401b03602435818116810362000655579260443582811681036200065557926064358381168103620006555792608435908116810362000655579160a4359160c435916103e411620006555760e490565b6103e0600319820112620006555760043563ffffffff811681036200065557916001600160401b03906024358281168103620006555792604435838116810362000655579260643590811681036200065557916084359160a4359160c4356001600160a01b03811681036200065557916103e411620006555760e490565b6001600160401b0381116200062b57604052565b606081019081106001600160401b038211176200062b57604052565b608081019081106001600160401b038211176200062b57604052565b602081019081106001600160401b038211176200062b57604052565b604081019081106001600160401b038211176200062b57604052565b90601f801991011681019081106001600160401b038211176200062b57604052565b6001600160401b0381116200062b57601f01601f191660200190565b81601f820112156200065557803590620044f582620044bf565b926200450560405194856200449d565b828452602083830101116200065557815f926020809301838601378301015290565b5f5b838110620045395750505f910152565b818101518382015260200162004529565b90602091620045658151809281855285808601910162004527565b601f01601f1916010190565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b63ffffffff809116908114620009825760010190565b9062093a806001600160401b03809316019182116200098257565b9190916001600160401b03808094169116019182116200098257565b6001600160401b0391821690821603919082116200098257565b906060600291620046326001600160401b03808351166001600160401b03198754161786556020830151168562004571565b604081015160018501550151910155565b60060180546001600160401b038160801c166200465e575050565b6001600160801b03169055565b908160209103126200065557518015158103620006555790565b9291906103208401936103008092823701905f915b60018310620046a857505050565b6001908251815260208091019201920191906200469a565b818102929181159184041417156200098257565b9293919495909796975f92620046ea856200537f565b906001600160401b03948560078801541693868a1694851062002ab157861690811562004b22575085600688015460801c16811162004b10575f526004860160205260405f2091856002840154935460401c16840362004afe575b8588169786821689111562002a7b5760205f620047836200476e8e88859f878f918b92620054af565b60409d8e519282848094519384920162004527565b8101039060025afa1562004af457620047ff60208b9c9d9e60019b9c7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000015f5106918c6001600160a01b039e8f9101541691805193620047e28562004465565b845251809581948293634890ed4560e11b84526004840162004685565b03915afa90811562004aea575f9162004ac6575b501562004ab557620048f8916200482a91620045e6565b948a620048446200483a620053ed565b89891690620046c0565b918151928b602085019263a9059cbb60e01b8452166024850152604484015260448352620048728362004449565b5f808c7f00000000000000000000000000000000000000000000000000000000000000001692845195620048a68762004481565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d1562004aab573d620027be620048ee82620044bf565b945194856200449d565b80518062004a32575b50509167ffffffffffffffff60401b93916200496d620049a194898954165f52608360205263ffffffff8c5f205416928c51916200493f8362004449565b89421683528560208401528d8301526060820152825f52607e6020528b5f208b5f526020528b5f2062004600565b5f52607e6020526001895f2001905f52602052875f20876001600160401b03198254161790556085549384891c16620045ca565b77ffffffffffffffffffffffffffffffff00000000000000001990921691861b16174260801b67ffffffffffffffff60801b1617608555541692833b15620006555760645f9283855196879485936332c2d15360e01b8552600485015260248401523360448401525af190811562004a29575062004a1c5750565b62004a279062004419565b565b513d5f823e3d90fd5b9060208062004a469383010191016200466b565b1562004a54575f8062004901565b885162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b6060925062004b57565b8a516309bde33960e01b8152600490fd5b62004ae3915060203d60201162002a585762002a4881836200449d565b5f62004813565b8c513d5f823e3d90fd5b89513d5f823e3d90fd5b604051632bd2e3e760e01b8152600490fd5b60405163bb14c20560e01b8152600490fd5b6040915084815260028801602052205491821562002a9f578581168411156200474557604051630f2b74f160e11b8152600490fd5b9192901562004bbc575081511562004b6d575090565b3b1562004b775790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501562004bd05750805190602001fd5b60405162461bcd60e51b81526020600482015290819062004bf69060248301906200454a565b0390fd5b60068101546001600160401b0390818160801c169060c01c80821162004c21575b50505050565b600101918211620009825762004c388284620053ba565b62004c45575b8062004c1b565b677fffffffffffffff62004c5e8362004c6993620045e6565b60011c1682620045ca565b9062004c768284620053ba565b1562004c91575062004c889162004ca4565b5f808062004c3e565b905062004c9e9162004ca4565b62004c88565b60068101908154916001600160401b039283851690848160c01c83119182159262004da4575b505062004d92577f581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b936001600160a01b03600163ffffffff6080965f868152600482016020526040948592838320978854851c169962004d2b8b8262004571565b600289019d8e548c865260028401602052868620556001600160c01b038254916001600160401b0360c01b9060c01b1691161790555416815260836020522054169762004d78896200563a565b5492015491815194855260208501528301526060820152a2565b60405163d086b70b60e01b8152600490fd5b60801c1682119050845f62004cca565b9192939496905f976001600160401b039889600786015416928a881693841062002ab1578a1690811562004fb7575089600686015460801c16811162004b1057805f526004850160205260405f20928a6002850154945460401c160362004afe575b8960068601549616958a8160801c16871191821562004fab575b50811562004f9c575b5062004f8a576004840196855f5260209688885260409a808c5f20548d1c169083160362004f79578988945f9462004e8a9462004e77948a620054af565b838b519282848094519384920162004527565b8101039060025afa1562004f6f5762004efe9184916001600160a01b0360017f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000015f5106930154169089519262004ee08462004465565b83528951809581948293634890ed4560e11b84526004840162004685565b03915afa90811562004f6f575f9162004f4d575b501562004f3c575f52526002825f2001541462004f2c5750565b5163a47276bd60e01b8152600490fd5b84516309bde33960e01b8152600490fd5b62004f689150833d851162002a585762002a4881836200449d565b5f62004f12565b86513d5f823e3d90fd5b8a516332a2a77f60e01b8152600490fd5b60405163bfa7079f60e01b8152600490fd5b905060c01c8511155f62004e39565b87111591505f62004e30565b80846040925260028701602052205492831562002a9f578a600687015460401c16101562004e1657604051630f2b74f160e11b8152600490fd5b919082039182116200098257565b81156200500a570490565b634e487b7160e01b5f52601260045260245ffd5b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1562000655576040518091632072f6c560e01b82528160045f948580945af1801562000e7957620050b4575b50606f549060ff82166200174a5760017f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a54979260ff191617606f5580a1565b620050bf9062004419565b5f62005076565b6001600160401b0381116200062b5760051b60200190565b90620050ea82620050c6565b620050f960405191826200449d565b82815280926200510c601f1991620050c6565b0190602036910137565b5f198114620009825760010190565b80518210156200513a5760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b63ffffffff8060815416801562005379576200516a81620050de565b915f5b82811062005344575050905f6020915b600190818514620052cd578185811c9516918286018096116200098257620051a586620050de565b915f915f1994858901898111945b8a8110620051f657505050505050939160408051906020820192808452818301528152620051e1816200442d565b5190209280156200098257019190926200517d565b85620009825781811480620052c3575b156200526a5780841b908082046002148115171562000982576200522f62005264928762005125565b516040805190602082019283528c8183015281526200524e816200442d565b5190206200525d828a62005125565b5262005116565b620051b3565b80841b9080820460021481151715620009825762005289828762005125565b518583018093116200098257620052a562005264938862005125565b516040805191602083019384528183015281526200524e816200442d565b5083831462005206565b919350508051156200513a5760200151915f905b828210620052ef5750505090565b6200533d90604094855160208101918252828782015286815262005313816200442d565b51902094805190602082019280845281830152815262005333816200442d565b5190209162005116565b90620052e1565b60018101908181116200098257826200537392165f526082602052600560405f2001546200525d828762005125565b6200516d565b50505f90565b60068101546001600160401b0391608082901c83168015620053b157600492505f520160205260405f205460401c1690565b505060401c1690565b6004906001600160401b038093165f5201602052620053e58160405f205416826086541690620045ca565b429116111590565b6040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156200298b575f9162005478575b50608554620054626001600160401b039182808260401c169116620045e6565b1680156200537957620054759162004fff565b90565b906020823d8211620054a6575b8162005494602093836200449d565b81010312620003de5750515f62005442565b3d915062005485565b92939094916001600160401b03948587165f9481865260038701602052604096878720549989861688528888205493151580620055cc575b620055bb578315620055aa57620054fe83620058ed565b1562005599579060749160018254920154918a519c8d993360601b60208c015260348b015260548a01526001600160401b0360c01b98899485809460c01b1691015260201b16607c8c015260201b1660848a0152608c89015260ac88015260cc87015260c01b1660ec85015260d48452610100840192848410908411176200558557505290565b634e487b7160e01b81526041600452602490fd5b88516305dae44f60e21b8152600490fd5b88516366385b5160e01b8152600490fd5b885163340c614f60e11b8152600490fd5b508a15620054e7565b60405190620055e48262004449565b5f6060838281528260208201528260408201520152565b906040516200560a8162004449565b60606002829480546001600160401b0390818116865260401c166020850152600181015460408501520154910152565b63ffffffff16905f9180835260209060828252604091828520607e825283862060028101966001600160401b039182895416986001998a80840191808652828952868b87205416808752858a526005888d892054169901986001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016935b620056fe575b5050505050505050507f728d0f7d9b5ea5d935d30470f2ebf8f73e4c87870780094fffb000337b1a1fa494955054908351928352820152a1565b610258810180821162001b9b5742119081620057d8575b5015620057d157871690818752858a528d8c882001548955620057376200514e565b833b1562000651578c51906333d6247d60e01b82526004820152878160248183885af18015620057c757918f95949392918a92620057b5575b506001600160401b03199283885416178755168752828a528b87209081541690558684541690818752878c88205416808852868b52888d8920541690919394620056be565b620057c09062004419565b5f62005770565b8d513d8a823e3d90fd5b80620056c4565b905015155f62005715565b335f9081527fc17b14a573f65366cdad721c7c0a0f76536bb4a86b935cdac44610e4f010b52a602052604090205460ff16156200063f57565b335f9081527f3b72b60ab7bcf7b6e5f79e30b28eb2a39d66a2a278905c2bfcfdd4e1c0a1907b602052604090205460ff16156200063f57565b5f52603460205260405f20335f5260205260ff60405f205416156200063f57565b905f9180835260346020526001600160a01b036040842092169182845260205260ff604084205416620058a857505050565b8083526034602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4565b67ffffffff000000016001600160401b0391818382161092836200594c575b836200593a575b50826200592c575b5050156200592857600190565b5f90565b60c01c1090505f806200591b565b818392945060801c1610915f62005913565b925081838260401c1610926200590c56fe604060a081526106128038038061001581610234565b92833981016060828203126102305761002d8261026d565b9060209161003c83850161026d565b8585015190946001600160401b038211610230570182601f820112156102305780519061007061006b83610281565b610234565b938285528583830101116102305784905f5b83811061021c5750505f9184010152803b156101fb577f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b03199081166001600160a01b038481169182179093558751929691948794929390917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28151156101da5750905f80838861014195519101845af43d156101d2573d9161013361006b84610281565b9283523d5f8985013e61029c565b505b8060805216917f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f857fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103958587549483519286168352820152a182156101bb5716179055516103129081610300823960805181600f0152f35b8451633173bdd160e11b81525f6004820152602490fd5b60609161029c565b9293505050346101ec57508390610143565b63b398979f60e01b8152600490fd5b8451634c9c8ce360e01b81526001600160a01b039091166004820152602490fd5b818101830151868201840152869201610082565b5f80fd5b6040519190601f01601f191682016001600160401b0381118382101761025957604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361023057565b6001600160401b03811161025957601f01601f191660200190565b906102c357508051156102b157805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806102f6575b6102d4575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156102cc56fe60806040526001600160a01b03337f00000000000000000000000000000000000000000000000000000000000000008216036100df575f9081356001600160e01b03191663278f794360e11b141582036100df57366004116100db5760403660031901126100db576004359081168091036100db576024359067ffffffffffffffff82116100d757366023830112156100d75781600401356100a86100a382610167565b61013c565b9281845236602483830101116100d357816100d195926024602093018387013784010152610183565b005b8480fd5b8280fd5b5080fd5b6001600160a01b037f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54165f808092368280378136915af43d82803e15610124573d90f35b3d90fd5b634e487b7160e01b5f52604160045260245ffd5b6040519190601f01601f1916820167ffffffffffffffff81118382101761016257604052565b610128565b67ffffffffffffffff811161016257601f01601f191660200190565b90813b1561021f576001600160a01b0382167f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc816bffffffffffffffffffffffff60a01b8254161790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156102045761020191610240565b50565b50503461020d57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b5f8061027293602081519101845af43d15610275573d916102636100a384610167565b9283523d5f602085013e610279565b90565b6060915b906102a0575080511561028e57805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806102d3575b6102b1575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156102a956fea26469706673582212206d39c8873bc04eb5abedc786a08c96eb91def62cbd719acb6ff60ade219421f664736f6c634300081400332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4a2646970667358221220635d1d65cf617052c441804f344e66727686078f61f820880b3ba8908de7555564736f6c63430008140033",
}

// PolygonrollupmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagerMetaData.ABI instead.
var PolygonrollupmanagerABI = PolygonrollupmanagerMetaData.ABI

// PolygonrollupmanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagerMetaData.Bin instead.
var PolygonrollupmanagerBin = PolygonrollupmanagerMetaData.Bin

// DeployPolygonrollupmanager deploys a new Ethereum contract, binding an instance of Polygonrollupmanager to it.
func DeployPolygonrollupmanager(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonrollupmanager, error) {
	parsed, err := PolygonrollupmanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagerBin), backend, _globalExitRootManager, _pol, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupmanager{PolygonrollupmanagerCaller: PolygonrollupmanagerCaller{contract: contract}, PolygonrollupmanagerTransactor: PolygonrollupmanagerTransactor{contract: contract}, PolygonrollupmanagerFilterer: PolygonrollupmanagerFilterer{contract: contract}}, nil
}

// Polygonrollupmanager is an auto generated Go binding around an Ethereum contract.
type Polygonrollupmanager struct {
	PolygonrollupmanagerCaller     // Read-only binding to the contract
	PolygonrollupmanagerTransactor // Write-only binding to the contract
	PolygonrollupmanagerFilterer   // Log filterer for contract events
}

// PolygonrollupmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupmanagerSession struct {
	Contract     *Polygonrollupmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PolygonrollupmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupmanagerCallerSession struct {
	Contract *PolygonrollupmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PolygonrollupmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupmanagerTransactorSession struct {
	Contract     *PolygonrollupmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonrollupmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupmanagerRaw struct {
	Contract *Polygonrollupmanager // Generic contract binding to access the raw methods on
}

// PolygonrollupmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupmanagerCallerRaw struct {
	Contract *PolygonrollupmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupmanagerTransactorRaw struct {
	Contract *PolygonrollupmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupmanager creates a new instance of Polygonrollupmanager, bound to a specific deployed contract.
func NewPolygonrollupmanager(address common.Address, backend bind.ContractBackend) (*Polygonrollupmanager, error) {
	contract, err := bindPolygonrollupmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanager{PolygonrollupmanagerCaller: PolygonrollupmanagerCaller{contract: contract}, PolygonrollupmanagerTransactor: PolygonrollupmanagerTransactor{contract: contract}, PolygonrollupmanagerFilterer: PolygonrollupmanagerFilterer{contract: contract}}, nil
}

// NewPolygonrollupmanagerCaller creates a new read-only instance of Polygonrollupmanager, bound to a specific deployed contract.
func NewPolygonrollupmanagerCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupmanagerCaller, error) {
	contract, err := bindPolygonrollupmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerCaller{contract: contract}, nil
}

// NewPolygonrollupmanagerTransactor creates a new write-only instance of Polygonrollupmanager, bound to a specific deployed contract.
func NewPolygonrollupmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupmanagerTransactor, error) {
	contract, err := bindPolygonrollupmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerTransactor{contract: contract}, nil
}

// NewPolygonrollupmanagerFilterer creates a new log filterer instance of Polygonrollupmanager, bound to a specific deployed contract.
func NewPolygonrollupmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupmanagerFilterer, error) {
	contract, err := bindPolygonrollupmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerFilterer{contract: contract}, nil
}

// bindPolygonrollupmanager binds a generic wrapper to an already deployed contract.
func bindPolygonrollupmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanager *PolygonrollupmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanager.Contract.PolygonrollupmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanager *PolygonrollupmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.PolygonrollupmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanager *PolygonrollupmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.PolygonrollupmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanager *PolygonrollupmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanager.Contract.DEFAULTADMINROLE(&_Polygonrollupmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanager.Contract.DEFAULTADMINROLE(&_Polygonrollupmanager.CallOpts)
}

// FINALITYPERIOD is a free data retrieval call binding the contract method 0x7c52fc1d.
//
// Solidity: function FINALITY_PERIOD() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) FINALITYPERIOD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "FINALITY_PERIOD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FINALITYPERIOD is a free data retrieval call binding the contract method 0x7c52fc1d.
//
// Solidity: function FINALITY_PERIOD() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) FINALITYPERIOD() (uint64, error) {
	return _Polygonrollupmanager.Contract.FINALITYPERIOD(&_Polygonrollupmanager.CallOpts)
}

// FINALITYPERIOD is a free data retrieval call binding the contract method 0x7c52fc1d.
//
// Solidity: function FINALITY_PERIOD() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) FINALITYPERIOD() (uint64, error) {
	return _Polygonrollupmanager.Contract.FINALITYPERIOD(&_Polygonrollupmanager.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanager.Contract.BridgeAddress(&_Polygonrollupmanager.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanager.Contract.BridgeAddress(&_Polygonrollupmanager.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanager.Contract.CalculateRewardPerBatch(&_Polygonrollupmanager.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanager.Contract.CalculateRewardPerBatch(&_Polygonrollupmanager.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanager.Contract.ChainIDToRollupID(&_Polygonrollupmanager.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanager.Contract.ChainIDToRollupID(&_Polygonrollupmanager.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanager.Contract.GetBatchFee(&_Polygonrollupmanager.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanager.Contract.GetBatchFee(&_Polygonrollupmanager.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanager.Contract.GetForcedBatchFee(&_Polygonrollupmanager.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanager.Contract.GetForcedBatchFee(&_Polygonrollupmanager.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanager.Contract.GetInputSnarkBytes(&_Polygonrollupmanager.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanager.Contract.GetInputSnarkBytes(&_Polygonrollupmanager.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanager.Contract.GetLastVerifiedBatch(&_Polygonrollupmanager.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanager.Contract.GetLastVerifiedBatch(&_Polygonrollupmanager.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanager.Contract.GetRoleAdmin(&_Polygonrollupmanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanager.Contract.GetRoleAdmin(&_Polygonrollupmanager.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanager.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanager.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanager.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanager.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanager.Contract.GetRollupExitRoot(&_Polygonrollupmanager.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanager.Contract.GetRollupExitRoot(&_Polygonrollupmanager.CallOpts)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetRollupPendingStateTransitions(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getRollupPendingStateTransitions", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesPendingState)).(*LegacyZKEVMStateVariablesPendingState)

	return out0, err

}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanager.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanager.CallOpts, rollupID, batchNum)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanager.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanager.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanager.Contract.GetRollupSequencedBatches(&_Polygonrollupmanager.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanager.Contract.GetRollupSequencedBatches(&_Polygonrollupmanager.CallOpts, rollupID, batchNum)
}

// GetRollupVerifiedBatchData is a free data retrieval call binding the contract method 0x30490fdd.
//
// Solidity: function getRollupVerifiedBatchData(uint32 rollupID, uint64 verifiedBatch) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GetRollupVerifiedBatchData(opts *bind.CallOpts, rollupID uint32, verifiedBatch uint64) (PolygonRollupManagerStateRootData, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "getRollupVerifiedBatchData", rollupID, verifiedBatch)

	if err != nil {
		return *new(PolygonRollupManagerStateRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(PolygonRollupManagerStateRootData)).(*PolygonRollupManagerStateRootData)

	return out0, err

}

// GetRollupVerifiedBatchData is a free data retrieval call binding the contract method 0x30490fdd.
//
// Solidity: function getRollupVerifiedBatchData(uint32 rollupID, uint64 verifiedBatch) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GetRollupVerifiedBatchData(rollupID uint32, verifiedBatch uint64) (PolygonRollupManagerStateRootData, error) {
	return _Polygonrollupmanager.Contract.GetRollupVerifiedBatchData(&_Polygonrollupmanager.CallOpts, rollupID, verifiedBatch)
}

// GetRollupVerifiedBatchData is a free data retrieval call binding the contract method 0x30490fdd.
//
// Solidity: function getRollupVerifiedBatchData(uint32 rollupID, uint64 verifiedBatch) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GetRollupVerifiedBatchData(rollupID uint32, verifiedBatch uint64) (PolygonRollupManagerStateRootData, error) {
	return _Polygonrollupmanager.Contract.GetRollupVerifiedBatchData(&_Polygonrollupmanager.CallOpts, rollupID, verifiedBatch)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanager.Contract.GlobalExitRootManager(&_Polygonrollupmanager.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanager.Contract.GlobalExitRootManager(&_Polygonrollupmanager.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanager.Contract.HasRole(&_Polygonrollupmanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanager.Contract.HasRole(&_Polygonrollupmanager.CallOpts, role, account)
}

// IsAllowedToRevertBatches is a free data retrieval call binding the contract method 0x7fea0864.
//
// Solidity: function isAllowedToRevertBatches() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) IsAllowedToRevertBatches(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "isAllowedToRevertBatches")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToRevertBatches is a free data retrieval call binding the contract method 0x7fea0864.
//
// Solidity: function isAllowedToRevertBatches() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) IsAllowedToRevertBatches() (bool, error) {
	return _Polygonrollupmanager.Contract.IsAllowedToRevertBatches(&_Polygonrollupmanager.CallOpts)
}

// IsAllowedToRevertBatches is a free data retrieval call binding the contract method 0x7fea0864.
//
// Solidity: function isAllowedToRevertBatches() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) IsAllowedToRevertBatches() (bool, error) {
	return _Polygonrollupmanager.Contract.IsAllowedToRevertBatches(&_Polygonrollupmanager.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanager.Contract.IsEmergencyState(&_Polygonrollupmanager.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanager.Contract.IsEmergencyState(&_Polygonrollupmanager.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) IsPendingStateConsolidable(opts *bind.CallOpts, rollupID uint32, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "isPendingStateConsolidable", rollupID, pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanager.Contract.IsPendingStateConsolidable(&_Polygonrollupmanager.CallOpts, rollupID, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanager.Contract.IsPendingStateConsolidable(&_Polygonrollupmanager.CallOpts, rollupID, pendingStateNum)
}

// IsRevertBatchesExecuted is a free data retrieval call binding the contract method 0xb40a15c3.
//
// Solidity: function isRevertBatchesExecuted() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) IsRevertBatchesExecuted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "isRevertBatchesExecuted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevertBatchesExecuted is a free data retrieval call binding the contract method 0xb40a15c3.
//
// Solidity: function isRevertBatchesExecuted() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) IsRevertBatchesExecuted() (bool, error) {
	return _Polygonrollupmanager.Contract.IsRevertBatchesExecuted(&_Polygonrollupmanager.CallOpts)
}

// IsRevertBatchesExecuted is a free data retrieval call binding the contract method 0xb40a15c3.
//
// Solidity: function isRevertBatchesExecuted() view returns(bool)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) IsRevertBatchesExecuted() (bool, error) {
	return _Polygonrollupmanager.Contract.IsRevertBatchesExecuted(&_Polygonrollupmanager.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanager.Contract.LastAggregationTimestamp(&_Polygonrollupmanager.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanager.Contract.LastAggregationTimestamp(&_Polygonrollupmanager.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanager.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanager.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanager.Contract.MultiplierBatchFee(&_Polygonrollupmanager.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanager.Contract.MultiplierBatchFee(&_Polygonrollupmanager.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanager.Contract.PendingStateTimeout(&_Polygonrollupmanager.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanager.Contract.PendingStateTimeout(&_Polygonrollupmanager.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanager.Contract.Pol(&_Polygonrollupmanager.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanager.Contract.Pol(&_Polygonrollupmanager.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanager.Contract.RollupAddressToID(&_Polygonrollupmanager.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanager.Contract.RollupAddressToID(&_Polygonrollupmanager.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanager.Contract.RollupCount(&_Polygonrollupmanager.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanager.Contract.RollupCount(&_Polygonrollupmanager.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	outstruct := new(struct {
		RollupContract                 common.Address
		ChainID                        uint64
		Verifier                       common.Address
		ForkID                         uint64
		LastLocalExitRoot              [32]byte
		LastBatchSequenced             uint64
		LastVerifiedBatch              uint64
		LastPendingState               uint64
		LastPendingStateConsolidated   uint64
		LastVerifiedBatchBeforeUpgrade uint64
		RollupTypeID                   uint64
		RollupCompatibilityID          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RollupContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainID = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Verifier = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ForkID = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.LastLocalExitRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.LastBatchSequenced = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.LastPendingState = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.LastPendingStateConsolidated = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatchBeforeUpgrade = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.RollupTypeID = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.RollupCompatibilityID = *abi.ConvertType(out[11], new(uint8)).(*uint8)

	return *outstruct, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RollupIDToRollupData(rollupID uint32) (struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}, error) {
	return _Polygonrollupmanager.Contract.RollupIDToRollupData(&_Polygonrollupmanager.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) RollupIDToRollupData(rollupID uint32) (struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}, error) {
	return _Polygonrollupmanager.Contract.RollupIDToRollupData(&_Polygonrollupmanager.CallOpts, rollupID)
}

// RollupIdToStateRootData is a free data retrieval call binding the contract method 0xdbebee19.
//
// Solidity: function rollupIdToStateRootData(uint32 rollupID) view returns(uint64 nextLocalExitRootBatch)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) RollupIdToStateRootData(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "rollupIdToStateRootData", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RollupIdToStateRootData is a free data retrieval call binding the contract method 0xdbebee19.
//
// Solidity: function rollupIdToStateRootData(uint32 rollupID) view returns(uint64 nextLocalExitRootBatch)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RollupIdToStateRootData(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanager.Contract.RollupIdToStateRootData(&_Polygonrollupmanager.CallOpts, rollupID)
}

// RollupIdToStateRootData is a free data retrieval call binding the contract method 0xdbebee19.
//
// Solidity: function rollupIdToStateRootData(uint32 rollupID) view returns(uint64 nextLocalExitRootBatch)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) RollupIdToStateRootData(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanager.Contract.RollupIdToStateRootData(&_Polygonrollupmanager.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanager.Contract.RollupTypeCount(&_Polygonrollupmanager.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanager.Contract.RollupTypeCount(&_Polygonrollupmanager.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

	outstruct := new(struct {
		ConsensusImplementation common.Address
		Verifier                common.Address
		ForkID                  uint64
		RollupCompatibilityID   uint8
		Obsolete                bool
		Genesis                 [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusImplementation = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ForkID = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RollupCompatibilityID = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Obsolete = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Genesis = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanager.Contract.RollupTypeMap(&_Polygonrollupmanager.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanager.Contract.RollupTypeMap(&_Polygonrollupmanager.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanager.Contract.TotalSequencedBatches(&_Polygonrollupmanager.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanager.Contract.TotalSequencedBatches(&_Polygonrollupmanager.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanager.Contract.TotalVerifiedBatches(&_Polygonrollupmanager.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanager.Contract.TotalVerifiedBatches(&_Polygonrollupmanager.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanager.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanager.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanager.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanager.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanager.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanager.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanager.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanager.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanager.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ActivateEmergencyState(&_Polygonrollupmanager.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ActivateEmergencyState(&_Polygonrollupmanager.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.AddExistingRollup(&_Polygonrollupmanager.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.AddExistingRollup(&_Polygonrollupmanager.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.AddNewRollupType(&_Polygonrollupmanager.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.AddNewRollupType(&_Polygonrollupmanager.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) ConsolidatePendingState(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "consolidatePendingState", rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ConsolidatePendingState(&_Polygonrollupmanager.TransactOpts, rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ConsolidatePendingState(&_Polygonrollupmanager.TransactOpts, rollupID, pendingStateNum)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.CreateNewRollup(&_Polygonrollupmanager.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.CreateNewRollup(&_Polygonrollupmanager.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.DeactivateEmergencyState(&_Polygonrollupmanager.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.DeactivateEmergencyState(&_Polygonrollupmanager.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.GrantRole(&_Polygonrollupmanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.GrantRole(&_Polygonrollupmanager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) Initialize(opts *bind.TransactOpts, trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "initialize", trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.Initialize(&_Polygonrollupmanager.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.Initialize(&_Polygonrollupmanager.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ObsoleteRollupType(&_Polygonrollupmanager.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ObsoleteRollupType(&_Polygonrollupmanager.TransactOpts, rollupTypeID)
}

// OnExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x41ec66dc.
//
// Solidity: function onExecuteAllForcedTransactions(uint64 newSequencedBatches, bytes32 newAccInputHash, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) OnExecuteAllForcedTransactions(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "onExecuteAllForcedTransactions", newSequencedBatches, newAccInputHash, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// OnExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x41ec66dc.
//
// Solidity: function onExecuteAllForcedTransactions(uint64 newSequencedBatches, bytes32 newAccInputHash, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) OnExecuteAllForcedTransactions(newSequencedBatches uint64, newAccInputHash [32]byte, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.OnExecuteAllForcedTransactions(&_Polygonrollupmanager.TransactOpts, newSequencedBatches, newAccInputHash, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// OnExecuteAllForcedTransactions is a paid mutator transaction binding the contract method 0x41ec66dc.
//
// Solidity: function onExecuteAllForcedTransactions(uint64 newSequencedBatches, bytes32 newAccInputHash, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) OnExecuteAllForcedTransactions(newSequencedBatches uint64, newAccInputHash [32]byte, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.OnExecuteAllForcedTransactions(&_Polygonrollupmanager.TransactOpts, newSequencedBatches, newAccInputHash, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.OnSequenceBatches(&_Polygonrollupmanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.OnSequenceBatches(&_Polygonrollupmanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) OverridePendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "overridePendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.OverridePendingState(&_Polygonrollupmanager.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.OverridePendingState(&_Polygonrollupmanager.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "proveNonDeterministicPendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanager.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanager.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.RenounceRole(&_Polygonrollupmanager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.RenounceRole(&_Polygonrollupmanager.TransactOpts, role, account)
}

// RevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xe8d2b073.
//
// Solidity: function revertLastVerifiedBatch() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) RevertLastVerifiedBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "revertLastVerifiedBatch")
}

// RevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xe8d2b073.
//
// Solidity: function revertLastVerifiedBatch() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RevertLastVerifiedBatch() (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.RevertLastVerifiedBatch(&_Polygonrollupmanager.TransactOpts)
}

// RevertLastVerifiedBatch is a paid mutator transaction binding the contract method 0xe8d2b073.
//
// Solidity: function revertLastVerifiedBatch() returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) RevertLastVerifiedBatch() (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.RevertLastVerifiedBatch(&_Polygonrollupmanager.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.RevokeRole(&_Polygonrollupmanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.RevokeRole(&_Polygonrollupmanager.TransactOpts, role, account)
}

// SetAllowedToRevertBatches is a paid mutator transaction binding the contract method 0xea1a7dde.
//
// Solidity: function setAllowedToRevertBatches(bool _status) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetAllowedToRevertBatches(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setAllowedToRevertBatches", _status)
}

// SetAllowedToRevertBatches is a paid mutator transaction binding the contract method 0xea1a7dde.
//
// Solidity: function setAllowedToRevertBatches(bool _status) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetAllowedToRevertBatches(_status bool) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetAllowedToRevertBatches(&_Polygonrollupmanager.TransactOpts, _status)
}

// SetAllowedToRevertBatches is a paid mutator transaction binding the contract method 0xea1a7dde.
//
// Solidity: function setAllowedToRevertBatches(bool _status) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetAllowedToRevertBatches(_status bool) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetAllowedToRevertBatches(&_Polygonrollupmanager.TransactOpts, _status)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetBatchFee(&_Polygonrollupmanager.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetBatchFee(&_Polygonrollupmanager.TransactOpts, newBatchFee)
}

// SetLastLocalExitRoot is a paid mutator transaction binding the contract method 0xdf6058c0.
//
// Solidity: function setLastLocalExitRoot(uint32 _rollupID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetLastLocalExitRoot(opts *bind.TransactOpts, _rollupID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setLastLocalExitRoot", _rollupID)
}

// SetLastLocalExitRoot is a paid mutator transaction binding the contract method 0xdf6058c0.
//
// Solidity: function setLastLocalExitRoot(uint32 _rollupID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetLastLocalExitRoot(_rollupID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetLastLocalExitRoot(&_Polygonrollupmanager.TransactOpts, _rollupID)
}

// SetLastLocalExitRoot is a paid mutator transaction binding the contract method 0xdf6058c0.
//
// Solidity: function setLastLocalExitRoot(uint32 _rollupID) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetLastLocalExitRoot(_rollupID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetLastLocalExitRoot(&_Polygonrollupmanager.TransactOpts, _rollupID)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetMultiplierBatchFee(&_Polygonrollupmanager.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetMultiplierBatchFee(&_Polygonrollupmanager.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetPendingStateTimeout(&_Polygonrollupmanager.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetPendingStateTimeout(&_Polygonrollupmanager.TransactOpts, newPendingStateTimeout)
}

// SetRevertBatchesExecuted is a paid mutator transaction binding the contract method 0x280a4774.
//
// Solidity: function setRevertBatchesExecuted(bool _status) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetRevertBatchesExecuted(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setRevertBatchesExecuted", _status)
}

// SetRevertBatchesExecuted is a paid mutator transaction binding the contract method 0x280a4774.
//
// Solidity: function setRevertBatchesExecuted(bool _status) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetRevertBatchesExecuted(_status bool) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetRevertBatchesExecuted(&_Polygonrollupmanager.TransactOpts, _status)
}

// SetRevertBatchesExecuted is a paid mutator transaction binding the contract method 0x280a4774.
//
// Solidity: function setRevertBatchesExecuted(bool _status) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetRevertBatchesExecuted(_status bool) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetRevertBatchesExecuted(&_Polygonrollupmanager.TransactOpts, _status)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanager.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanager.TransactOpts, newTrustedAggregatorTimeout)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanager.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanager.TransactOpts, newVerifyBatchTimeTarget)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.UpdateRollup(&_Polygonrollupmanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.UpdateRollup(&_Polygonrollupmanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) VerifyBatches(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "verifyBatches", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.VerifyBatches(&_Polygonrollupmanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.VerifyBatches(&_Polygonrollupmanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanager *PolygonrollupmanagerTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanager.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// PolygonrollupmanagerAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerAddExistingRollupIterator struct {
	Event *PolygonrollupmanagerAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerAddExistingRollup)
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
		it.Event = new(PolygonrollupmanagerAddExistingRollup)
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
func (it *PolygonrollupmanagerAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerAddExistingRollup represents a AddExistingRollup event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerAddExistingRollup struct {
	RollupID                       uint32
	ForkID                         uint64
	RollupAddress                  common.Address
	ChainID                        uint64
	RollupCompatibilityID          uint8
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterAddExistingRollup is a free log retrieval operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerAddExistingRollupIterator{contract: _Polygonrollupmanager.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerAddExistingRollup)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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

// ParseAddExistingRollup is a log parse operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseAddExistingRollup(log types.Log) (*PolygonrollupmanagerAddExistingRollup, error) {
	event := new(PolygonrollupmanagerAddExistingRollup)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerAddNewRollupTypeIterator struct {
	Event *PolygonrollupmanagerAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerAddNewRollupType)
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
		it.Event = new(PolygonrollupmanagerAddNewRollupType)
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
func (it *PolygonrollupmanagerAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerAddNewRollupType represents a AddNewRollupType event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerAddNewRollupType struct {
	RollupTypeID            uint32
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Genesis                 [32]byte
	Description             string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAddNewRollupType is a free log retrieval operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagerAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerAddNewRollupTypeIterator{contract: _Polygonrollupmanager.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerAddNewRollupType)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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

// ParseAddNewRollupType is a log parse operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseAddNewRollupType(log types.Log) (*PolygonrollupmanagerAddNewRollupType, error) {
	event := new(PolygonrollupmanagerAddNewRollupType)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerAllowToRevertBatchesIterator is returned from FilterAllowToRevertBatches and is used to iterate over the raw logs and unpacked data for AllowToRevertBatches events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerAllowToRevertBatchesIterator struct {
	Event *PolygonrollupmanagerAllowToRevertBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerAllowToRevertBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerAllowToRevertBatches)
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
		it.Event = new(PolygonrollupmanagerAllowToRevertBatches)
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
func (it *PolygonrollupmanagerAllowToRevertBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerAllowToRevertBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerAllowToRevertBatches represents a AllowToRevertBatches event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerAllowToRevertBatches struct {
	Status bool
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllowToRevertBatches is a free log retrieval operation binding the contract event 0x4bd13f042de7fee85a224d6bbb9efaab00a2e510a67c51f907c4870182192701.
//
// Solidity: event AllowToRevertBatches(bool _status, address _sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterAllowToRevertBatches(opts *bind.FilterOpts) (*PolygonrollupmanagerAllowToRevertBatchesIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "AllowToRevertBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerAllowToRevertBatchesIterator{contract: _Polygonrollupmanager.contract, event: "AllowToRevertBatches", logs: logs, sub: sub}, nil
}

// WatchAllowToRevertBatches is a free log subscription operation binding the contract event 0x4bd13f042de7fee85a224d6bbb9efaab00a2e510a67c51f907c4870182192701.
//
// Solidity: event AllowToRevertBatches(bool _status, address _sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchAllowToRevertBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerAllowToRevertBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "AllowToRevertBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerAllowToRevertBatches)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "AllowToRevertBatches", log); err != nil {
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

// ParseAllowToRevertBatches is a log parse operation binding the contract event 0x4bd13f042de7fee85a224d6bbb9efaab00a2e510a67c51f907c4870182192701.
//
// Solidity: event AllowToRevertBatches(bool _status, address _sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseAllowToRevertBatches(log types.Log) (*PolygonrollupmanagerAllowToRevertBatches, error) {
	event := new(PolygonrollupmanagerAllowToRevertBatches)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "AllowToRevertBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerConsolidatePendingStateIterator struct {
	Event *PolygonrollupmanagerConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerConsolidatePendingState)
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
		it.Event = new(PolygonrollupmanagerConsolidatePendingState)
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
func (it *PolygonrollupmanagerConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerConsolidatePendingState struct {
	RollupID        uint32
	NumBatch        uint64
	StateRoot       [32]byte
	ExitRoot        [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerConsolidatePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerConsolidatePendingStateIterator{contract: _Polygonrollupmanager.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerConsolidatePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerConsolidatePendingState)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonrollupmanagerConsolidatePendingState, error) {
	event := new(PolygonrollupmanagerConsolidatePendingState)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerCreateNewRollupIterator struct {
	Event *PolygonrollupmanagerCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerCreateNewRollup)
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
		it.Event = new(PolygonrollupmanagerCreateNewRollup)
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
func (it *PolygonrollupmanagerCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerCreateNewRollup represents a CreateNewRollup event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerCreateNewRollup struct {
	RollupID        uint32
	RollupTypeID    uint32
	RollupAddress   common.Address
	ChainID         uint64
	GasTokenAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewRollup is a free log retrieval operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerCreateNewRollupIterator{contract: _Polygonrollupmanager.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerCreateNewRollup)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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

// ParseCreateNewRollup is a log parse operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseCreateNewRollup(log types.Log) (*PolygonrollupmanagerCreateNewRollup, error) {
	event := new(PolygonrollupmanagerCreateNewRollup)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerEmergencyStateActivatedIterator struct {
	Event *PolygonrollupmanagerEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerEmergencyStateActivated)
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
		it.Event = new(PolygonrollupmanagerEmergencyStateActivated)
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
func (it *PolygonrollupmanagerEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonrollupmanagerEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerEmergencyStateActivatedIterator{contract: _Polygonrollupmanager.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerEmergencyStateActivated)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonrollupmanagerEmergencyStateActivated, error) {
	event := new(PolygonrollupmanagerEmergencyStateActivated)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerEmergencyStateDeactivatedIterator struct {
	Event *PolygonrollupmanagerEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerEmergencyStateDeactivated)
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
		it.Event = new(PolygonrollupmanagerEmergencyStateDeactivated)
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
func (it *PolygonrollupmanagerEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonrollupmanagerEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerEmergencyStateDeactivatedIterator{contract: _Polygonrollupmanager.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerEmergencyStateDeactivated)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonrollupmanagerEmergencyStateDeactivated, error) {
	event := new(PolygonrollupmanagerEmergencyStateDeactivated)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerInitializedIterator struct {
	Event *PolygonrollupmanagerInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerInitialized)
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
		it.Event = new(PolygonrollupmanagerInitialized)
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
func (it *PolygonrollupmanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerInitialized represents a Initialized event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupmanagerInitializedIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerInitializedIterator{contract: _Polygonrollupmanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerInitialized)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseInitialized(log types.Log) (*PolygonrollupmanagerInitialized, error) {
	event := new(PolygonrollupmanagerInitialized)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerLastLocalExitRootIterator is returned from FilterLastLocalExitRoot and is used to iterate over the raw logs and unpacked data for LastLocalExitRoot events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerLastLocalExitRootIterator struct {
	Event *PolygonrollupmanagerLastLocalExitRoot // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerLastLocalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerLastLocalExitRoot)
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
		it.Event = new(PolygonrollupmanagerLastLocalExitRoot)
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
func (it *PolygonrollupmanagerLastLocalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerLastLocalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerLastLocalExitRoot represents a LastLocalExitRoot event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerLastLocalExitRoot struct {
	RollupID          uint32
	LastLocalExitRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLastLocalExitRoot is a free log retrieval operation binding the contract event 0x728d0f7d9b5ea5d935d30470f2ebf8f73e4c87870780094fffb000337b1a1fa4.
//
// Solidity: event LastLocalExitRoot(uint32 _rollupID, bytes32 _lastLocalExitRoot)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterLastLocalExitRoot(opts *bind.FilterOpts) (*PolygonrollupmanagerLastLocalExitRootIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "LastLocalExitRoot")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerLastLocalExitRootIterator{contract: _Polygonrollupmanager.contract, event: "LastLocalExitRoot", logs: logs, sub: sub}, nil
}

// WatchLastLocalExitRoot is a free log subscription operation binding the contract event 0x728d0f7d9b5ea5d935d30470f2ebf8f73e4c87870780094fffb000337b1a1fa4.
//
// Solidity: event LastLocalExitRoot(uint32 _rollupID, bytes32 _lastLocalExitRoot)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchLastLocalExitRoot(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerLastLocalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "LastLocalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerLastLocalExitRoot)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "LastLocalExitRoot", log); err != nil {
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

// ParseLastLocalExitRoot is a log parse operation binding the contract event 0x728d0f7d9b5ea5d935d30470f2ebf8f73e4c87870780094fffb000337b1a1fa4.
//
// Solidity: event LastLocalExitRoot(uint32 _rollupID, bytes32 _lastLocalExitRoot)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseLastLocalExitRoot(log types.Log) (*PolygonrollupmanagerLastLocalExitRoot, error) {
	event := new(PolygonrollupmanagerLastLocalExitRoot)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "LastLocalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerObsoleteRollupTypeIterator struct {
	Event *PolygonrollupmanagerObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerObsoleteRollupType)
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
		it.Event = new(PolygonrollupmanagerObsoleteRollupType)
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
func (it *PolygonrollupmanagerObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerObsoleteRollupType represents a ObsoleteRollupType event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagerObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerObsoleteRollupTypeIterator{contract: _Polygonrollupmanager.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerObsoleteRollupType)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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

// ParseObsoleteRollupType is a log parse operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseObsoleteRollupType(log types.Log) (*PolygonrollupmanagerObsoleteRollupType, error) {
	event := new(PolygonrollupmanagerObsoleteRollupType)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerOnSequenceBatchesIterator struct {
	Event *PolygonrollupmanagerOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerOnSequenceBatches)
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
		it.Event = new(PolygonrollupmanagerOnSequenceBatches)
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
func (it *PolygonrollupmanagerOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerOnSequenceBatches represents a OnSequenceBatches event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerOnSequenceBatchesIterator{contract: _Polygonrollupmanager.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerOnSequenceBatches)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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

// ParseOnSequenceBatches is a log parse operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseOnSequenceBatches(log types.Log) (*PolygonrollupmanagerOnSequenceBatches, error) {
	event := new(PolygonrollupmanagerOnSequenceBatches)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerOverridePendingStateIterator struct {
	Event *PolygonrollupmanagerOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerOverridePendingState)
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
		it.Event = new(PolygonrollupmanagerOverridePendingState)
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
func (it *PolygonrollupmanagerOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerOverridePendingState represents a OverridePendingState event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerOverridePendingState struct {
	RollupID   uint32
	NumBatch   uint64
	StateRoot  [32]byte
	ExitRoot   [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterOverridePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerOverridePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerOverridePendingStateIterator{contract: _Polygonrollupmanager.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerOverridePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerOverridePendingState)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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

// ParseOverridePendingState is a log parse operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseOverridePendingState(log types.Log) (*PolygonrollupmanagerOverridePendingState, error) {
	event := new(PolygonrollupmanagerOverridePendingState)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerProveNonDeterministicPendingStateIterator struct {
	Event *PolygonrollupmanagerProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerProveNonDeterministicPendingState)
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
		it.Event = new(PolygonrollupmanagerProveNonDeterministicPendingState)
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
func (it *PolygonrollupmanagerProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonrollupmanagerProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerProveNonDeterministicPendingStateIterator{contract: _Polygonrollupmanager.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerProveNonDeterministicPendingState)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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

// ParseProveNonDeterministicPendingState is a log parse operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonrollupmanagerProveNonDeterministicPendingState, error) {
	event := new(PolygonrollupmanagerProveNonDeterministicPendingState)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerRevertLastVerifiedBatchIterator is returned from FilterRevertLastVerifiedBatch and is used to iterate over the raw logs and unpacked data for RevertLastVerifiedBatch events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRevertLastVerifiedBatchIterator struct {
	Event *PolygonrollupmanagerRevertLastVerifiedBatch // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerRevertLastVerifiedBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerRevertLastVerifiedBatch)
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
		it.Event = new(PolygonrollupmanagerRevertLastVerifiedBatch)
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
func (it *PolygonrollupmanagerRevertLastVerifiedBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerRevertLastVerifiedBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerRevertLastVerifiedBatch represents a RevertLastVerifiedBatch event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRevertLastVerifiedBatch struct {
	PrevlastVerifiedBatch uint64
	LastVerifiedBatch     uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRevertLastVerifiedBatch is a free log retrieval operation binding the contract event 0xc2ac6a39960d89d285c7e09e8ba997a82779fa1d2d673b6ca813cdbf35309ddc.
//
// Solidity: event RevertLastVerifiedBatch(uint64 _prevlastVerifiedBatch, uint64 lastVerifiedBatch)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterRevertLastVerifiedBatch(opts *bind.FilterOpts) (*PolygonrollupmanagerRevertLastVerifiedBatchIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "RevertLastVerifiedBatch")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerRevertLastVerifiedBatchIterator{contract: _Polygonrollupmanager.contract, event: "RevertLastVerifiedBatch", logs: logs, sub: sub}, nil
}

// WatchRevertLastVerifiedBatch is a free log subscription operation binding the contract event 0xc2ac6a39960d89d285c7e09e8ba997a82779fa1d2d673b6ca813cdbf35309ddc.
//
// Solidity: event RevertLastVerifiedBatch(uint64 _prevlastVerifiedBatch, uint64 lastVerifiedBatch)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchRevertLastVerifiedBatch(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerRevertLastVerifiedBatch) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "RevertLastVerifiedBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerRevertLastVerifiedBatch)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "RevertLastVerifiedBatch", log); err != nil {
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

// ParseRevertLastVerifiedBatch is a log parse operation binding the contract event 0xc2ac6a39960d89d285c7e09e8ba997a82779fa1d2d673b6ca813cdbf35309ddc.
//
// Solidity: event RevertLastVerifiedBatch(uint64 _prevlastVerifiedBatch, uint64 lastVerifiedBatch)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseRevertLastVerifiedBatch(log types.Log) (*PolygonrollupmanagerRevertLastVerifiedBatch, error) {
	event := new(PolygonrollupmanagerRevertLastVerifiedBatch)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "RevertLastVerifiedBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRoleAdminChangedIterator struct {
	Event *PolygonrollupmanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerRoleAdminChanged)
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
		it.Event = new(PolygonrollupmanagerRoleAdminChanged)
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
func (it *PolygonrollupmanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonrollupmanagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerRoleAdminChangedIterator{contract: _Polygonrollupmanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerRoleAdminChanged)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonrollupmanagerRoleAdminChanged, error) {
	event := new(PolygonrollupmanagerRoleAdminChanged)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRoleGrantedIterator struct {
	Event *PolygonrollupmanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerRoleGranted)
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
		it.Event = new(PolygonrollupmanagerRoleGranted)
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
func (it *PolygonrollupmanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerRoleGranted represents a RoleGranted event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerRoleGrantedIterator{contract: _Polygonrollupmanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerRoleGranted)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseRoleGranted(log types.Log) (*PolygonrollupmanagerRoleGranted, error) {
	event := new(PolygonrollupmanagerRoleGranted)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRoleRevokedIterator struct {
	Event *PolygonrollupmanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerRoleRevoked)
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
		it.Event = new(PolygonrollupmanagerRoleRevoked)
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
func (it *PolygonrollupmanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerRoleRevoked represents a RoleRevoked event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerRoleRevokedIterator{contract: _Polygonrollupmanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerRoleRevoked)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseRoleRevoked(log types.Log) (*PolygonrollupmanagerRoleRevoked, error) {
	event := new(PolygonrollupmanagerRoleRevoked)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetBatchFeeIterator struct {
	Event *PolygonrollupmanagerSetBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerSetBatchFee)
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
		it.Event = new(PolygonrollupmanagerSetBatchFee)
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
func (it *PolygonrollupmanagerSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerSetBatchFee represents a SetBatchFee event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagerSetBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerSetBatchFeeIterator{contract: _Polygonrollupmanager.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerSetBatchFee)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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

// ParseSetBatchFee is a log parse operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseSetBatchFee(log types.Log) (*PolygonrollupmanagerSetBatchFee, error) {
	event := new(PolygonrollupmanagerSetBatchFee)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetMultiplierBatchFeeIterator struct {
	Event *PolygonrollupmanagerSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerSetMultiplierBatchFee)
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
		it.Event = new(PolygonrollupmanagerSetMultiplierBatchFee)
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
func (it *PolygonrollupmanagerSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagerSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerSetMultiplierBatchFeeIterator{contract: _Polygonrollupmanager.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerSetMultiplierBatchFee)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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

// ParseSetMultiplierBatchFee is a log parse operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonrollupmanagerSetMultiplierBatchFee, error) {
	event := new(PolygonrollupmanagerSetMultiplierBatchFee)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetPendingStateTimeoutIterator struct {
	Event *PolygonrollupmanagerSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerSetPendingStateTimeout)
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
		it.Event = new(PolygonrollupmanagerSetPendingStateTimeout)
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
func (it *PolygonrollupmanagerSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagerSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerSetPendingStateTimeoutIterator{contract: _Polygonrollupmanager.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerSetPendingStateTimeout)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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

// ParseSetPendingStateTimeout is a log parse operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonrollupmanagerSetPendingStateTimeout, error) {
	event := new(PolygonrollupmanagerSetPendingStateTimeout)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagerSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerSetTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagerSetTrustedAggregator)
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
func (it *PolygonrollupmanagerSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonrollupmanagerSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerSetTrustedAggregatorIterator{contract: _Polygonrollupmanager.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerSetTrustedAggregator)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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

// ParseSetTrustedAggregator is a log parse operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonrollupmanagerSetTrustedAggregator, error) {
	event := new(PolygonrollupmanagerSetTrustedAggregator)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonrollupmanagerSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonrollupmanagerSetTrustedAggregatorTimeout)
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
func (it *PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerSetTrustedAggregatorTimeoutIterator{contract: _Polygonrollupmanager.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerSetTrustedAggregatorTimeout)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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

// ParseSetTrustedAggregatorTimeout is a log parse operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonrollupmanagerSetTrustedAggregatorTimeout, error) {
	event := new(PolygonrollupmanagerSetTrustedAggregatorTimeout)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonrollupmanagerSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonrollupmanagerSetVerifyBatchTimeTarget)
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
func (it *PolygonrollupmanagerSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonrollupmanagerSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerSetVerifyBatchTimeTargetIterator{contract: _Polygonrollupmanager.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerSetVerifyBatchTimeTarget)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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

// ParseSetVerifyBatchTimeTarget is a log parse operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonrollupmanagerSetVerifyBatchTimeTarget, error) {
	event := new(PolygonrollupmanagerSetVerifyBatchTimeTarget)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerUpdateRollupIterator struct {
	Event *PolygonrollupmanagerUpdateRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerUpdateRollup)
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
		it.Event = new(PolygonrollupmanagerUpdateRollup)
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
func (it *PolygonrollupmanagerUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerUpdateRollup represents a UpdateRollup event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerUpdateRollupIterator{contract: _Polygonrollupmanager.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerUpdateRollup)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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

// ParseUpdateRollup is a log parse operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseUpdateRollup(log types.Log) (*PolygonrollupmanagerUpdateRollup, error) {
	event := new(PolygonrollupmanagerUpdateRollup)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerVerifyBatchesIterator struct {
	Event *PolygonrollupmanagerVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerVerifyBatches)
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
		it.Event = new(PolygonrollupmanagerVerifyBatches)
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
func (it *PolygonrollupmanagerVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerVerifyBatches represents a VerifyBatches event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerVerifyBatches struct {
	RollupID   uint32
	NumBatch   uint64
	StateRoot  [32]byte
	ExitRoot   [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterVerifyBatches(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagerVerifyBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerVerifyBatchesIterator{contract: _Polygonrollupmanager.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerVerifyBatches, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerVerifyBatches)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseVerifyBatches(log types.Log) (*PolygonrollupmanagerVerifyBatches, error) {
	event := new(PolygonrollupmanagerVerifyBatches)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagerVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagerVerifyBatchesTrustedAggregator)
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
func (it *PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonrollupmanager contract.
type PolygonrollupmanagerVerifyBatchesTrustedAggregator struct {
	RollupID   uint32
	NumBatch   uint64
	StateRoot  [32]byte
	ExitRoot   [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerVerifyBatchesTrustedAggregatorIterator{contract: _Polygonrollupmanager.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanager.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerVerifyBatchesTrustedAggregator)
				if err := _Polygonrollupmanager.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanager *PolygonrollupmanagerFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonrollupmanagerVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonrollupmanagerVerifyBatchesTrustedAggregator)
	if err := _Polygonrollupmanager.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
