package requests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Get(url string) (response string, getError error) {
	getResponse, getError := http.Get(url)
	if getError != nil {
		return "", fmt.Errorf("unsuccessfull GET request to %v; %w", url, getError)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(getResponse.Body)

	stringWriter := &strings.Builder{}
	writeError := getResponse.Write(stringWriter)
	if writeError != nil {
		return "", fmt.Errorf("could not write http response; %w", writeError)
	}

	return stringWriter.String(), nil
}
