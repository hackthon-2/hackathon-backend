package util

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

type Email interface {
	SetMsg(url string, port string, username string, passwd string)
	SendTLS(rec string, subject string, body string) error
}
type Mail struct {
	Host     string
	Port     string
	Username string
	Password string
	Receiver string
}

func (m *Mail) SetMsg(host string, port string, username string, passwd string) {
	m.Host = host
	m.Port = port
	m.Username = username
	m.Password = passwd
}

func (m *Mail) SendTLS(rec string, subject string, body string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)
	msg := []byte("From:" + m.Username + "\r\n" +
		"To:" + rec + "\r\n" +
		"Subject:" + subject + "\r\n" +
		"Content-Type: text/plain;charset=UTF-8" +
		"\r\n\r\n" +
		body + "\r\n")
	tlsConfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: m.Host,
	}
	//关键代码
	conn, err := tls.Dial("tcp", m.Host+":"+m.Port, tlsConfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, m.Host)
	if err != nil {
		log.Panic(err)
	}
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	if err = c.Mail(m.Username); err != nil {
		log.Panic(err)
	}
	if err = c.Rcpt(rec); err != nil {
		log.Panic(err)
	}
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(msg)
	if err != nil {
		log.Panic(err)
	}
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	err=c.Quit()
	return err
}
