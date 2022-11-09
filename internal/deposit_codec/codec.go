package depositcodec

import (
	"encoding/json"
	"sbit-emitter/domain/model"
)

type DepositCodec struct{}

func (c *DepositCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *DepositCodec) Decode(data []byte) (interface{}, error) {
	var m model.Deposit
	return &m, json.Unmarshal(data, &m)
}

type DepositListCodec struct{}

func (c *DepositListCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *DepositListCodec) Decode(data []byte) (interface{}, error) {
	var m []model.Deposit
	err := json.Unmarshal(data, &m)
	return m, err
}
