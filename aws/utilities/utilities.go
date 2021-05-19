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

// Package utilities provides common helper functions for the various lambda functions
// in the bgurt project.
//
package utilities

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/profburke/bgurt/microbadge"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// TODO: better name for struct
type Notification struct {
	Message string `json:"message"`
}

func GetEnvOrDie(key string) (value string) {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Environment variable %s not set.", key)
	}
	return
}

func GetEnvOrDefault(key, defaultValue string) (result string) {
	if value, ok := os.LookupEnv(key); ok {
		result = value
	} else {
		result = defaultValue
	}

	return
}

// TODO: make this "generic"
func Picksome(original []microbadge.Microbadge, n int) (picked []microbadge.Microbadge) {
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

// Local Variables:
// compile-command: "GOOS=linux go build"
// End:
