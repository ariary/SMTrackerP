package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ariary/go-utils/pkg/check"
	"github.com/ariary/go-utils/pkg/color"
	"github.com/ariary/smtrackerp/pkg/smtrackerp"
	"github.com/spf13/cobra"
)

func main() {

	//CMD FILELESS-XEC
	var url string
	var target string
	var verbose bool

	var rootCmd = &cobra.Command{
		Use:   "smtrackerp",
		Short: "Track if mail has been read",
		Long:  `Track if mail has been read by inserting malicious and transparent image within`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &smtrackerp.Config{Url: url, Target: target, Verbose: verbose}

			if cfg.Target != "" {
				smtrackerp.GeneratePayload(cfg)
				fmt.Println(color.Bold("✉️ Generate HTML payload:"), color.Italic("payload.html"))
			}

			smtrackerp.Serve(cfg)
		},
	}
	// flag handling
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "external address of the HTTP server waiting for proof of reading")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")
	rootCmd.Flags().StringVarP(&target, "generate", "t", "", "generate HTML template to include within th email for a specific target")

	// SEND command
	var recipients []string
	var subject string
	var bodyFile string
	var smtpHost, smtpPort string
	var serve bool

	var sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send mail with tracker",
		Long:  `Send mail with the tracker`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// retrieve smtp credentials
			username := os.Getenv("SMTP_USERNAME")
			password := os.Getenv("SMTP_PASSWORD")
			if username == "" || password == "" {
				fmt.Println("❌ Set your environment variable SMTP_USERNAME and SMTP_PASSWORD to use send command.")
				os.Exit(1)
			}

			//read body
			body, err := ioutil.ReadFile(bodyFile)
			check.CheckAndExit(err, "❌ Failed reading body template for mail")

			// send mails
			cfg := &smtrackerp.Config{Url: url, Target: target, Verbose: verbose, Username: username, Password: password, Subject: subject, Body: body, SmtpHost: smtpHost, SmtpPort: smtpPort, Recipents: recipients}
			smtrackerp.SendMail(cfg)

			//Also track?
			if serve {
				smtrackerp.Serve(cfg)
			}
		},
	}
	//flag handling
	sendCmd.PersistentFlags().StringSliceVarP(&recipients, "recipients", "r", []string{}, "list of recipients to send mail")
	sendCmd.PersistentFlags().StringVarP(&subject, "subject", "s", "", "mail subject")
	sendCmd.PersistentFlags().StringVarP(&bodyFile, "body", "b", "", "file containing the mail body")
	sendCmd.PersistentFlags().StringVarP(&smtpHost, "smtphost", "a", "", "smtp server address")
	sendCmd.PersistentFlags().StringVarP(&smtpPort, "smtpport", "p", "", "smtp server port")
	sendCmd.PersistentFlags().BoolVarP(&serve, "track", "t", false, "wait for proof of reading after sending mails")
	sendCmd.MarkPersistentFlagRequired("url")
	sendCmd.MarkPersistentFlagRequired("recipients")
	sendCmd.MarkPersistentFlagRequired("body")
	sendCmd.MarkPersistentFlagRequired("smtphost")
	sendCmd.MarkPersistentFlagRequired("smtpport")

	rootCmd.AddCommand(sendCmd)
	rootCmd.Execute()

}
