package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ComputerKeeda/junctionAPI/admin"
	"github.com/ComputerKeeda/junctionAPI/api"
	"github.com/ComputerKeeda/junctionAPI/chain"
	"github.com/ComputerKeeda/junctionAPI/config"
)

func main() {

	// connect to junction && create admin wallet if do not exists
	client, account, addr, ctx, sAPI := config.JunctionConnection()

	// connect to levelDB
	dbIPaddress := config.LevelDB()

	// check admin balance
	isBalance, amount, err := chain.CheckBalance(ctx, addr, client)
	if err != nil {
		fmt.Println("Error in checking balance", err)
		return
	}

	// call faucet (if balance is 0)
	if amount < 3 || !isBalance {
		fmt.Println("admin currently don't have faucet, requesting faucet from this address:", addr)

		// calling faucet api
		err = admin.RequestFaucet(addr)
		if err != nil {
			fmt.Println("Error in calling faucet api", err)
			return
		}

		fmt.Println("Faucet request successful!")

		// check admin balance
		_, amount, err := chain.CheckBalance(ctx, addr, client)
		if err != nil {
			fmt.Println("Error in checking balance", err)
			return
		}
		fmt.Println("Admin have", amount, "tokens \nstarting api...")

		// await 2 seconds
		time.Sleep(2 * time.Second)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go api.StartAPI(&wg, client, ctx, account, addr, dbIPaddress, sAPI)
	go admin.AdminBalanceCheckerTimer(&wg, ctx, client, account, addr, dbIPaddress)

	// Wait for both functions to finish
	wg.Wait()

	// wg Crashed. send report to admin
	fmt.Println("wg Crashed: faucet api stopped")
}
