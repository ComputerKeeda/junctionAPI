package handler

/*
Sample GET:
http://localhost:8080/get_batch

Sample Request Body:
{
    "batch_number":1,
    "chain_id": "f0722463-03f1-485e-8d91-f592cad02d23"
}

Sample Response Body:
{
    "status": true,
    "data": "{\"batch\":{\"batchNumber\":\"1\",\"chainId\":\"f0722463-03f1-485e-8d91-f592cad02d23\",\"prevMerkleRootHash\":\"\",\"merkleRootHash\":\"0xMerkleRoot\",\"zkProof\":\"zkproof\",\"inputs\":[\"input1\",\"input2\"],\"verified\":\"true\",\"batchSubmitter\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\",\"batchVerifier\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"}}",
    "description": "get batch details successfully"
}
*/

import (
	"context"

	"github.com/airchains-network/settlement_layer_calls_api/chain"
	"github.com/airchains-network/settlement_layer_calls_api/model"
	"github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandleGetBatch(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyGetBatch
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	// check if chain id  is empty
	chainId := requestBody.ChainId
	batchNumber := requestBody.BatchNumber

	// data can not be empty
	if len(chainId) == 0 || batchNumber == 0 {
		respondWithError(c, "Invalid JSON format")
		return
	}

	success, batchDetils := chain.GetBatch(chainId, batchNumber, sAPI)
	if success {
		respondWithSuccess(c, batchDetils, "get batch details successfully")
		return
	} else {
		respondWithError(c, "error in getting batch details")
		return
	}

}
