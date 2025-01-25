package exceptions

import "fmt"

type (
	InventoryFilling struct {
		PlayerID string
		ItemID   uint64
	}
)

func (e *InventoryFilling) Error() string {
	return fmt.Sprintf("inventory filling for PlayerID: %s and ItemID: %d failed", e.PlayerID, e.ItemID)
}