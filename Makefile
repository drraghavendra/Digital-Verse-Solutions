build:
	go build -o ./bin/app ./cmd/server
.PHONY: build

docker-build:
	docker build -t test -f Dockerfile .
.PHONY: docker-build

eth_contract:
	abigen --abi=./solidity/DvNft.abi --pkg=blockchain --out=./pkg/blockchain/smart.go
.PHONY: eth_contract

eth_contract_market:
	abigen --abi=./solidity/DvNftMarket.abi --pkg=ethmarket --out=./pkg/blockchain/ethmarket/smart.go
.PHONY: eth_contract_market