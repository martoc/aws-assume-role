package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
	"fmt"
	"flag"
)

func main() {
	var region string
	var roleArn string
	var operation string
	flag.StringVar(&region, "region", "us-east-1", "-region us-east-1")
	flag.StringVar(&roleArn, "role-arn", "", "-role-arn <ARN>")
	flag.StringVar(&operation, "operation", "assume", "-operation assume")
	flag.Parse()
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	service := sts.New(session)
	credentials, _ := service.AssumeRole(&sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(3600),
		RoleSessionName: aws.String("aws-assume-role"),
		RoleArn: aws.String(roleArn),
	})
	if operation == "assume" {
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", *credentials.Credentials.AccessKeyId)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", *credentials.Credentials.SecretAccessKey)
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", *credentials.Credentials.SessionToken)
	} else {
		fmt.Printf("unset AWS_ACCESS_KEY_ID\n")
		fmt.Printf("unset AWS_SECRET_ACCESS_KEY\n")
		fmt.Printf("unset AWS_SESSION_TOKEN\n")
	}
}
