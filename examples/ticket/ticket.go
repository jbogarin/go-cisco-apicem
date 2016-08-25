package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/jbogarin/go-apic-em/apic-em"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	apicEMClient := apicem.NewClient(client)
	user := &apicem.User{
		Username: "<USERNAME>",
		Password: "<PASSWORD>",
	}
	apicEMClient.BaseURL, _ = url.Parse("<URL>")

	fmt.Println("==========================================")

	newTicket, _, err := apicEMClient.Ticket.AddTicket(user)
	if err != nil {
		log.Fatal(err)
	}

	ticket := newTicket.Response.ServiceTicket
	fmt.Println("Ticket", ticket)

	fmt.Println("==========================================")

	apicEMClient.Authorization = ticket

	idleTimeout, _, err := apicEMClient.Ticket.GetIdleTimeout()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(idleTimeout.Response.Attribute, idleTimeout.Response.Value)

	fmt.Println("==========================================")

	sessionTimeout, _, err := apicEMClient.Ticket.GetSessionTimeout()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sessionTimeout.Response.Attribute, sessionTimeout.Response.Value)

	fmt.Println("==========================================")

	root, _, err := apicEMClient.Ticket.DeleteTicket(ticket)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(root.Response)

}
