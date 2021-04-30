package mailer

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

var MailCh chan *gomail.Message

func New() chan *gomail.Message {
	MailCh = make(chan *gomail.Message)

	go func() {
		d := gomail.NewDialer(viper.GetString("host"), viper.GetInt("port"), viper.GetString("username"), viper.GetString("password"))

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-MailCh:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()

	// Use the channel in your program to send emails.

	return MailCh
}
