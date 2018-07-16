package subject

import "testing"

type mockSubject struct{}

func (ms mockSubject) GetSubject() string {
	return "mock Subject"
}

func TestGetSubject(t *testing.T) {

	sj := Subject(mockSubject{})

	expected := "mock Subject"
	if sj.GetSubject() != expected {
		t.Errorf("GetSubject() should be equal %s\n", expected)
	}
}
