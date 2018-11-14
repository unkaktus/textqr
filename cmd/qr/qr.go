// qr.go - make an QR from arguments
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to byteqr, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/nogoegst/byteqr"
)

func main() {
	var large = flag.Bool("large", false, "produce x4 larger qr codes")
	var inverted = flag.Bool("inverted", false, "inverse colors")
	flag.Parse()
	data := strings.Join(flag.Args(), " ")

	_, err := byteqr.Write(os.Stdout, data, byteqr.L, !*large, *inverted)
	if err != nil {
		log.Fatal(err)
	}
}
