package etherscan

import (
    "net/http"
    "net/url"
)

const (
    baseUrl     = "https://api.etherscan.io/api"
    apiKeyParam = "apikey"
)

type Configuration struct {
    apiKey        string
    printResponse bool
    replay        bool
    replayPath    string
}

func Client(apiKey string) *Configuration {
    c := &Configuration{
        apiKey:        apiKey,
        printResponse: false,
        replay:        false,
    }
    return c
}

func (c *Configuration) PrintResponse(printResponse bool) *Configuration {
    c.printResponse = printResponse
    return c
}

func doGetRequest(q url.Values, c *Configuration) ([]byte, error) {
    req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
    if err != nil {
        return nil, err
    }

    q.Add(apiKeyParam, c.apiKey)

    req.Header.Set("Accepts", "application/json")
    req.URL.RawQuery = q.Encode()

    resBody, requestError := doRequest(req, c.replay, c.replayPath, []string{}, []string{apiKeyParam})
    if requestError != nil {
        return nil, requestError
    }

    if c.printResponse && resBody != nil {
        responseString := string(resBody)
        println("Response: ", responseString)
    }

    return resBody, nil
}
