package exception

type (
	PlayerCreating struct {
		PlayerID string
	}
)

func (e *PlayerCreating) Error() string {
	return "creating playerID: " + e.PlayerID + " failed"
}