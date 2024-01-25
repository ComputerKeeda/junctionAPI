package chain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/ComputerKeeda/junctionAPI/model"
)

func GetStationByAddress(address string, sAPI string) (success bool, chainId string) {

	apiURL := sAPI + "/ComputerKeeda/junction/junction/get_station_details_by_address/" + address

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
	var stationResponse model.StationResponseBody
	if err := json.Unmarshal(body, &stationResponse); err == nil {
		if len(stationResponse.Station.Id) == 0 {
			return false, ""
		} else {
			return true, stationResponse.Station.Id
		}
	}

	// code may not reach here... but just in case
	var stationErrResponse model.StationErrorResponseBody
	if err := json.Unmarshal(body, &stationErrResponse); err == nil {
		return false, ""
	}

	// if not both data type
	return false, ""

}
