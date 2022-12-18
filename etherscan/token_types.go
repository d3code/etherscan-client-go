package etherscan

type Transactions struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Result  []struct {
        BlockNumber       int    `json:"blockNumber"`
        TimeStamp         int    `json:"timeStamp"`
        Hash              string `json:"hash"`
        Nonce             string `json:"nonce"`
        BlockHash         string `json:"blockHash"`
        TransactionIndex  string `json:"transactionIndex"`
        From              string `json:"from"`
        To                string `json:"to"`
        Value             string `json:"value"`
        Gas               string `json:"gas"`
        GasPrice          string `json:"gasPrice"`
        IsError           bool   `json:"isError"`
        TxreceiptStatus   string `json:"txreceipt_status"`
        Input             string `json:"input"`
        ContractAddress   string `json:"contractAddress"`
        CumulativeGasUsed string `json:"cumulativeGasUsed"`
        GasUsed           string `json:"gasUsed"`
        Confirmations     int    `json:"confirmations"`
        MethodId          string `json:"methodId"`
        FunctionName      string `json:"functionName"`
    } `json:"result"`
}

type TokenTransactions struct {
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
