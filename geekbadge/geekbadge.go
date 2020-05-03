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

// Package geekbadge contains methods to set and retrieve geek- and uber-badges.
//
package geekbadge

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/profburke/bgg/bggclient"
)

type Box struct {
	Text       string
	Background color.RGBA
	TextColor  color.RGBA
	TextStart  uint
}

// TODO: need to better differentiate between geek- and uber- badges.
//       need to implement handling of image settings for uber-geekbadge.

type Geekbadge struct {
	OuterBorder, InnerBorder color.RGBA
	BarPosition              uint
	LeftBox, RightBox        Box
}

// TODO: include more details...

func (gb Geekbadge) String() string {
	return fmt.Sprintf("[%s | %s]", gb.LeftBox.Text, gb.RightBox.Text)
}

var geekbadgeRegEx *regexp.Regexp
var getGeekbadgeURL *url.URL
var setGeekbadgeURL *url.URL

func init() {
	geekbadgeRegEx = regexp.MustCompile("<img src=\"/button.php\\?(.+)\">")
	var err error

	getGeekbadgeURL, err = url.Parse("geekaccount/edit/geekbadge")
	if err != nil {
		log.Fatal("geekbadge: could not create get geekbadge URL")
	}

	setGeekbadgeURL, err = url.Parse("geekaccount.php")
	if err != nil {
		log.Fatal("geekbadge: could not create set geekbadge URL")
	}
}

// Get retrieves the currently set badge.
//
func Get() (gb Geekbadge, err error) {
	page, err := bggclient.Get(getGeekbadgeURL)
	if err != nil {
		message := fmt.Sprintf("geekbadge.Get: could not retrieve the geekbadge edit form: %v", err)
		return Geekbadge{}, errors.New(message)
	}

	pieces := geekbadgeRegEx.FindStringSubmatch(page)
	if len(pieces) != 2 {
		return Geekbadge{}, errors.New("geekbadge.Get: could not parse geekbadge description")
	}

	gb = Geekbadge{}
	lb := Box{}
	rb := Box{}

	keyvalues := strings.Split(pieces[1], "&amp;")
	for _, keyvalue := range keyvalues {
		pair := strings.Split(keyvalue, "=")
		// TODO: error handling for all the following type conversions
		//       (colorFromString, ParseUint)
		switch pair[0] {
		case "outerBorder":
			color, _ := colorFromString(pair[1])
			gb.OuterBorder = color
		case "innerBorder":
			color, _ := colorFromString(pair[1])
			gb.InnerBorder = color
		case "barPosition":
			val, _ := strconv.ParseUint(pair[1], 10, 64)
			gb.BarPosition = uint(val)
		case "leftFill":
			color, _ := colorFromString(pair[1])
			lb.Background = color
		case "rightFill":
			color, _ := colorFromString(pair[1])
			rb.Background = color
		case "leftTextColor":
			color, _ := colorFromString(pair[1])
			lb.TextColor = color
		case "rightTextColor":
			color, _ := colorFromString(pair[1])
			rb.TextColor = color
		case "leftText":
			lb.Text = pair[1]
		case "rightText":
			rb.Text = pair[1]
		case "leftTextPosition":
			val, _ := strconv.ParseUint(pair[1], 10, 64)
			lb.TextStart = uint(val)
		case "rightTextPosition":
			val, _ := strconv.ParseUint(pair[1], 10, 64)
			rb.TextStart = uint(val)
		}

		gb.LeftBox = lb
		gb.RightBox = rb
	}

	return
}

// Set takes a structure describing the desired badge and posts it to
// boardgamegeek.
//
func Set(gb Geekbadge) (success bool, err error) {
	data := url.Values{}
	data.Set("action", "savebadge")
	data.Set("outerBorder", hexify(gb.OuterBorder))
	data.Set("innerBorder", hexify(gb.InnerBorder))
	data.Set("barPosition", fmt.Sprintf("%d", gb.BarPosition))
	data.Set("leftText", gb.LeftBox.Text)
	data.Set("leftFill", hexify(gb.LeftBox.Background))
	data.Set("leftTextColor", hexify(gb.LeftBox.TextColor))
	data.Set("leftTextPosition", fmt.Sprintf("%d", gb.LeftBox.TextStart))
	data.Set("rightText", gb.RightBox.Text)
	data.Set("rightFill", hexify(gb.RightBox.Background))
	data.Set("rightTextColor", hexify(gb.RightBox.TextColor))
	data.Set("rightTextPosition", fmt.Sprintf("%d", gb.RightBox.TextStart))

	resp, err := bggclient.Post(setGeekbadgeURL, data)
	if err != nil {
		return false, err
	}

	success = resp.StatusCode == 200
	return
}

// Local Variables:
// compile-command: "go build"
// End:
