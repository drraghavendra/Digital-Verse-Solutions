package generic

import (
	"fmt"
	"math/big"

	"github.com/rs/zerolog/log"
	"gitlab.com/digitalverse/blockchain/pkg/blockchain"
)

func NewGenericNetwork(
	urlTemplate, endpointURL, deployWalletPk, nftContractAddress string,
	shardID int64,
) blockchain.Network {
	return &genericNetwork{
		urlTemplate:        urlTemplate,
		endpointURL:        endpointURL,
		deployWalletPk:     deployWalletPk,
		nftContractAddress: nftContractAddress,
		shardID:            shardID,
	}
}

type genericNetwork struct {
	urlTemplate        string
	endpointURL        string
	deployWalletPk     string
	nftContractAddress string
	shardID            int64
}

func (n *genericNetwork) Mint(metadataURI string) (*blockchain.Resp, error) {
	instance, txOptions, err := blockchain.CreateEthNftInstance(
		n.endpointURL,
		n.deployWalletPk,
		n.nftContractAddress,
		big.NewInt(n.shardID),
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create contract instance")
		return nil, err
	}

	txHash, err := blockchain.MintNft(instance, txOptions, metadataURI)
	if err != nil {
		log.Error().Err(err).Msg("MintNft")
		return nil, err
	}

	return &blockchain.Resp{
		TransactionHash: txHash,
		TransactionURL:  fmt.Sprintf(n.urlTemplate, txHash),
		FileURL:         blockchain.CreateFileUrl(metadataURI),
	}, nil
}
