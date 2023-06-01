package swish

type MakePaymentRequest struct {
	PayeePaymentReference string `json:"payeePaymentReference"`
	CallbackUrl           string `json:"callbackUrl"`
	PayerAlias            string `json:"payerAlias"`
	PayeeAlias            string `json:"payeeAlias"`
	Amount                string `json:"amount"`
	Currency              string `json:"currency"`
	Message               string `json:"message"`
}
