package notifications

type Notifier interface {
	Send(customer string) error
}
