package soapbubblemachineapplication

const SwitchOnSoapBubbleMachineCommandType = "switch-on-soap-bubble-machine-command.soap-bubble-operator.local"

type SwitchOnSoapBubbleMachineCommand struct {
	SoapBubbleMachineID string
	Speed               int
	Namespace           string
}

func NewSwitchOnSoapBubbleMachineCommand(
	soapBubbleMachineID string,
	speed int,
	namespace string,
) *SwitchOnSoapBubbleMachineCommand {
	return &SwitchOnSoapBubbleMachineCommand{
		SoapBubbleMachineID: soapBubbleMachineID,
		Speed:               speed,
		Namespace:           namespace,
	}
}

func (c *SwitchOnSoapBubbleMachineCommand) Type() string {
	return SwitchOnSoapBubbleMachineCommandType
}
