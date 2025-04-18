package soapbubblemachineinfra

import (
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
	resp, err := h.httpClient.Get(soapBubbleMachine.StartURL())
	if err != nil {
		return fmt.Errorf("failed to send start request: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func (h *HttpSoapBubbleMachineRemoteController) SwitchOff(ctx context.Context, soapBubbleMachine soapbubblemachinedomain.SoapBubbleMachine) error {
	resp, err := h.httpClient.Get(soapBubbleMachine.StopURL())
	if err != nil {
		return fmt.Errorf("failed to send stop request: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
