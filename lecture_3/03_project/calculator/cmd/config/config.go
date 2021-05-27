package config

import "github.com/kelseyhightower/envconfig"

var Cfg Config

type Config struct {
	Rabbit         rabbitConfig `split_words:"true"`
	SqliteDatabase string       `split_words:"true"`
}

type rabbitConfig struct {
	ConsumerEventUpdateQueue    string `split_words:"true" required:"true"`
	ConsumerBetQueue            string `split_words:"true" required:"true"`
	ConsumerEventUpdateName     string `split_words:"true" default:"calculatorEvent"`
	ConsumerBetName             string `split_words:"true" default:"calculatorBet"`
	ConsumerAutoAck             bool   `split_words:"true" default:"true"`
	ConsumerExclusive           bool   `split_words:"true" default:"false"`
	ConsumerNoLocal             bool   `split_words:"true" default:"false"`
	ConsumerNoWait              bool   `split_words:"true" default:"false"`
	PublisherBetCalculatedQueue string `split_words:"true" required:"true"`
	PublisherDeclareDurable     bool   `split_words:"true" default:"true"`
	PublisherDeclareAutoDelete  bool   `split_words:"true" default:"false"`
	PublisherDeclareExclusive   bool   `split_words:"true" default:"false"`
	PublisherDeclareNoWait      bool   `split_words:"true" default:"false"`
	PublisherExchange           string `split_words:"true" default:""`
	PublisherMandatory          bool   `split_words:"true" default:"false"`
	PublisherImmediate          bool   `split_words:"true" default:"false"`
	DeclareDurable              bool   `split_words:"true" default:"true"`
	DeclareAutoDelete           bool   `split_words:"true" default:"false"`
	DeclareExclusive            bool   `split_words:"true" default:"false"`
	DeclareNoWait               bool   `split_words:"true" default:"false"`
}

func Load() {
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}