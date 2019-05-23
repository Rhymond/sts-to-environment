package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

var (
	RoleArn      string
	SerialNumber string
	Profile      string
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("only two argument is required: go run main.go <mfa>")
	}

	err := os.Setenv("AWS_PROFILE", Profile)
	if err != nil {
		log.Fatalf("failed to set aws profile env: %s", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(endpoints.EuWest2RegionID),
	})
	if err != nil {
		log.Fatalf("failed to create aws session: %s", err)
	}

	s := sts.New(sess)
	req, resp := s.AssumeRoleRequest(&sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(3600),
		RoleArn:         aws.String(RoleArn),
		RoleSessionName: aws.String(strconv.Itoa(time.Now().Nanosecond())),
		SerialNumber:    aws.String(SerialNumber),
		TokenCode:       aws.String(args[1]),
	})
	err = req.Send()
	if err != nil {
		log.Fatalf("failed to sts assume role: %s", err)
	}

	var command string
	command += fmt.Sprintf(`export AWS_SECRET_ACCESS_KEY="%s"`, *resp.Credentials.SecretAccessKey)
	command += " && "
	command += fmt.Sprintf(`export AWS_ACCESS_KEY_ID="%s"`, *resp.Credentials.AccessKeyId)
	command += " && "
	command += fmt.Sprintf(`export AWS_SESSION_TOKEN="%s"`, *resp.Credentials.SessionToken)
	fmt.Print(command)
}
