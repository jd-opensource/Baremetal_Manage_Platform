package object

// LogObject type of Event log from dell lifecyclelog or IPMI
type LogObject struct {
	SeqNumber string //664
	Messages  []string
}
