package balanceinput

import (
	"fmt"
	"sbit-emitter/domain/model"
	logging "sbit-emitter/infrastructure/log"

	"github.com/lovoo/goka"
)

type BalanceInput struct{}

// func NewBalanceServer(bs *goka.Processor) *BalanceServer {
// 	return &BalanceServer{bs: bs}
// }

func NewBalanceInput() *BalanceInput {
	return &BalanceInput{}
}

func (b *BalanceInput) BalanceInputCB(ctx goka.Context, msg interface{}) {
	var (
		request *model.Deposit
		wallet  *model.Wallet
		err     error
	)

	if err != nil {
		logging.WithFields(logging.Fields{"component": "controller", "action": "balance input cb"}).
			Errorf("error proccess deposit to wallet: %v", err)
		return
	}

	//get the message deposit request
	request = msg.(*model.Deposit)
	if val := ctx.Value(); val != nil {
		wallet = val.(*model.Wallet)
	} else {
		wallet = new(model.Wallet)
	}

	if wallet != nil && request != nil {
		//sum
		wallet.Amount += request.Amount
		//set value
		ctx.SetValue(wallet)
		//log
		fmt.Printf("[proc] key: %s wallet_id: %d amount: %f msg: %v\n",
			ctx.Key(),
			wallet.WalletID,
			wallet.Amount,
			msg,
		)
	} else {
		fmt.Printf("either existing: %v or request: %v is nil\n",
			wallet,
			request,
		)
	}

}
