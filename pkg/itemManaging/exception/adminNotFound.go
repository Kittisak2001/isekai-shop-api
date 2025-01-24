package exception

type AdminNotFound struct{}

func (e *AdminNotFound) Error() string {
	return "admin not found"
}