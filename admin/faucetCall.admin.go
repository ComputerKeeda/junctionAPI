package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RequestFaucet(address string) (err error) {
	url := "https://biqmfxzlbn.airchains.co/faucet-api/faucet"
	data := map[string]string{"address": address, "denom": "dair"}

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	// Make POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", resp.StatusCode)
		return err
	}
	return nil

}
