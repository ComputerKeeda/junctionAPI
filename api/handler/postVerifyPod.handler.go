package handler

import (
	"context"
	"github.com/ComputerKeeda/junctionAPI/chain"
	"github.com/ComputerKeeda/junctionAPI/model"
	"github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

func HandlePostVerifyPod(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyVerifyPod
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	stationId := requestBody.StationId
	podNumber := requestBody.PodNumber
	merkleRootHash := requestBody.MerkleRootHash
	previousMerkleRootHash := requestBody.PreviousMerkleRootHash
	zkProof := requestBody.ZkProof

	// print all: for testing purpose
	// fmt.Println("Station ID: ", stationId)
	// fmt.Println("Pod Number: ", podNumber)
	// fmt.Println("Merkle Root Hash: ", merkleRootHash)
	// fmt.Println("Previous Merkle Root Hash: ", previousMerkleRootHash)
	// fmt.Println("ZK Proof: ", zkProof)

	if len(stationId) == 0 || podNumber < 1 || len(merkleRootHash) == 0 || len(previousMerkleRootHash) == 0 || len(string(zkProof)) == 0 {
		respondWithError(c, "stationId, podNumber, merkleRootHash, previousMerkleRootHash, and zkProof can not be empty")
		return
	}

	success, data, error_msg := chain.VerifyPod(requestBody, client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	respondWithSuccess(c, data, "verify batch successfully")
	return
}
