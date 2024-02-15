package rest

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/JMjirapat/qrthrough-api/pkg/utils"
	"github.com/bytedance/sonic"
)

func PrepareHttpRequest(method, url string, headers map[string]string, data interface{}) (*http.Request, error) {
	var req *http.Request
	var err error
	if data != nil {
		switch v := data.(type) {
		case string:
			req, err = http.NewRequest(method, url, bytes.NewBufferString(v))
		default:
			if encodedData, err := sonic.Marshal(data); err != nil {
				return nil, err
			} else {
				bufferedData := bytes.NewBuffer(encodedData)
				req, err = http.NewRequest(method, url, bufferedData)
			}
		}
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Content-Type", "application/json")
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	log.Print(req.Body)

	return req, nil
}

func PrepareHttpResponse[R interface{}](req *http.Request) (*R, int, error) {
	// request to endpoint
	client := &http.Client{}
	res, err := client.Do(req)
	log.Print(res.StatusCode)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	// read a bytes of data
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// check empty response body
	if len(resBody) == 0 {
		return nil, res.StatusCode, nil
	}

	// decode response body
	var result R
	if err := utils.Recast(resBody, &result); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &result, res.StatusCode, nil
}

func HttpPost[R any](url string, headers map[string]string, data interface{}) (*R, int, error) {

	req, err := PrepareHttpRequest(http.MethodPost, url, headers, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return PrepareHttpResponse[R](req)
}
