package internal

type CommunicationInterface interface {
	Read() (string, error)
	Write(string) error
}
