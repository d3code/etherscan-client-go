package etherscan

import "time"

type TokenTransactions struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Result  []struct {
        BlockNumber       int       `json:"blockNumber"`
        TimeStamp         time.Time `json:"timeStamp"`
        Hash              string    `json:"hash"`
        Nonce             string    `json:"nonce"`
        BlockHash         string    `json:"blockHash"`
        TransactionIndex  string    `json:"transactionIndex"`
        From              string    `json:"from"`
        To                string    `json:"to"`
        Value             string    `json:"value"`
        Gas               string    `json:"gas"`
        GasPrice          string    `json:"gasPrice"`
        IsError           bool      `json:"isError"`
        TxreceiptStatus   string    `json:"txreceipt_status"`
        Input             string    `json:"input"`
        ContractAddress   string    `json:"contractAddress"`
        CumulativeGasUsed string    `json:"cumulativeGasUsed"`
        GasUsed           string    `json:"gasUsed"`
        Confirmations     int       `json:"confirmations"`
        MethodId          string    `json:"methodId"`
        FunctionName      string    `json:"functionName"`
    } `json:"result"`
}

type TokenInfoByContractAddress struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Result  []struct {
        ContractAddress string `json:"contractAddress"`
        TokenName       string `json:"tokenName"`
        Symbol          string `json:"symbol"`
        Divisor         string `json:"divisor"`
        TokenType       string `json:"tokenType"`
        TotalSupply     string `json:"totalSupply"`
        BlueCheckmark   string `json:"blueCheckmark"`
        Description     string `json:"description"`
        Website         string `json:"website"`
        Email           string `json:"email"`
        Blog            string `json:"blog"`
        Reddit          string `json:"reddit"`
        Slack           string `json:"slack"`
        Facebook        string `json:"facebook"`
        Twitter         string `json:"twitter"`
        Bitcointalk     string `json:"bitcointalk"`
        Github          string `json:"github"`
        Telegram        string `json:"telegram"`
        Wechat          string `json:"wechat"`
        Linkedin        string `json:"linkedin"`
        Discord         string `json:"discord"`
        Whitepaper      string `json:"whitepaper"`
        TokenPriceUSD   string `json:"tokenPriceUSD"`
    } `json:"result"`
}

type TokenTransfers struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Result  []struct {
        BlockNumber       string `json:"blockNumber"`
        TimeStamp         string `json:"timeStamp"`
        Hash              string `json:"hash"`
        Nonce             string `json:"nonce"`
        BlockHash         string `json:"blockHash"`
        From              string `json:"from"`
        ContractAddress   string `json:"contractAddress"`
        To                string `json:"to"`
        Value             string `json:"value"`
        TokenName         string `json:"tokenName"`
        TokenSymbol       string `json:"tokenSymbol"`
        TokenDecimal      string `json:"tokenDecimal"`
        TransactionIndex  string `json:"transactionIndex"`
        Gas               string `json:"gas"`
        GasPrice          string `json:"gasPrice"`
        GasUsed           string `json:"gasUsed"`
        CumulativeGasUsed string `json:"cumulativeGasUsed"`
        Input             string `json:"input"`
        Confirmations     string `json:"confirmations"`
    } `json:"result"`
}
