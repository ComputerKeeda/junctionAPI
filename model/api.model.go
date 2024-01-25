package model

// ResponseBody is the structure for the JSON response
type ResponseBody struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
	Description string `json:"description"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyGetStationByAddress struct {
	Address string `json:"address"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyGetStationById struct {
	ChainId string `json:"chain_id"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyGetVerificationKeyById struct {
	ChainId string `json:"chain_id"`
}

// RequestBody is the structure for the incoming JSON request
type RequestBodyAddStation struct {
	VerificationKey []byte `json:"verification_key"`
	StationInfo     string `json:"chain_info"`
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