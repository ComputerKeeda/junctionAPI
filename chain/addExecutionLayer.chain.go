package chain

import (
	"context"
	"fmt"

	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/google/uuid"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func AddExecutionLayer(verificationKey []byte, chainInfo string, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	newUUID := uuid.New().String()

	msg := &types.MsgInitStation{
		Creator:         addr,
		VerificationKey: verificationKey,
		StationInfo:     chainInfo,
		StationId:       newUUID,
	}

	res, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false, "error in transaction", error_msg
	}

	// // get execution layer by address
	// success, chainId := GetExecutionLayerByAddress(addr, sAPI)
	// if !success {
	// 	return false, "", "error in receiving execution layer"
	// }

	fmt.Println(res)
	fmt.Println(res.TxHash)

	return true, newUUID, "nil"
}
