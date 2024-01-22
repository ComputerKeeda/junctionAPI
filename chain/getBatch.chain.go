package chain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/airchains-network/settlement_layer_calls_api/model"
)

func GetBatch(chainId string, batchNumber uint64, sAPI string) (success bool, data string) {

	strBatchNumber := fmt.Sprint(batchNumber)
	apiURL := sAPI + "/airchains-network/airsettle/airsettle/get_batch/" + strBatchNumber + "/" + chainId

	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return false, ""
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, "Error in Requesting to Execution Layer Blockchain API"
	}

	// Check the structure of the response body to determine the appropriate struct
	var batchResponse model.BatchResponseBody
	if err := json.Unmarshal(body, &batchResponse); err == nil {
		if len(batchResponse.Batch.ChainId) == 0 { // check anyone data
			return false, ""
		} else {
			return true, string(body)
		}
	}

	// if not both data type
	return false, ""
}
