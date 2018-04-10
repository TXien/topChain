package crypto

import(
	"fmt"
	"testing"

)

func TestHash(t *testing.T){
	fmt.Println("HashStringToString")
	fmt.Println(HashStringToString("123"))
	fmt.Println("HashStringToByte")
	fmt.Println(HashStringToByte("123"))
	fmt.Println("HashByteToByte")
	fmt.Println(HashByteToByte([]byte("123")))
	fmt.Println("HashByte512ToByte")
	fmt.Println(HashByte512ToByte([]byte("123")))
	fmt.Println("SignResult")
	fmt.Println(VerifyTest())
	PrivatekeyToPublicKey([]byte("4d4e4ccf1a2ac238379251e5295ec041c3dfefedba07269baefd2d6bc6ec1dfe"))
}
