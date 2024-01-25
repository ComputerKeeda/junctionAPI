package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/ComputerKeeda/junctionAPI/model"
	fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func AddPod(podDetails model.RequestBodyAddPod, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	var witness fr.Vector

	stationID := podDetails.StationId
	podNumber := podDetails.PodNumber
	previousMerkleRootHash := podDetails.PreviousMerkleRootHash
	merkleRootHash := podDetails.MerkleRootHash
	timestamp := podDetails.Timestamp
	publicWitnessByte := podDetails.PublicWitness

	err := json.Unmarshal(publicWitnessByte, &witness)
	if err != nil {
		return false, "error in unmarshalling public witness bytes", err.Error()
	}

	//check if pod is already submitted
	//get the latest submitted pod number
	success, latestFetchedPodNumber := GetLatestSubmittedPod(sAPI, stationID)
	if !success {
		return false, "error in getting latest submitted pod number", "nil"
	}
	// check if latestPodNumber is equal to podNumber
	if podNumber != latestFetchedPodNumber+1 {
		return false, "pod already submitted", "nil"
	}

	msg := &types.MsgSubmitPod{
		Creator:                addr,
		StationId:              stationID,
		PodNumber:              podNumber,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		PublicWitness:          publicWitnessByte,
		Timestamp:              timestamp,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		errMsg := formatErrorMessage(err)
		return false, "error in transaction", errMsg
	}

	// get the latest submitted pod number
	success, latestFetchedPodNumber = GetLatestSubmittedPod(sAPI, stationID)
	if !success {
		return false, "error in getting latest submitted pod number after successful transaction", "nil"
	}
	// check if latestPodNumber is equal to podNumber
	if podNumber != latestFetchedPodNumber {
		return false, "pod submission failed, but transaction success", "nil"
	}

	fmt.Println("txHash:", txResp.TxHash)
	fmt.Println("pod number:", podNumber)

	data = txResp.TxHash
	return true, data, "nil"
}
