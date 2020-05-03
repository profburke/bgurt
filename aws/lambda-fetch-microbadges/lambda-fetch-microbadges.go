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

// This Lambda function fetches all the microbadges and stores the information as a JSON
// file at the S3 Bucket path specified in the environment variables BUCKETNAME and ITEMNAME.
//
package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/profburke/bgg/aws/utilities"
	"github.com/profburke/bgg/bggclient"
	"github.com/profburke/bgg/microbadge"
)

func HandleRequest() {
	user := utilities.GetEnvOrDie("BGGUSERNAME")
	passhash := utilities.GetEnvOrDie("BGGPASSHASH")
	bucketname := utilities.GetEnvOrDie("BUCKETNAME")
	itemname := utilities.GetEnvOrDie("ITEMNAME")

	bggclient.SetCredentials(bggclient.Credentials{user, passhash})

	log.Println("Fetching microbadges...")

	badges, err := microbadge.GetAll()
	if err != nil {
		log.Println("Could not get microbadges: ", err)
	}

	if badges != nil {
		var jsonData []byte
		jsonData, err := json.Marshal(badges)
		if err != nil {
			log.Println(err)
		} else {
			// TODO: handle error
			awsSession, _ := session.NewSession(&aws.Config{
				Region: aws.String("us-east-1"),
			})

			s3Uploader := s3manager.NewUploader(awsSession)

			reader := bytes.NewReader(jsonData)

			uploadInput := &s3manager.UploadInput{
				Bucket: aws.String(bucketname),
				Key:    aws.String(itemname),
				Body:   reader,
			}

			_, err := s3Uploader.Upload(uploadInput)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("...uploaded micrboadges to %s:%s\n",
					bucketname, itemname)
			}
		}
	} else {
		log.Println("...no badges")
	}
}

func main() {
	lambda.Start(HandleRequest)
}

// Local Variables:
// compile-command: "GOOS=linux go build"
// End:
