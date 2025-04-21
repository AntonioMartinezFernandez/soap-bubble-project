use std::sync::{Arc, Mutex};

use embedded_svc::{
    http::{Headers, Method},
    io::{Read, Write},
};
use esp_idf_hal::gpio::{PinDriver, Pins};
use esp_idf_svc::http::server::EspHttpServer;
use log::info;
use serde::Deserialize;

const STACK_SIZE: usize = 10240;
const MAX_LEN: usize = 128;

#[derive(Deserialize)]
struct SoapBubbleMachineConfig<'a> {
    status: &'a str,
    speed: u32,
}

pub fn start(pins: Pins) -> anyhow::Result<EspHttpServer<'static>> {
    let server_configuration = esp_idf_svc::http::server::Configuration {
        stack_size: STACK_SIZE,
        ..Default::default()
    };

    // Create the server
    let mut server = EspHttpServer::new(&server_configuration)?;

    // Register handlers
    server.fn_handler("/", Method::Get, |req| {
        req.into_ok_response()?
            .write_all("Hello from Soap Bubble Machine".as_bytes())
            .map(|_| ())
    })?;

    let mosfet = Arc::new(Mutex::new(PinDriver::output(pins.gpio13).unwrap()));

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

        if let Ok(form) = serde_json::from_slice::<SoapBubbleMachineConfig>(&buf) {
            // Print the form data to the console
            info!("STATUS: {} - SPEED: {}", form.status, form.speed);

            // Set the requested mosfet output
            if form.status == "on" {
                mosfet.lock().unwrap().set_high().unwrap();
            }
            if form.status == "off" {
                mosfet.lock().unwrap().set_low().unwrap();
            }

            // Write the form data to the response
            write!(
                resp,
                "Received status: {} - Received speed: {}",
                form.status, form.speed
            )?;
        } else {
            resp.write_all("JSON error".as_bytes())?;
        }

        Ok(())
    })?;

    Ok(server)
}
