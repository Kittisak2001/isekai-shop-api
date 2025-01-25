package exceptions

import "fmt"

type (
	PlayerItemRemoving struct {
		ItemID uint64
	}
)

func (e *PlayerItemRemoving) Error() string {
	return fmt.Sprintf("removing item ItemID: %d failed", e.ItemID)
}