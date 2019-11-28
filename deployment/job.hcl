job "quote-service" {
  datacenters = ["dc1"]
  group "quote" {
    task "server" {
      env {
        JAEGER_ENABLED = "true"
        JAEGER_AGENT_HOST = "127.0.0.1"
        JAEGER_AGENT_PORT = "6831"
      }
      driver = "docker"
      config {
        network_mode = "host"
        image = "saboteurkid/america-election-quote:latest"
      }
      resources {
        network {
          mbits = 1
          port "http" {
            static = 1323
          }
        }
      }
    }
  }
}