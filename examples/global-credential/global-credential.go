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

	newTicket, resp, err := apicEMClient.Ticket.AddTicket(user)
	if err != nil {
		fmt.Println(resp.Response.Status)
		log.Fatal(err)
	}

	ticket := newTicket.Response.ServiceTicket
	fmt.Println("Ticket", ticket)
	apicEMClient.Authorization = ticket

	fmt.Println("==========================================")

	credentialSubType := &apicem.CredentialQueryParams{
		CredentialSubType: "SNMPV2_READ_COMMUNITY",
	}
	globalCredentials, resp, err := apicEMClient.GlobalCredential.GetGlobalCredential("all", credentialSubType)
	if err != nil {
		log.Fatal(err)
	}

	for _, globalCredential := range globalCredentials.Response {
		fmt.Println(globalCredential.Description, globalCredential.InstanceUUID)
	}
}
