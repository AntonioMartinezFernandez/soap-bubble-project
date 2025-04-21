package soapbubblemachineapplication

const SwitchOffSoapBubbleMachineCommandType = "switch-off-soap-bubble-machine-command.soap-bubble-operator.local"

type SwitchOffSoapBubbleMachineCommand struct {
	SoapBubbleMachineID   string
	SoapBubbleMachineName string
	SoapBubbleMachineIP   string
}

func NewSwitchOffSoapBubbleMachineCommand(
	id,
	name,
	ip string,
) *SwitchOffSoapBubbleMachineCommand {
	return &SwitchOffSoapBubbleMachineCommand{
		SoapBubbleMachineID:   id,
		SoapBubbleMachineName: name,
		SoapBubbleMachineIP:   ip,
	}
}

func (c *SwitchOffSoapBubbleMachineCommand) Type() string {
	return SwitchOffSoapBubbleMachineCommandType
}
