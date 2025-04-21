# Soap Bubble Project

This project consists of two main components:

1. [Bubble Machine Firmware (Rust + ESP32)](./soap-bubble-firmware/README.md)
   A firmware written in Rust for an ESP32 microcontroller that controls a bubble machine. It exposes an HTTP endpoint to:

- Turn the bubble machine on or off
- Adjust the speed of operation

2. [Kubernetes Operator (Go + Kubebuilder)](./soap-bubble-operator/README.md)
   A Kubernetes operator written in Go using the Kubebuilder framework. It allows users to manage the bubble machine as a Kubernetes custom resource, enabling declarative control of the device from within a Kubernetes cluster.

Use this setup to integrate physical devices like a bubble machine into cloud-native infrastructure.

## Requirements

IDE:

- [VSCode](https://code.visualstudio.com/)
- [extension: Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [extension: rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer)
- [extension: Rust syntax](https://marketplace.visualstudio.com/items/?itemName=dustypomerleau.rust-syntax)

Containerization/Orchestration:

- [Kubernetes cluster (with Orbstack or Docker Desktop)](https://orbstack.dev/)

Programming languages compilers & tools:

- [Rust](https://www.rust-lang.org/tools/install)
- [Golang](https://go.dev/)
- [kubectl (configured)](https://kubernetes.io/docs/tasks/tools/)
- [kubebuilder](https://kubebuilder.io/quick-start)

Hardware:

- [ESP32 NodeMCU (w/ WiFi Antenna)](https://www.amazon.es/AZDelivery-NodeMCU-ESP-WROOM-32-Tablero-Desarrollo/dp/B071P98VTG)
- IRF520N module
- Dupont cables
- USB cable
- 5V 2A DC power supply
- Soap Bubble Machine

## Getting started

1. Install all the dependencies described above
1. Clone the repository
1. Set the WiFi SSID, WiFi password, and network configuration in the [config.toml](./soap-bubble-firmware/.cargo/config.toml) file
   - Instead of hardcoding WiFi SSID and password, you can define WIFI_SSID and WIFI_PASSWORD environment variables in your system
1. Flash the ESP32 NodeMCU with the soap bubble firmware
1. Mount the soap bubble machine electronic circuit
1. Install the soap bubble operator CRDs in the cluster
1. Power on the soap bubble machine
1. Start the soap bubble operator
1. Interact with the cluster creating/updating/deleting SoapBubbleMachine custom resources
1. Enjoy! 🎉

## Additional HOW-TOs

### How to create a Kubernetes Operator with kubebuilder (basic steps)

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

### How to create ESP32 Rust embedded project with esp-idf-template (basic steps)

```bash
# Install Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Install cargo sub-commands
cargo install cargo-generate
cargo install ldproxy
cargo install esp-generate
cargo install espflash
cargo install cargo-espflash
curl --proto '=https' --tlsv1.2 -LsSf https://github.com/probe-rs/probe-rs/releases/latest/download/probe-rs-tools-installer.sh | sh

# Install support for espressif SoCs
cargo install espup
espup install
# IMPORTANT: Copy the content of the $HOME/export-esp.sh file and paste it into the $HOME/.zshrc file

# Create project
cargo generate esp-rs/esp-idf-template cargo

# Build project
cargo build --release

# Flash project
cargo run --release # select /dev/tty.usbserial-0001 (or similar)
```

## Additional Resources

- [Setting up a MacBook for Rust + ESP32 (video)](https://www.youtube.com/watch?v=o4oTmUozaXA)
- [Rust on ESP32 course (video)](https://www.youtube.com/watch?v=o8yNNVFzNnM&list=PL0U7YUX2VnBFbwTi96wUB1nZzPVN3HzgS&index=1)
- [WiFi manager with ESP32 (C++)](https://randomnerdtutorials.com/esp32-wi-fi-manager-asyncwebserver/)
- [LittleFS with ESP32 (C++)](https://randomnerdtutorials.com/esp32-vs-code-platformio-littlefs/)
