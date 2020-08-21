package mpesa


type Opcode string

const (
	B2cPayment Opcode = "B2cPayment"
	B2bPayment Opcode = "B2bPayment"
	C2bPayment Opcode = "C2bPayment"
	Reversal Opcode = "Reversal"
	QueryTransactionStatus Opcode = "QueryTransactionStatus"
)

var (
	DefaultOperations = map[Opcode]interface{} {
		C2bPayment: map[string]interface{} {
			"method": "",
			"port": "",
			"path": "",
			"mapping": map[string]string {

			},
			"validation": map[string]string {

			},
			"required": []string{},
			"optional": []string{},
		},
		B2bPayment: map[string]interface{} {
			"method": "",
			"port": "",
			"path": "",
			"mapping": map[string]string {

			},
			"validation": map[string]string {

			},
			"required": []string{},
			"optional": []string{},
		},
		B2cPayment: map[string]interface{} {
			"method": "",
			"port": "",
			"path": "",
			"mapping": map[string]string {

			},
			"validation": map[string]string {

			},
			"required": []string{

			},
			"optional": []string{},
		},
		Reversal: map[string]interface{} {
			"method": "",
			"port": "",
			"path": "",
			"mapping": map[string]string {

			},
			"validation": map[string]string {

			},
			"required": []string{},
			"optional": []string{},
		},
		QueryTransactionStatus: map[string]interface{} {
			"method": "",
			"port": "",
			"path": "",
			"mapping": map[string]string {

			},
			"validation": map[string]string {

			},
			"required": []string{},
			"optional": []string{},
		},
	}
)