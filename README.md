# Soap Bubble Project

## Pre-requisites

### Hardware

- ESP32CAM
- USB-UART adapter
- USB cable
- Dupont cables

### Software

IDE:

- [VSCode](https://code.visualstudio.com/)

Containerization/Orchestration:

- [Orbstack (or Docker Desktop)](https://orbstack.dev/)

Programming languages compilers & tools:

- [Rust](https://www.rust-lang.org/tools/install)
- [Golang](https://go.dev/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [kubebuilder](https://kubebuilder.io/quick-start)

Embedded development tools:

- [PlatformIO](https://platformio.org/install/ide?install=vscode)
- [espup](https://github.com/esp-rs/espup)
- [esp-generate](https://github.com/esp-rs/esp-generate)
- [espflash](https://github.com/esp-rs/espflash/blob/main/cargo-espflash/README.md)
- [espflash](https://github.com/esp-rs/espflash/blob/main/espflash/README.md)
- [probe-rs](https://probe.rs/docs/getting-started/installation/)

## How to create a Kubernetes Operator with kubebuilder (basic steps)

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

## Resources

- [LittleFS with ESP32](https://randomnerdtutorials.com/esp32-vs-code-platformio-littlefs/)
- [WiFi manager with ESP32](https://randomnerdtutorials.com/esp32-wi-fi-manager-asyncwebserver/)
- [Setting up a MacBook for Rust + ESP32 (video)](https://www.youtube.com/watch?v=o4oTmUozaXA)
