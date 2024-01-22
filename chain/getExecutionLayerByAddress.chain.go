package chain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/airchains-network/settlement_layer_calls_api/model"
)

func GetExecutionLayerByAddress(address string, sAPI string) (success bool, chainId string) {

	apiURL := sAPI + "/airchains-network/airsettle/airsettle/show_execution_layer_by_address/" + address

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
	var executionLayerResponse model.ExecutionLayerTrueResponseBody
	if err := json.Unmarshal(body, &executionLayerResponse); err == nil {
		if len(executionLayerResponse.ExeLayer.ID) == 0 {
			return false, ""
		} else {
			return true, executionLayerResponse.ExeLayer.ID
		}
	}

	// code may not reach here... but just in case
	var executionLayerErrResponse model.ExecutionLayerErrorResponseBody
	if err := json.Unmarshal(body, &executionLayerErrResponse); err == nil {
		// Successfully unmarshaled into ExecutionLayerErrorResponseBody
		return false, ""
	}

	// if not both data type
	return false, ""
}
