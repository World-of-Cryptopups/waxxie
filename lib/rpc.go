package lib

import "github.com/World-of-Cryptopups/eosrpc.go"

var WAX = eosrpc.New("https://wax.greymass.com")
var CHAIN = WAX.NewChainAPI()
