package main

import (
	"flag"
	"fmt"
	"github.com/lnashier/usps-imb"
)

func main() {
	flag.Parse()
	track := flag.Arg(0)
	route := flag.Arg(1)

	barcode, statusCode := usps.IMb(track, route)

	if statusCode != usps.StatusEncoderApiSuccess {
		fmt.Printf("Error generating barcode: %d, %s\n", statusCode, usps.StatusText(statusCode))
		return
	}

	fmt.Println("Generated Barcode:", barcode)
}
