# Geocoord

### Usage
#### Initialize your module
```sh
go mod init example.com/my-golib-demo
```

#### Get the geocoord package
```sh
go get github.com/matbragan/geocoord
```

#### Example of usage
```go
package main

import (
    "fmt"

    "github.com/matbragan/geocoord"
)

func main() {
    zip_code := "87109"
    coordinates, err := getCoordinates(zip_code)
    
    if err != nil {
        fmt.Printf("Error: %v/n", err)
        return
    }
    
    fmt.Printf("Latitude: %f, Longitude: %f\n", coordinates.Latitude, coordinates.Longitude)
}
```
