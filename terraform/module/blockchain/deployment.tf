resource "kubernetes_deployment" "app" {
  depends_on = [
    kubernetes_secret.secret-docker-registry,
    kubernetes_secret.app,
  ]

  metadata {
    namespace = var.namespace
    name      = local.name
    labels    = local.labels
  }

  spec {
    replicas = var.deployment.replicas

    selector {
      match_labels = local.labels
    }

    template {
      metadata {
        labels = local.labels
      }

      spec {
        image_pull_secrets {
          name = kubernetes_secret.secret-docker-registry.metadata.0.name
        }

        container {
          name = "app"
          image = join(":", [
            var.image.name,
            var.image.tag
          ])
          image_pull_policy = var.image.tag == "latest" ? "Always" : "IfNotPresent"

          port {
            container_port = var.app.listen_port
          }

          env_from {
            secret_ref {
              name = kubernetes_secret.app.metadata.0.name
            }
          }

          resources {
            requests = {
              cpu    = var.deployment.requests.cpu
              memory = var.deployment.requests.memory
            }

            limits = {
              cpu    = var.deployment.limits.cpu
              memory = var.deployment.limits.memory
            }
          }
        }
      }
    }
  }
}
