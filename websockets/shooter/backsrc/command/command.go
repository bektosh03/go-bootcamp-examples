package command

const (
	WaitForOpponent = "wait_for_opponent"
	Play            = "play"
	Start           = "start"
	Shoot           = "shoot"
)

type Command interface {
	Name() string
	Payload() []byte
}
