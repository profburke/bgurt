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

// Package microbadge contains functions to retrieve and set microbadges.
//
package microbadge

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"

	"github.com/profburke/bgurt/bggclient"
)

type Group struct {
	GroupNumber uint
	Name        string
}

type Microbadge struct {
	BadgeNumber    uint
	Name           string
	Category       Group
	Subcategory    Group
	Subsubcategory Group
	NumberOfOwners uint
	Mouseover      string
	Creator        string
	Description    string
	ImageFilename  string
	// TODO: (low priority) related badge info
}

func (mb Microbadge) String() string {
	return fmt.Sprintf("Badge #%d: %s (%v > %v; %d owners)",
		mb.BadgeNumber, mb.Name, mb.Category.Name, mb.Subcategory.Name, mb.NumberOfOwners)
}

const TotalSlots = 5

var microbadgeListRegEx *regexp.Regexp
var metadataRegEx *regexp.Regexp
var microbadgeListURL *url.URL
var setSlotURL *url.URL

func init() {
	microbadgeListRegEx = regexp.MustCompile("<div id='badgename_(\\d+)'>([^<]+)</div>")
	metadataRegEx = regexp.MustCompile("(?s:<td>Group</td>\\s*<td>\\s*<a \\s*href=\"/microbadges/group/(\\d+)\"\\s*>(.*?)</a>\\s*<div class='ml10'>\\s*<a \\s*href=\"/microbadges/group/(\\d+)\"\\s*>(.*?)</a>.*?<td>Num Owners</td>\\s*<td>(\\d+)</td>)")
	var err error

	microbadgeListURL, err = url.Parse("microbadge/edit")
	if err != nil {
		log.Fatal("microbadge: could not create list URL")
	}

	setSlotURL, err = url.Parse("geekmicrobadge.php")
	if err != nil {
		log.Fatal("microbadge: could not create set slot URL")
	}
}

// SetAll takes a collection of microbadge IDs and sends them to the server to set
// as the displayed microbadges.
//
func SetAll(badgeNumbers []uint) (success bool, err error) {
	if len(badgeNumbers) != TotalSlots {
		message := fmt.Sprintf("microbadge.SetAll: must be called with %d badge numbers, received %d",
			TotalSlots, len(badgeNumbers))
		return false, errors.New(message)
	}

	for i, b := range badgeNumbers {
		// TODO: fire these off as Go routines ...
		// TODO: halt on first error? or try them all and return
		// a summary of successes?
		success, err = SetSlot(uint(i+1), b)
		if err != nil {
			return false, err
		}
	}

	return
}

// ValidSlot checks whether the specified number is a valid slot number,
// i.e. an integer between and 1 and the number of slots.
//
func ValidSlot(n uint) (result bool) {
	return 1 <= n && n <= TotalSlots
}

// TODO: any additional validation we can do for badge number?

// SetSlot updates the specified microbadge display slot with the microbadge specified
// by the ID passed in.
//
func SetSlot(slot, badgeNumber uint) (success bool, err error) {
	if !ValidSlot(slot) {
		message := fmt.Sprintf("microbadge.SetSlot: %d is an invalid slot number",
			slot)
		return false, errors.New(message)
	}

	data := url.Values{}
	data.Set("badgeid", fmt.Sprintf("%d", badgeNumber))
	data.Set("slot", fmt.Sprintf("%d", slot))
	data.Set("action", "setslot")
	data.Set("ajax", "1")

	resp, err := bggclient.Post(setSlotURL, data)
	if err != nil {
		return false, err
	}
	success = resp.StatusCode == 200
	return
}

// GetSlot returns the microbadge in the specified slot.
//
func GetSlot(slot uint) (mb Microbadge, err error) {
	if !ValidSlot(slot) {
		message := fmt.Sprintf("microbadge.GetSlot: %s is an invalid slot number",
			slot)
		return Microbadge{}, errors.New(message)
	}
	// TODO: implement function
	log.Fatal("not currently implemented")

	return
}

// ClearSlot clears the specified slot.
//
func ClearSlot(slot uint) (err error) {
	if !ValidSlot(slot) {
		message := fmt.Sprintf("microbadge.ClearSlot: %s is an invalid slot number",
			slot)
		return errors.New(message)
	}
	// TODO: implement function
	log.Fatal("not currently implemented")

	return
}

// TODO: look for subsubgroup...

func getMetadata(id uint) (category Group, subcategory Group, numOwners uint, err error) {
	metadataURL, err := url.Parse(fmt.Sprintf("microbadge/%d", id))
	if err != nil {
		message := fmt.Sprintf("microbadge.metadataURL: error creating URL: %v", err)
		return Group{}, Group{}, 0, errors.New(message)
	}

	page, err := bggclient.Get(metadataURL)
	if err != nil {
		message := fmt.Sprintf("microbadge.getMetadata (for badge# %d): could not get page: %v", id, err)
		return Group{}, Group{}, 0, errors.New(message)
	}

	pieces := metadataRegEx.FindStringSubmatch(page)
	if len(pieces) != 6 {
		return Group{}, Group{}, 0, errors.New("could not parse metadata")
	}

	var catid, subid uint
	if v, err := strconv.ParseUint(pieces[1], 10, 64); err == nil {
		catid = uint(v)
	} else {
		// TODO: error
	}

	if v, err := strconv.ParseUint(pieces[3], 10, 64); err == nil {
		subid = uint(v)
	} else {
		// TODO: error
	}

	if v, err := strconv.ParseUint(pieces[5], 10, 64); err == nil {
		numOwners = uint(v)
	} else {
		// TODO: error
	}

	category = Group{catid, pieces[2]}
	subcategory = Group{subid, pieces[4]}

	return
}

// TODO: a lot more error handling

// GetAll returns a collection of all the user's microbadges.
//
func GetAll() (badges []Microbadge, err error) {
	page, err := bggclient.Get(microbadgeListURL)
	if err != nil {
		message := fmt.Sprintf("microbadge.GetAll: could not get microbadge list: %v", err)
		return nil, errors.New(message)
	}

	matches := microbadgeListRegEx.FindAllStringSubmatch(page, -1)

	for _, match := range matches {
		mb := Microbadge{}
		mb.Name = string(match[2])

		var badgeNumber uint
		if v, err := strconv.Atoi(match[1]); err == nil {
			badgeNumber = uint(v)
		} else {
			// TODO: error handling
		}
		mb.BadgeNumber = badgeNumber

		category, subcategory, numOwners, err := getMetadata(mb.BadgeNumber)
		if err != nil {
			// TODO: error
		} else {
			mb.NumberOfOwners = numOwners
			mb.Category = category
			mb.Subcategory = subcategory
		}
		badges = append(badges, mb)
	}

	return
}

// Local Variables:
// compile-command: "go build"
// End:
