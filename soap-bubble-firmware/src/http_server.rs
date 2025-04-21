use embedded_svc::{
    http::{Headers, Method},
    io::{Read, Write},
};
use esp_idf_hal::prelude::*;
use esp_idf_hal::{
    gpio::Pins,
    ledc::{config::TimerConfig, LedcDriver, LedcTimerDriver, LEDC},
};
use esp_idf_svc::http::server::EspHttpServer;
use log::info;
use serde::Deserialize;
use std::sync::{Arc, Mutex};

const STACK_SIZE: usize = 10240;
const MAX_LEN: usize = 128;

#[derive(Deserialize)]
struct SoapBubbleMachineConfig<'a> {
    status: &'a str,
    speed: u32,
}

pub fn start(pins: Pins, ledc: LEDC) -> anyhow::Result<EspHttpServer<'static>> {
    let server_configuration = esp_idf_svc::http::server::Configuration {
        stack_size: STACK_SIZE,
        ..Default::default()
    };

    // Create the server
    let mut server = EspHttpServer::new(&server_configuration)?;

    // Configure the PWM output (at 10kHz)
    let timer_driver = LedcTimerDriver::new(
        ledc.timer0,
        &TimerConfig::default().frequency(10.kHz().into()),
    )?;
    let mosfet_driver = LedcDriver::new(ledc.channel0, timer_driver, pins.gpio13).unwrap();
    let max_duty = mosfet_driver.get_max_duty();
    let mosfet = Arc::new(Mutex::new(mosfet_driver));

    // Register handlers
    server.fn_handler("/", Method::Get, |req| {
        req.into_ok_response()?
            .write_all("Hello from Soap Bubble Machine".as_bytes())
            .map(|_| ())
    })?;

    server.fn_handler::<anyhow::Error, _>("/status", Method::Post, move |mut req| {
        let len = req.content_len().unwrap_or(0) as usize;

        if len > MAX_LEN {
            req.into_status_response(413)?
                .write_all("Request too big".as_bytes())?;
            return Ok(());
        }

        let mut buf = vec![0; len];
        req.read_exact(&mut buf)?;
        let mut resp = req.into_ok_response()?;

        if let Ok(received_data) = serde_json::from_slice::<SoapBubbleMachineConfig>(&buf) {
            // Send the received data to serial port
            info!(
                "STATUS: {} - SPEED: {}",
                received_data.status, received_data.speed
            );

            // Set the PWM output (at mosfet pin)
            if received_data.status == "on" && received_data.speed <= 100 {
                mosfet
                    .lock()
                    .unwrap()
                    .set_duty(max_duty * received_data.speed / 100)
                    .unwrap();
            }
            if received_data.status == "off" {
                mosfet.lock().unwrap().set_duty(0).unwrap();
            }

            // Write the received data into the http response
            write!(
                resp,
                "Received status: {} - Received speed: {}",
                received_data.status, received_data.speed
            )?;
        } else {
            resp.write_all("JSON error".as_bytes())?;
        }

        Ok(())
    })?;

    Ok(server)
}
