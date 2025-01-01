package rpc

import (
	"context"
	"github.com/kitshy/go-service/protobuf/wallet"
	"github.com/kitshy/go-service/service/address"
	"strconv"
)

func (*RpcServer) GetSupportCoins(ctx context.Context, in *wallet.SupportCoinsRequest) (*wallet.SupportCoinsResponse, error) {

	return &wallet.SupportCoinsResponse{
		Code:    strconv.Itoa(200),
		Msg:     "success",
		Support: true,
	}, nil
}

func (*RpcServer) GetWalletAddress(ctx context.Context, in *wallet.WalletAddressRequest) (*wallet.WalletAddressResponse, error) {

	addressInfo, err := address.CreateAddressFromPrivateKey()
	if err != nil {
		return &wallet.WalletAddressResponse{
			Code:      strconv.Itoa(400),
			Msg:       "create eth address fail",
			Address:   "",
			PublicKey: "",
		}, err
	}

	return &wallet.WalletAddressResponse{
		Code:      strconv.Itoa(200),
		Msg:       "success",
		Address:   addressInfo.Address,
		PublicKey: addressInfo.PublicKey,
	}, nil
}
