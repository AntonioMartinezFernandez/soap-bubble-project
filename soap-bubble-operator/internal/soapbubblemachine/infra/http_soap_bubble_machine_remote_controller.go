package soapbubblemachineinfra

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
)

var _ soapbubblemachinedomain.SoapBubbleMachineRemoteController = (*HttpSoapBubbleMachineRemoteController)(nil)

type HttpSoapBubbleMachineRemoteController struct {
	httpClient http.Client
}

func NewHttpSoapBubbleMachineRemoteController() *HttpSoapBubbleMachineRemoteController {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	return &HttpSoapBubbleMachineRemoteController{
		httpClient: httpClient,
	}
}

func (h *HttpSoapBubbleMachineRemoteController) SwitchOn(ctx context.Context, soapBubbleMachine soapbubblemachinedomain.SoapBubbleMachine) error {
	// Create the request
	jsonBody := []byte(fmt.Sprintf(`{"status": "on", "speed": %d}`, soapBubbleMachine.Speed()))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url(soapBubbleMachine.IP()), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create start request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send start request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to start soap bubble machine, status code: %d", resp.StatusCode)
	}

	return nil
}

func (h *HttpSoapBubbleMachineRemoteController) SwitchOff(ctx context.Context, soapBubbleMachine soapbubblemachinedomain.SoapBubbleMachine) error {
	// Create the request
	jsonBody := []byte(fmt.Sprintf(`{"status": "off", "speed": %d}`, soapBubbleMachine.Speed()))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url(soapBubbleMachine.IP()), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create stop request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send stop request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to stop soap bubble machine, status code: %d", resp.StatusCode)
	}

	return nil
}

func url(ip string) string {
	return "http://" + ip + ":80/status"
}
