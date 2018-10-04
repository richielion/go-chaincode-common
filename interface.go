package go_chaincode_common

type PayerInterface interface {
	//TODO: is callback designed needed?
	GetTokens(auth MemberAuth, params []string) (tokenVerify, tokenPay string)
	Propose(auth ClinicAuth, params []string) (feeForm string)
	Modify(auth ClinicAuth, params []string) (extraFee string)
	Revert(auth ClinicAuth, params []string)
	Settlement(auth PayerAuth, params []string)
}

const (
	Payer_fcn_getTokens  = "getTokens"
	Payer_fcn_propose    = "propose"
	Payer_fcn_modify     = "modify"
	Payer_fcn_revert     = "revert"
	Payer_fcn_settlement = "settlement"
)

type ClinicAuth func(transient map[string][]byte) bool
type MemberAuth func(transient map[string][]byte) bool
type PayerAuth func(transient map[string][]byte) bool

type VisitData struct {
	Token            string //provided by QRCode
	ClinicID         string //provided by clinic
	Doctor           Doctor //provided by clinic
	MedicalNetworkID string //provided by clinic
}
type Doctor string //doctor code
