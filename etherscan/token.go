package etherscan

import (
    "encoding/json"
    "net/url"
)

func (c *client) GetTokenInfoByContractAddress(contractAddress string) (*TokenInfoByContractAddress, []byte, error) {

    values := url.Values{}
    values.Add("module", "token")
    values.Add("action", "tokeninfo")
    values.Add("contractaddress", contractAddress)

    resBody, err := doGetRequest(values, c)
    if err != nil {
        return nil, resBody, err
    }

    var etherscanResponse TokenInfoByContractAddress

    unmarshalError := json.Unmarshal(resBody, &etherscanResponse)
    if unmarshalError != nil {
        return nil, resBody, unmarshalError
    }

    return &etherscanResponse, resBody, nil
}
