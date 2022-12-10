package etherscan

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
