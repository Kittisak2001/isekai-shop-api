package exceptions

import "fmt"

type (
	InventoryListing struct {
		PlayerID string
	}
)

func (e *InventoryListing) Error() string {
	return fmt.Sprintf("inventory listing of PlayerID: %s failed", e.PlayerID)
}