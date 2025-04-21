package soapbubblemachineapplication

const SwitchOnSoapBubbleMachineCommandType = "switch-on-soap-bubble-machine-command.soap-bubble-operator.local"

type SwitchOnSoapBubbleMachineCommand struct {
	SoapBubbleMachineID   string
	SoapBubbleMachineName string
	SoapBubbleMachineIP   string
	Speed                 int
}

func NewSwitchOnSoapBubbleMachineCommand(
	soapBubbleMachineID,
	soapBubbleMachineName,
	ip string,
	speed int,
) *SwitchOnSoapBubbleMachineCommand {
	return &SwitchOnSoapBubbleMachineCommand{
		SoapBubbleMachineID:   soapBubbleMachineID,
		SoapBubbleMachineName: soapBubbleMachineName,
		SoapBubbleMachineIP:   ip,
		Speed:                 speed,
	}
}

func (c *SwitchOnSoapBubbleMachineCommand) Type() string {
	return SwitchOnSoapBubbleMachineCommandType
}
