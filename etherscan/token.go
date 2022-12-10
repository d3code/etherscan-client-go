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

func (c *client) GetTokenTransfersByContractAddress(contractAddress string, address string,
    page string, offset string, startBlock string, endBlock string, sort string) (*TokenInfoByContractAddress, []byte, error) {

    values := url.Values{}
    values.Add("module", "account")
    values.Add("action", "tokentx")
    values.Add("contractaddress", contractAddress)
    values.Add("address", address)
    values.Add("page", page)
    values.Add("offset", offset)
    values.Add("startblock", startBlock)
    values.Add("endblock", endBlock)
    values.Add("sort", sort)

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
