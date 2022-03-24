package smtrackerp

import (
	"bytes"
	"fmt"
	"net/smtp"
)

func SendMail(cfg *Config) {
	fmt.Println("📨 Sending mail...")
	for i := 0; i < len(cfg.Recipents); i++ {
		recipient := cfg.Recipents[i]
		// include tracker
		tracker, err := GetTracker(cfg.Url, recipient)
		if err != nil {
			fmt.Println("❌ Failed to build tracker (don't send mail to", recipient, "):", err)
		} else {
			cfg.Body = append(cfg.Body, []byte(tracker)...)
			// send mail
			sendMail(cfg.Username, recipient, cfg.Password, cfg.Subject, cfg.Body, cfg.SmtpHost, cfg.SmtpPort)
		}
	}
}

func sendMail(from string, to string, password string, subject string, body []byte, smtpHost string, smtpPort string) {
	// authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var mailContent bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	mailContent.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))

	mailContent.Write(body)
	// sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, mailContent.Bytes())
	if err != nil {
		fmt.Println("❌ Failed sending mail to", to)
		return
	} else {
		fmt.Println("\t• email send to:", to)
	}

}
