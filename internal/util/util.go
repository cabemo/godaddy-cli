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

func PrintAvailability(availability godaddygo.DomainAvailability) {
	dlen := len(availability.Domain)
	plen := 5
	price := fmt.Sprintf("%s %s", strconv.FormatFloat(float64(availability.Price)/1000000, 'f', 2, 64), availability.Currency)

	if len(price) > plen {
		plen = len(price)
	}

	fmt.Printf("%-"+strconv.Itoa(dlen)+"s\tAvailable\t%-"+strconv.Itoa(plen)+"s\tPeriod\n", "Domain", "Price")
	fmt.Printf("%-"+strconv.Itoa(dlen)+"s\t%-9t\t%s\t%d year(s)\n", availability.Domain, availability.Available, price, availability.Period)
}

// PrintRecords prints the listed records of a domain
func PrintRecords(records []godaddygo.Record) {
	// Get longest name (DNS) for padding
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

// PrintRecord prints a single record with a specified message
func PrintRecord(record godaddygo.Record, message string) {
	fmt.Printf("%s\t%-5s\t%s\n", message, record.Type, record.Name)
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

/// StringToRecord takes a string and returns a oze4/godaddygo.RecordType
func StringToRecord(r string) godaddygo.RecordType {
	var recordType godaddygo.RecordType

	switch strings.ToUpper(r) {
	case "A":
		recordType = godaddygo.RecordTypeA
		break
	case "AAAA":
		recordType = godaddygo.RecordTypeAAAA
		break
	case "CNAME":
		recordType = godaddygo.RecordTypeCNAME
		break
	case "MX":
		recordType = godaddygo.RecordTypeMX
		break
		//						case "NS":
		//							recordType = godaddygo.RecordTypeNS
		//							break
		//						case "SOA":
		//							recordType = godaddygo.RecordTypeSOA
		//							break
		//						case "SRV":
		//							recordType = godaddygo.RecordTypeSRV
		//							break
	case "TXT":
		recordType = godaddygo.RecordTypeTXT
		break
	default:
		recordType = godaddygo.RecordTypeA
	}

	return recordType
}
