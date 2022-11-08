package balance

import (
	"context"
	"log"
	cd "sbit-emitter/adapter/balance/codec"
	bi "sbit-emitter/adapter/balance/input"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/spf13/viper"
)

var balance *goka.Processor

func InitBalance() {
	var (
		err error
	)

	//crete config first
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	//define group
	group := goka.Group(viper.GetString("processor.balance.group"))
	topic := goka.Stream(viper.GetString("brocker.topic"))

	b := bi.NewBalanceInput()
	g := goka.DefineGroup(
		group,
		goka.Input(topic, new(codec.String), b.BalanceInputCB),
		goka.Persist(new(cd.DepositCodec)),
	)
	balance, err = goka.NewProcessor(
		viper.GetStringSlice("brocker.url"),
		g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)

	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
}

func GetBalanceProcessor() *goka.Processor {
	return balance
}

func RunBallanceProcessor() error {
	return balance.Run(context.Background())
}
