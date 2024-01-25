package chain

import (
	"context"
	"fmt"

	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/google/uuid"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func AddStation(verificationKey []byte, stationInfo string, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	newSatationId := uuid.New().String()

	msg := &types.MsgInitStation{
		Creator:         addr,
		VerificationKey: verificationKey,
		StationInfo:     stationInfo,
		StationId:       newSatationId,
	}

	res, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		fmt.Println(res.TxHash)

		error_msg := formatErrorMessage(err)
		return false, "error in transaction", error_msg
	}
	// fmt.Println(res.TxHash)

	// get station ID by address
	success, chainId := GetStationByAddress(addr, sAPI)
	if !success {
		return false, "", "error in receiving station ID"
	}
	fmt.Println(chainId)

	if chainId != newSatationId {
		// check if chainID and newUUID are same
		return false, "", "error in receiving station ID"
	}

	// print tx hash
	fmt.Println(res.TxHash)

	return true, newSatationId, "nil"
}
