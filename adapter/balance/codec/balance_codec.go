package codec

import (
	"encoding/json"
	"fmt"
	"sbit-emitter/domain/model"
)

type DepositCodec model.Wallet

// Encode a wallet into []byte
func (dc *DepositCodec) Encode(value interface{}) ([]byte, error) {
	if _, isWallet := value.(*model.Wallet); !isWallet {
		return nil, fmt.Errorf("codec requires value *wallet, got %T", value)
	}

	return json.Marshal(value)
}

// Decode a wallet from []byte to it's go representation.
func (dc *DepositCodec) Decode(data []byte) (interface{}, error) {
	var (
		mw  model.Wallet
		err error
	)

	err = json.Unmarshal(data, &mw)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling wallet: %v", err)
	}

	return &mw, nil
}
