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

	stationId := podDetails.StationId
	podNumber := podDetails.PodNumber
	merkleRootHash := podDetails.MerkleRootHash
	previousMerkleRootHash := podDetails.PreviousMerkleRootHash
	zkProof := podDetails.ZkProof

	var proof *bls12381.Proof
	err := json.Unmarshal(zkProof, &proof)
	if err != nil {
		return false, "error in Unmarshal proof", err.Error()
	}

	// check if the pod is already verified
	success, latestVerifiedPodNumber := GetLaterstVerifiedPod(sAPI, stationId)
	if !success {
		return false, "error in getting latest verified pod number", "nil"
	}
	// check if latestVerifiedPodNumber is equal to podNumber
	if podNumber != latestVerifiedPodNumber+1 {
		return false, "pod already verified or wrong", "nil"
	}

	// call verify pod in blockchain
	msg := &types.MsgVerifyPod{
		Creator:                addr,
		StationId:              stationId,
		PodNumber:              podNumber,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		ZkProof:                zkProof,
	}
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {

		fmt.Println(txResp.TxHash)

		errMsg := formatErrorMessage(err)
		return false, "error in transaction", errMsg
		
	}

	// get the latest verified pod number
	success, latestVerifiedPodNumber = GetLaterstVerifiedPod(sAPI, stationId)
	if !success {
		return false, "error in getting latest verified pod number after transaction is success", "nil"
	}
	// check if latestVerifiedPodNumber is equal to podNumber
	if podNumber != latestVerifiedPodNumber {
		return false, "pod not verified but transaction success", "nil"
	}
	//
	//fmt.Println("txHash:", txResp.TxHash)
	//fmt.Println("pod number:", podNumber)

	data = txResp.TxHash


	return true, data, "verify batch successfully"
}
