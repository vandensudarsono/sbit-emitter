package emitter

import (
	"fmt"
	"log"
	dc "sbit-emitter/internal/deposit_codec"

	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

var emitter *goka.Emitter

func InitEmitter() {
	var err error
	stream := goka.Stream(viper.GetString("broker.topic"))
	broker := []string{fmt.Sprintf("%s:%s", viper.GetString("broker.host"), viper.GetString("broker.port"))}

	emitter, err = goka.NewEmitter(
		broker,
		stream,
		new(dc.DepositCodec),
	)

	if err != nil {
		log.Panicf("error creating emitter: %v", err)
	}

	//ßdefer emitter.Finish()

}

func GetEmitter() *goka.Emitter {
	return emitter
}

func CloseEmitter() {
	err := emitter.Finish()
	if err != nil {
		log.Printf("Error finisih emitter: %v", err)
	}
}
