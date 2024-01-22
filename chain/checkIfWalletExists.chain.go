package chain 

import (
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"fmt"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func CheckIfAccountExists(accountName string, client cosmosclient.Client , addressPrefix string,accountPath string) (bool,string) {

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		fmt.Println(err)
		return false,""
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		fmt.Println(err)
		return false,""
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		fmt.Println("Failed to get the Address:", err)
		return false,""
	}

	return true,addr
}