// Copyright (c) 2020 BlueDino Software (http://bluedino.net)
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation and/or
//    other materials provided with the distribution.
// 3. Neither the name of the copyright holder nor the names of its contributors may be
//    used to endorse or promote products derived from this software without specific prior
//    written permission.
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT
// OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR
// TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// The ot-fetch program is a command line tool to retrieve your overtext. You can specify
// which overtext to retrieve (avatar, badge, or both) and whether to print the results
// as plain text or JSON via flags on the command line.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/profburke/bgurt/cli/utilities"
	"github.com/profburke/bgurt/overtext"
)

func main() {
	var verbose, jsonOutput bool
	var display string

	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")
	flag.BoolVar(&jsonOutput, "json", false, "format output as a JSON object; overrides the display flag")
	flag.StringVar(&display, "display", "both", "`display` <avatar, badge, both>")

	flag.Parse()

	displayChoices := map[string]bool{"avatar": true, "badge": true, "both": true}
	if _, validChoice := displayChoices[display]; !validChoice {
		fmt.Fprintf(os.Stderr, "'%s' is not a valid choice for the display flag.\n", display)
		fmt.Fprintln(os.Stderr, "Usage of ot-fetch:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	utilities.SetCredentials()

	if verbose {
		fmt.Println("fetching overtext...")
	}

	overtext, err := overtext.Get()
	if err != nil {
		log.Fatalf("ot-fetch: %v.", err)
	}

	if jsonOutput {
		json, err := json.Marshal(overtext)
		if err != nil {
			log.Fatalf("ot-fetch: %v", err)
		}
		fmt.Println(string(json))
	} else {
		var message string
		switch display {
		case "avatar":
			message = fmt.Sprintf("'%s'", *overtext.Avatar)
		case "badge":
			message = fmt.Sprintf("'%s'", *overtext.Badge)
		default:
			message = fmt.Sprintf("avatar overtext: '%s'\nbadge ovetext: '%s'",
				*overtext.Avatar, *overtext.Badge)
		}
		fmt.Println(message)
	}
}

// Local Variables:
// compile-command: "go build"
// End:
