package chain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

// chain.VerifyBatch(batchNumber, chainId, merkleRootHash, prevMerkleRoot, zkProof, client, ctx, account, addr, sAPI)

func VerifyBatch(batchNumber uint64, chainId string, merkleRootHash string, prevMerkleRoot string, proof_byte []byte, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	var proof *bls12381.Proof
	err := json.Unmarshal(proof_byte, &proof)
	if err != nil {
		return false, "error in Unmarshal proof", err.Error()
	}

	/*
			Proof formate
			{
		    	"Ar": {
		    	    "X": "495870166108798604731145358879824227970170045696996199523236680845476654815869684153408461772683268328332709022310",
		    	    "Y": "2199802072378822310073473713870395975565027856969244190053942673480459904522913862113008044348169493922449712975913"
		    	},
		    	"Krs": {
		    	    "X": "539750921760002630548813613206526862540204475257961757993588672436949520379398195286762180074471491363466784634742",
		    	    "Y": "3459123025362292051697880916654237811698644064117443475262676468786174999765244816088282059614498427524465200172523"
		    	},
		    	"Bs": {
		    	    "X": {
		    	        "A0": "2709896886096357918498787920332521367932027196967051133316872948604035726991137766264967008075317076798233104472765",
		    	        "A1": "2683505342519834873175921422613102645188984122619400055933487372311374644826745894396833879226591465565692940425610"
		    	    },
		    	    "Y": {
		    	        "A0": "3205805920896809604597172693164176535464114471543850094119068022375287384195748478895666038570566265331949712051106",
		    	        "A1": "23319190191518536297419377591197501597626416189614389834773981705258609877047403996185938875319708205727773504960"
		    	    }
		    	},
		    	"Commitments": [],
		    	"CommitmentPok": {
		    	    "X": 0,
		    	    "Y": 0
		    	}
			}
			required formate:
			 "{\"Ar\":{\"X\":\"495870166108798604731145358879824227970170045696996199523236680845476654815869684153408461772683268328332709022310\",\"Y\":\"2199802072378822310073473713870395975565027856969244190053942673480459904522913862113008044348169493922449712975913\"},\"Krs\":{\"X\":\"539750921760002630548813613206526862540204475257961757993588672436949520379398195286762180074471491363466784634742\",\"Y\":\"3459123025362292051697880916654237811698644064117443475262676468786174999765244816088282059614498427524465200172523\"},\"Bs\":{\"X\":{\"A0\":\"2709896886096357918498787920332521367932027196967051133316872948604035726991137766264967008075317076798233104472765\",\"A1\":\"2683505342519834873175921422613102645188984122619400055933487372311374644826745894396833879226591465565692940425610\"},\"Y\":{\"A0\":\"3205805920896809604597172693164176535464114471543850094119068022375287384195748478895666038570566265331949712051106\",\"A1\":\"23319190191518536297419377591197501597626416189614389834773981705258609877047403996185938875319708205727773504960\"}},\"Commitments\":[],\"CommitmentPok\":{\"X\":0,\"Y\":0}}"
	*/

	commitmentPokString := "["

	currentcommitmentPok := proof.Commitments

	for i := 0; i < len(currentcommitmentPok); i++ {
		commitmentPokString = commitmentPokString + "\"" + currentcommitmentPok[i].String() + "\""
		if i != len(currentcommitmentPok)-1 {
			commitmentPokString = commitmentPokString + ","
		}
	}

	commitmentPokString = commitmentPokString + "]"

	proof_string := `{\"Ar\": {\"X\": \"` + proof.Ar.X.String() + `\",\"Y\": \"` + proof.Ar.Y.String() + `\",},\"Krs\": {\"X\": \"` + proof.Krs.X.String() + `\",\"Y\": \"` + proof.Krs.Y.String() + `\",},\"Bs\": {\"X\": {\"A0\": \"` + proof.Bs.X.A0.String() + `\",\"A1\": \"` + proof.Bs.X.A1.String() + `\"},\"Y\": {\"A0\": \"` + proof.Bs.Y.A0.String() + `\",\"A1\": \"` + proof.Bs.Y.A1.String() + `\"}},\"Commitments\": ` + commitmentPokString + `,\"CommitmentPok\": {\"X\": ` + proof.CommitmentPok.X.String() + `,\"Y\": ` + proof.CommitmentPok.Y.String() + `}}`

	msg := &types.MsgVerifyBatch{
		Creator:        addr,
		BatchNumber:    batchNumber,
		ChainId:        chainId,
		MerkleRootHash: merkleRootHash,
		PrevMerkleRoot: prevMerkleRoot,
		ZkProof:        proof_string,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false, "error in transaction", error_msg
	}

	data = fmt.Sprintf(`{"txDetails":"%s"}`, txResp)
	return true, data, "verify batch successfully"
}
