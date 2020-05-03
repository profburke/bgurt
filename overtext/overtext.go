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

// Package overtext enables setting and retrieving avatar and badge overtext.
// The way bgg's form is set up, you must send the overtext for both the badge and avatar
// (assuming the user has both). Otherwise, the one that is not sent gets erased.
//
package overtext

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/profburke/bgg/bggclient"
)

type Overtext struct {
	Avatar *string
	Badge  *string
}

func (o Overtext) String() string {
	return fmt.Sprintf("avatar: %s\nbadge: %s", o.Avatar, o.Badge)
}

var avatarOvertextRegEx *regexp.Regexp
var badgeOvertextRegEx *regexp.Regexp
var editOvertextURL *url.URL
var overtextFormURL *url.URL

func init() {
	avatarOvertextRegEx = regexp.MustCompile("name=\"overtext\\[avatar\\]\"\\s*?value=\"(.*)\"")
	badgeOvertextRegEx = regexp.MustCompile("name=\"overtext\\[badge\\]\"\\s*?value=\"(.*)\"")
	var err error

	editOvertextURL, err = url.Parse("geekaccount/edit/overtext")
	if err != nil {
		log.Fatal("overtext: could not create edit overtext URL")
	}

	overtextFormURL, err = url.Parse("geekaccount.php")
	if err != nil {
		log.Fatal("overtext: could not create overtext form URL")
	}
}

// TODO: what happens if you send one of these texts but user hasn't purchased that feature?

// Set user's overtext. NOTE: If the Overtext struct has empty or nil for either of its
// fields, then the corresponding overtext will be deleted by the server. This function
// is "dumb" in that regards. That is, it's the calling functions responsibility to handle
// not deleting/overwriting overtext. Set will just set exactly what it's called with.
//
func Set(overtext Overtext) (success bool, err error) {
	data := url.Values{}
	data.Set("action", "saveovertext")
	data.Set("overtext[avatar]", *overtext.Avatar)
	data.Set("overtext[badge]", *overtext.Badge)

	resp, err := bggclient.Post(overtextFormURL, data)
	if err != nil {
		return false, err
	}
	success = resp.StatusCode == 200
	return
}

// Get user's overtext.
//
func Get() (overtext Overtext, err error) {
	page, err := bggclient.Get(editOvertextURL)
	if err != nil {
		message := fmt.Sprintf("overtext.Get: could not get edit avatar page: %v", err)
		return Overtext{}, errors.New(message)
	}

	pieces := avatarOvertextRegEx.FindStringSubmatch(page)

	if len(pieces) != 2 {
		return Overtext{}, errors.New("overtext.Get: could not parse avatar overtext")
	}

	avatarOvertext := pieces[1]

	pieces = badgeOvertextRegEx.FindStringSubmatch(page)
	if len(pieces) != 2 {
		return Overtext{}, errors.New("overtext.Get: could not parse badge overtext")
	}

	badgeOvertext := pieces[1]

	return Overtext{
		Avatar: &avatarOvertext,
		Badge:  &badgeOvertext,
	}, nil
}

// Local Variables:
// compile-command: "go build"
// End:
