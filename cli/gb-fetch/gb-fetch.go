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

// The gb-fetch program is a command line tool to retrieve the user's geekbadge (uberbadge).
// The data is written to standard out, or, if a filename was specified, saved to a file.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/profburke/bgg/cli/utilities"
	"github.com/profburke/bgg/geekbadge"
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
		fmt.Println("fetching geekbadge...")
	}

	gb, err := geekbadge.Get()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var jsonData []byte
	jsonData, err = json.Marshal(gb)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		if outputFilename != "" {
			utilities.WriteToFile(outputFilename, "gb-fetch", force, jsonData)
		} else {
			fmt.Println(string(jsonData))
		}
	}
}

// Local Variables:
// compile-command: "go build"
// End:
