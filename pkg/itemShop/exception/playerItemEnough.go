package exception

type (
	PlayerItemNotEnough struct{}
)

func (e *PlayerItemNotEnough) Error() string {
	return "player have item not enough"
}
