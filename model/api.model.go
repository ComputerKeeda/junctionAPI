package model

// ResponseBody is the structure for the JSON response
type ResponseBody struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
	Description string `json:"description"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyGetExecutionLayerByAddress struct {
	Address string `json:"address"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyGetExecutionLayerById struct {
	ChainId string `json:"chain_id"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyGetVerificationKeyById struct {
	ChainId string `json:"chain_id"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyAddExecutionLayer struct {
	VerificationKey string `json:"verification_key"`
	ChainInfo       string `json:"chain_info"`
}

// RequestBodyAddBatch
type RequestBodyAddBatch struct {
	ChainId string `json:"chain_id"`
	BatchNumber uint64 `json:"batch_number"`
	Witness []byte `json:"witness"`
}


type RequestBodyVerifyBatch struct {
	BatchNumber    uint64 `json:"batch_number"`
	ChainId        string `json:"chain_id"`
	MerkleRootHash string `json:"merkle_root_hash"`
	PrevMerkleRoot string `json:"prev_merkle_root"`
	ZkProof        []byte `json:"zk_proof"`
}

type RequestBodyGetBatch struct {
	BatchNumber    uint64 `json:"batch_number"`
	ChainId        string `json:"chain_id"`
}