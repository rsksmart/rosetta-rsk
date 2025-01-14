// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rsk

import (
	"context"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

const (
	// NodeVersion is the version of rskj we are using.
	NodeVersion = "2.2.0"

	// Blockchain is RSK.
	Blockchain string = "RSK"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "Mainnet"

	// TestnetNetwork is the value of the network
	// in TestnetNetworkIdentifier.
	TestnetNetwork string = "Testnet"

	// MinerRewardOpType is used to describe
	// a miner block reward.
	MinerRewardOpType = "MINER_REWARD"

	// UncleRewardOpType is used to describe
	// an uncle block reward.
	UncleRewardOpType = "UNCLE_REWARD"

	// FeeOpType is used to represent fee operations.
	FeeOpType = "FEE"

	// CallOpType is used to represent CALL trace operations.
	CallOpType = "CALL"

	// CreateOpType is used to represent CREATE trace operations.
	CreateOpType = "CREATE"

	// Create2OpType is used to represent CREATE2 trace operations.
	Create2OpType = "CREATE2"

	// SelfDestructOpType is used to represent SELFDESTRUCT trace operations.
	SelfDestructOpType = "SELFDESTRUCT"

	// CallCodeOpType is used to represent CALLCODE trace operations.
	CallCodeOpType = "CALLCODE"

	// DelegateCallOpType is used to represent DELEGATECALL trace operations.
	DelegateCallOpType = "DELEGATECALL"

	// StaticCallOpType is used to represent STATICCALL trace operations.
	StaticCallOpType = "STATICCALL"

	// DestructOpType is a synthetic operation used to represent the
	// deletion of suicided accounts that still have funds at the end
	// of a transaction.
	DestructOpType = "DESTRUCT"

	// SuccessStatus is the status of any
	// Ethereum operation considered successful.
	SuccessStatus = "SUCCESS"

	// FailureStatus is the status of any
	// Ethereum operation considered unsuccessful.
	FailureStatus = "FAILURE"

	// HistoricalBalanceSupported is whether
	// historical balance is supported.
	HistoricalBalanceSupported = true

	// GenesisBlockIndex is the index of the
	// genesis block.
	GenesisBlockIndex = int64(0)

	// TransferGasLimit is the gas limit
	// of a transfer.
	TransferGasLimit = int64(21000) //nolint:gomnd

	// MainnetArguments are the arguments to start a mainnet rskj instance.
	MainnetArguments = `-Xmx4G -Drsk.conf.file=/app/rsk/rsk.conf -cp /app/rskj-core-latest.jar co.rsk.Start`

	// IncludeMempoolCoins does not apply to rosetta-rsk as it is not UTXO-based.
	IncludeMempoolCoins = false

	BridgeTransactionDestinationAddress = "0x0000000000000000000000000000000001000006"
	RemascTransactionDestinationAddress = "0x0000000000000000000000000000000001000008"
	RskNormalTransactionCode            = "0x"
	RskRemascTransactionType            = "remasc"
	RskBridgeTransactionType            = "bridge"
	RskContractCallTransactionType      = "contract call"
	RskContractCreationTransactionType  = "contract creation"
	RskNormalTransactionType            = "normal"
)

var (
	// TestnetRskjArguments are the arguments to start a testnet rskj instance.
	TestnetRskjArguments = fmt.Sprintf("%s --testnet", MainnetArguments)

	// MainnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the mainnet genesis block.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  "0xf88529d4ab262c0f4d042e9d8d3f2472848eaafe1a9b7213f57617eb40a9f9e0",
		Index: GenesisBlockIndex,
	}

	MainnetChainID = big.NewInt(30)

	// TestnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the testnet genesis block.
	TestnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  "0xcabb7fbe88cd6d922042a32ffc08ce8b1fbb37d650b9d4e7dbfe2a7469adfa42",
		Index: GenesisBlockIndex,
	}

	TestnetChainID = big.NewInt(31)

	// DefaultCurrency is the *types.Currency for all
	// Ethereum networks.
	DefaultCurrency = &types.Currency{
		Symbol:   "RBTC",
		Decimals: 18,
	}
	RIFCurrency = &types.Currency{
		Symbol:   "RIF",
		Decimals: 18,
	}
	RDOCCurrency = &types.Currency{
		Symbol:   "RDOC",
		Decimals: 18,
	}
	DOCCurrency = &types.Currency{
		Symbol:   "DOC",
		Decimals: 18,
	}

	AvailableCurrencies = []*types.Currency{DefaultCurrency, RIFCurrency, RDOCCurrency, DOCCurrency}

	AddressByTokenSymbol = map[string]string{
		RIFCurrency.Symbol:  "0x2acc95758f8b5f583470ba265eb685a8f45fc9d5",
		RDOCCurrency.Symbol: "0xE700691Da7B9851F2F35f8b8182C69C53ccad9DB",
		DOCCurrency.Symbol:  "0x2d919f19D4892381d58EdEbEcA66D5642ceF1A1F",
	}

	DecimalsByCurrencySymbol = map[string]int32{
		DefaultCurrency.Symbol: DefaultCurrency.Decimals,
		RIFCurrency.Symbol:     18,
		RDOCCurrency.Symbol:    18,
		DOCCurrency.Symbol:     18,
	}

	// OperationTypes are all supported operation types.
	OperationTypes = []string{
		MinerRewardOpType,
		UncleRewardOpType,
		FeeOpType,
		CallOpType,
		CreateOpType,
		Create2OpType,
		SelfDestructOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
		DestructOpType,
	}

	// OperationStatuses are all supported operation statuses.
	OperationStatuses = []*types.OperationStatus{
		{
			Status:     SuccessStatus,
			Successful: true,
		},
		{
			Status:     FailureStatus,
			Successful: false,
		},
	}

	// CallMethods are all supported call methods.
	CallMethods = []string{
		"eth_getTransactionReceipt",
	}
)

// JSONRPC is the interface for accessing go-ethereum's JSON RPC endpoint.
type JSONRPC interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	Close()
}

// CallType returns a boolean indicating
// if the provided trace type is a call type.
func CallType(t string) bool {
	callTypes := []string{
		CallOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
	}

	for _, callType := range callTypes {
		if callType == t {
			return true
		}
	}

	return false
}

// CreateType returns a boolean indicating
// if the provided trace type is a create type.
func CreateType(t string) bool {
	createTypes := []string{
		CreateOpType,
		Create2OpType,
	}

	for _, createType := range createTypes {
		if createType == t {
			return true
		}
	}

	return false
}

type Block struct {
	Number       string         `json:"number"`
	Hash         string         `json:"hash"`
	ParentHash   string         `json:"parentHash"`
	Timestamp    string         `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	Hash             string `json:"hash"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	GasPrice         string `json:"gasPrice"`
	Gas              string `json:"gas"`
}

type Receipt struct {
	To              string `json:"to"`
	ContractAddress string `json:"contractAddress"`
	GasUsed         string `json:"gasUsed"`
}

type Trace struct {
	Result          string          `json:"result,omitempty"`
	Error           string          `json:"error,omitempty"`
	Reverted        bool            `json:"reverted,omitempty"`
	ProgramInvoke   interface{}     `json:"programInvoke,omitempty"`
	TransferInvoke  *TransferInvoke `json:"transferInvoke,omitempty"`
	ContractAddress string          `json:"contractAddress,omitempty"`
	SubTraces       []*SubTrace     `json:"subtraces,omitempty"`
	StructLogs      []*StructLog    `json:"structLogs,omitempty"`
}

//type Trace struct {
//	SubTraces []*SubTrace `json:"subtraces"`
//	Result    string             `json:"result"`
//	Error     string             `json:"error"`
//	Reverted  bool               `json:"reverted"`
//}

//type SummarizedProgramTrace struct {
//	Trace
//	ProgramInvoke  interface{}     `json:"programInvoke"`
//	TransferInvoke *TransferInvoke `json:"transferInvoke"`
//}
//
//type DetailedProgramTrace struct {
//	// storage properties are not parsed in purpose, since the rskj node doesn't support options to disable storage yet.
//	Trace
//	ContractAddress string
//	StructLogs      []*StructLog
//}

type SubTrace struct {
	TraceType      string         `json:"traceType"`
	CallType       string         `json:"callType"`
	CreationData   interface{}    `json:"creationData"`
	CreationMethod string         `json:"creationMethod"`
	InvokeData     *InvokeData    `json:"invokeData"`
	ProgramResult  *ProgramResult `json:"programResult"`
	CodeAddress    string         `json:"codeAddress"`
	SubTraces      []*SubTrace    `json:"subtraces"`
}

type ProgramResult struct {
	Revert  bool  `json:"revert"`
	GasUsed int64 `json:"gasUsed"`
}

type TransferInvoke struct {
	OwnerAddress  string `json:"ownerAddress"`
	CallerAddress string `json:"callerAddress"`
	Gas           int64  `json:"gas"`
	CallValue     string `json:"callValue"`
}

type InvokeData struct {
	TransferInvoke
	DataValue string `json:"dataValue"`
	DataSize  string `json:"dataSize"`
}

type StructLog struct {
	Op      string `json:"op"`
	Depth   int64  `json:"depth"`
	Pc      int64  `json:"pc"`
	Gas     int64  `json:"gas"`
	GasCost int64  `json:"gasCost"`
}
