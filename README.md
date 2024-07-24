# Trustpilot Client
A Trustpilot client for Go programs. Useful to retrieve data from Trustpilot, such as the reviews statistic and star rating SVG.

### Installation
> go get github.com/z3ntl3/trustpilot/client

#### Posibilities
- Option to thoroughly customize the underlying HTTP client this package uses.

    > Useful to proxy or to set a custom dial timeout or anything alike etc.

#### Model
The model you can easily acquire from Trustpilot, using this client package:

```
{
    Company: "IKEA"  
    Reviews: {
        Total: "1.5" 
        Text: "24,080   •   Bad"
    } 
    ImageSRC: "https://cdn.trustpilot.net/brand-assets/4.1.0/stars/stars-1.5.svg"
    
}
```

#### API example
```go
package main

import (
	"fmt"
	"log"

	"github.com/z3ntl3/trustpilot/client"
)

var URL = "https://trustpilot.com/review/www.ikea.com"

func main() {
	client := client.New()

	result, err := client.Execute(URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Resulting date model: %+v \n", result)
}
```
