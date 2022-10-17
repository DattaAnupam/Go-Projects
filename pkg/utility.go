package pkg

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type domainModel struct {
	Domain      string
	HasMx       bool
	HasSPF      bool
	SpfRecord   string
	HasDMARC    bool
	DmarcRecord string
}

func VerifyEmailDomain() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		checkDomain(scanner)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

// Verify input domain
func checkDomain(scanner *bufio.Scanner) {
	var dmod domainModel

	// Take input
	dmod.Domain = scanner.Text()

	// Check for MxRecords
	mxRecords, err := net.LookupMX(dmod.Domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		dmod.HasMx = true
	}

	// Check for SpfRecords
	txtRecords, err := net.LookupTXT(dmod.Domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Get and Set Spf value
	getSpf(txtRecords, &dmod)

	// Check for DmarcRecord
	dmarcRecords, err := net.LookupTXT("_dmarc." + dmod.Domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Get and set Dmarc value
	getDmarc(dmarcRecords, &dmod)

	// Print Received data
	dmod.showResult()
}

// Set spf value
func getSpf(txtRecords []string, dmod *domainModel) {
	for _, txtRecord := range txtRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			dmod.HasSPF = true
			dmod.SpfRecord = txtRecord
			break
		}
	}
}

// Set Dmarc value
func getDmarc(dmarcRecords []string, dmod *domainModel) {
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			dmod.HasDMARC = true
			dmod.DmarcRecord = record
			break
		}
	}
}

// Show Received information
func (d domainModel) showResult() {
	fmt.Printf("Domain: %s\n", d.Domain)
	fmt.Printf("Has Mx: %v\n", d.HasMx)
	fmt.Printf("Has Spf: %v\n", d.HasSPF)
	fmt.Printf("Spf Record: %s\n", d.SpfRecord)
	fmt.Printf("Has DMARC: %v\n", d.HasDMARC)
	fmt.Printf("Dmarc Record: %s\n", d.DmarcRecord)
}
