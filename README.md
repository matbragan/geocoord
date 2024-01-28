# Geocoord

## Usage
#### Initialize your module
```sh
go mod init example.com/demo
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
    coordinates, err := geocoord.GetCoordinates(zip_code)
    
    if err != nil {
        fmt.Printf("Error: %v/n", err)
        return
    }
    
    fmt.Printf("Latitude: %f, Longitude: %f\n", coordinates.Latitude, coordinates.Longitude)
}
```

## Documentation

You can access the full documentation here: <br>
[![GoDoc](https://godoc.org/github.com/matbragan/geocoord?status.svg)](https://godoc.org/github.com/matbragan/geocoord)
