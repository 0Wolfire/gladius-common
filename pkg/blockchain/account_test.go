package blockchain_test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-common/pkg/blockchain"
	"testing"
)

// Test for printing out the balance struct, does not assert any values
func TestGetAccountBalance(t *testing.T) {
	address := common.HexToAddress("0xddf4192f04856aa4afe906e3d811e71a93eeeb24")
	glaBalance, err := blockchain.GetAccountBalance(address, blockchain.GLA)
	if err != nil {
		t.Error(err)
	}

	ethBalance, err := blockchain.GetAccountBalance(address, blockchain.ETH)
	if err != nil {
		t.Error(err)
	}

	glaBalanceJson, _ := json.Marshal(glaBalance)
 	fmt.Println(string(glaBalanceJson))

	ethBalanceJson, _ := json.Marshal(ethBalance)
	fmt.Println(string(ethBalanceJson))
}

func TestGetAccountTransactions(t *testing.T) {
	address := common.HexToAddress("0xddf4192f04856aa4afe906e3d811e71a93eeeb24")
	filter := blockchain.TransactionOptions{
		Filters: &blockchain.TransactionFilter{
			EthTransfer:true,
		},
	}
	response, _ := blockchain.GetAccountTransactions(address, filter)
	ethBalanceJson, _ := json.Marshal(response)
	fmt.Println(string(ethBalanceJson))
}