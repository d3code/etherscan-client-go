package etherscan

import (
    "io"
    "net/http"
    "net/url"
    "time"
)

const (
    cmcBaseUrl = "https://api.etherscan.io/api"
)

type client struct {
    apiKey        string
    printResponse bool
}

func Client(apiKey string) *client {
    c := &client{
        apiKey:        apiKey,
        printResponse: false,
    }
    return c
}

func (c *client) PrintResponse(printResponse bool) *client {
    c.printResponse = printResponse
    return c
}

func doGetRequest(q url.Values, c *client) ([]byte, error) {
    req, err := http.NewRequest(http.MethodGet, cmcBaseUrl, nil)
    if err != nil {
        return nil, err
    }

    q.Add("apiKey", c.apiKey)

    req.Header.Set("Accepts", "application/json")
    req.URL.RawQuery = q.Encode()

    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    res, requestError := httpClient.Do(req)

    if requestError != nil {
        return nil, requestError
    }

    resBody, responseError := io.ReadAll(res.Body)
    if responseError != nil {
        return nil, responseError
    }

    if c.printResponse && resBody != nil {
        responseString := string(resBody)
        println("Response: ", responseString)
    }

    return resBody, nil
}
