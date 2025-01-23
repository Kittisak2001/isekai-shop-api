package exception

type (
	AdminCreating struct {
		AdminID string
	}
)

func (e *AdminCreating) Error() string {
	return "creating adminID: " + e.AdminID + " failed"
}