# gRPC-Hello-Server

This Go-based gRPC server provides two endpoints: SayHello, returning a simple greeting, and WhoAmI, which mimics Traefik’s /whoami by echoing the caller's IP address, user-agent, and hostname. The server is compiled into a statically linked binary and runs from a minimal scratch Docker image with no extra dependencies, making it ideal for secure and lightweight deployments in containerized environments like Kubernetes.

[Gateway API Example](https://github.com/nmdra/K8s-Learn/tree/fa50d02a43d4d976be1e03c0a4c28a63c109ea74/16-Gateway-API)