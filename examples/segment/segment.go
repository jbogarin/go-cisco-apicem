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

	getSegmentInfoQueryParams := &apicem.GetSegmentInfoQueryParams{
		Type: "wireless",
	}
	segments, resp, err := apicEMClient.Segment.GetSegmentInfo(getSegmentInfoQueryParams)
	if err != nil {
		log.Fatal(err)
	}

	for _, segment := range segments.Response {
		for _, networkDevice := range segment.NetworkDevices {
			fmt.Println(segment.Name, networkDevice.HostName)
		}
	}
}
