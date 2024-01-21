package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    response, err := http.Get("http://example.com")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    response.Body.Close()
}

