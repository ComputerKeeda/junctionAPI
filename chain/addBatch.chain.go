package chain

import (
	"context"
	"encoding/json"
	fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func AddBatch(batchNumber uint64, chainId string, witness_byte []byte, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	var witness fr.Vector
 
	err := json.Unmarshal(witness_byte, &witness)
	if err != nil {
		panic(err)
	}

	var stringOfWitness string
	stringOfWitness = "["
	for i := 0; i < len(witness); i++ {
		if i == len(witness)-1 {
			stringOfWitness += "\"" + witness[i].String() + "\""
		} else {
			stringOfWitness += "\"" + witness[i].String() + "\","
		}
	}
	stringOfWitness += "]"

	msg := &types.MsgAddBatch{
		Creator:     addr,
		BatchNumber: batchNumber,
		ChainId:     chainId,
		Witness:     stringOfWitness,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false, "error in transaction", error_msg
	}

	data = txResp.TxHash // fmt.Sprintf(`{"txDetails":"%s"}`, txResp.TxHash)
	return true, data, "nil"
}
