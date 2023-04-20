package main

//"errors"

type Message struct {
	MessageID               string
	Date                    string
	From                    string
	To                      string
	Subject                 string
	CC                      string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	BCC                     string
	XFrom                   string
	Xto                     string
	XCC                     string
	XBCC                    string
	XFolder                 string
	XOrigin                 string
	XFilename               string
	Content                 string
}

func resetMessage(message *Message) {
	message.MessageID = ""
	message.Date = ""
	message.From = ""
	message.To = ""
	message.Subject = ""
	message.CC = ""
	message.MimeVersion = ""
	message.ContentType = ""
	message.ContentTransferEncoding = ""
	message.BCC = ""
	message.XFrom = ""
	message.Xto = ""
	message.XCC = ""
	message.XBCC = ""
	message.XFolder = ""
	message.XOrigin = ""
	message.XFilename = ""
	message.Content = ""
}
