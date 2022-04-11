resource "kubernetes_secret" "app" {
  metadata {
    namespace = var.namespace
    name      = local.name
    labels    = local.labels
  }

  data = {
    LISTEN_PORT = var.app.listen_port
    SENTRY_DSN = var.app.sentry_dsn

    RINKEBY_DEPLOY_WALLET_PK = var.app.rinkeby_deploy_wallet_pk
    RINKEBY_ENDPOINT_URL = var.app.rinkeby_endpoint_url
    RINKEBY_NFT_CONTRACT_ADDRESS = var.app.rinkeby_nft_contract_address
    RINKEBY_NFT_MARKET_CONTRACT_ADDRESS = var.app.rinkeby_nft_market_contract_address

    HECO_DEPLOY_WALLET_PK = var.app.heco_deploy_wallet_pk
    HECO_ENDPOINT_URL = var.app.heco_endpoint_url
    HECO_NFT_CONTRACT_ADDRESS = var.app.heco_nft_contract_address
    HECO_NFT_MARKET_CONTRACT_ADDRESS = var.app.heco_nft_market_contract_address

    AURORA_DEPLOY_WALLET_PK = var.app.aurora_deploy_wallet_pk
    AURORA_ENDPOINT_URL = var.app.aurora_endpoint_url
    AURORA_NFT_CONTRACT_ADDRESS = var.app.aurora_nft_contract_address
    AURORA_NFT_MARKET_CONTRACT_ADDRESS = var.app.aurora_nft_market_contract_address
    AURORA_SHARD_ID = var.app.aurora_shard_id

    POLYGON_DEPLOY_WALLET_PK = var.app.polygon_deploy_wallet_pk
    POLYGON_ENDPOINT_URL = var.app.polygon_endpoint_url
    POLYGON_NFT_CONTRACT_ADDRESS = var.app.polygon_nft_contract_address
    POLYGON_NFT_MARKET_CONTRACT_ADDRESS = var.app.polygon_nft_market_contract_address
    POLYGON_SHARD_ID = var.app.polygon_shard_id

    NFT_STORAGE_KEY = var.app.nft_storage_key

    SOLANA_TESTNET_DEPLOY_WALLET_PK = var.app.solana_testnet_deploy_wallet_pk
    SOLANA_TESTNET_ENDPOINT_URL = var.app.solana_testnet_endpoint_url
    SOLANA_TESTNET_NFT_CONTRACT_ADDRESS = var.app.solana_testnet_nft_contract_address

    HARMONY_TESTNET_DEPLOY_WALLET_PK = var.app.harmony_testnet_deploy_wallet_pk
    HARMONY_TESTNET_ENDPOINT_URL = var.app.harmony_testnet_endpoint_url
    HARMONY_TESTNET_NFT_CONTRACT_ADDRESS = var.app.harmony_testnet_nft_contract_address
    HARMONY_SHARD_ID = var.app.harmony_shard_id

    CELO_TESTNET_DEPLOY_WALLET_PK = var.app.celo_testnet_deploy_wallet_pk
    CELO_TESTNET_ENDPOINT_URL = var.app.celo_testnet_endpoint_url
    CELO_TESTNET_NFT_CONTRACT_ADDRESS = var.app.celo_testnet_nft_contract_address
    CELO_TESTNET_SHARD_ID = var.app.celo_testnet_shard_id

    CRONOS_TESTNET_DEPLOY_WALLET_PK = var.app.cronos_testnet_deploy_wallet_pk
    CRONOS_TESTNET_ENDPOINT_URL = var.app.cronos_testnet_endpoint_url
    CRONOS_TESTNET_NFT_CONTRACT_ADDRESS = var.app.cronos_testnet_nft_contract_address
    CRONOS_TESTNET_SHARD_ID = var.app.cronos_testnet_shard_id

    THETA_TESTNET_DEPLOY_WALLET_PK = var.app.theta_testnet_deploy_wallet_pk
    THETA_TESTNET_ENDPOINT_URL = var.app.theta_testnet_endpoint_url
    THETA_TESTNET_NFT_CONTRACT_ADDRESS = var.app.theta_testnet_nft_contract_address
    THETA_TESTNET_SHARD_ID = var.app.theta_testnet_shard_id

    COINEX_TESTNET_DEPLOY_WALLET_PK = var.app.coinex_testnet_deploy_wallet_pk
    COINEX_TESTNET_ENDPOINT_URL = var.app.coinex_testnet_endpoint_url
    COINEX_TESTNET_NFT_CONTRACT_ADDRESS = var.app.coinex_testnet_nft_contract_address
    COINEX_TESTNET_SHARD_ID = var.app.coinex_testnet_shard_id

    OKEX_TESTNET_DEPLOY_WALLET_PK = var.app.okex_testnet_deploy_wallet_pk
    OKEX_TESTNET_ENDPOINT_URL = var.app.okex_testnet_endpoint_url
    OKEX_TESTNET_NFT_CONTRACT_ADDRESS = var.app.okex_testnet_nft_contract_address
    OKEX_TESTNET_SHARD_ID = var.app.okex_testnet_shard_id
  }
}
