use core::convert::TryInto;

use std::net::Ipv4Addr;
use std::str::FromStr;

use embedded_svc::wifi::{AuthMethod, ClientConfiguration, Configuration as WifiConfiguration};

use esp_idf_svc::ipv4::{
    ClientConfiguration as IpClientConfiguration, ClientSettings as IpClientSettings,
    Configuration as IpConfiguration, Mask, Subnet,
};
use esp_idf_svc::netif::{EspNetif, NetifConfiguration, NetifStack};
use esp_idf_svc::wifi::{BlockingWifi, EspWifi, WifiDriver};

use log::info;

// The SSID and password of the access point to connect to
const SSID: &str = env!("WIFI_SSID");
const PASSWORD: &str = env!("WIFI_PASSWORD");
// Expects IPv4 address as device IP
const DEVICE_IP: &str = env!("DEVICE_IP");
// Expects IPv4 address as gateway IP
const GATEWAY_IP: &str = env!("GATEWAY_IP");
// Expects a number between 0 and 32 (24 is equivalent to 255.255.255.0)
const GATEWAY_NETMASK: &str = env!("GATEWAY_NETMASK");

pub fn configure(wifi: WifiDriver) -> anyhow::Result<EspWifi> {
    let netmask = u8::from_str(GATEWAY_NETMASK)?;
    let gateway_addr = Ipv4Addr::from_str(GATEWAY_IP)?;
    let static_ip = Ipv4Addr::from_str(DEVICE_IP)?;

    let mut wifi = EspWifi::wrap_all(
        wifi,
        EspNetif::new_with_conf(&NetifConfiguration {
            ip_configuration: Some(IpConfiguration::Client(IpClientConfiguration::Fixed(
                IpClientSettings {
                    ip: static_ip,
                    subnet: Subnet {
                        gateway: gateway_addr,
                        mask: Mask(netmask),
                    },
                    // Can also be set to Ipv4Addrs if you need DNS
                    dns: None,
                    secondary_dns: None,
                },
            ))),
            ..NetifConfiguration::wifi_default_client()
        })?,
        #[cfg(esp_idf_esp_wifi_softap_support)]
        EspNetif::new(NetifStack::Ap)?,
    )?;

    let wifi_configuration = WifiConfiguration::Client(ClientConfiguration {
        ssid: SSID.try_into().unwrap(),
        bssid: None,
        auth_method: AuthMethod::WPA2Personal,
        password: PASSWORD.try_into().unwrap(),
        channel: None,
        ..Default::default()
    });
    wifi.set_configuration(&wifi_configuration)?;

    Ok(wifi)
}

pub fn connect(wifi: &mut BlockingWifi<EspWifi<'static>>) -> anyhow::Result<()> {
    wifi.start()?;
    info!("Wifi started");

    wifi.connect()?;
    info!("Wifi connected");

    wifi.wait_netif_up()?;
    info!("Wifi netif up");

    Ok(())
}
