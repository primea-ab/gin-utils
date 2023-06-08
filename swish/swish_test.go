package swish

import (
	"context"
	"github.com/google/uuid"
	"os"
	"testing"
)

func TestImpl_MakePayment(t *testing.T) {
	cert, err := os.ReadFile("Swish_Merchant_TestCertificate_1234679304.pem")
	if err != nil {
		panic(err)
	}
	id := uuid.New().String()
	sw := NewSwish(false, "https://eat-easy-backend.fly.dev/nope", cert, "swish")
	err = sw.MakePayment(context.Background(), "4671234768", 10, id, "For testing")
	if err != nil {
		t.Error(err)
	}
}
