package admin

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/airchains-network/settlement_layer_calls_api/chain"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

func AdminBalanceCheckerTimer(wg *sync.WaitGroup, ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, addr string, db *leveldb.DB) {
	defer wg.Done()
	_ = db

	// check admin balance every minute
	for {
		// check client connection
		_, err := client.Status(ctx)
		if err != nil {
			InformAdmin("error", "Blockchain is offline")
		}

		// check if admin have enough balance to send tokens
		isBalance, amount, _ := chain.CheckBalance(ctx, addr, client)

		// no balance in admin wallet
		if isBalance == false {
			InformAdmin("error", "No balance in execution layer faucet/admin wallet")
		}

		// admin have less than 10 tokens (not enough balance)
		amountStr := strconv.FormatInt(amount, 10)
		if amount < 10 {
			InformAdmin("error", "Admin dont have enough tokens: "+amountStr)
		} else {
			// inform admin
			InformAdmin("information", "Admin have Currently "+amountStr+" tokens")
		}

		// wait 5 minutes
		time.Sleep(5 * 60 * time.Second)
	}
}

func InformAdmin(msgtype string, message string) {
	switch msgtype {
	case "error":
		fmt.Println("Error: ", message)
	case "information":
		fmt.Println("Information: ", message)
	default: // default not possible in out case, but still just for safety
		fmt.Println("Error: ", message)
	}
}
