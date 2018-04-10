package crypto

import (
	//"encoding/hex"
	"golang.org/x/crypto/ed25519"
	"fmt"
	"crypto/rand"
	//"encoding/hex"
)
/*
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
*/

func key(){
	//pub,pri,_ := /*ed25519.*/GenerateKey(rand.Reader)

}

func publickey(data []byte)([]byte){
	return ed25519.PublicKey(data)
}

func VerifyTest() bool{
        pub,pri,_ := /*ed25519.*/GenerateKey(rand.Reader)
	fmt.Println("123pub:",pub)
	fmt.Println("123pub:",ed25519.PublicKey(pri[2:]))
        tmp := /*ed25519.*/Sign(pri,[]byte(""))
	//fmt.Println(hex.EncodeToString(pri))
	return Verify(pub,[]byte(""),tmp)

}

func Verify(pub []byte, content []byte, hash []byte) bool {
	//pubKey := ToECDSAPub(signerAddress.Bytes())
	//fmt.Println(ed25519.Verify(pub,content,hash))
	return ed25519.Verify(pub,content,hash)
}
