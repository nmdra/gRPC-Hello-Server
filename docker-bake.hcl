
variable "TAG" {
  default = "latest"
}

group "default" {
  targets = ["gRPC-Hello-Server"]
}

target "gRPC-Hello-Server" {
  dockerfile = "Dockerfile"
  context    = "."
  tags = [
    "ghcr.io/nmdra/grpc-hello-server:latest",
    "ghcr.io/nmdra/grpc-hello-server:${TAG}",
  ]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  labels = {
    "org.opencontainers.image.title"       = "gRPC-Hello-Server"
    "org.opencontainers.image.created" = "${timestamp()}"
    "org.opencontainers.image.version"     = "${TAG}"
    "org.opencontainers.image.source"      = "https://github.com/nmdra/gRPC-Hello-Server"
    "org.opencontainers.image.licenses"    = "MIT"
    "org.opencontainers.image.description" = "Minimal gRPC semantic search server written in Go."
  }
  args = {
    CGO_ENABLED = "0"
  }
}