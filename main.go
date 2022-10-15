package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type domainModel struct {
	domain      string
	hasMx       bool
	hasSPF      bool
	spfRecord   string
	hasDMARC    bool
	dmarcRecord string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Let's check for domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	var domModel *domainModel

	for scanner.Scan() {
		// checkDomain(scanner.Text())
		checkDomain(scanner, domModel)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(scanner *bufio.Scanner, domain *domainModel) {
	var dmod domainModel

	// Take input
	dmod.domain = scanner.Text()

	// Check for MxRecords
	mxRecords, err := net.LookupMX(dmod.domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		dmod.hasMx = true
	}

	// Check for SpfRecords
	txtRecords, err := net.LookupTXT(dmod.domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, txtRecord := range txtRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			dmod.hasSPF = true
			dmod.spfRecord = txtRecord
			break
		}
	}

	// Check for DmarcRecord
	dmarcRecords, err := net.LookupTXT("_dmarc." + dmod.domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			dmod.hasDMARC = true
			dmod.dmarcRecord = record
			break
		}
	}

	// Print Received data
	dmod.showResult()
}

func (d domainModel) showResult() {
	fmt.Printf("Domain: %s\n", d.domain)
	fmt.Printf("Has Mx: %v\n", d.hasMx)
	fmt.Printf("Has Spf: %v\n", d.hasSPF)
	fmt.Printf("Spf Record: %s\n", d.spfRecord)
	fmt.Printf("Has DMARC: %v\n", d.hasDMARC)
	fmt.Printf("Dmarc Record: %s\n", d.dmarcRecord)
}
