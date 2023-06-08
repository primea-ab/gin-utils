package swish

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	PROD_URL    = "https://cpc.getswish.net/swish-cpcapi/api/v2/"
	TEST_URL    = "https://mss.cpc.getswish.net/swish-cpcapi/api/v2/"
	certificate = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURyekNDQXBlZ0F3SUJBZ0lRQ0R2Z1ZwQkNSckdoZFdySldaSEhTakFOQmdrcWhraUc5dzBCQVFVRkFEQmgKTVFzd0NRWURWUVFHRXdKVlV6RVZNQk1HQTFVRUNoTU1SR2xuYVVObGNuUWdTVzVqTVJrd0Z3WURWUVFMRXhCMwpkM2N1WkdsbmFXTmxjblF1WTI5dE1TQXdIZ1lEVlFRREV4ZEVhV2RwUTJWeWRDQkhiRzlpWVd3Z1VtOXZkQ0JEClFUQWVGdzB3TmpFeE1UQXdNREF3TURCYUZ3MHpNVEV4TVRBd01EQXdNREJhTUdFeEN6QUpCZ05WQkFZVEFsVlQKTVJVd0V3WURWUVFLRXd4RWFXZHBRMlZ5ZENCSmJtTXhHVEFYQmdOVkJBc1RFSGQzZHk1a2FXZHBZMlZ5ZEM1agpiMjB4SURBZUJnTlZCQU1URjBScFoybERaWEowSUVkc2IySmhiQ0JTYjI5MElFTkJNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUE0anZoRVhMZXFLVFRvMWVxVUtLUEMzZVF5YUtsN2hMT2xsc0IKQ1NETUFaT25UakMzVS9kRHhHa0FWNTNpalNMZGh3WkFBSUVKenM0Ymc3L2Z6VHR4UnVMV1pzY0ZzM1luRm85NwpuaDZWZmU2M1NLTUkydGF2ZWd3NUJtVi9TbDBmdkJmNHE3N3VLTmQwZjNwNG1WbUZhRzVjSXpKTHYwN0E2RnB0CjQzQy9keEMvL0FIMmhkbW9SQkJZTXFsMUdOWFJvcjVINGlkcTlKb3orRWtJWUl2VVg3UTZoTCtocWtwTWZUN1AKVDE5c2RsNmdTemVSbnR3aTVtM09GQnFPYXN2K3piTVVaQmZIV3ltZU1yL3k3dnJUQzBMVXE3ZEJNdG9NMU8vNApnZFc3alZnL3RSdm9TU2lpY05veEJOMzNzaGJ5VEFwT0I2anRTajFldFgramtNT3ZKd0lEQVFBQm8yTXdZVEFPCkJnTlZIUThCQWY4RUJBTUNBWVl3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFkQmdOVkhRNEVGZ1FVQTk1UU5WYlIKVEx0bThLUGlHeHZEbDdJOTBWVXdId1lEVlIwakJCZ3dGb0FVQTk1UU5WYlJUTHRtOEtQaUd4dkRsN0k5MFZVdwpEUVlKS29aSWh2Y05BUUVGQlFBRGdnRUJBTXVjTjZwSUV4SUsrdDFFbkU5U3NQVGZyZ1QxZVhrSW95UVkvRXNyCmhNQXR1ZFhIL3ZUQkgxakx1RzJjZW5Ubm1DbXJFYlhqY0tDaHpVeUltWk9Na1hEaXF3OGN2cE9wLzJQVjVBZGcKMDZPL25Wc0o4ZFdPNDFQMGptUDZQNmZidEdiZlltYlcwVzVCamZJdHRlcDNTcCtkV09JcldjQkFJKzB0S0lKRgpQbmxVa2lhWTRJQklxRGZ2OE5aNVlCYmVyT2dPelc2c1JCYzRMMG5hNFVVK0tyazJVODg2VUFiM0x1akVWMGxzCllTRVkxUVN0ZUR3c09vQnJwK3V2RlJUcDJJbkJ1VGhzNHBGc2l2OWt1WGNsVnpEQUd5U2o0ZHpwMzBkOHRiUWsKQ0FVdzdDMjlDNzlGdjFDNXFmUHJtQUVTcmNpSXhwZzBYNDBLUE1icDFaV1ZiZDQ9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURqakNDQW5hZ0F3SUJBZ0lRQXpyeDVxY1JxYUM3S0dTeEhRbjY1VEFOQmdrcWhraUc5dzBCQVFzRkFEQmgKTVFzd0NRWURWUVFHRXdKVlV6RVZNQk1HQTFVRUNoTU1SR2xuYVVObGNuUWdTVzVqTVJrd0Z3WURWUVFMRXhCMwpkM2N1WkdsbmFXTmxjblF1WTI5dE1TQXdIZ1lEVlFRREV4ZEVhV2RwUTJWeWRDQkhiRzlpWVd3Z1VtOXZkQ0JICk1qQWVGdzB4TXpBNE1ERXhNakF3TURCYUZ3MHpPREF4TVRVeE1qQXdNREJhTUdFeEN6QUpCZ05WQkFZVEFsVlQKTVJVd0V3WURWUVFLRXd4RWFXZHBRMlZ5ZENCSmJtTXhHVEFYQmdOVkJBc1RFSGQzZHk1a2FXZHBZMlZ5ZEM1agpiMjB4SURBZUJnTlZCQU1URjBScFoybERaWEowSUVkc2IySmhiQ0JTYjI5MElFY3lNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF1emZOTk54N2E4bXlhSkN0U25YL1Jyb2hDZ2lOOVJsVXlmdUkKMi9PdThqcUprVHg2NXFzR0dtdlByQzNvWGdra1JMcGltbjdXbzZoKzRGUjFJQVdzVUxlY1l4cHNNTnphSHhteAoxeDdlL2RmZ3k1U0RONjdzSDBOTzNYc3MwcjB1cFMva3FiaXRPdFNacExZbDZadHJBR0NTWVA5UElVa1k5MmVRCnEyRUduSS95dXVtMDZaSXlhN1h6VitoZEc4Mk1IYXVWQkpWSjh6VXRsdU5KYmQxMzQvdEpTN1NzVlFlcGo1V3oKdENPN1RHMUY4UGFwc3BVd3RQMU1WWXduU2xjVWZJS2R6WE9TMHhaS0JneU1VTkdQSGdtK0Y2SG1JY3I5ZytVUQp2SU9sQ3NSbktQWnpGQlE5Um5iRGh4U0pJVFJOcnc5RkRLWkpvYnE3bk1XeE00TXBoUUlEQVFBQm8wSXdRREFQCkJnTlZIUk1CQWY4RUJUQURBUUgvTUE0R0ExVWREd0VCL3dRRUF3SUJoakFkQmdOVkhRNEVGZ1FVVGlKVUlCaVYKNXVOdTVnLzYrcmtTN1FZWGp6a3dEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBR0JuS0pSdkRraGo2ekhkNm1jWQoxWWw5UE1XTFNuL3B2dHNyRjkrd1gzTjNLaklUT1lGblFvUWo4a1ZuTmV5SXYvaVBzR0VNTktTdUlFeUV4dHY0Ck5lRjIyZCttUXJ2SFJBaUdmelowSkZyYWJBMFVXVFc5OGtuZHRoL0pzdzFIS2oyWkw3dGN1N1hVSU9HWlgxTkcKRmR0b20vRHpNTlUrTWVLTmhKN2ppdHJhbGo0MUU2VmY4UGx3VUhCSFFSRlhHVTdBajY0R3hKVVRGeThiSlo5MQo4ckdPbWFGdkU3RkJjZjZJS3NoUEVDQlYxL01VUmVYZ1JQVHFoNVV5a3c3K1UwYjZMSjMvaXlLNVM5a0pSYVRlCnBMaWFXTjBiZlZLZmpsbERpSUdrbmliVmI2M2REY1kzZmUwRGtodmxkMTkyN2p5TnhGMVdXNkxaWm02ek5UZmwKTXJZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0t"
)

type Swish interface {
	MakePayment(ctx context.Context, phoneNumber string, amount int, reference string, message string) error
}

func NewSwish(prod bool, callbackUrl string, clientCert []byte, clientCertPassword string) Swish {
	s := &Impl{}
	if prod {
		s.url = PROD_URL
	} else {
		s.url = TEST_URL
	}
	s.callbackUrl = callbackUrl
	rootCert, err := base64.StdEncoding.DecodeString(certificate)
	if err != nil {
		log.Fatalf("could not decode the certificate. %v", err)
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(rootCert) {
		log.Fatalf("could not append CA Certificate to pool. Invalid rootCertBase64")
	}

	cert, err := tls.X509KeyPair(clientCert, clientCert)
	if err != nil {
		log.Fatalf("unable to load pkcs12 %v", err)
	}
	tlsConf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caPool,
		RootCAs:      caPool,
	}

	s.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConf,
		},
	}

	return s
}

type Impl struct {
	url         string
	callbackUrl string
	client      *http.Client
}

func (s *Impl) MakePayment(
	ctx context.Context,
	phoneNumber string,
	amountCent int,
	reference string,
	message string,
) error {
	reference = strings.ToUpper(strings.ReplaceAll(reference, "-", ""))
	payload := &MakePaymentRequest{
		PayeePaymentReference: reference,
		CallbackUrl:           s.callbackUrl,
		PayeeAlias:            "1230434126",
		PayerAlias:            phoneNumber,
		Amount:                strconv.Itoa(amountCent / 100),
		Currency:              "SEK",
		Message:               message,
	}

	b, _ := json.Marshal(payload)
	url := s.url + "paymentrequests/" + reference
	req, err := http.NewRequestWithContext(ctx, http.MethodPut,
		url,
		bytes.NewReader(b))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Printf("Failed to create request %v", err)
		return err
	}
	res, err := s.client.Do(req)
	if strings.TrimSpace(res.Status) == "201" {
		return nil
	} else {
		log.Printf("Swish status %s", res.Status)
		log.Printf("Got response from swish %v", res)

		rb, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Body %v", string(rb))
		return err
	}
}
