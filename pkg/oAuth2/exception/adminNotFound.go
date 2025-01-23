package exception

type (
	AdminNotFound struct {
		AdminID string
	}
)

func (e *AdminNotFound) Error() string {
	return "not found adminID: " + e.AdminID
}