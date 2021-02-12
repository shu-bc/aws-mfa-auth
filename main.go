package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide your token")
		return
	}
	token := os.Args[1]

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := sts.New(sess)
	identity, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Println(err)
		return
	}
	serial := strings.ReplaceAll(*identity.Arn, "user", "mfa")

	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
		SerialNumber:    &serial,
		TokenCode:       &token,
	}
	result, err := svc.GetSessionToken(input)
	if err != nil {
		fmt.Printf("failed get session token: %s", err)
		return
	}

	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", *result.Credentials.AccessKeyId)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", *result.Credentials.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%s\n", *result.Credentials.SessionToken)
}