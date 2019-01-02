package client

// untested, still...

import (
	"encoding/json"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaClient struct {
	artisanalinteger.Client
	service     *lambda.Lambda
	lambda_func string
}

type Integer struct {
	Integer int64 `json:"integer"`
}

func NewLambdaClient(sess *session.Session, lambda_func string) (artisanalinteger.Client, error) {

	service := lambda.New(sess)

	cl := LambdaClient{
		service:     service,
		lambda_func: lambda_func,
	}

	return &cl, nil
}

func (cl *LambdaClient) NextInt() (int64, error) {

	input := &lambda.InvokeInput{
		FunctionName:   aws.String(cl.lambda_func),
		InvocationType: aws.String("RequestResponse"),
		LogType:        aws.String("Tail"),
	}

	rsp, err := cl.service.Invoke(input)

	if err != nil {
		return -1, err
	}

	/*

	if *rsp.StatusCode != 200 {
		return -1, errors.New(string(result))
	}
		enc_result := *rsp.LogResult

		result, err := base64.StdEncoding.DecodeString(enc_result)

		if err != nil {
			return -1, err
		}
	*/

	var int Integer

	err = json.Unmarshal(rsp.Payload, &int)

	if err != nil {
		return -1, err
	}

	return int.Integer, nil
}
