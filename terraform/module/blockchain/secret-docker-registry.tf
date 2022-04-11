variable "docker-registry" {
  type = object({
    server   = string
    username = string
    password = string
  })
}

locals {
  registry_creds = jsonencode({
    "auths" : {
      (var.docker-registry.server) : {
        username = var.docker-registry.username
        password = var.docker-registry.password
        auth = base64encode(join(":", [
          var.docker-registry.username,
          var.docker-registry.password
        ]))
      }
    }
  })
}

resource "kubernetes_secret" "secret-docker-registry" {
  metadata {
    name      = "${local.name}-docker-registry"
    namespace = var.namespace
    labels    = local.labels
  }

  type = "kubernetes.io/dockerconfigjson"
  data = {
    ".dockerconfigjson" = local.registry_creds
  }
}
