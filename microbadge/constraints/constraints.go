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

// Package constraints implements a data structure to describe which microbadges can
// be shown with in which slots and other constraints.
//
package constraints

import "github.com/profburke/bgg/microbadge"

type Constraints interface {
	Pick(badges []microbadge.Microbadge) (chosenBadges []microbadge.Microbadge)
}

type ConstraintsData struct {
}

func LoadConstraintsData(filepath string) (cd ConstraintsData, err error) {
	return ConstraintsData{}, nil
}

// TODO: I'm going back and forth between badge as int and badge as microbadge.Microbadge;
// need to decide
type badgeSlotPair struct {
	badge uint
	slot  uint
}

type defaultConstraints struct {
	disallowed map[badgeSlotPair]bool
}

func (dc defaultConstraints) IsAllowed(microbadge microbadge.Microbadge, slot uint) (result bool) {
	return !dc.IsNotAllowed(microbadge, slot)
}

// I have a hunch that there will be fewer disallowed pairs than allowed,
// so, if true, it makes sense to have this one do work while IsAllowed is just negating
// the result here...

func (dc defaultConstraints) IsNotAllowed(m microbadge.Microbadge, slot uint) (result bool) {
	_, result = dc.disallowed[badgeSlotPair{m.BadgeNumber, slot}]

	return
}

func (dc defaultConstraints) Pick(badges []microbadge.Microbadge) (chosenBadges []microbadge.Microbadge) {
	// create an empty set of possible microbadges for each slot
	// possibles := ...

	for mb := range badges {
		for slot := 1; slot < microbadge.TotalSlots; slot++ {
			if dc.IsAllowed(mb, slot) {
				// put badge in set for slot in the possibles data structure
			}
		}
	}

	for slot := 1; slot < microbadge.TotalSlots; slot++ {
		// p := randomly pick a badge from its set
		// chosenBadges = append(chosenBadges, p)

		for otherSlot := slot + 1; otherSlot < microbadge.TotalSlots; otherSlot++ {
			// remove p from otheSlot
		}
	}

	return
}

func New(constraintsData ConstraintsData) (constraints Constraints) {
	return defaultConstraints{}
}

// Local Variables:
// compile-command: "go build"
// End:
