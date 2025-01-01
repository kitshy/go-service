package address

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ETHAddress struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

func CreateAddressFromPrivateKey() (*ETHAddress, error) {
	prvKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := &ETHAddress{
		PrivateKey: hex.EncodeToString(crypto.FromECDSA(prvKey)),
		PublicKey:  hex.EncodeToString(crypto.FromECDSAPub(&prvKey.PublicKey)),
		Address:    crypto.PubkeyToAddress(prvKey.PublicKey).String(),
	}
	return address, nil
}

func PublicKeyToAddress(publicKey string) (string, error) {

	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	addressCommom := common.BytesToAddress(crypto.Keccak256(publicKeyBytes[1:])[12:])
	return addressCommom.String(), nil

}
