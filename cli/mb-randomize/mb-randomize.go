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

// The mb-randomize program is a command line tool to set your microbadges randomly according
// to constraints specified in XYZ. To build the XYZ file you should run the TUV tool.
//
// TODO: more details
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/profburke/bgg/cli/utilities"
	"github.com/profburke/bgg/microbadge"
)

func badgeIDs(badges []microbadge.Microbadge) (result []uint) {
	for _, badge := range badges {
		result = append(result, badge.BadgeNumber)
	}

	return
}

// TODO: dedup this -- also implemented in aws/utilities

func pick(original []microbadge.Microbadge, n int) (picked []microbadge.Microbadge) {
	picked = make([]microbadge.Microbadge, len(original))
	for i, value := range original {
		picked[i] = value
	}

	const passes = 6
	for j := 0; j < passes; j++ {
		for i := 0; i < len(picked); i++ {
			j := rand.Intn(len(picked))
			picked[i], picked[j] = picked[j], picked[i]
		}
	}

	return picked[0:n]
}

// TODO: try to read in badges from $CONFIG_DIR/badges.json
//       if it doesn't exist, download them and then proceed
//       overridden by filename on command line?

func main() {
	var verbose bool

	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")

	flag.Parse()

	// 2. get filename from command line and read it in as allBadges
	args := flag.Args()
	if len(args) != 1 {
		log.Println("usage: mb-randomize <filename>")
		os.Exit(1)
	}

	jsonData, err := ioutil.ReadFile(args[0])
	if err != nil {
	}

	var allBadges []microbadge.Microbadge
	_ = json.Unmarshal(jsonData, &allBadges)

	utilities.SetCredentials()

	rand.Seed(time.Now().UnixNano())
	newBadges := pick(allBadges, microbadge.TotalSlots)

	fmt.Printf("badges: %v", newBadges)

	badgeNumbers := badgeIDs(newBadges)

	fmt.Printf("badge numbers: %v", badgeNumbers)

	if verbose {
		fmt.Println("sending new badges to server")
	}

	_, err = microbadge.SetAll(badgeNumbers)

	if err != nil {
		fmt.Fprintf(os.Stderr, "mb-randomize: could not set badges: %v\n", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Printf("badges set")
	}
}

// Local Variables:
// compile-command: "go build"
// End:
