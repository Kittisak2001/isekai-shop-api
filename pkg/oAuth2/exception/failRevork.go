package exception

type FailRevoke struct{
}

func (e *FailRevoke) Error() string {
	return "failed to revoke token "
}