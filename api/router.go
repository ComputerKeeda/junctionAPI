package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/airchains-network/settlement_layer_calls_api/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/joho/godotenv"
	"github.com/syndtr/goleveldb/leveldb"
)

func StartAPI(wg *sync.WaitGroup, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, db *leveldb.DB, sAPI string) {
	defer wg.Done()

	// Create a new router
	router := gin.Default()

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the PORT value
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	// * Register the Handlers / Routers

	// add execution layer
	router.POST("/addexelayer", func(c *gin.Context) {
		handler.HandlePostAddExecutionLayer(c, client, ctx, account, addr, db, sAPI)
	})

	// get execution layer by address
	// router.GET("/get_admin_exelayer", func(c *gin.Context) {
	// 	handler.HandleGetAdminExecutionLayer(c, client, ctx, account, addr, db, sAPI)
	// })

	// get execution layer by address
	router.GET("/getexelayer_by_address", func(c *gin.Context) {
		handler.HandleGetExecutionLayerByAddress(c, client, ctx, account, addr, db, sAPI)
	})

	// get execution layer by chain id
	router.GET("/getexelayer_by_id", func(c *gin.Context) {
		handler.HandleGetExecutionLayerById(c, client, ctx, account, addr, db, sAPI)
	})

	// get all execution layers
	router.GET("/get_all_exelayer", func(c *gin.Context) {
		handler.HandleGetAllExecutionLayers(c, client, ctx, account, addr, db, sAPI)
	})

	// get verification key
	router.GET("/get_vkey", func(c *gin.Context) {
		handler.HandleGetVerificationKeyById(c, client, ctx, account, addr, db, sAPI)
	})

	// delete execution layer
	router.POST("/delete_exelayer", func(c *gin.Context) {
		handler.HandlePostDeleteExecutionLayer(c, client, ctx, account, addr, db, sAPI)
	})

	// add batch
	router.POST("/add_batch", func(c *gin.Context) {
		fmt.Println("add_batch called")
		handler.HandlePostAddBatch(c, client, ctx, account, addr, db, sAPI)
	})

	// verify batch
	router.POST("/verify_batch", func(c *gin.Context) {
		handler.HandlePostVerifyBatch(c, client, ctx, account, addr, db, sAPI)
	})

	// get batch
	router.GET("/get_batch", func(c *gin.Context) {
		handler.HandleGetBatch(c, client, ctx, account, addr, db, sAPI)
	})

	// * future updates for multinode sequencer.
	// // add validator  $chainid
	// router.POST("/add_validator", func(c *gin.Context) {
	// 	handler.HandlePostAddValidator(c,client,ctx, account,addr,db, sAPI)
	// })
	// // list-polls $chainid
	// router.GET("/get_polls_list", func(c *gin.Context) {
	// 	handler.HandleGetPollsList(c,client,ctx, account,addr,db, sAPI)
	// })
	// // get-poll $chainid $pollid
	// router.GET("/get_poll", func(c *gin.Context) {
	// 	handler.HandleGetPoll(c,client,ctx, account,addr,db, sAPI)
	// })
	// // vote-poll $chainid $pollid $vote
	// router.POST("/vote-poll", func(c *gin.Context) {
	// 	handler.HandlePostVotePoll(c,client,ctx, account,addr,db, sAPI)
	// })

	// Run the server on given port
	if err := router.Run(port); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
