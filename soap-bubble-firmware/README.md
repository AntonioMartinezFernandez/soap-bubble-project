# Soap Bubble Firmware

## Hardware Requirements

- ESP32 NodeMCU (w/ antenna)
- IRF520N module
- Dupont cables
- USB cable
- 5V 2A DC power supply
- Soap Bubble Machine

## Software Requirements

- [Rust](https://www.rust-lang.org/tools/install)
- [espup](https://github.com/esp-rs/espup)
- [esp-generate](https://github.com/esp-rs/esp-generate)
- [espflash](https://github.com/esp-rs/espflash/blob/main/espflash/README.md)
- [probe-rs](https://probe.rs/docs/getting-started/installation/)
- [esp-rs/esp-idf-template](https://github.com/esp-rs/esp-idf-template)
- [VSCode](https://code.visualstudio.com/)
- [rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer)
- [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb)
- Other VSCode extensions: 'Code Runner', 'Dependi' and 'Rust Syntax'

## Install development environment

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

# Install support for Espressif SoCs
cargo install espup
espup install
# IMPORTANT: Copy the content of the $HOME/export-esp.sh file and paste it into the $HOME/.zshrc file
```

## Entering ESP32NodeMCU in flash mode

1. Connect to USB
2. Start press both RESET (EN) and BOOT buttons
3. Release first the RESET (EN) button and then also the BOOT button

## Compile and flash firmware

```bash
# Build project
cargo build --release

# Flash project
cargo run --release # Select /dev/tty.usbserial-0001 (or similar), where the ESP32 is connected
```

## Control the device with http requests

Switch on the device:

```bash
curl -X POST \
  http://192.168.1.200/status \
  -H 'content-type: application/json' \
  -d '{"status":"on", "speed": 100}'
```

Switch off the device:

```bash
curl -X POST \
 http://192.168.1.200/status \
 -H 'content-type: application/json' \
 -d '{"status":"off", "speed": 0}'
```
