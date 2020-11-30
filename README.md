# network-scheduler-for-edge

## Build the network-scheduler

1. Use the Makefile to build your own Docker image.
2. Push the Docker image to a container registry.

## Deply the network-scheduler

1. Deploy the network-scheduler policy configuration:
  kubectl create -f deployments/scheduler-policy-config.yaml
2. Deploy the network-scheduler:
  kubectl create -f deployments/extender.yaml
3. Deploy your own pod. See deployments/test-pod.yaml and deployments/test-pod2.yaml
