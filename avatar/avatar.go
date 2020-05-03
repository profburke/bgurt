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

// Package avatar provides functions for setting and retrieving the user's avatar.
// The avatar must be a PNG, GIF, or JPG file and is 64x64 pixels.
//
package avatar

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"

	"github.com/profburke/bgg/bggclient"
)

var avatarRegEx *regexp.Regexp
var getAvatarURL *url.URL
var setAvatarURL *url.URL

func init() {
	avatarRegEx = regexp.MustCompile("(https://cf.geekdo-static.com/avatars/avatar_id\\d+.(?:(?i)jpg|png|gif))")
	var err error

	getAvatarURL, err = url.Parse("myprofile")
	if err != nil {
		log.Fatal("avatar: could not create get avatar URL")
	}

	setAvatarURL, err = url.Parse("geekaccount/edit/avatar")
	if err != nil {
		log.Fatal("avatar: could not create edit avatar URL")
	}
}

// Get retrieves the user's avatar and writes it to the file specified by the passed
// in parameter.
//
func Get(filepath string) (err error) {
	page, err := bggclient.Get(getAvatarURL)
	if err != nil {
		message := fmt.Sprint("avatar.Get could not get page: %v", err)
		return errors.New(message)
	}

	pieces := avatarRegEx.FindStringSubmatch(page)
	if len(pieces) != 2 {
		message := fmt.Sprintf("avatar.Get could not parse avatar url: %v", err)
		return errors.New(message)
	}

	avatarURL, err := url.Parse(pieces[1])
	if err != nil {
		message := fmt.Sprintf("avatar.Get could not create avatar url: %v", err)
		return errors.New(message)
	}

	bytes, err := bggclient.Download(avatarURL)
	if err != nil {
		return
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(bytes)

	return
}

// Set reads the avatar from the passed in file name and uploads it to the server.
//
func Set(filepath string) (err error) {
	fields := map[string]string{
		"action":     "saveavatar",
		"domainname": "boardgamegeek.com",
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	err = bggclient.Upload(setAvatarURL, data, filepath, "filename", fields)

	return
}

// Local Variables:
// compile-command: "go build"
// End:
