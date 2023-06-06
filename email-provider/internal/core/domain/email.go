package domain

type Email struct {
	From     string
	Subject  string
	To       string
	Receiver string
	Content  string
}

type EmailSettings struct {
	APIKey        string
	APIPrivateKey string
	SenderName    string
	SenderEmail   string
}
