package main

import (
	"github.com/matcornic/hermes"
	"io/ioutil"
)

func main() {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Bibelschule-Stephanus",
			Link: "https://bibelschule-stephanus.de/",
			// Optional product logo
			Logo: "https://bibelschule-stephanus.de/wp-content/uploads/2017/07/logo_200.png",
			Copyright:"Copyright © 2017 Bibelschule-Stephanus.",
			TroubleText:"Wenn du Schwierigkeiten mit der Schaltfläche '{ACTION}' hast, kopiere und füge die URL unten in deinen Webbrowser ein.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Greeting: "Hallo",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Optionally, preview the generated HTML e-mail by writing it to a local file
	err = ioutil.WriteFile("preview.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	err = ioutil.WriteFile("preview.txt", []byte(emailText), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
}
