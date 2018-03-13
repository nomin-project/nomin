package main

import (
	"encoding/json"
	"fmt"

	"github.com/nomin-project/nomin/pkg/sender"
	"github.com/pkg/browser"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
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

		fmt.Println("opening window")
		sendErr := bootstrap.SendMessage(w, "sending.success", "message has been sent")
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
