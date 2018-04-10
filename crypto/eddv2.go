package crypto

import (
//	"bytes"
//	"crypto"
	"crypto/rand"
	cryptorand "crypto/rand"
	"crypto/sha512"
//	"errors"
	"fmt"
	"encoding/hex"
	"io"
	"strconv"
//"github.com/topchain/crypto/crypto"
	"github.com/topchain/crypto/edward"
//	"golang.org/x/crypto/ed25519/internal/edwards25519"
)

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize = 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize = 64
)

type PublicKey []byte
type PrivateKey []byte

func main(){
	pub,pri,_ := GenerateKey(rand.Reader)
	fmt.Println(hex.EncodeToString(pri))
	fmt.Println(hex.EncodeToString(pub))
}

func GenerateKey(rand io.Reader) (publicKey PublicKey, privateKey PrivateKey, err error) {
	if rand == nil {
		rand = cryptorand.Reader
	}

	privateKey = make([]byte, PrivateKeySize)
	publicKey = make([]byte, PublicKeySize)
	_, err = io.ReadFull(rand, privateKey[:32])
	if err != nil {
		return nil, nil, err
	}

	digest := sha512.Sum512(privateKey[:32])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], digest[:])
	edwards25519.GeScalarMultBase(&A, &hBytes)
	var publicKeyBytes [32]byte
	A.ToBytes(&publicKeyBytes)

	copy(privateKey[32:], publicKeyBytes[:])
	copy(publicKey, publicKeyBytes[:])

	return publicKey, privateKey, nil
}

func PrivatekeyToPublicKey(data []byte)(publicKey PublicKey, err error){
	/*if rand == nil {
            rand = cryptorand.Reader
	}*/
	privateKey := data
	publicKey = make([]byte, PublicKeySize)
	/*_, err = io.ReadFull(rand, privateKey[:32])*/
	if err != nil {
			return nil, err
	}

	digest := data
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64


	
	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], digest[:])
	edwards25519.GeScalarMultBase(&A, &hBytes)
	var publicKeyBytes [32]byte
	A.ToBytes(&publicKeyBytes)
	
	copy(privateKey[32:], publicKeyBytes[:])
	copy(publicKey, publicKeyBytes[:])
	
	return publicKey, nil

}

// Sign signs the message with privateKey and returns a signature. It will
// panic if len(privateKey) is not PrivateKeySize.
func Sign(privateKey PrivateKey, message []byte) []byte {
	if l := len(privateKey); l != PrivateKeySize {
		panic("ed25519: bad private key length: " + strconv.Itoa(l))
	}

	h := sha512.New()
	h.Write(privateKey[:32])

	var digest1, messageDigest, hramDigest [64]byte
	var expandedSecretKey [32]byte
	h.Sum(digest1[:0])
	copy(expandedSecretKey[:], digest1[:])
	expandedSecretKey[0] &= 248
	expandedSecretKey[31] &= 63
	expandedSecretKey[31] |= 64

	h.Reset()
	h.Write(digest1[32:])
	h.Write(message)
	h.Sum(messageDigest[:0])

	var messageDigestReduced [32]byte
	edwards25519.ScReduce(&messageDigestReduced, &messageDigest)
	var R edwards25519.ExtendedGroupElement
	edwards25519.GeScalarMultBase(&R, &messageDigestReduced)

	var encodedR [32]byte
	R.ToBytes(&encodedR)

	h.Reset()
	h.Write(encodedR[:])
	h.Write(privateKey[32:])
	h.Write(message)
	h.Sum(hramDigest[:0])
	var hramDigestReduced [32]byte
	edwards25519.ScReduce(&hramDigestReduced, &hramDigest)

	var s [32]byte
	edwards25519.ScMulAdd(&s, &hramDigestReduced, &expandedSecretKey, &messageDigestReduced)

	signature := make([]byte, SignatureSize)
	copy(signature[:], encodedR[:])
	copy(signature[32:], s[:])

	return signature
}
