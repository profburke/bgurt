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

// The av-set program is a command line tool to set the user's avatar.
// The image in the file named on the command line is sent to the server.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/profburke/bgurt/avatar"
	"github.com/profburke/bgurt/cli/utilities"
)

func main() {
	var verbose bool
	var logfile string

	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")

	flag.StringVar(&logfile, "log", "", "filename for log")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("usage: av-set <filename>")
		os.Exit(1)
	}

	var lf *os.File
	var logger *log.Logger

	if len(logfile) > 0 {
		lf, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		logger = log.New(lf, "av-set: ", log.LstdFlags)
	}
	defer func() {
		if lf != nil {
			lf.Close()
		}
	}()

	utilities.SetCredentials()
	err := avatar.Set(args[0])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else if verbose {
		// TODO: logging shouldn't be dependent on verbose flag
		if len(logfile) > 0 {
			logger.Printf("avatar set to %s\n", args[0])
		}
		fmt.Printf("avatar set to %s\n", args[0])
	}
}

// Local Variables:
// compile-command: "go build"
// End:
