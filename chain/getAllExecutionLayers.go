package chain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/airchains-network/settlement_layer_calls_api/model"
)

func GetAllExecutionLayers(sAPI string) (success bool, chainId string) {

	apiURL := sAPI + "/airchains-network/airsettle/airsettle/list_all_execution_layers"

	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return false, "Blockchain API Error"
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, "Execution Layers Not Found"
	}

	// Check the structure of the response body to determine the appropriate struct
	var allExecutionLayerResponse model.AllExecutionLayersResponseBody
	if err := json.Unmarshal(body, &allExecutionLayerResponse); err == nil {
		if len(allExecutionLayerResponse.ExeLayer) == 0 {
			return false, "Execution Layers Not Found"
		} else {
			return true, string(body)
		}
	}

	// almost impossible case, but still
	return false, ""

}
