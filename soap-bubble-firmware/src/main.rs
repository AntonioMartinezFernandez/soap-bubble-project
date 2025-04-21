#![allow(unknown_lints)]
#![allow(unexpected_cfgs)]

use esp_idf_svc::hal::prelude::Peripherals;
use esp_idf_svc::log::EspLogger;
use esp_idf_svc::wifi::{BlockingWifi, WifiDriver};
use esp_idf_svc::{eventloop::EspSystemEventLoop, nvs::EspDefaultNvsPartition};
use log::info;

pub mod wifi_connection;
use wifi_connection::{configure, connect};

pub mod http_server;
use http_server::start;

fn main() -> anyhow::Result<()> {
    esp_idf_svc::sys::link_patches();
    EspLogger::initialize_default();

    // Initialize peripherals and system event loop
    let peripherals = Peripherals::take()?;
    let sys_loop = EspSystemEventLoop::take()?;
    let nvs = EspDefaultNvsPartition::take()?;

    // Connect to WiFi
    let wifi = WifiDriver::new(peripherals.modem, sys_loop.clone(), Some(nvs))?;
    let wifi = configure(wifi)?;
    let mut wifi = BlockingWifi::wrap(wifi, sys_loop)?;
    connect(&mut wifi)?;
    let ip_info = wifi.wifi().sta_netif().get_ip_info()?;
    info!("WiFi info: {:?}", ip_info);

    // Create HTTP server
    let started_server = start(peripherals.pins);
    if started_server.is_err() {
        info!("Failed to start HTTP server");
    }

    // Initialize main loop
    loop {
        std::thread::sleep(core::time::Duration::from_secs(10));
    }
}
