package address

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateAddressFromPrivateKey(*testing.T) {

	addressInfo, err := CreateAddressFromPrivateKey()
	if err != nil {
		return
	}

	// 将结构体转换为 JSON
	jsonData, err := json.Marshal(addressInfo)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
	fmt.Println("Address: ", addressInfo.Address, "PrivateKey: ", addressInfo.PrivateKey, "PublicKey: ", addressInfo.PublicKey)

}

/*
*
{"private_key":"b0be241ce4fb980b744f0a583ad92a238beb1f9af3885e3d3c7ddd4edd5a09fe",
"public_key":"04d57b7e1581627bf60f0e1d6239e4626deb0f03970b3529f85665aefdc3db7f033e637357c545192da78aa950174995277e0a34957b907bfd44c2cc8e80586419",
"address":"0xc6Dd43F53d6f4c4bcAD35D91C78Fb786509bb17c"}
*/
func TestPublicKeyToAddress(t *testing.T) {

	address, err := PublicKeyToAddress("04d57b7e1581627bf60f0e1d6239e4626deb0f03970b3529f85665aefdc3db7f033e637357c545192da78aa950174995277e0a34957b907bfd44c2cc8e80586419")

	if err != nil {
		return
	}

	fmt.Println("address: ", address)

}
