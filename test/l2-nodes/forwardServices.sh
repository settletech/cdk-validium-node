#!/bin/bash

# Define the services and ports to export
SERVICES=(
  "rpc:8123:8123"
  "dbs-postgres:5432:5432"
  "prover:50061:50061"
  "prover:50071:50071"
  "execution-layer:8545:8545"
  "dac:8484:8484"
)

# Loop through the services and export each one
pkill -f "kubectl port-forward"
for service in "${SERVICES[@]}"; do
  name=${service%%:*}
  ports=${service#*:}
  kubectl port-forward services/$name $ports &
done