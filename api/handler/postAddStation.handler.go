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

// HandlePostAPI handles the POST request
func HandlePostAddStation(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyAddStation
	if err := c.BindJSON(&requestBody); err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	verificationKey := requestBody.VerificationKey
	stationInfo := requestBody.StationInfo
	// fmt.Println(verificationKey)
	// fmt.Println(stationInfo)

	// Validate the verification_key (add your validation logic here)
	if len(verificationKey) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	// Validate the chain_info (add your validation logic here)
	if len(stationInfo) == 0 {
		respondWithError(c, "stationInfo cannot be empty")
		return
	}

	// get execution layer by address
	success, chainId := chain.GetStationByAddress(addr, sAPI)
	if success {
		respondWithSuccess(c, chainId, "chain already exists with this address")
		return
	}

	success, data, error_msg := chain.AddStation(verificationKey, stationInfo, client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	respondWithSuccess(c, data, "add execution layer successfully")
	return
}
