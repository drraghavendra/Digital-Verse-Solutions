terraform {
  backend "http" {}
}

provider "kubernetes" {
  host                   = var.k8s_host
  token                  = base64decode(var.k8s_token)
  cluster_ca_certificate = base64decode(var.k8s_ca_crt)
}
