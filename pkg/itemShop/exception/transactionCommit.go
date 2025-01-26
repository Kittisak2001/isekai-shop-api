package exception

type (
	TransactionCommit struct{}
)

func (e *TransactionCommit) Error() string {
	return "transaction commit failed"
}
