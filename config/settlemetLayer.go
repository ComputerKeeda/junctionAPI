package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ComputerKeeda/junctionAPI/chain"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/joho/godotenv"
)

func SettlementLayer() (client cosmosclient.Client, account cosmosaccount.Account, addr string, ctx context.Context, sAPI string) {

	// connect to blockchain (settlement layer)
	accountName := "admin"
	accountPath := "./accounts"
	ctx = context.Background()
	gasLimit := "70000000"
	addressPrefix := "air" // "cosmos"

	// sRPC := "https://devnet-jrpc.airchains.network:443" // tendermint
	// sAPI = "https://devnet.airchains.network"           // Blockchain API
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sRPC := os.Getenv("SETTLEMENT_RPC")
	sAPI := os.Getenv("SETTLEMENT_API")

	client, err := cosmosclient.New(ctx, cosmosclient.WithGas(gasLimit), cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(sRPC), cosmosclient.WithKeyringDir(accountPath))
	if err != nil {
		panic(err)
	}

	// check if admin account exists
	isAccountExists, _ := chain.CheckIfAccountExists(accountName, client, addressPrefix, accountPath)

	// create admin account if don't exists
	if !isAccountExists {
		chain.CreateAccount(accountName, accountPath)

		// now recheck, its impossible its not exists. if error occure it will panic and stop.
		isAccountExists, _ = chain.CheckIfAccountExists(accountName, client, addressPrefix, accountPath)
	}

	account, err = client.Account(accountName)
	if err != nil {
		panic(err)
	}

	addr, err = account.Address(addressPrefix)
	if err != nil {
		panic(err)
	}
	fmt.Println("admin address:", addr)

	return client, account, addr, ctx, sAPI
}
