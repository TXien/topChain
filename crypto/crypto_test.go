package crypto

import(
	"fmt"
	"testing"

)

func TestHash(t *testing.T){
	fmt.Println(HashStringToString("123"))
	fmt.Println(HashStringToByte("123"))
	fmt.Println(HashByteToByte([]byte("123")))
}
