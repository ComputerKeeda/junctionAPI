package handler

/*
Sample GET:
http://localhost:8080/get_all_exelayer

Sample Request Body: (no inputs required)

Sample Response Body:
{
    "status": true,
    "data": "{\"exelayer\":[{\"validator\":[\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"],\"votingPower\":[\"100\"],\"latestBatch\":\"1\",\"latestMerkleRootHash\":\"0xMerkleRoot\",\"verificationKey\":\"/verificationKey/f0722463-03f1-485e-8d91-f592cad02d23/\",\"chainInfo\":\"some information about the chain, e.g. its made for DeFi\",\"id\":\"f0722463-03f1-485e-8d91-f592cad02d23\",\"creator\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"}],\"pagination\":{\"next_key\":null,\"total\":\"1\"}}",
    "description": "get execution layer successfully"
}
*/

import (
	"context"

	"github.com/airchains-network/settlement_layer_calls_api/chain"
	"github.com/gin-gonic/gin"

	// "github.com/airchains-network/settlement_layer_calls_api/model"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandleGetAllExecutionLayers(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	success, chainDetails := chain.GetAllExecutionLayers(sAPI)
	if success {
		respondWithSuccess(c, chainDetails, "get execution layer successfully")
		return
	} else {
		respondWithError(c, "error in getting execution layers list")
		return
	}

}
