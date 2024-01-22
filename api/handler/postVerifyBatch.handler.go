package handler

/*
Sample Post:
http://localhost:8080/verify_batch

Sample Request Body:
{
    "batch_number":1,
    "chain_id": "f0722463-03f1-485e-8d91-f592cad02d23",
    "merkle_root_hash" : "0xMerkleRootHash" ,
    "prev_merkle_root" : "0",
    "zk_proof" : "zkproof"
}

Sample Response Body:
{
    "status": false,
    "data": "",
    "description": "error code: '18' msg: 'failed to execute message; message index: 0: invalid request'"
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
func HandlePostVerifyBatch(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyVerifyBatch
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	batchNumber := requestBody.BatchNumber
	chainId := requestBody.ChainId
	merkleRootHash := requestBody.MerkleRootHash
	prevMerkleRoot := requestBody.PrevMerkleRoot
	zkProof := requestBody.ZkProof

	// data can not be empty
	if batchNumber == 0 || chainId == "" || merkleRootHash == "" || prevMerkleRoot == "" || string(zkProof) == "" {
		respondWithError(c, "Invalid JSON format")
		return
	}

	success, data, error_msg := chain.VerifyBatch(batchNumber, chainId, merkleRootHash, prevMerkleRoot, zkProof, client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	respondWithSuccess(c, data, "verify batch successfully")
	return
}
