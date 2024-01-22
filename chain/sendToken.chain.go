package chain 

import (
	"context"
	cosmosBankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
)

func SendToken(toAddress string, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string ) (success bool,err error, message string ,txhash string) {

	// check if admin have enough balance to send tokens
	isBalance, amount, _ := CheckBalance(ctx, addr, client)

	// no balance in admin wallet
	if isBalance == false {
		return false,err,"No balance in admin wallet",""
	}

	// admin have less than 10 tokens (not enough balance)
	if amount < 10 {
		return false,err,"Admin dont have enough tokens",""
	}

	msg := &cosmosBankTypes.MsgSend{
			FromAddress: addr,
			ToAddress: toAddress,
			Amount: cosmosTypes.NewCoins(cosmosTypes.NewInt64Coin("token", 1)),
	}

	// http://0.0.0.0:1317/cosmos/tx/v1beta1/txs/{transaction_hash}
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		return false,err,"error in transaction",""
	}

	return true,nil,"Success",txResp.TxHash
}
