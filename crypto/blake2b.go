package crypto

import (
	"encoding/hex"
//	"fmt"
	"golang.org/x/crypto/blake2b")

func HashByteToByte(hashbyte []byte)([]byte){
	return HashByte256(hashbyte)
}

func HashStringToString(hashstring string)(string){
	return hex.EncodeToString(HashString256(hashstring))
}

func HashStringToByte(hashstring string)([]byte){
	return HashString256(hashstring)
}

func HashString256(hashstring string) ([]byte){
	test,_ := blake2b.New256([]byte(hashstring))
//	fmt.Println(hex.EncodeToString(test.Sum(nil)))
	return test.Sum(nil)
}

func HashByte256(hashbyte []byte) ([]byte){
        test,_ := blake2b.New256(hashbyte)
//      fmt.Println(hex.EncodeToString(test.Sum(nil)))
        return test.Sum(nil)
}

func HashByte512ToByte(hashbyte []byte)([]byte){
	return HashByte512(hashbyte)
}

func HashByte512(hashbyte []byte) ([]byte){
        test,_ := blake2b.New512(hashbyte)
//      fmt.Println(hex.EncodeToString(test.Sum(nil)))
        return test.Sum(nil)
}
