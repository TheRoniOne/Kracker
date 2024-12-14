package utils

import (
	"io"
	"net/http"
)

func ReadRespBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
