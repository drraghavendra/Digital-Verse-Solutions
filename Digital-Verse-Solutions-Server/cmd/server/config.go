package main

import (
	"gitlab.com/digitalverse/blockchain/pkg/blockchain"
	"gitlab.com/falaleev-golang/config"
)

type cfg struct {
	SecretKey  string `env:"AWS_SECRET_KEY"`
	AccessKey  string `env:"AWS_ACCESS_KEY"`
	Bucket     string `env:"AWS_BUCKET"`
	Region     string `env:"AWS_REGION"`
	ListenPort int    `env:"LISTEN_PORT"`
}

var c = new(cfg)
var bc = new(blockchain.Config)

func init() {
	config.LoadConfig(c)
	config.LoadConfig(bc)
}
