package main

import (
    "fmt"
    "net/http"
)

func retMsg(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("success"))
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/getsomething", retMsg)
    err := http.ListenAndServe(":7777", mux)
    if err != nil {
        fmt.Println("listen failed err is ", err)
    }

}
