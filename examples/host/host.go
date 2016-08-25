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

	getHostsQueryParams := &apicem.GetHostsQueryParams{
		Limit: "10",
	}
	hosts, resp, err := apicEMClient.Host.GetHosts("ALL", getHostsQueryParams)
	if err != nil {
		log.Fatal(err)
	}

	for _, host := range hosts.Response {
		fmt.Println(host.ID, host.HostName, host.HostType)
	}

	fmt.Println("==========================================")

	host, resp, err := apicEMClient.Host.GetHostByID("28af757c-f05f-46b7-b3c0-f4d87b6d705e", "ALL")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(host.Response.ID, host.Response.HostName, host.Response.HostType)

	fmt.Println("==========================================")

	hostCount, resp, err := apicEMClient.Host.GetHostCount("ALL")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hosts: ", hostCount.Response)

}
