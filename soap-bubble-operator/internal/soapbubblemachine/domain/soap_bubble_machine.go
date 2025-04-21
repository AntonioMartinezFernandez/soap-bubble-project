package soapbubblemachinedomain

import "context"

type SoapBubbleMachine struct {
	id            string
	name          string
	ip            string
	speed         int
	makingBubbles bool
}

func NewSoapBubbleMachine(id, name, ip string, makingBubbles bool, speed int) *SoapBubbleMachine {
	return &SoapBubbleMachine{
		id:            id,
		name:          name,
		ip:            ip,
		makingBubbles: makingBubbles,
		speed:         speed,
	}
}

func (s *SoapBubbleMachine) ID() string {
	return s.id
}

func (s *SoapBubbleMachine) Name() string {
	return s.name
}

func (s *SoapBubbleMachine) IP() string {
	return s.ip
}

func (s *SoapBubbleMachine) MakingBubbles() bool {
	return s.makingBubbles
}

func (s *SoapBubbleMachine) Speed() int {
	return s.speed
}

func (s *SoapBubbleMachine) SetSpeed(speed int) {
	s.speed = speed
}

func (s *SoapBubbleMachine) SwitchON(ctx context.Context, remoteController SoapBubbleMachineRemoteController) error {
	s.makingBubbles = true
	return remoteController.SwitchOn(ctx, *s)
}

func (s *SoapBubbleMachine) SwitchOFF(ctx context.Context, remoteController SoapBubbleMachineRemoteController) error {
	s.makingBubbles = false
	return remoteController.SwitchOff(ctx, *s)
}
