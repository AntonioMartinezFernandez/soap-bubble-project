package soapbubblemachineapplication

const SwitchOnSoapBubbleMachineCommandType = "switch-on-soap-bubble-machine-command.soap-bubble-operator.local"

type SwitchOnSoapBubbleMachineCommand struct {
	SoapBubbleMachineID   string
	SoapBubbleMachineName string
	StartURL              string
	StopURL               string
	MakingBubbles         bool
}

func NewSwitchOnSoapBubbleMachineCommand(
	soapBubbleMachineID,
	soapBubbleMachineName,
	startURL, stopURL string,
	makingBubbles bool,
) *SwitchOnSoapBubbleMachineCommand {
	return &SwitchOnSoapBubbleMachineCommand{
		SoapBubbleMachineID:   soapBubbleMachineID,
		SoapBubbleMachineName: soapBubbleMachineName,
		StartURL:              startURL,
		StopURL:               stopURL,
		MakingBubbles:         makingBubbles,
	}
}

func (c *SwitchOnSoapBubbleMachineCommand) Type() string {
	return SwitchOnSoapBubbleMachineCommandType
}
