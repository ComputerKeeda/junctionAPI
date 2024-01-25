package model

// ResponseBody is the structure for the JSON response
type ResponseBody struct {
	Status      bool   `json:"status"`
	Data        string `json:"data"`
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

type RequestBodyAddPod struct {
	StationId              string `json:"station_id"`
	PodNumber              uint64 `json:"pod_number"`
	MerkleRootHash         string `json:"merkle_root_hash"`
	PreviousMerkleRootHash string `json:"previous_merkle_root_hash"`
	PublicWitness          []byte `json:"public_witness"`
	Timestamp              uint64 `json:"timestamp"`
}

/*
message MsgVerifyPod {
  string stationId = 2;
  uint64 podNumber = 3;
  string merkleRootHash = 4;
  string previousMerkleRootHash = 5;
  bytes zkProof = 6;
}
*/

type RequestBodyVerifyPod struct {
	StationId              string `json:"station_id"`
	PodNumber              uint64 `json:"pod_number"`
	MerkleRootHash         string `json:"merkle_root_hash"`
	PreviousMerkleRootHash string `json:"previous_merkle_root_hash"`
	ZkProof                []byte `json:"zk_proof"`
}

type RequestBodyGetPod struct {
	BatchNumber uint64 `json:"batch_number"`
	ChainId     string `json:"chain_id"`
}
