package handler

import (
	"context"
	"fmt"
	"github.com/ComputerKeeda/junctionAPI/chain"
	"github.com/ComputerKeeda/junctionAPI/model"
	"github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandlePostAddPod(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {
	// Parse the request body into a struct
	var requestBody model.RequestBodyAddPod
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	podNumber := requestBody.PodNumber
	stationID := requestBody.StationId
	publicWitness := requestBody.PublicWitness
	timeStamp := requestBody.Timestamp
	pMrh := requestBody.PreviousMerkleRootHash
	mrh := requestBody.MerkleRootHash

	// print all: for testing purpose
	// fmt.Println("Pod Number: ", podNumber)
	// fmt.Println("Station ID: ", stationID)
	// fmt.Println("Public Witness: ", publicWitness)
	// fmt.Println("Timestamp: ", timeStamp)
	// fmt.Println("Previous Merkle Root Hash: ", pMrh)
	// fmt.Println("Merkle Root Hash: ", mrh)

	// batchNumber, stationID, and witness length can not be 0
	if podNumber < 1 || len(stationID) == 0 || len(string(publicWitness)) == 0 || len(pMrh) == 0 || len(mrh) == 0 || timeStamp < 1 {
		respondWithError(c, "Pod Number, stationID, witness, previous merkle root hash, merkle root hash, and timestamp can not be empty")
		return
	}

	success, data, errMsg := chain.AddPod(requestBody, client, ctx, account, addr, sAPI)

	fmt.Println(success, data, errMsg)

	if !success {
		respondWithError(c, errMsg)
		return
	}

	respondWithSuccess(c, data, "batch add successfully")
	return
}
