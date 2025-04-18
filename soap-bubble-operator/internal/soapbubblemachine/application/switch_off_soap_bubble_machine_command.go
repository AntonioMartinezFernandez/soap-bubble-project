package soapbubblemachineapplication

const SwitchOffSoapBubbleMachineCommandType = "switch-off-soap-bubble-machine-command.soap-bubble-operator.local"

type SwitchOffSoapBubbleMachineCommand struct {
	SoapBubbleMachineID   string
	SoapBubbleMachineName string
	StartURL              string
	StopURL               string
	MakingBubbles         bool
}

func NewSwitchOffSoapBubbleMachineCommand(
	soapBubbleMachineID,
	soapBubbleMachineName,
	startURL, stopURL string,
	makingBubbles bool,
) *SwitchOffSoapBubbleMachineCommand {
	return &SwitchOffSoapBubbleMachineCommand{
		SoapBubbleMachineID:   soapBubbleMachineID,
		SoapBubbleMachineName: soapBubbleMachineName,
		StartURL:              startURL,
		StopURL:               stopURL,
		MakingBubbles:         makingBubbles,
	}
}

func (c *SwitchOffSoapBubbleMachineCommand) Type() string {
	return SwitchOffSoapBubbleMachineCommandType
}
