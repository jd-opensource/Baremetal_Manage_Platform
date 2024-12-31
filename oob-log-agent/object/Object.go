package object

// Object ,the base interface for different object
type Object interface {
	JSON() (string, error)
}
