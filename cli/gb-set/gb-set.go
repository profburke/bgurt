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

// The gb-set program is a command line tool to set the user's geekbadge (uberbadge).
// The data is read from the specified file.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/profburke/bgg/cli/utilities"
	"github.com/profburke/bgg/geekbadge"
)

func main() {
	var verbose bool

	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: gb-set <filename>")
		os.Exit(1)
	}

	utilities.SetCredentials()

	jsonData, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	var gb geekbadge.Geekbadge
	err = json.Unmarshal(jsonData, &gb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't decode geekbadge: %v\n", err)
		os.Exit(1)
	}

	_, err = geekbadge.Set(gb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	} else if verbose {
		fmt.Println("new geekbadge set")
	}
}

// Local Variables:
// compile-command: "go build"
// End:
