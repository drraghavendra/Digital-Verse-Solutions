package registry

import (
	"gitlab.com/digitalverse/blockchain/pkg/blockchain"
	"gitlab.com/digitalverse/blockchain/pkg/blockchain/generic"
)

func CreateRegistry(c *blockchain.Config) map[string]blockchain.Network {
	return map[string]blockchain.Network{
		"eth": generic.NewGenericNetwork(
			"https://rinkeby.etherscan.io/tx/%s",
			c.RinkebyEndpointUrl,
			c.RinkebyDeployWalletPk,
			c.RinkebyNftContractAddress,
			-1,
		),
		"near": generic.NewGenericNetwork(
			"https://explorer.testnet.aurora.dev/tx/%s",
			c.AuroraEndpointUrl,
			c.AuroraDeployWalletPk,
			c.AuroraNftContractAddress,
			c.AuroraShardID,
		),
	}
}
