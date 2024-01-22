package handler

/*
Sample POST: http://localhost:8080/addexelayer

Sample Request Body:
{
    "verification_key": "the long long verification_key",
    "chain_info": "some information about the chain, e.g. its made for DeFi"
}

Sample Response Body:
{
    "status": true,
    "data": "{\"txDetails\":\"code: 0\ncodespace: \"\"\ndata: 12330A312F616972736574746C652E616972736574746C652E4D7367416464457865637574696F6E4C61796572526573706F6E7365\nevents:\n- attributes:\n  - index: true\n    key: fee\n    value: \"\"\n  - index: true\n    key: fee_payer\n    value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\n  type: tx\n- attributes:\n  - index: true\n    key: acc_seq\n    value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t/16\n  type: tx\n- attributes:\n  - index: true\n    key: signature\n    value: KsMf/BgjpmcZ5ZE60S6q4UMhLHOGhWuuBcHnqRzLcn0bcn6IpjWpgTaWlnT5/tcmA6fYzDaflP7GfryoQ/L/GA==\n  type: tx\n- attributes:\n  - index: true\n    key: action\n    value: /airsettle.airsettle.MsgAddExecutionLayer\n  - index: true\n    key: sender\n    value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\n  - index: true\n    key: module\n    value: airsettle\n  type: message\ngas_used: \"52115\"\ngas_wanted: \"300000\"\nheight: \"9829\"\ninfo: \"\"\nlogs:\n- events:\n  - attributes:\n    - key: action\n      value: /airsettle.airsettle.MsgAddExecutionLayer\n    - key: sender\n      value: air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\n    - key: module\n      value: airsettle\n    type: message\n  log: \"\"\n  msg_index: 0\nraw_log: '[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/airsettle.airsettle.MsgAddExecutionLayer\"},{\"key\":\"sender\",\"value\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"},{\"key\":\"module\",\"value\":\"airsettle\"}]}]}]'\ntimestamp: \"\"\ntx: null\ntxhash: 1ED3B216F212847DF010B4CD93026A7FD93CBAE4E7DED3A03A57FED272E7B920\n\",\"chainDetails\":\"{\"exelayer\":{\"validator\":[\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"],\"votingPower\":[\"100\"],\"latestBatch\":\"0\",\"latestMerkleRootHash\":\"0\",\"verificationKey\":\"/verificationKey/f0722463-03f1-485e-8d91-f592cad02d23/\",\"chainInfo\":\"some information about the chain, e.g. its made for DeFi\",\"id\":\"f0722463-03f1-485e-8d91-f592cad02d23\",\"creator\":\"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t\"}}\"}",
    "description": "add execution layer successfully"
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
func HandlePostAddExecutionLayer(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyAddExecutionLayer
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	verificationKey := requestBody.VerificationKey
	chainInfo := requestBody.ChainInfo

	// Validate the verification_key (add your validation logic here)
	if len(verificationKey) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	// Validate the chain_info (add your validation logic here)
	if len(chainInfo) == 0 {
		respondWithError(c, "ChainInfo cannot be empty")
		return
	}

	// get execution layer by address
	success, _ := chain.GetExecutionLayerByAddress(addr, sAPI)
	if success {
		respondWithSuccess(c, "exist", "chain already exists with this address")
		return
	}

	success, data, error_msg := chain.AddExecutionLayer(verificationKey, chainInfo, client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	respondWithSuccess(c, data, "add execution layer successfully")
	return
}
