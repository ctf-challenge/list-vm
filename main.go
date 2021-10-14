package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)


func main() {
    secret_token := "secretsecretsecret"  // アカウントに払い出された鍵情報を設定してください。
    url := fmt.Sprintf("https://cloudnet-security-challenge.co.jp/list?token=%s", secret_token)

    resp, err := http.Get(url)
    if err != nil {
        panic("fail")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic("fail")
    }
    fmt.Println(string(body))
}

