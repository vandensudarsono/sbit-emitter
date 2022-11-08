package model

type Wallet struct {
	WalletID int     `json:"wallet_id"`
	Amount   float32 `json:"amount"`
}
