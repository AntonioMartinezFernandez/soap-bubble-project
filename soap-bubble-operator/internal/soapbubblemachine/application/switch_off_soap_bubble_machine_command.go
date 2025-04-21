package soapbubblemachineapplication

const SwitchOffSoapBubbleMachineCommandType = "switch-off-soap-bubble-machine-command.soap-bubble-operator.local"

type SwitchOffSoapBubbleMachineCommand struct {
	SoapBubbleMachineID string
	Namespace           string
}

func NewSwitchOffSoapBubbleMachineCommand(
	id,
	namespace string,
) *SwitchOffSoapBubbleMachineCommand {
	return &SwitchOffSoapBubbleMachineCommand{
		SoapBubbleMachineID: id,
		Namespace:           namespace,
	}
}

func (c *SwitchOffSoapBubbleMachineCommand) Type() string {
	return SwitchOffSoapBubbleMachineCommandType
}
