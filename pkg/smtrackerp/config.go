package smtrackerp

type Config struct {
	Url       string
	Target    string
	Verbose   bool
	Recipents []string
	Subject   string
	Body      []byte
	SmtpHost  string
	SmtpPort  string
	Username  string
	Password  string
}
