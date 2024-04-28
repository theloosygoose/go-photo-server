package main


import (
    "net/http"
    "log"
)

func main() {
    r := http.FileServer(http.Dir("/mnt/usb/"))


    log.Fatal(http.ListenAndServe(":5000", r))
}
