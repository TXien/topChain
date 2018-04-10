package crypto

import (
	"encoding/hex"
//	"fmt"
	"golang.org/x/crypto/blake2b")

func HashByteToByte(hashbyte []byte)([]byte){
	return HashByte(hashbyte)
}

func HashStringToString(hashstring string)(string){
	return hex.EncodeToString(HashString(hashstring))
}

func HashStringToByte(hashstring string)([]byte){
	return HashString(hashstring)
}

func HashString(hashstring string) ([]byte){
	test,_ := blake2b.New256([]byte(hashstring))
//	fmt.Println(hex.EncodeToString(test.Sum(nil)))
	return test.Sum(nil)
}

func HashByte(hashbyte []byte) ([]byte){
        test,_ := blake2b.New256(hashbyte)
//      fmt.Println(hex.EncodeToString(test.Sum(nil)))
        return test.Sum(nil)
}
