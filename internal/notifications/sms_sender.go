package notifications

import "fmt"

type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(customer string) error {
	fmt.Printf("SMS уведомление отправлено клиенту %s\n", customer)
	return nil
}
