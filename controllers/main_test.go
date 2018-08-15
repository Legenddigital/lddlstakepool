package controllers

import (
	"testing"

	"github.com/Legenddigital/lddld/chaincfg"
)

func TestGetNetworkName(t *testing.T) {
	// First test that "testnet3" is translated to "testnet"
	mc := MainController{
		params: &chaincfg.TestNet3Params,
	}

	netName := mc.getNetworkName()
	if netName != "testnet" {
		t.Errorf("Incorrect network name: expected %s, got %s", "testnet",
			netName)
	}

	// ensure "mainnet" is unaltered
	mc.params = &chaincfg.MainNetParams
	netName = mc.getNetworkName()
	if netName != "mainnet" {
		t.Errorf("Incorrect network name: expected %s, got %s", "mainnet",
			netName)
	}
}
