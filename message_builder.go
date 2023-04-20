package main

import (
	"strings"
)

type MessageBuilder struct {
	message      Message
	currentField string
}

func (mb *MessageBuilder) isParsingHeader() bool {
	return mb.currentField != "Content"
}

func (mb *MessageBuilder) ParseLine(line string) {
	line = strings.TrimSpace(line)

	if strings.HasPrefix(line, "Message-ID:") && mb.isParsingHeader() {
		mb.currentField = "Message-ID"
		mb.message.MessageID = strings.TrimPrefix(line, "Message-ID: <")
		mb.message.MessageID = strings.TrimSuffix(mb.message.MessageID, ">")
	} else if strings.HasPrefix(line, "Date:") && mb.isParsingHeader() {
		mb.currentField = "Date"
		mb.message.Date = strings.TrimPrefix(line, "Date:")
	} else if strings.HasPrefix(line, "From:") && mb.isParsingHeader(){
		mb.currentField = "From"
		mb.message.From = strings.TrimPrefix(line, "From:")
	} else if strings.HasPrefix(line, "To:") && mb.isParsingHeader() {
		mb.currentField = "To"
		mb.message.To = strings.TrimPrefix(line, "To:")
	} else if strings.HasPrefix(line, "Subject:") && mb.isParsingHeader(){
		mb.currentField = "Subject"
		mb.message.Subject = strings.TrimPrefix(line, "Subject:")
	} else if strings.HasPrefix(line, "CC:") && mb.isParsingHeader(){
		mb.currentField = "CC"
		mb.message.CC = strings.TrimPrefix(line, "CC:")
	}else if strings.HasPrefix(line, "Mime-Version:") && mb.isParsingHeader(){
		mb.currentField = "Mime-Version"
		mb.message.MimeVersion = strings.TrimPrefix(line, "Mime-Version:")
	} else if strings.HasPrefix(line, "Content-Type:") && mb.isParsingHeader(){
		mb.currentField = "Content-Type"
		mb.message.ContentType = strings.TrimPrefix(line, "Content-Type:")
	} else if strings.HasPrefix(line, "Content-Transfer-Encoding:") && mb.isParsingHeader(){
		mb.currentField = "Content-Transfer-Encoding"
		mb.message.ContentTransferEncoding = strings.TrimPrefix(line, "Content-Transfer-Encoding:")
	} else if strings.HasPrefix(line, "Bcc:") && mb.isParsingHeader(){
		mb.currentField = "Bcc"
		mb.message.BCC = strings.TrimPrefix(line, "Bcc:")
	} else if strings.HasPrefix(line, "X-From:") && mb.isParsingHeader(){
		mb.currentField = "X-From"
		mb.message.XFrom = strings.TrimPrefix(line, "X-From:")
	} else if strings.HasPrefix(line, "X-To:") && mb.isParsingHeader(){
		mb.currentField = "X-To"
		mb.message.Xto = strings.TrimPrefix(line, "X-To:")
	}else if strings.HasPrefix(line, "X-cc:") && mb.isParsingHeader(){
		mb.currentField = "X-cc"
		mb.message.XCC = strings.TrimPrefix(line, "X-cc:")
	}else if strings.HasPrefix(line, "X-bcc:") && mb.isParsingHeader(){
		mb.currentField = "X-bcc"
		mb.message.XBCC = strings.TrimPrefix(line, "X-bcc:")
	}else if strings.HasPrefix(line, "X-Folder:") && mb.isParsingHeader(){
		mb.currentField = "X-Folder"
		mb.message.XFolder = strings.TrimPrefix(line, "X-Folder:")
	}else if strings.HasPrefix(line, "X-Origin:") && mb.isParsingHeader(){
		mb.currentField = "X-Origin"
		mb.message.XOrigin = strings.TrimPrefix(line, "X-Origin:")
	}else if strings.HasPrefix(line, "X-FileName:") && mb.isParsingHeader(){
		mb.currentField = "X-FileName"
		mb.message.XFilename = strings.TrimPrefix(line, "X-FileName:")
	} else {
		switch mb.currentField {
		case "Subject":
			mb.message.Subject += " " + line
		case "Date":
			mb.message.Date += " " + line
		case "To":
			mb.message.To += " " + line
		case "CC":
			mb.message.CC += " " + line
		case "Bcc":
			mb.message.BCC += " " + line
		case "Content":
			mb.message.Content += " " + line
		default:
			mb.currentField = "Content"
		}
	}
}


func (mb *MessageBuilder) Build() Message {
	return mb.message
}

func (mb *MessageBuilder) Reset() {
	resetMessage(&mb.message)
	mb.currentField = ""
}
