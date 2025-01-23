package exception

type FailCallback struct{
	Role string
}

func (e *FailCallback) Error() string {
	return "failed to callback " + e.Role
}