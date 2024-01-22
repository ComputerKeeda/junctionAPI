package chain

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	// "fmt"
)

func AddExecutionLayer(verificationKey string,chainInfo string , client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string ) (status bool,data string, error string) {

	msg := &types.MsgAddExecutionLayer{
		Creator:         addr,
		VerificationKey: verificationKey,
		ChainInfo:       chainInfo,
	}

	_, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false,"error in transaction",error_msg
	}

	// get execution layer by address
	success, chainId := GetExecutionLayerByAddress(addr,sAPI) 
	if !success {
		return false,"","error in receiving execution layer"
	}

	return true, chainId, "nil" 
}