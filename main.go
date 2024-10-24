// pemsplit allows to split a PEM file containing multiple certificates into separate files.
// For example, if bundle.pem contains 3 certificates, running 'pemsplit bundle.pem'
// will create bundle01.pem, bundle02.pem, and bundle03.pem in the current directory,
// with one certificate per file.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: splitcert <certbundle.pem>")
	}
	inputFile := os.Args[1]
	base := filepath.Base(inputFile)
	filenameBase := strings.TrimSuffix(base, filepath.Ext(base))
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content by the certificate delimiter
	certificates := strings.Split(string(content), "-----END CERTIFICATE-----\n")

	// Remove any empty strings from the list
	var cleanedCertificates []string
	for _, cert := range certificates {
		if strings.TrimSpace(cert) != "" {
			cleanedCertificates = append(cleanedCertificates, cert+"-----END CERTIFICATE-----\n")
		}
	}

	// Write each certificate to a separate file
	for i, cert := range cleanedCertificates {
		outputFile := fmt.Sprintf(filenameBase+"%02d.pem", i+1)
		err := os.WriteFile(outputFile, []byte(cert), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Printf("Written %s\n", outputFile)
	}
}
