# CoinMarketCap API Client

A Go client for the [CoinMarketCap API](https://coinmarketcap.com/api/documentation/v1/#section/Quick-Start-Guide)

The CoinMarketCap API is a suite of high-performance RESTful JSON endpoints that are specifically designed to meet the mission-critical demands of application developers, data scientists, and enterprise business platforms.

### Install

```shell
go get github.com/d3code/etherscan-client-go
```

### Import

```go
import (
    "github.com/d3code/cmc-client-go/cmc"
)
```

## Client

To use the client, create a CMC Client.

```go
package main

import (
    "github.com/d3code/cmc-client-go/cmc"
)

func main() {
    client := cmc.Client("<your-api-key>")
}
```

### Testing

To use the CMC sandbox environment, call `Test(true)` on the client. This will replace the API Key and the Base URL for the requests to use the sandbox environment.

You do not need to provide your API Key, though this can be provided and calling `Test(false)` will revert to using the supplied API Key and the production CMC environment.

```go
cmc.Client("").Test(true)
```

### Logging

There is no logging in this library, an `error` is returned if there are request / unmarshalling errors and can be handled accordingly.

To print the raw response, call `PrintResponse(true)` on the client.

```go
cmc.Client("").PrintResponse(true)
```