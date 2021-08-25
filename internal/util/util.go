// util package contains util functions for godaddy-cli
package util

import (
    "fmt"
    "strconv"
    "github.com/oze4/godaddygo"
)

// PrintDomains formats the output for listing domains
func PrintDomains(domains []godaddygo.DomainSummary) {
    longestDomain:= 0
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
       fmt.Printf("%-" + strconv.Itoa(longestDomain)+ "s\t%-" + strconv.Itoa(longestStatus) + "s\n", domain.Domain, domain.Status)
    }
}
