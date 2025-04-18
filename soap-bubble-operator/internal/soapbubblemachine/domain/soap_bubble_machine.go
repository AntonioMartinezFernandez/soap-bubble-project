package soapbubblemachinedomain

type SoapBubbleMachine struct {
	id            string
	name          string
	startURL      string
	stopURL       string
	makingBubbles bool
}

func NewSoapBubbleMachine(id, name, startURL, stopURL string, makingBubbles bool) *SoapBubbleMachine {
	return &SoapBubbleMachine{
		id:            id,
		name:          name,
		startURL:      startURL,
		stopURL:       stopURL,
		makingBubbles: makingBubbles,
	}
}

func (s *SoapBubbleMachine) ID() string {
	return s.id
}

func (s *SoapBubbleMachine) Name() string {
	return s.name
}

func (s *SoapBubbleMachine) StartURL() string {
	return s.startURL
}

func (s *SoapBubbleMachine) StopURL() string {
	return s.stopURL
}

func (s *SoapBubbleMachine) MakingBubbles() bool {
	return s.makingBubbles
}

func (s *SoapBubbleMachine) SwitchON() {
	s.makingBubbles = true
}

func (s *SoapBubbleMachine) SwitchOFF() {
	s.makingBubbles = false
}
