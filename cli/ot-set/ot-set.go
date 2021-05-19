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

// The ot-set program is a command line tool to set your overtext. You can specify
// which overtext to set (avatar, badge, or both).
//
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/profburke/bgurt/cli/utilities"
	"github.com/profburke/bgurt/overtext"
)

func main() {
	var verbose bool
	var avatarOvertext, badgeOvertext string

	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")
	flag.StringVar(&avatarOvertext, "avatar", "", "specify avatar overtext")
	flag.StringVar(&badgeOvertext, "badge", "", "specify badge overtext")

	flag.Parse()

	utilities.SetCredentials()

	// if either avatarOvertext or badgeOvertext are empty,
	// fetch them from server ,,, how do we then explicitly reset one of the text

	_, err := overtext.Set(overtext.Overtext{
		Avatar: &avatarOvertext,
		Badge:  &badgeOvertext,
	})
	if err != nil {
		log.Fatalf("ot-set: %v", err)
	}

	if verbose {
		fmt.Println("overtext updated.")
	}
}

// Local Variables:
// compile-command: "go build"
// End:
