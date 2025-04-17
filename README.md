# Soap Bubble Project

## Pre-requisites

### Hardware

- ESP32CAM
- USB-UART adapter
- USB cable
- Dupont cables

### Software

- VSCode
- PlatformIO
- Kubernetes Cluster (available with Orbstack or Docker Desktop)
- Golang
- kubectl
- kubebuilder

## Resources

1. https://randomnerdtutorials.com/esp32-vs-code-platformio-littlefs/
1. https://randomnerdtutorials.com/esp32-wi-fi-manager-asyncwebserver/

## Kubernetes Operator Creation

1. Scaffold the operator with kubebuilder

```bash
mkdir soap-bubble-operator
cd soap-bubble-operator
kubebuilder init --domain soap-bubble-operator.local --repo github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator
kubebuilder create api --group soap-bubble-operator --version v1alpha1 --kind SoapBubbleMachine

# Answer both questions with "yes"
```

2. Edit the `api/v1alpha1/soapbubblemachine_types.go` file and set the CRD schema
3. Edit the `config/samples/soap-bubble-operator_v1alpha1_soapbubblemachine.yaml` file and set the sample values for the CR
4. Execute `make manifests`
5. Execute `make build`
6. With the kubernetes cluster running, execute `make install`
