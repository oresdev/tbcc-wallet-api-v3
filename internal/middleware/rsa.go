package rsa

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	b64 "encoding/base64"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// CheckRSA ...
func CheckRSA(r *http.Request) error {
	// calculate time marker ( 1 minute )
	marker := time.Now().UnixNano() / 60000000000

	// prepare rsa public key
	modulus, _ := new(big.Int).SetString("24978937929039079698900217277211139741298503225304940708471500632998807940594098094279715760007837567758140271648213160782665090862429674989595786786706158407718132602243073350573105466295259402015968964428697348933647799166252832445412109186606552830657937566913634244602248448168263595041609788376844207040134685246988517369175423733578630513935247876944612547038369213440965707990434304216134994872988178442403703838375961614530418130049520601503784368593256331382120905216278087273411514276175298853530167166932291726699385905365795710542400924502550360760823220499686311360010761311412959233442539203726425962389", 0)

	publickey := rsa.PublicKey{N: modulus, E: 65537}

	// prepare message
	body, _ := ioutil.ReadAll(r.Body)
	message := r.URL.Path + string(body) + "47519363023227045875440997723988" + strconv.FormatInt(marker, 10)

	// prepare signature from request
	signature := r.Header.Get("x-signature")
	sDec, _ := b64.StdEncoding.DecodeString(signature)

	// calculate sha256 hash of message
	digest := sha256.New()
	digest.Write([]byte(message))

	hashed := digest.Sum(nil)

	// verify signature
	err := rsa.VerifyPKCS1v15(&publickey, crypto.SHA256, hashed, sDec)

	return err
}
