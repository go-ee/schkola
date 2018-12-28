package mailer

import (
	"gopkg.in/gomail.v2"
	"fmt"
	"log"
)

type User struct {
	Name    string
	Address string
}

type Mailer struct {
}

func (o *Mailer) SendRegistration() {
	// The list of recipients.
	var list []User
	list = make([]User, 0)

	list = append(list, User{Name: "Eugen Eisler", Address: "eoeisler@gmail.com"})
	list = append(list, User{Name: "Oxana Eisler", Address: "oxana.eisler@gmail.com"})
	list = append(list, User{Name: "Maria Eisler", Address: "maria.eisler@gmail.com"})

	//d := gomail.NewDialer("localhost", 1025, "user", "123456")
	d := gomail.NewDialer("localhost", 1025, "user", "123456")
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	for _, r := range list {
		m.SetHeader("From", "no-reply@example.com")
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", "Newsletter #1")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!", r.Name))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
}
