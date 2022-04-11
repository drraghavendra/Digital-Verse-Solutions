locals {
  name = "blockchain-${var.slug}"
  labels = {
    app  = "blockchain"
    slug = var.slug
  }
}
