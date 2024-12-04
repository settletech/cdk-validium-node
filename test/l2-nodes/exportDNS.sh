#!/bin/bash
services=("execution-layer" "dbs-postgres" "dbs-redis" "prover" "rpc" "dac")
namespace="local"
kubeconfig_path="/home/nachofq/.kube/config"

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo)"
  exit 1
fi

cp /etc/hosts /etc/hosts.bak

for service in "${services[@]}"; do
  sed -i.bak "/[[:space:]]$service$/d" /etc/hosts
done

for service in "${services[@]}"; do
    echo "127.0.0.1 $service" >> /etc/hosts
done

echo "Kubernetes hosts updated successfully."
cat /etc/hosts