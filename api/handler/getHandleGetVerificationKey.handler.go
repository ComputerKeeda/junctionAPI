package handler

/*
Sample GET:
http://localhost:8080/get_vkey

Sample Request Body:
{
    "chain_id": "f0722463-03f1-485e-8d91-f592cad02d23"
}

Sample Response Body:
{
    "status": true,
    "data": "{\"vkey\":\"the long long verification_key\"}",
    "description": "get verification key successfully"
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
func HandleGetVerificationKeyById(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyGetVerificationKeyById
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	// check if chain id  is empty
	chainId := requestBody.ChainId
	if len(chainId) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	success, chainDetails := chain.GetVerificationKeyById(chainId, sAPI)

	if success {
		respondWithSuccess(c, chainDetails, "get verification key successfully")
		return
	} else {
		respondWithError(c, "Verification key not found in this chain id")
		return
	}

}
