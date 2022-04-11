terraform {
  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "2.0.2"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.0.2"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.0.1"
    }
  }
  required_version = ">= 0.13"
}
