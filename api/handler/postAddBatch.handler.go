package handler

/*
Sample POST:
http://localhost:8080/add_batch

Sample Request Body:
{
    "chain_id": "f0722463-03f1-485e-8d91-f592cad02d23",
    "batch_number":1,
    "witness":"[witness numberss]"
}

Sample Response Body:
{clear

    "status": true,
    "data": "{\"txDetails\":\"code: 0\ncodespace: \"\"\ndata: 122A0A282F616972736574746C652E616972736574746C652E4D73674164644261746368526573706F6E7365\nevents:\n- attributes:\n  - index: true\n    key: fee\n    value: \"\"\n  - index: true\n    key: fee_payer\n    value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\n  type: tx\n- attributes:\n  - index: true\n    key: acc_seq\n    value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t/18\n  type: tx\n- attributes:\n  - index: true\n    key: signature\n    value: IYtSBKyUsthXCAeXWl9XPW4lvVQAa2m+8TPqSZRxUxVpHutnTJ4lQGr+aGJI/5FyPsXP1wQsA4RnP+ZPALtM5g==\n  type: tx\n- attributes:\n  - index: true\n    key: action\n    value: /airsettle.airsettle.MsgAddBatch\n  - index: true\n    key: sender\n    value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\n  - index: true\n    key: module\n    value: airsettle\n  type: message\ngas_used: \"32543\"\ngas_wanted: \"300000\"\nheight: \"14185\"\ninfo: \"\"\nlogs:\n- events:\n  - attributes:\n    - key: action\n      value: /airsettle.airsettle.MsgAddBatch\n    - key: sender\n      value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\n    - key: module\n      value: airsettle\n    type: message\n  log: \"\"\n  msg_index: 0\nraw_log: '[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/airsettle.airsettle.MsgAddBatch\"},{\"key\":\"sender\",\"value\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"},{\"key\":\"module\",\"value\":\"airsettle\"}]}]}]'\ntimestamp: \"\"\ntx: null\ntxhash: ADF9A6C353FEF42BD2A9B6F85BFAD6C296C189EF23D506B05E3819C23F4C3483\n\"}",
    "description": "add execution layer successfully"
}
*/

import (
	"context"
	"fmt"
	"time"

	"github.com/airchains-network/settlement_layer_calls_api/chain"
	"github.com/airchains-network/settlement_layer_calls_api/model"
	"github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandlePostAddBatch(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {
	// Parse the request body into a struct
	var requestBody model.RequestBodyAddBatch
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	batchNumber := requestBody.BatchNumber
	chainId := requestBody.ChainId
	witness := requestBody.Witness
	fmt.Println(time.Now().Unix())
	fmt.Println(batchNumber)
	fmt.Println(chainId)
	// fmt.Println(witness)

	// batchNumber, chainId, and witness length can not be 0
	if batchNumber == 0 || len(chainId) == 0 || len(string(witness)) == 0 {
		respondWithError(c, "BatchNumber, ChainId, and Witness cannot be empty")
		return
	}

	success, data, error_msg := chain.AddBatch(batchNumber, chainId, witness, client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	fmt.Println(time.Now().Unix())
	respondWithSuccess(c, data, "batch add successfully")
	return
}
