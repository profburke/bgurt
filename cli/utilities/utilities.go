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

// Package utilities provides several functions for use by the various cli programs.
//
package utilities

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/profburke/bgg/bggclient"
	"github.com/profburke/bgg/microbadge"
)

// TODO: create a report function that writes to stderr, stdout (if verbose),
//       and logfile (if configured) ???

const configFilename = "config.toml"
const badgeFilename = "badges.json"
const AppName = "bgurt"

func ConfigDir() (string, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, AppName), nil
}

func ConfigFilename() (string, error) {
	dirname, err := ConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dirname, configFilename), nil
}

func BadgeFilename() (string, error) {
	dirname, err := ConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dirname, badgeFilename), nil
}

// TODO: refactor the next two functions

func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// FileExists returns a bool indicating whether the specified file exists (and is
// not a directory).
//
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func LoadBadges(filepath string) (badges []microbadge.Microbadge, err error) {
	return nil, errors.New("not implemented yet")
}

// PrintErrorAndDie prints the specified message to standard error and exits
// the program with a return code of 1.
//
func PrintErrorAndDie(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

// SetCredentials retrieves the username and password hash and configures the
// bggclient object. Environment variables take priority. If they are not set,
// try and load from config file. If still not set, emit error and quit program.
//
// TODO: allow override from command line
// TODO: this function does too much; refactor needed
//
func SetCredentials() (credentials bggclient.Credentials) {
	var username, passhash string

	var configCredentials bggclient.Credentials

	cfile, err := ConfigFilename()
	_, err = toml.DecodeFile(cfile, &configCredentials)
	if err != nil {
		// TODO: figure out how to handle error
	}

	if v, ok := os.LookupEnv("BGGUSERNAME"); ok {
		username = v
	}

	if v, ok := os.LookupEnv("BGGPASSHASH"); ok {
		passhash = v
	}

	envCredentials := bggclient.Credentials{username, passhash}

	if !envCredentials.IsSet() && !configCredentials.IsSet() {
		message := `
Username or password hash missing.
Either set the BGGUSERNAME and BGGPASSHASH environment variables.
Or set username and password hash in the configuration file.
`
		PrintErrorAndDie(message)
	}

	if envCredentials.IsSet() {
		credentials = envCredentials
	} else {
		credentials = configCredentials
	}
	bggclient.SetCredentials(credentials)

	return
}

// WriteToFile writes data to file. If force is false and the file exists, returns error
// rather than overwriting the file.
//
func WriteToFile(filename, toolname string, force bool, jsonData []byte) {
	if !force && FileExists(filename) {
		fmt.Fprintf(os.Stderr, "%s: '%s' exists", toolname, filename)
		return
	}

	err := ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatalf("%s: error writing to file: %v", toolname, err)
	}
}

// Local Variables:
// compile-command: "go build"
// End:
