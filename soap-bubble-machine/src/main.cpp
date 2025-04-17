#include <Arduino.h>

#include <WiFi.h>
#include <ESPAsyncWebServer.h>

// WiFi credentials
#define WIFI_SSID "WIFI_SSID"
#define WIFI_PASSWORD "WIFI_PASSWORD"

// OUTPUT pin AND state
#define OUTPUT_PIN 2
int OUTPUT_STATE = LOW;

// Web server object
AsyncWebServer server(80);

// Function to generate the HTML and CSS code for the web page
String getHTML()
{
  String html = "<!DOCTYPE HTML>";
  html += "<html>";
  html += "<head>";
  html += "<style>";
  html += "body {background-color: #F0F0F0; font-family: Arial, Helvetica, sans-serif;}";
  html += "h1 {color: #333333; text-align: center;}";
  html += "button {width: 150px; height: 50px; font-size: 20px; margin: 10px;}";
  html += "</style>";
  html += "</head>";
  html += "<body>";
  html += "<h1>SOAP BUBBLE MACHINE</h1>";
  html += "<p>OUTPUT state: <span style='color: red;'>";
  if (OUTPUT_STATE == LOW)
    html += "OFF";
  else
    html += "ON";
  html += "</span></p>";
  html += "<button onclick=\"window.location.href='/output/on'\">Turn ON</button>";
  html += "<button onclick=\"window.location.href='/output/off'\">Turn OFF</button>";
  html += "</body>";
  html += "</html>";
  return html;
}

// Function to connect to WiFi network
void connectWiFi()
{
  Serial.print("Connecting to WiFi...");
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println();
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

// Function to handle HTTP requests
void handleRequest(AsyncWebServerRequest *request)
{
  // Get the request path
  String path = request->url();
  // Check if the request is to turn the OUTPUT on
  if (path == "/output/on")
  {
    // Set the OUTPUT pin to HIGH
    digitalWrite(OUTPUT_PIN, HIGH);
    // Update the OUTPUT state
    OUTPUT_STATE = HIGH;
    // Send a confirmation message
    request->send(200, "text/plain", "OUTPUT turned on");
  }
  // Check if the request is to turn the OUTPUT off
  else if (path == "/output/off")
  {
    // Set the OUTPUT pin to LOW
    digitalWrite(OUTPUT_PIN, LOW);
    // Update the OUTPUT state
    OUTPUT_STATE = LOW;
    // Send a confirmation message
    request->send(200, "text/plain", "OUTPUT turned off");
  }
  // Otherwise, send the web page
  else
  {
    // Get the HTML and CSS code
    String html = getHTML();
    // Send the web page
    request->send(200, "text/html", html);
  }
}

void setup()
{
  // Initialize the serial monitor
  Serial.begin(115200);

  // Initialize the OUTPUT pin
  pinMode(OUTPUT_PIN, OUTPUT);
  digitalWrite(OUTPUT_PIN, OUTPUT_STATE);

  // Connect to WiFi network
  connectWiFi();

  // Start the web server
  server.onNotFound(handleRequest);
  server.begin();
}

void loop()
{
  // Nothing to do here
}