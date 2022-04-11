locals {
  docker-registry = {
    password = var.docker_registry_password
    username = var.docker_registry_username
    server   = var.docker_registry_server
  }
}
