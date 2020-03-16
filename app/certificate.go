package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/spacemonkeygo/openssl"
)

func issueCertificate() []byte {

	key, err := openssl.GenerateRSAKey(1024)
	if err != nil {
		fmt.Println(err)
	}
	info := &openssl.CertificateInfo{
		Serial:       big.NewInt(int64(1)),
		Issued:       0,
		Expires:      24 * time.Hour,
		Country:      "BR",
		Organization: "BB",
		CommonName:   "bb.com.br",
	}
	cert, err := openssl.NewCertificate(info, key)
	if err != nil {
		fmt.Println(err)
	}
	if err := cert.Sign(key, openssl.EVP_SHA256); err != nil {
		fmt.Println(err)
	}

	pk, _ := cert.PublicKey()
	pem, _ := pk.MarshalPKIXPublicKeyPEM()

	fmt.Println("cert: ", string(pem))

	return pem

}
