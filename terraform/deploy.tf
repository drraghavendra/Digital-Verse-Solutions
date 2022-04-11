module "blockchain" {
  source          = "./module/blockchain"
  namespace       = var.namespace
  slug            = var.slug
  docker-registry = local.docker-registry
  image           = {
    name = "${var.image}/app"
    tag  = var.image_version
  }
  app             = {
    listen_port    = var.listen_port
    sentry_dsn     = var.sentry_dsn
    aws_secret_key = var.aws_secret_key
    aws_access_key = var.aws_access_key
    aws_bucket     = var.aws_bucket
    aws_region     = var.aws_region
    listen_port    = var.listen_port

    rinkeby_deploy_wallet_pk            = var.rinkeby_deploy_wallet_pk
    rinkeby_endpoint_url                = var.rinkeby_endpoint_url
    rinkeby_nft_contract_address        = var.rinkeby_nft_contract_address
    rinkeby_nft_market_contract_address = var.rinkeby_nft_market_contract_address

    heco_deploy_wallet_pk            = var.heco_deploy_wallet_pk
    heco_endpoint_url                = var.heco_endpoint_url
    heco_nft_contract_address        = var.heco_nft_contract_address
    heco_nft_market_contract_address = var.heco_nft_market_contract_address

    aurora_deploy_wallet_pk            = var.aurora_deploy_wallet_pk
    aurora_endpoint_url                = var.aurora_endpoint_url
    aurora_nft_contract_address        = var.aurora_nft_contract_address
    aurora_nft_market_contract_address = var.aurora_nft_market_contract_address
    aurora_shard_id                    = var.aurora_shard_id

    polygon_deploy_wallet_pk            = var.polygon_deploy_wallet_pk
    polygon_endpoint_url                = var.polygon_endpoint_url
    polygon_nft_contract_address        = var.polygon_nft_contract_address
    polygon_nft_market_contract_address = var.polygon_nft_market_contract_address
    polygon_shard_id                    = var.polygon_shard_id

    nft_storage_key = var.nft_storage_key

    solana_testnet_deploy_wallet_pk     = var.solana_testnet_deploy_wallet_pk
    solana_testnet_endpoint_url         = var.solana_testnet_endpoint_url
    solana_testnet_nft_contract_address = var.solana_testnet_nft_contract_address

    harmony_testnet_deploy_wallet_pk     = var.harmony_testnet_deploy_wallet_pk
    harmony_testnet_endpoint_url         = var.harmony_testnet_endpoint_url
    harmony_testnet_nft_contract_address = var.harmony_testnet_nft_contract_address
    harmony_shard_id                     = var.harmony_shard_id

    celo_testnet_deploy_wallet_pk     = var.celo_testnet_deploy_wallet_pk
    celo_testnet_endpoint_url         = var.celo_testnet_endpoint_url
    celo_testnet_nft_contract_address = var.celo_testnet_nft_contract_address
    celo_testnet_shard_id             = var.celo_testnet_shard_id

    cronos_testnet_deploy_wallet_pk     = var.cronos_testnet_deploy_wallet_pk
    cronos_testnet_endpoint_url         = var.cronos_testnet_endpoint_url
    cronos_testnet_nft_contract_address = var.cronos_testnet_nft_contract_address
    cronos_testnet_shard_id             = var.cronos_testnet_shard_id

    theta_testnet_deploy_wallet_pk     = var.theta_testnet_deploy_wallet_pk
    theta_testnet_endpoint_url         = var.theta_testnet_endpoint_url
    theta_testnet_nft_contract_address = var.theta_testnet_nft_contract_address
    theta_testnet_shard_id             = var.theta_testnet_shard_id

    coinex_testnet_deploy_wallet_pk     = var.coinex_testnet_deploy_wallet_pk
    coinex_testnet_endpoint_url         = var.coinex_testnet_endpoint_url
    coinex_testnet_nft_contract_address = var.coinex_testnet_nft_contract_address
    coinex_testnet_shard_id             = var.coinex_testnet_shard_id

    okex_testnet_deploy_wallet_pk     = var.okex_testnet_deploy_wallet_pk
    okex_testnet_endpoint_url         = var.okex_testnet_endpoint_url
    okex_testnet_nft_contract_address = var.okex_testnet_nft_contract_address
    okex_testnet_shard_id             = var.okex_testnet_shard_id
  }
  ingress         = {
    host = var.host
  }
}
