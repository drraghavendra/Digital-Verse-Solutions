package blockchain

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

func CreateEthNftInstance(endpointUrl string, walletPk string, smartContractAddress string, chainId *big.Int) (instance *EthNft, txOptions *bind.TransactOpts, err error) {
	client, err := ethclient.Dial(endpointUrl)
	if err != nil {
		log.Error().Err(err).Msgf("ethclient.Dial = %s", endpointUrl)
		return
	}

	privateKey, err := crypto.HexToECDSA(walletPk)
	if err != nil {
		log.Error().Err(err).Msg("crypto.HexToECDSA")
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Err(err).Msg("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error().Err(err).Msg("client.PendingNonceAt")
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("client.SuggestGasPrice")
		return
	}

	// NewKeyedTransactorWithChainID
	if chainId.Int64() != -1 {
		txOptions, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			log.Error().Err(err).Msg("bind.NewKeyedTransactorWithChainID")
			return
		}
	} else {
		txOptions = bind.NewKeyedTransactor(privateKey)
	}

	txOptions.Nonce = big.NewInt(int64(nonce))
	txOptions.Value = big.NewInt(0)     // in wei
	txOptions.GasLimit = uint64(300000) // in units
	txOptions.GasPrice = gasPrice

	address := common.HexToAddress(smartContractAddress)
	instance, err = NewEthNft(address, client)
	if err != nil {
		log.Error().Err(err).Msg("NewEthNft")
		return
	}
	return instance, txOptions, nil
}

func MintNft(instance *EthNft, txOptions *bind.TransactOpts, metadataURI string) (txHash string, err error) {
	tx, err := instance.MintToken(txOptions, metadataURI)
	if err != nil {
		log.Error().Err(err).Msg("instance.MintToken")
		return
	}
	txHash = tx.Hash().Hex()
	log.Info().Msgf("tx sent: %s", txHash)
	return txHash, nil
}
