package chain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ComputerKeeda/junctionAPI/model"
)

func GetLatestSubmittedPod(sAPI string, stationID string) (success bool, podNumber uint64) {
	apiURL := sAPI + "ComputerKeeda/junction/junction/get_latest_submitted_pod_number/" + stationID
	fmt.Println(apiURL)
	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return false, 0
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, 0
	}

	// Check the structure of the response body to determine the appropriate struct
	var podResponse model.GetLatestSubmittedPodResponseBody
	err = json.Unmarshal(body, &podResponse)
	if err != nil {
		fmt.Println("error in unmarshal pod response")
		return false, 0
	}
	podN, err := strconv.ParseUint(podResponse.PodNumber, 10, 64)
	if err != nil {
		fmt.Println("error in parsing pod number")
		return false, 0
	}
	return true, podN
}

func GetLaterstVerifiedPod(sAPI string, stationID string) (status bool, podNumber uint64) {
	apiURL := sAPI + "/ComputerKeeda/junction/junction/get_latest_verified_pod_number/" + stationID
	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return false, 0
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, 0
	}

	// Check the structure of the response body to determine the appropriate struct
	var podResponse model.GetLatestVerifiedPodResponseBody
	err = json.Unmarshal(body, &podResponse)
	if err != nil {
		fmt.Println("error in unmarshal pod response")
		return false, 0
	}
	podN, err := strconv.ParseUint(podResponse.PodNumber, 10, 64)
	if err != nil {
		fmt.Println("error in parsing pod number")
		return false, 0
	}
	
	return true, podN
}
