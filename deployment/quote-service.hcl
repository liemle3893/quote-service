job "quote-service" {
  group "quote" {
    task "server" {
      driver = "docker"
      config {
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