package handler

/*
Sample GET:
http://localhost:8080/delete_exelayer

Sample Request Body: no inputs required

Sample Response Body:
{
    "status": false,
    "data": "",
    "description": "Cannot delete a chain with batch number greater than 0"
}
*/

import (
	"context"

	"github.com/airchains-network/settlement_layer_calls_api/chain"
	"github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandlePostDeleteExecutionLayer(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	success, data, error_msg := chain.DeleteExecutionLayer(client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	respondWithSuccess(c, data, "delete execution layer successfully")
	return
}
