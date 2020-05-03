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

// This Lambda function is a simple, default implementation for randomly selecting
// microbadges to display. It downloads a list of microbadges from the specified S3 Bucket
// path and then randomly picks <TotalSlots> badges. The function outputs the list of selected
// IDs as a json object.
//
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/profburke/bgg/aws/utilities"
	"github.com/profburke/bgg/microbadge"
)

func downloadMicrobadges(bucketname, itemname string) (badges []microbadge.Microbadge, err error) {
	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	buffer := aws.NewWriteAtBuffer([]byte{})

	s3Downloader := s3manager.NewDownloader(awsSession)

	objectInput := &s3.GetObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(itemname),
	}

	_, err = s3Downloader.Download(buffer, objectInput)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not download microbadges from S3: %v", err))
	}

	err = json.Unmarshal(buffer.Bytes(), &badges)

	return badges, err
}

func HandleRequest() (response []byte, err error) {
	bucketname := utilities.GetEnvOrDie("MICROBADGE_BUCKETNAME")
	itemname := utilities.GetEnvOrDie("MICROBADGE_ITEMNAME")

	badges, err := downloadMicrobadges(bucketname, itemname)
	if err != nil {
		message := fmt.Sprintf("Could not retrieve microbadges: %v", err)
		log.Fatalf(message)
		return nil, errors.New(message)
	}

	chosen := utilities.Picksome(badges, microbadge.TotalSlots)

	log.Printf("Chosen microbadges: %v.", chosen)

	response, err = json.Marshal(chosen)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not convert badge list to JSON: %v.", err))
	}

	return response, err
}

func main() {
	lambda.Start(HandleRequest)
}

// Local Variables:
// compile-command: "GOOS=linux go build"
// End:
