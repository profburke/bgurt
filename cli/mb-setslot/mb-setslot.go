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

// The mb-setslot program is a command line tool to set the microbadge for a specific
// slot. Specify the slot number and microbadge ID as flags on the command line.
//
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/profburke/bgg/cli/utilities"
	"github.com/profburke/bgg/microbadge"
)

func main() {
	var verbose bool
	var slot, badgeNumber uint

	flag.BoolVar(&verbose, "verbose", false, "makes execution verbose")
	flag.BoolVar(&verbose, "v", false, "makes execution verbose (shorthand)")
	flag.UintVar(&slot, "slot", 0, "slot number to set (required)")
	flag.UintVar(&badgeNumber, "microbadge", 0, "microbadge ID (required)")

	flag.Parse()

	utilities.SetCredentials()

	if !microbadge.ValidSlot(slot) {
		fmt.Fprintf(os.Stderr, "mb-setslot: slot number must be between 1 and %d.\n",
			microbadge.TotalSlots)
		os.Exit(1)
	}

	if badgeNumber < 1 {
		fmt.Fprintf(os.Stderr, "mb-setslot: badge number must be a positive integer.\n")
		os.Exit(1)
	}

	_, err := microbadge.SetSlot(slot, badgeNumber)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mb-setslot: %v\n", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Printf("Set slot %d to badge ID %d.\n", slot, badgeNumber)
	}
}

// Local Variables:
// compile-command: "go build"
// End:
