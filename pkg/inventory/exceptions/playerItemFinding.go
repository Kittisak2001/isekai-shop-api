package exceptions

import "fmt"

type (
	PlayerItemsFinding struct {
		PlayerID string
	}
)

func (e *PlayerItemsFinding) Error() string {
	return fmt.Sprintf("finding player items for PlayerID: %s failed", e.PlayerID)
}