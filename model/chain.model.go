package model

type StationResponseBody struct {
	Station struct {
		Tracks               []string `json:"tracks"`
		VotingPower          []string `json:"votingPower"`
		LatestPod            string   `json:"latestPod"`
		LatestMerkleRootHash string   `json:"latestMerkleRootHash"`
		VerificationKey      string   `json:"verificationKey"`
		StationInfo          string   `json:"stationInfo"`
		Id                   string   `json:"id"`
		Creator              string   `json:"creator"`
	} `json:"station"`
}

type GetLatestSubmittedPodResponseBody struct {
	PodNumber string `json:"podNumber"`
	Message   string `json:"message"`
}

type GetLatestVerifiedPodResponseBody struct {
	PodNumber string `json:"podNumber"`
	Message   string `json:"message"`
}

type StationErrorResponseBody struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

type ExecutionLayerErrorResponseBody struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

type VerificationKeyResponseBody struct {
	Vkey string `json:"vkey"`
}

type AllExecutionLayersResponseBody struct {
	ExeLayer []struct {
		Validator            []string `json:"validator"`
		VotingPower          []string `json:"votingPower"`
		LatestBatch          string   `json:"latestBatch"`
		LatestMerkleRootHash string   `json:"latestMerkleRootHash"`
		VerificationKey      string   `json:"verificationKey"`
		ChainInfo            string   `json:"chainInfo"`
		ID                   string   `json:"id"`
		Creator              string   `json:"creator"`
	} `json:"exelayer"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}

type BatchResponseBody struct {
	Batch struct {
		BatchNumber        string `json:"batchNumber"`
		ChainId            string `json:"chainId"`
		PrevMerkleRootHash string `json:"prevMerkleRootHash"`
		MerkleRootHash     string `json:"merkleRootHash"`
		ZkProof            string `json:"zkProof"`
		Witness            string `json:"witness"`
		Verified           string `json:"verified"`
		BatchSubmitter     string `json:"batchSubmitter"`
		BatchVerifier      string `json:"batchVerifier"`
	} `json:"batch"`
}

type PodResponseBody struct {
	Pods struct {
		PodNumber              uint64 `json:"pod_number"`
		MerkleRootHash         string `json:"merkle_root_hash"`
		PreviousMerkleRootHash string `json:"previous_merkle_root_hash"`
		ZkProof                []byte `json:"zk_proof"`
		Witness                []byte `json:"witness"`
		Timestamp              uint64 `json:"timestamp"`
		IsVerified             bool   `json:"is_verified"`
	} `json:"pods"`
}
