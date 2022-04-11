variable "aws_secret_key" { type = string }
variable "aws_access_key" { type = string }
variable "aws_bucket" { type = string }
variable "aws_region" { type = string }

variable "sentry_dsn" { type = string }

variable "listen_port" {
  type = number
  default = 4000
}

variable "rinkeby_deploy_wallet_pk" { type = string }
variable "rinkeby_endpoint_url" { type = string }
variable "rinkeby_nft_contract_address" { type = string }
variable "rinkeby_nft_market_contract_address" { type = string }

variable "heco_deploy_wallet_pk" { type = string }
variable "heco_endpoint_url" { type = string }
variable "heco_nft_contract_address" { type = string }
variable "heco_nft_market_contract_address" { type = string }

variable "aurora_deploy_wallet_pk" { type = string }
variable "aurora_endpoint_url" { type = string }
variable "aurora_nft_contract_address" { type = string }
variable "aurora_nft_market_contract_address" { type = string }
variable "aurora_shard_id" { type = number }

variable "polygon_deploy_wallet_pk" { type = string }
variable "polygon_endpoint_url" { type = string }
variable "polygon_nft_contract_address" { type = string }
variable "polygon_nft_market_contract_address" { type = string }
variable "polygon_shard_id" { type = number }

variable "nft_storage_key" { type = string }

variable "solana_testnet_deploy_wallet_pk" { type = string }
variable "solana_testnet_endpoint_url" { type = string }
variable "solana_testnet_nft_contract_address" { type = string }

variable "harmony_testnet_deploy_wallet_pk" { type = string }
variable "harmony_testnet_endpoint_url" { type = string }
variable "harmony_testnet_nft_contract_address" { type = string }
variable "harmony_shard_id" { type = string }

variable "celo_testnet_deploy_wallet_pk" { type = string }
variable "celo_testnet_endpoint_url" { type = string }
variable "celo_testnet_nft_contract_address" { type = string }
variable "celo_testnet_shard_id" { type = string }

variable "cronos_testnet_deploy_wallet_pk" { type = string }
variable "cronos_testnet_endpoint_url" { type = string }
variable "cronos_testnet_nft_contract_address" { type = string }
variable "cronos_testnet_shard_id" { type = string }

variable "theta_testnet_deploy_wallet_pk" { type = string }
variable "theta_testnet_endpoint_url" { type = string }
variable "theta_testnet_nft_contract_address" { type = string }
variable "theta_testnet_shard_id" { type = string }

variable "coinex_testnet_deploy_wallet_pk" { type = string }
variable "coinex_testnet_endpoint_url" { type = string }
variable "coinex_testnet_nft_contract_address" { type = string }
variable "coinex_testnet_shard_id" { type = string }

variable "okex_testnet_deploy_wallet_pk" { type = string }
variable "okex_testnet_endpoint_url" { type = string }
variable "okex_testnet_nft_contract_address" { type = string }
variable "okex_testnet_shard_id" { type = string }
