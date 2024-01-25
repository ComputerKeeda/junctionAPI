package chain

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"

// 	"github.com/ComputerKeeda/junctionAPI/model"
// 	"github.com/airchains-network/airsettle/x/airsettle/types"
// 	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
// 	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
// )

// func DeleteExecutionLayer(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (success bool, data string, error_msg string) {

// 	// check if there is a chain id under this account
// 	apiURL := sAPI + "/ComputerKeeda/junction/junction/get_station_details_by_address/" + addr

// 	// Make the GET request
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		return false, "", "Error in Requesting to Junction Blockchain API"
// 	}
// 	defer response.Body.Close()

// 	// Read the response body
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return false, "", "Error in Requesting to Junction Blockchain API"
// 	}

// 	// Check the structure of the response body to determine the appropriate struct
// 	var stationResponse model.StationResponseBody
// 	err = json.Unmarshal(body, &stationResponse)

// 	if err != nil {
// 		return false, "", "Error in Requesting to Junction Blockchain API"
// 	}

// 	latestPodString := (stationResponse.Station.LatestPod)
// 	// string to uint64 conversion
// 	latestPod, err := strconv.ParseUint(latestPodString, 10, 64)
// 	if err != nil {
// 		return false, "", "Error in Requesting to Junction Blockchain API"
// 	}

// 	if latestPod > 10 {
// 		return false, "", "Cannot delete a station with pod number greater than 0"
// 	}

// 	// ! make changes from here to delete execution layer when function is created

// 	// delete the execution layer associated with this address
// 	msg := &types.MsgDeleteExecutionLayer{
// 		Creator: addr,
// 	}

// 	txResp, err := client.BroadcastTx(ctx, account, msg)
// 	if err != nil {
// 		error_msg := formatErrorMessage(err)
// 		return false, "error in transaction", error_msg
// 	}

// 	data = fmt.Sprintf(`{"txDetails":"%s"}`, txResp)
// 	return true, data, "execution layer deleted successfully"

// }
