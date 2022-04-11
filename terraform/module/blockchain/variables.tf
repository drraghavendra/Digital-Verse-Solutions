variable "namespace" {
  type = string
}

variable "slug" {
  type = string
}

variable "image" {
  type = object({
    name = string
    tag = string
  })
}

variable "ingress" {
  type = object({
    host = string
  })
}

variable "deployment" {
  type = object({
    replicas = number,
    requests = object({
      cpu = string
      memory = string
    })
    limits = object({
      cpu = string
      memory = string
    })
  })
  default = {
    replicas = 1,
    requests = {
      cpu = "2"
      memory = "512Mi"
    }
    limits = {
      cpu = "2"
      memory = "512Mi"
    }
  }
}

variable "app" {
  type = object({
    listen_port = number
    sentry_dsn = string

    rinkeby_deploy_wallet_pk = string
    rinkeby_endpoint_url = string
    rinkeby_nft_contract_address = string
    rinkeby_nft_market_contract_address = string

    heco_deploy_wallet_pk = string
    heco_endpoint_url = string
    heco_nft_contract_address = string
    heco_nft_market_contract_address = string

    aurora_deploy_wallet_pk = string
    aurora_endpoint_url = string
    aurora_nft_contract_address = string
    aurora_nft_market_contract_address = string
    aurora_shard_id = number

    polygon_deploy_wallet_pk = string
    polygon_endpoint_url = string
    polygon_nft_contract_address = string
    polygon_nft_market_contract_address = string
    polygon_shard_id = number

    nft_storage_key = string

    solana_testnet_deploy_wallet_pk = string
    solana_testnet_endpoint_url = string
    solana_testnet_nft_contract_address = string

    harmony_testnet_deploy_wallet_pk = string
    harmony_testnet_endpoint_url = string
    harmony_testnet_nft_contract_address = string
    harmony_shard_id = string

    celo_testnet_deploy_wallet_pk = string
    celo_testnet_endpoint_url = string
    celo_testnet_nft_contract_address = string
    celo_testnet_shard_id = string

    cronos_testnet_deploy_wallet_pk = string
    cronos_testnet_endpoint_url = string
    cronos_testnet_nft_contract_address = string
    cronos_testnet_shard_id = string

    theta_testnet_deploy_wallet_pk = string
    theta_testnet_endpoint_url = string
    theta_testnet_nft_contract_address = string
    theta_testnet_shard_id = string

    coinex_testnet_deploy_wallet_pk = string
    coinex_testnet_endpoint_url = string
    coinex_testnet_nft_contract_address = string
    coinex_testnet_shard_id = string

    okex_testnet_deploy_wallet_pk = string
    okex_testnet_endpoint_url = string
    okex_testnet_nft_contract_address = string
    okex_testnet_shard_id = string
  })
}
