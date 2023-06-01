package swish

import (
	"context"
	"os"
	"testing"
)

func TestImpl_MakePayment(t *testing.T) {
	cert, err := os.ReadFile("Swish_Merchant_TestCertificate_1234679304.p12")
	if err != nil {
		panic(err)
	}
	sw := NewSwish(false, "https://eat-easy-backend.fly.dev/nope", cert, "swish")
	_, err = sw.MakePayment(context.Background(), "4671234768", 10, "abcdefgqwertyuiop", "For testing")
	if err != nil {
		t.Error(err)
	}
}
