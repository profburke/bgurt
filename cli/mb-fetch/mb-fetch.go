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

// The mb-fetch program is a command line tool to retrieve your microbadges. By defaut,
// the data is written to standard out in JSON format. You can use a command line flag
// to specify a file name instead.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/profburke/bgg/cli/utilities"
	"github.com/profburke/bgg/microbadge"
)

func main() {
	var verbose, force bool
	var outputFilename string

	flag.BoolVar(&force, "force", false, "overwrite output file if it exists")
	flag.BoolVar(&force, "f", false, "overwrite output file if it exists (shorthand)")
	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")
	flag.StringVar(&outputFilename, "output", "", "filename for output")
	flag.StringVar(&outputFilename, "o", "", "filename for output (shorthand)")

	flag.Parse()

	utilities.SetCredentials()

	if verbose {
		fmt.Println("fetching microbadges...")
	}

	badges, err := microbadge.GetAll()
	if err != nil {
		log.Fatalf("mb-fetch: could not get microbadges: ", err)
	}

	if badges != nil {
		var jsonData []byte
		jsonData, err := json.Marshal(badges)
		if err != nil {
			log.Fatalf("mb-fetch: %v", err)
		}
		if outputFilename != "" {
			utilities.WriteToFile(outputFilename, "mb-fetch", force, jsonData)
		} else {
			fmt.Println(string(jsonData))
		}
	} else {
		if verbose {
			// NOTE: no error, just no badges
			fmt.Println("no badges")
		}
	}
}

// Local Variables:
// compile-command: "go build"
// End:
