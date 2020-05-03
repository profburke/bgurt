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

// This Lambda function expects a notification as input and sends the notification's
// message to the SNS topic specified by the TOPICARN environment variable.
//
// The notification type is defined in the package github.com/profburke/bgg/aws/utilities
//
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/profburke/bgg/aws/utilities"
)

func HandleRequest(notification utilities.Notification) (err error) {
	topicARN := utilities.GetEnvOrDie("TOPICARN")
	awsRegion := utilities.GetEnvOrDefault("AWSREGION", "us-east-1")

	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	snsService := sns.New(awsSession)

	params := &sns.PublishInput{
		Message:  aws.String(notification.Message),
		TopicArn: aws.String(topicARN),
	}

	resp, err := snsService.Publish(params)
	if err != nil {
		log.Printf("error from call to snsService.Publish: %v", err)
		return errors.New(fmt.Sprintf("error from call to snsService.Publish: %v", err))
	}

	log.Printf("response: %v", resp)
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}

// Local Variables:
// compile-command: "GOOS=linux go build"
// End:
