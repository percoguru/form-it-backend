package constants

type ResponseStatus int64

const (
	Empty ResponseStatus = iota
	Started
	Complete
)

func (s ResponseStatus) String() string {
	switch s {
	case Empty:
		return "Empty"
	case Started:
		return "Started"
	case Complete:
		return "Complete"

	}
	return "unkown"
}
