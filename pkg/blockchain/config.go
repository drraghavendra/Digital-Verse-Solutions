package blockchain

// type GenericNetworkConfig struct {
// 	EndpointUrl              string `env:"ENDPOINT_URL"`
// 	DeployWalletPk           string `env:"DEPLOY_WALLET_PK"`
// 	DeployWalletAddress      string `env:"DEPLOY_WALLET_ADDRESS"`
// 	ShardID                  int64  `env:"SHARD_ID"`
// 	NftContractAddress       string `env:"NFT_CONTRACT_ADDRESS"`
// 	NftMarketContractAddress string `env:"NFT_MARKET_CONTRACT_ADDRESS"`
// }

// type NewConfig struct {
// 	Rinkeby GenericNetworkConfig `env:"RINKEBY"`
// 	Heco    GenericNetworkConfig `env:"HECO"`
// 	Aurora  GenericNetworkConfig `env:"AURORA"`
// 	Polygon GenericNetworkConfig `env:"POLYGON"`
// 	Ropsten GenericNetworkConfig `env:"ROPSTEN"`
// 	Solana  GenericNetworkConfig `env:"SOLANA"`
// 	Harmony GenericNetworkConfig `env:"HARMONY"`
// 	Velos   GenericNetworkConfig `env:"VELOS"`
// 	Coinex  GenericNetworkConfig `env:"COINEX"`
// 	Theta   GenericNetworkConfig `env:"THETA"`
// 	Chronos GenericNetworkConfig `env:"CHRONOS"`
// 	OkEx    GenericNetworkConfig `env:"OKEX"`
// 	Celo    GenericNetworkConfig `env:"CELO"`
// }

type Config struct {
	NftStorageKey string `env:"NFT_STORAGE_KEY"`

	RinkebyDeployWalletPk           string `env:"RINKEBY_DEPLOY_WALLET_PK"`
	RinkebyEndpointUrl              string `env:"RINKEBY_ENDPOINT_URL"`
	RinkebyNftContractAddress       string `env:"RINKEBY_NFT_CONTRACT_ADDRESS"`
	RinkebyNftMarketContractAddress string `env:"RINKEBY_NFT_MARKET_CONTRACT_ADDRESS"`

	HecoDeployWalletPk           string `env:"HECO_DEPLOY_WALLET_PK"`
	HecoEndpointUrl              string `env:"HECO_ENDPOINT_URL"`
	HecoNftContractAddress       string `env:"HECO_NFT_CONTRACT_ADDRESS"`
	HecoNftMarketContractAddress string `env:"HECO_NFT_MARKET_CONTRACT_ADDRESS"`

	AuroraDeployWalletPk           string `env:"AURORA_DEPLOY_WALLET_PK"`
	AuroraEndpointUrl              string `env:"AURORA_ENDPOINT_URL"`
	AuroraNftContractAddress       string `env:"AURORA_NFT_CONTRACT_ADDRESS"`
	AuroraNftMarketContractAddress string `env:"AURORA_NFT_MARKET_CONTRACT_ADDRESS"`
	AuroraShardID                  int64  `env:"AURORA_SHARD_ID"`

	PolygonDeployWalletPk           string `env:"POLYGON_DEPLOY_WALLET_PK"`
	PolygonEndpointUrl              string `env:"POLYGON_ENDPOINT_URL"`
	PolygonNftContractAddress       string `env:"POLYGON_NFT_CONTRACT_ADDRESS"`
	PolygonNftMarketContractAddress string `env:"POLYGON_NFT_MARKET_CONTRACT_ADDRESS"`
	PolygonShardID                  int64  `env:"POLYGON_SHARD_ID"`

	SolanaTestnetDeployWalletPk     string `env:"SOLANA_TESTNET_DEPLOY_WALLET_PK"`
	SolanaTestnetEndpointUrl        string `env:"SOLANA_TESTNET_ENDPOINT_URL"`
	SolanaTestnetNftContractAddress string `env:"SOLANA_TESTNET_NFT_CONTRACT_ADDRESS"`

	HarmonyTestnetDeployWalletPk     string `env:"HARMONY_TESTNET_DEPLOY_WALLET_PK"`
	HarmonyTestnetEndpointUrl        string `env:"HARMONY_TESTNET_ENDPOINT_URL"`
	HarmonyTestnetNftContractAddress string `env:"HARMONY_TESTNET_NFT_CONTRACT_ADDRESS"`
	HarmonyShardID                   int64  `env:"HARMONY_SHARD_ID"`

	VelasTestnetDeployWalletPk     string `env:"VELAS_TESTNET_DEPLOY_WALLET_PK"`
	VelasTestnetEndpointUrl        string `env:"VELAS_TESTNET_ENDPOINT_URL"`
	VelasTestnetNftContractAddress string `env:"VELAS_TESTNET_NFT_CONTRACT_ADDRESS"`
	VelasShardID                   int64  `env:"VELAS_SHARD_ID"`

	OKExChainTestnetDeployWalletPk     string `env:"OKEXChain_TESTNET_DEPLOY_WALLET_PK"`
	OKExChainTestnetEndpointUrl        string `env:"OKEXChain_TESTNET_ENDPOINT_URL"`
	OKExChainTestnetNftContractAddress string `env:"OKEXChain_TESTNET_NFT_CONTRACT_ADDRESS"`
	OKExChainTestnetShardID            int64  `env:"OKEXChain_TESTNET_SHARD_ID"`

	CronosTestnetDeployWalletPk     string `env:"CRONOS_TESTNET_DEPLOY_WALLET_PK"`
	CronosTestnetEndpointUrl        string `env:"CRONOS_TESTNET_ENDPOINT_URL"`
	CronosTestnetNftContractAddress string `env:"CRONOS_TESTNET_NFT_CONTRACT_ADDRESS"`
	CronosTestnetShardID            int64  `env:"CRONOS_SHARD_TESTNET_ID"`

	CeloTestnetDeployWalletPk     string `env:"CELO_TESTNET_DEPLOY_WALLET_PK"`
	CeloTestnetEndpointUrl        string `env:"CELO_TESTNET_ENDPOINT_URL"`
	CeloTestnetNftContractAddress string `env:"CELO_TESTNET_NFT_CONTRACT_ADDRESS"`
	CeloTestnetShardID            int64  `env:"CELO_SHARD_TESTNET_ID"`

	ThetaTestnetDeployWalletPk     string `env:"THETA_TESTNET_DEPLOY_WALLET_PK"`
	ThetaTestnetEndpointUrl        string `env:"THETA_TESTNET_ENDPOINT_URL"`
	ThetaTestnetNftContractAddress string `env:"THETA_TESTNET_NFT_CONTRACT_ADDRESS"`
	ThetaTestnetShardID            int64  `env:"THETA_SHARD_TESTNET_ID"`

	CoinExTestnetDeployWalletPk     string `env:"COINEX_TESTNET_DEPLOY_WALLET_PK"`
	CoinExTestnetEndpointUrl        string `env:"COINEX_TESTNET_ENDPOINT_URL"`
	CoinExTestnetNftContractAddress string `env:"COINEX_TESTNET_NFT_CONTRACT_ADDRESS"`
	CoinExTestnetShardID            int64  `env:"COINEX_SHARD_TESTNET_ID"`
}
