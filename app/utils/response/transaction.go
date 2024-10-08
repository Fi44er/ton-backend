package response

type Transaction struct {
	Hash            string       `json:"hash"`
	Lt              int64        `json:"lt"`
	Account         Account      `json:"account"`
	Success         bool         `json:"success"`
	Utime           int64        `json:"utime"`
	OrigStatus      string       `json:"orig_status"`
	EndStatus       string       `json:"end_status"`
	TotalFees       int64        `json:"total_fees"`
	EndBalance      int64        `json:"end_balance"`
	TransactionType string       `json:"transaction_type"`
	StateUpdateOld  string       `json:"state_update_old"`
	StateUpdateNew  string       `json:"state_update_new"`
	InMsg           InMsg        `json:"in_msg"`
	OutMsgs         []OutMsg     `json:"out_msgs"`
	Block           string       `json:"block"`
	PrevTransHash   string       `json:"prev_trans_hash"`
	PrevTransLt     int64        `json:"prev_trans_lt"`
	ComputePhase    ComputePhase `json:"compute_phase"`
	StoragePhase    StoragePhase `json:"storage_phase"`
	ActionPhase     ActionPhase  `json:"action_phase"`
	Aborted         bool         `json:"aborted"`
	Destroyed       bool         `json:"destroyed"`
	Raw             string       `json:"raw"`
}

type Account struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	IsScam   bool   `json:"is_scam"`
	IsWallet bool   `json:"is_wallet"`
}

type InMsg struct {
	MsgType       string      `json:"msg_type"`
	CreatedLt     int64       `json:"created_lt"`
	IhrDisabled   bool        `json:"ihr_disabled"`
	Bounce        bool        `json:"bounce"`
	Bounced       bool        `json:"bounced"`
	Value         int64       `json:"value"`
	FwdFee        int64       `json:"fwd_fee"`
	IhrFee        int64       `json:"ihr_fee"`
	Destination   Account     `json:"destination"`
	Source        Account     `json:"source"`
	ImportFee     int64       `json:"import_fee"`
	CreatedAt     int64       `json:"created_at"`
	OpCode        string      `json:"op_code"`
	Hash          string      `json:"hash"`
	RawBody       string      `json:"raw_body"`
	DecodedOpName string      `json:"decoded_op_name"`
	DecodedBody   DecodedBody `json:"decoded_body"`
}

type DecodedBody struct {
	Text string `json:"text"`
}

type OutMsg struct {
	// Поля OutMsg не указаны в примере JSON, поэтому они не включены в эту структуру
}

type ComputePhase struct {
	Skipped             bool   `json:"skipped"`
	Success             bool   `json:"success"`
	GasFees             int64  `json:"gas_fees"`
	GasUsed             int64  `json:"gas_used"`
	VmSteps             int64  `json:"vm_steps"`
	ExitCode            int64  `json:"exit_code"`
	ExitCodeDescription string `json:"exit_code_description"`
}

type StoragePhase struct {
	FeesCollected int64  `json:"fees_collected"`
	StatusChange  string `json:"status_change"`
}

type ActionPhase struct {
	Success        bool  `json:"success"`
	ResultCode     int64 `json:"result_code"`
	TotalActions   int64 `json:"total_actions"`
	SkippedActions int64 `json:"skipped_actions"`
	FwdFees        int64 `json:"fwd_fees"`
	TotalFees      int64 `json:"total_fees"`
}

type CreditPhase struct {
	FeesCollected int64 `json:"fees_collected"`
	Credit        int64 `json:"credit"`
}

type Response struct {
	Transaction Transaction `json:"transaction"`
	Interfaces  []string    `json:"interfaces"`
	Children    []Child     `json:"children"`
}

type Child struct {
	Transaction Transaction `json:"transaction"`
	Interfaces  []string    `json:"interfaces"`
}
