package main

import (
  "fmt"
  "net/http"
  )

func Home(w http.responsewriter, r *http.request) {
    fmt.print(w, "Home Page")
  }
func Handleresquest() {
  http.Handlefunc("/", Home)
  log.Fatal(http.ListenAndServe(":8000",nil))
  }

func main() {
  fmt.PrintIn(Iniciando o servidor Rest com Go")
        Handleresquest()
  }
