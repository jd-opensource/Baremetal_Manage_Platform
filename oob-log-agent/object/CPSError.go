package object

//
const (
	MonitorSuccess = iota
	MonitorOOBError
	MonitorUnimplement
	MonitorUnknownCPS
	MonitorCPSiLOBlank
	MonitorCPSFiltered
	MonitorOtherError
)

// NewMonitorCode new a ParseCode
func NewMonitorCode(code int, text string) error {
	return &MonitorCode{code, text}
}

// MonitorCode Parser process error
type MonitorCode struct {
	c int
	s string
}

func (e *MonitorCode) Error() string {
	return e.s
}

// Code return result code
func (e *MonitorCode) Code() int {
	return e.c
}
