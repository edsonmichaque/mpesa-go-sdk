# M-Pesa SDK for Go

M-Pesa SDK for Go is an unofficial library aiming to help businesses integrating every [M-Pesa](https://developer.mpesa.vm.co.mz) operations to their Go applications.

## Contents

- [Features](#features)
- [Usage](#usage)
   - [Quickstart](#usage/scenario-1)
   - [Receive Money from a Mobile Account](#usage/scenario-1)
   - [Send Money to a Mobile Account](#usage/scenario-2)
   - [Send Money to a Business Account](#usage/scenario-3)
   - [Revert a Transaction](#usage/scenario-4)
   - [Query the Status of a Transaction](#usage/scenario-5)
   - [Examples](#usage/scenario-6)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Related Projects](#related-projects)
   - [Dependencies](#related-projects/dependencies)
   - [Friends](#related-projects/friends)
- [Contributing](#contributing)
- [Changelog](#changelog)
- [Authors](#authors)
- [Credits](#credits)
- [License](#license)

## Features <a name="features"></a>

- Receive money from a mobile account to a business account
- Send money from a business account to a mobile account
- Send money from a business account to a another business account
- Revert a transaction
- Query the status of a transaction

## Usage <a name="usage"></a>

### Quickstart <a name="#usage/scenario-1"></a>

### Receive Money from a Mobile Account <a name="#usage/scenario-2"></a>

```go
package main

import (
    "fmt"
    mpesa "github.com/edsonmichaque/mpesa-go-sdk"
)

func main() {
    client = mpesa.New(map[string]string {
        "apiKey": "<REPLACE>",              // API Key
        "publicKey": "<REPLACE>",           // Public Key
        "serviceProviderCode": "<REPLACE>", // input_ServiceProviderCode
    })

    paymentData := map[string]string {
        "from": "841234567",           // input_CustomerMSISDN
        "reference": "11114",          // input_ThirdPartyReference
        "transation": "T12344CC",      // input_TransactionReference
        "amount": "10"                 // input_Amountss
    }

    if response, error := client.Receive(paymentData); error != nil {
        fmt.Println(response.Data)
    }
}
```

### Send Money to a Mobile Account <a name="#usage/scenario-3"></a>

### Send Money to a Business Account <a name="#usage/scenario-4"></a>

### Revert a Transaction <a name="#usage/scenario-5"></a>

### Query the Status of a Transaction <a name="#usage/scenario-6"></a>

### Examples <a name="usage/scenario-7"></a>

## Prerequisites <a name="prerequisites"></a>

- [Go 1.10+](https://golang.org) 

## Installation <a name="installation"></a>

## Configuration <a name="configuration"></a>

## Related Projects <a name="related-projects"></a>

### Dependencies <a name="related-projects/dependencies"></a>

#### Production Dependencies

#### Development Dependencies

### Friends <a name="related-projects/friends"></a>


## Contributing <a name="contributing"></a>

## Changelog <a name="changelog"></a>

## Authors <a name="authors"></a>

- [Edson Michaque](https://github.com/edsonmichaque)

## Credits <a name="credits"></a>

## License <a name="license"></a>

Copyright 2020 Edson Michaque

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

