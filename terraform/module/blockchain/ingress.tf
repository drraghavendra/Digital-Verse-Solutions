#resource "kubernetes_ingress" "ingress" {
#  depends_on = [
#    kubernetes_service.app
#  ]
#
#  metadata {
#    namespace = var.namespace
#    name      = local.name
#    labels    = local.labels
#    annotations = {
#      "kubernetes.io/ingress.allow-http"            = "false"
#      "kubernetes.io/ingress.class"                 = "nginx"
#      "kubernetes.io/tls-acme"                      = "true"
#      "acme.cert-manager.io/http01-edit-in-place"   = "true"
#      "nginx.ingress.kubernetes.io/proxy-body-size" = "500m"
#    }
#  }
#
#  spec {
#    rule {
#      host = var.ingress.host
#
#      http {
#        path {
#          path = "/graphql"
#          backend {
#            service_name = kubernetes_service.app.metadata.0.name
#            service_port = var.app.listen_port
#          }
#        }
#
#        path {
#          path = "/graphql-play"
#          backend {
#            service_name = kubernetes_service.app.metadata.0.name
#            service_port = var.app.listen_port
#          }
#        }
#
#        path {
#          path = "/debug"
#          backend {
#            service_name = kubernetes_service.debug.metadata.0.name
#            service_port = 6060
#          }
#        }
#      }
#    }
#
#    tls {
#      secret_name = "graphql-cert-${var.slug}"
#      hosts = [
#        var.ingress.host
#      ]
#    }
#  }
#}
