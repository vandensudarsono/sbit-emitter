package model

type (
	Deposit struct {
		WalletID int64   `json:"wallet_id"`
		Amount   float32 `json:"amount"`
	}

	DepositResponse struct {
		Status Status `json:"status"`
	}

	Status struct {
		Code          int32  `json:"code"`
		MessageClient string `json:"message_client"`
	}
)
