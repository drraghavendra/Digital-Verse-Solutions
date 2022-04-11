resource "kubernetes_service" "app" {
  depends_on = [
    kubernetes_deployment.app
  ]

  metadata {
    namespace = var.namespace
    name      = local.name
    labels    = local.labels
  }

  spec {
    type     = "ClusterIP"
    selector = local.labels

    port {
      port        = var.app.listen_port
      target_port = var.app.listen_port
    }
  }
}
