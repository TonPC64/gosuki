package subject

type Subject interface {
	GetSubject() string
}

type subject struct {
	name string
}

// GetSubject func
func (s subject) GetSubject() string {
	return s.name
}

func NewSubject() Subject {
	return Subject(subject{
		name: "Golang",
	})
}
