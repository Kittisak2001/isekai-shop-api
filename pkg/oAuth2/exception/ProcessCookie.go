package exception

type ProcessCookie struct{}

func (e *ProcessCookie) Error() string{
	return "failed to processing cookie"
}