package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/ComputerKeeda/junctionAPI/model"

	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func VerifyPod(podDetails model.RequestBodyVerifyPod, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	var proof *bls12381.Proof
	err := json.Unmarshal(proof_byte, &proof)
	if err != nil {
		return false, "error in Unmarshal proof", err.Error()
	}

	msg := &types.MsgVerifyPod{
		Creator:                addr,
		StationId:              podDetails.StationId,
		PodNumber:              podDetails.PodNumber,
		MerkleRootHash:         podDetails.MerkleRootHash,
		PreviousMerkleRootHash: podDetails.PreviousMerkleRootHash,
		ZkProof:                podDetails.ZkProof,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false, "error in transaction", error_msg
	}

	// get pod details, and check if its verified
	GetPod(podDetails.PodNumber, client, ctx, account, addr, sAPI)

	data = fmt.Sprintf(`{"txDetails":"%s"}`, txResp)
	return true, data, "verify batch successfully"
}
