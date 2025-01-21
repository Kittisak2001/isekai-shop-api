package exception

import "fmt"

type ItemNotfound struct {
	ItemID uint64
}

func (e *ItemNotfound) Error() string {
	return fmt.Sprintf("itemID: %d not found", e.ItemID)
}