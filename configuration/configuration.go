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

package configuration

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/rsksmart/rosetta-rsk/rsk"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// Mode is the setting that determines if
// the implementation is "online" or "offline".
type Mode string

const (
	// Online is when the implementation is permitted
	// to make outbound connections.
	Online Mode = "ONLINE"

	// Offline is when the implementation is not permitted
	// to make outbound connections.
	Offline Mode = "OFFLINE"

	Mainnet string = "MAINNET"

	Testnet string = "TESTNET"

	// DataDirectory is the default location for all
	// persistent data.
	DataDirectory = "/data"

	// ModeEnv is the environment variable read
	// to determine mode.
	ModeEnv = "MODE"

	// NetworkEnv is the environment variable
	// read to determine network.
	NetworkEnv = "NETWORK"

	// PortEnv is the environment variable
	// read to determine the port for the Rosetta
	// implementation.
	PortEnv = "PORT"

	// RskjEnv is an optional environment variable
	// used to connect rosetta-rsk to an already
	// running rskj node.
	RskjEnv = "RSKJ"

	// DefaultRskjURL is the default URL for
	// a running rskj node. This is used
	// when RskjEnv is not populated.
	DefaultRskjURL = "http://localhost:4444"

	// MiddlewareVersion is the version of rosetta-rsk.
	MiddlewareVersion = "0.1.0"
)

// Configuration determines how
type Configuration struct {
	Mode                   Mode
	Network                *types.NetworkIdentifier
	GenesisBlockIdentifier *types.BlockIdentifier
	RskjURL                string
	RemoteRskj             bool
	Port                   int
	RskjArguments          string
	ChainID                *big.Int
}

// LoadConfiguration attempts to create a new Configuration
// using the ENVs in the environment.
func LoadConfiguration() (*Configuration, error) {
	config := &Configuration{}

	modeValue := Mode(os.Getenv(ModeEnv))
	switch modeValue {
	case Online:
		config.Mode = Online
	case Offline:
		config.Mode = Offline
	case "":
		return nil, errors.New("MODE must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid mode", modeValue)
	}

	networkValue := os.Getenv(NetworkEnv)
	switch networkValue {
	case Mainnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: rsk.Blockchain,
			Network:    rsk.MainnetNetwork,
		}
		config.GenesisBlockIdentifier = rsk.MainnetGenesisBlockIdentifier
		config.ChainID = rsk.MainnetChainID
		config.RskjArguments = rsk.MainnetArguments
	case Testnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: rsk.Blockchain,
			Network:    rsk.TestnetNetwork,
		}
		config.GenesisBlockIdentifier = rsk.TestnetGenesisBlockIdentifier
		config.ChainID = rsk.TestnetChainID
		config.RskjArguments = rsk.TestnetRskjArguments
	case "":
		return nil, errors.New("NETWORK must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid network", networkValue)
	}

	config.RskjURL = DefaultRskjURL
	envRskjURL := os.Getenv(RskjEnv)
	if len(envRskjURL) > 0 {
		config.RemoteRskj = true
		config.RskjURL = envRskjURL
	}

	portValue := os.Getenv(PortEnv)
	if len(portValue) == 0 {
		return nil, errors.New("PORT must be populated")
	}

	port, err := strconv.Atoi(portValue)
	if err != nil || len(portValue) == 0 || port <= 0 {
		return nil, fmt.Errorf("%w: unable to parse port %s", err, portValue)
	}
	config.Port = port

	return config, nil
}
