package main

import (
	"fmt"

	"github.com/anupam/email-verifier/pkg"
)

func main() {
	fmt.Println("Let's check for domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	pkg.VerifyEmailDomain()
}
