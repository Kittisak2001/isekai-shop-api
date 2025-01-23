package exception

type (
	PlayerNotFound struct {
		PlayerID string
	}
)

func (e *PlayerNotFound) Error() string {
	return "not found playerID: " + e.PlayerID
}