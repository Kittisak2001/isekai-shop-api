package exception

type PlayerNotFound struct{}

func (e *PlayerNotFound)Error()string{
	return "player not found"
}