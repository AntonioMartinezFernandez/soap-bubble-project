package soapbubblemachinedomain

import "strings"

type SoapBubbleMachineID string

func NewSoapBubbleMachineID(namespace, name string) SoapBubbleMachineID {
	return SoapBubbleMachineID(namespace + "-" + name)
}

func (id SoapBubbleMachineID) Namespace() string {
	parts := strings.Split(string(id), "-")
	if len(parts) < 2 {
		return ""
	}
	return parts[0]
}

func (id SoapBubbleMachineID) Name() string {
	parts := strings.Split(string(id), "-")
	if len(parts) < 2 {
		return ""
	}
	return strings.Join(parts[1:], "-")
}

func (id SoapBubbleMachineID) String() string {
	return string(id)
}
