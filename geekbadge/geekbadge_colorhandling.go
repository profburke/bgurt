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

package geekbadge

import (
	"encoding/hex"
	"fmt"
	"image/color"
)

// colorFromString takes a 6 character Hex string and converts it into
// a color.RGBA. Alpha is assumed to be FF.
//
func colorFromString(s string) (col color.RGBA, err error) {
	b, err := hex.DecodeString(s + "FF")
	if err == nil {
		col = color.RGBA{b[0], b[1], b[2], b[3]}
	}

	return
}

// colorString converts a color.RGBA into a string representation.
// Not sure why I wrote this...
//
func colorString(c color.RGBA) (result string) {
	r, g, b, a := c.RGBA()
	result = fmt.Sprintf("color.RGBA{%d, %d, %d, %d}", r, g, b, a)

	return
}

// hexify converts a color.RGBA into a 6 character Hex string.
//
func hexify(c color.RGBA) (result string) {
	r, g, b, _ := c.RGBA()
	result = fmt.Sprintf("%02x%02x%02x", uint8(r), uint8(g), uint8(b))

	return
}

// Local Variables:
// compile-command: "go build"
// End:
