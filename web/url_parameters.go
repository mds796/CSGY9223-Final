package web

import (
	"errors"
	"io"
	"io/ioutil"
	"net/url"
)

func getParameters(reader io.Reader) (url.Values, error) {
	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	parameters := string(bytes)
	values, err := url.ParseQuery(parameters)

	return values, err
}

func getKey(parameters map[string][]string, key string) (string, error) {
	values, ok := parameters[key]
	if !ok || len(values) == 0 || len(values[0]) == 0 {
		return "", errors.New("Please provide a " + key + ".")
	}

	return values[0], nil
}
