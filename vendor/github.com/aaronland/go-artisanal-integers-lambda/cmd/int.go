package main

import (
	"flag"
	"github.com/aaronland/go-artisanal-integers-lambda/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func main() {

	var lambda_func = flag.String("lambda-func", "NextInt", "...")

	flag.Parse()

	var aws_sess *session.Session // fix me...

	cl, err := client.NewLambdaClient(aws_sess, *lambda_func)

	if err != nil {
		log.Fatal(err)
	}

	int, err := cl.NextInt()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(int)
}
