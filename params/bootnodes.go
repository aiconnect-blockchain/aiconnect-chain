// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// TODO: use the true node
	// "enode://527a7280361271ab4ff5ce0fd36bfc15800a0f629b8f23fc9436905b1efa5d9976286d8a038d304df409febd42d56f4ca63d49f0054f2659469f69b5edead857@3.137.172.33:33101",
	
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the testnet.
var TestnetBootnodes = []string{
	// TODO: use the true node
	// "enode://10638ced559d8f4fa7e569097c312b930ea95d7be6d99d74a7ddfce7c141ae3c7d55d8485581d4260dedf96325e4ac89a904ff9928996c60fb0270c309233f03@3.144.195.4:33601",
}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
