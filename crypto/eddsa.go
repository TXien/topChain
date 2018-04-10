package crypto

import (
	"golang.org/x/crypto/ed25519"
	"fmt"
	"crypto/rand"
	//"encoding/hex"
)

func main(){
	//pub,pri,_ := ed25519.GenerateKey(rand.Reader)
	//fmt.Println(hex.EncodeToString(pri))
	//fmt.Println(hex.EncodeToString(pub))
	//tmp := ed25519.Sign(pri,[]byte(""))
	//fmt.Println(hex.EncodeToString(tmp))
	//tr:= ed25519.Verify(pub,[]byte(""),tmp)
	//fmt.Println(tr)
	//fmt.Println(Verify(pub,[]byte(""),tmp))
	fmt.Println(VerifyTest())
}

func VerifyTest() bool{
        pub,pri,_ := ed25519.GenerateKey(rand.Reader)
        tmp := ed25519.Sign(pri,[]byte(""))
	return Verify(pub,[]byte(""),tmp)

}

func Verify(pub []byte, content []byte, hash []byte) bool {
	//pubKey := ToECDSAPub(signerAddress.Bytes())
	return ed25519.Verify(pub,content,hash)
}
