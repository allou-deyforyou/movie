package handler

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func ServeResponse(data interface{}) (Response, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(data)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
	}
	return resp, nil
}
