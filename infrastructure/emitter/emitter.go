package emitter

import (
	"log"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/spf13/viper"
)

var emitter *goka.Emitter

func InitEmitter() {
	var err error
	topic := goka.Stream(viper.GetString("brocker.topic"))
	emitter, err = goka.NewEmitter(
		viper.GetStringSlice("brocker.url"),
		topic,
		new(codec.String),
	)

	if err != nil {
		log.Panicf("error creating emitter: %v", err)
	}

	//defer emitter.Finish()
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
