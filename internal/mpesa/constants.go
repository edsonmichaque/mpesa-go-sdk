package mpesa

import (
	"regexp"
)

type Opcode string

const (
	B2cPayment Opcode = "B2cPayment"
	B2bPayment Opcode = "B2bPayment"
	C2bPayment Opcode = "C2bPayment"
	Reversal Opcode = "Reversal"
	QueryTransactionStatus Opcode = "QueryTransactionStatus"
)

var  (
	PatternPhoneNumber, _ = regexp.Compile("^((00|\\+)?(258))?8[45][0-9]{7}$")
	PatternMoneyAmount, _ = regexp.Compile("^[1-9][0-9]*(\\.[0-9]+)?$/")
	PatternWord, _ = regexp.Compile("^\\w+$")
	PatternServiceProviderCode, _ = regexp.Compile("^[0-9]{5,6}$")
)

var (
	DefaultOperations = map[Opcode]interface{} {
		C2bPayment: map[string]interface{} {
			"method": "post",
			"port": "18352",
			"path": "/ipg/v1x/c2bPayment/singleStage/",
			"mapping": map[string]string {
				"from": "input_CustomerMSISDN",
				"to": "input_ServiceProviderCode",
				"amount": "input_Amount",
				"transaction": "input_TransactionReference",
				"reference": "input_ThirdPartyReference",
			},
			"validation": map[string]*regexp.Regexp {
				"from": PatternPhoneNumber,
				"to": PatternServiceProviderCode,
				"amount": PatternMoneyAmount,
				"transaction": PatternWord,
				"reference": PatternWord,
			},
			"required": []string{
				"to", 
				"from", 
				"amount", 
				"transaction", 
				"reference",
			},
			"optional": []string{
				"from",
			},
		},
		B2bPayment: map[string]interface{} {
			"method": "post",
			"port": "18349",
			"path": "/ipg/v1x/b2bPayment/",
			"mapping": map[string]string {
				"from": "input_PrimaryPartyCode",
				"to": "input_ReceiverPartyCode",
				"amount": "input_Amount",
				"transaction": "input_TransactionReference",
				"reference": "input_ThirdPartyReference",
			},
			"validation": map[string]*regexp.Regexp {
				"from": PatternServiceProviderCode,
				"to": PatternServiceProviderCode,
				"amount": PatternMoneyAmount,
				"transaction": PatternWord,
				"reference": PatternWord,
			},
			"required": []string{
				"to", 
				"from", 
				"amount", 
				"transaction", 
				"reference",
			},
			"optional": []string{
				"from",
			},
		},
		B2cPayment: map[string]interface{} {
			"method": "post",
			"port": "18345",
			"path": "/ipg/v1x/b2cPayment/",
			"mapping": map[string]string {
				"number": "input_CustomerMSISDN",
				"to": "input_CustomerMSISDN",
				"from": "input_ServiceProviderCode",
				"amount": "input_Amount",
				"transaction": "input_TransactionReference",
				"reference": "input_ThirdPartyReference",
			},
			"validation": map[string]*regexp.Regexp {
				"to": PatternPhoneNumber,
				"from": PatternServiceProviderCode,
				"amount": PatternMoneyAmount,
				"transaction": PatternWord,
				"reference": PatternWord,
			},
			"required": []string{
				"to", 
				"from", 
				"amount", 
				"transaction", 
				"reference",
			},
			"optional": []string{
				"to",
			},
		},
		Reversal: map[string]interface{} {
			"method": "put",
			"port": "18354",
			"path": "/ipg/v1x/reversal/",
			"mapping": map[string]string {
				"to": "input_ServiceProviderCode",
				"amount": "input_Amount",
				"reference": "input_ThirdPartyReference",
				"transaction": "input_TransactionID",
				"securityCredential": "input_SecurityCredential",
				"initiatorIdentifier": "input_InitiatorIdentifier",
			},
			"validation": map[string]*regexp.Regexp {
				"to": PatternServiceProviderCode,
				"amount": PatternMoneyAmount,
				"reference": PatternWord,
				"transaction": PatternWord,
				"securityCredential": PatternWord,
				"initiatorIdentifier": PatternWord,
			},
			"required": []string{
				"to",
				"amount",
				"reference",
				"transaction",
				"securityCredential",
				"initiatorIdentifier",
			},
			"optional": []string{
				"to",
				"securityCredential",
				"initiatorIdentifier",
			},
		},
		QueryTransactionStatus: map[string]interface{} {
			"method": "get",
			"port": "18353",
			"path": "/ipg/v1x/queryTransactionStatus/",
			"mapping": map[string]string {
				"from": "input_ServiceProviderCode",
				"subject": "input_QueryReference",
				"reference": "input_ThirdPartyReference",
			},
			"validation": map[string]*regexp.Regexp {
				"from": PatternServiceProviderCode,
				"subject": PatternWord,
				"reference": PatternWord,
			},
			"required": []string{
				"from",
				"subject",
				"reference",				
			},
			"optional": []string{
				"from",
			},
		},
	}
)

