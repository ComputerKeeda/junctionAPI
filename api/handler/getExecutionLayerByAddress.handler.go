package handler

/*
Sample GET:
http://localhost:8080/getexelayer_by_address

Sample Request Body:
{
    "address": "air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"
}

Sample Response Body:
{
    "status": true,
    "data": "{\"exelayer\":{\"validator\":[\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"],\"votingPower\":[\"100\"],\"latestBatch\":\"0\",\"latestMerkleRootHash\":\"0\",\"verificationKey\":\"/verificationKey/f0722463-03f1-485e-8d91-f592cad02d23/\",\"chainInfo\":\"some information about the chain, e.g. its made for DeFi\",\"id\":\"f0722463-03f1-485e-8d91-f592cad02d23\",\"creator\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"}}",
    "description": "get execution layer successfully"
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
func HandleGetExecutionLayerByAddress(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyGetExecutionLayerByAddress
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	// check if address is empty
	Address := requestBody.Address // "air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"
	if len(Address) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	success, chainDetails := chain.GetExecutionLayerByAddress(Address, sAPI)
	if success {
		respondWithSuccess(c, chainDetails, "get execution layer successfully")
		return
	} else {
		respondWithError(c, "chain id not found in this address")
		return
	}

}
