// Package util contains util functions for godaddy-cli
package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/oze4/godaddygo"
)

// PrintDomains formats the output for listing domains
func PrintDomains(domains []godaddygo.DomainSummary) {
	longestDomain := 0
	longestStatus := 0

	for _, domain := range domains {
		if len(domain.Domain) > longestDomain {
			longestDomain = len(domain.Domain)
		}
		if len(domain.Status) > longestStatus {
			longestStatus = len(domain.Status)
		}
	}

	for _, domain := range domains {
		fmt.Printf("%-"+strconv.Itoa(longestDomain)+"s\t%-"+strconv.Itoa(longestStatus)+"s\n", domain.Domain, domain.Status)
	}
}

// PrintRecords prints the listed records of a domain
func PrintRecords(records []godaddygo.Record) {
	// Get longest godaddygo.Record.Name
	longestName := 0
	for _, record := range records {
		if len(record.Name) > longestName {
			longestName = len(record.Name)
		}
	}
	for _, record := range records {
		fmt.Printf("%-5s\t%4d\t%"+strconv.Itoa(longestName)+"s\t%s\n", record.Type, record.TTL, record.Name, record.Data)
	}
}

type goDaddyError struct {
	Message string `json:"message"`
}

// PrintError prints the error received by oze4/godaddygo
func PrintError(err error) {
	errStr := err.Error()
	// Find where the json starts
	i := strings.Index(errStr, "{")
	body := errStr[i:]
	data := new(goDaddyError)
	json.Unmarshal([]byte(body), &data)

	fmt.Println(data.Message)
}
