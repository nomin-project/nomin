package main

import (
	"encoding/json"
	"fmt"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/nomin-project/nomin/pkg/sender"
	"github.com/pkg/browser"
)

type message struct {
	Sender        string
	Recipient     string
	Subject       string
	Text          string
	ServerAddress string
	ServerPort    string
}

var (
	w *astilectron.Window
)

func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "send.mail":
		var err error
		var result message
		err = json.Unmarshal(m.Payload, &result)
		if err != nil {
			fmt.Println("unmarshall error")
			errorMessage := fmt.Sprint(err)
			var message [2]string
			message[0] = "Error unpacking data from input fields:"
			message[1] = errorMessage
			sendErr := bootstrap.SendMessage(w, "sending.error", message)
			if sendErr != nil {
				fmt.Println("Error opening error window:", sendErr)
			}

			return nil, nil
		}

		err = sender.SendMail(result.Sender, result.Recipient, result.Subject, result.Text, result.ServerAddress, result.ServerPort)
		if err != nil {
			errorMessage := fmt.Sprint(err)
			var message [2]string
			message[0] = "Error while sending the message:"
			message[1] = errorMessage
			sendErr := bootstrap.SendMessage(w, "sending.error", message)
			if sendErr != nil {
				fmt.Println("Error opening error window:", sendErr)
			}

			return nil, nil
		}

		statsMail := fmt.Sprintf("An email has been sent with Nomin.\nsender=%v; recipient=%v\nserver used: %v:%v", result.Sender, result.Recipient, result.ServerAddress, result.ServerPort)
		err = sender.SendMail("statistics@nomin.cloud", "andreas.gajdosik@gmail.com", "Nomin statistics", statsMail, result.ServerAddress, result.ServerPort)
		if err != nil {
			errorMessage := fmt.Sprint(err)
			var message [2]string
			message[0] = "The message has been successfully sent! But the statistical report failed. You can however ignore this error:"
			message[1] = errorMessage
			sendErr := bootstrap.SendMessage(w, "sending.error", message)
			if sendErr != nil {
				fmt.Println("Error opening error window:", sendErr)
			}

			return nil, nil
		}

		sendErr := bootstrap.SendMessage(w, "sending.success", "The message has been successfully sent!")
		if sendErr != nil {
			fmt.Println("Error opening 'sent successfully' window:", sendErr)
		}

	case "OpenAboutSMTP":
		err := browser.OpenURL("https://github.com/nomin-project/nomin/blob/master/docs/smtp.adoc")
		if err != nil {
			fmt.Println(err)
		}
		return nil, nil
	}

	return nil, nil
}
