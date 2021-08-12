# Bitwin SDK for Go (v3)

🌍 *[English](README.md) ∙ [繁體中文](README_zh-TW.md) ∙ [简体中文](README_zh-CN.md)*

[![Build Status](https://travis-ci.com/xup6m6fu04/bitwin-sdk-go.svg?branch=master)](https://travis-ci.com/xup6m6fu04/bitwin-sdk-go)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fxup6m6fu04%2Fbitwin-sdk-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fxup6m6fu04%2Fbitwin-sdk-go?ref=badge_shield)
[![codecov](https://codecov.io/gh/xup6m6fu04/bitwin-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/xup6m6fu04/bitwin-sdk-go)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/xup6m6fu04/bitwin-sdk-go/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/xup6m6fu04/bitwin-sdk-go)](https://goreportcard.com/report/github.com/xup6m6fu04/bitwin-sdk-go)

## 簡介
為了減少使用者串接 API 所耗費時間過多，你可以使用這個 SDK 輕鬆的建立與串接 BITWIN 的商戶 API

## 文件

查看官方 API 文件已獲得更多細節

- 簡體中文: https://bitwin.ai/api_manual_latest.html

## 需求

至少需要 Go 1.13 或以上版本.

## 安裝 ##

```sh
$ go get -u github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin
```

## 構建 ##

```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    ...
}

```

## 範例 ##
### 建立付款訂單

```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    // true 是正式環境, false 是測試環境, 預設為 false
    client.CreateCryptoPayOrder.SetIsProdEnvironment(false)
    client.CreateCryptoPayOrder.SetMerchantUserID("YOZERO_USER_0000001")
    client.CreateCryptoPayOrder.SetMerchantOrderID("YOZERO_ORDER_0000001")
    client.CreateCryptoPayOrder.SetOrderDescription("YOZERO_DESC_0000001")
    client.CreateCryptoPayOrder.SetAmount("700000000") // 7 USDT
    client.CreateCryptoPayOrder.SetMerchantRMB("45.38")
    client.CreateCryptoPayOrder.SetSymbol("USDT_ERC20")
    client.CreateCryptoPayOrder.SetCallBackURL("https://yourdomain/api/callback")
    client.CreateCryptoPayOrder.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
    // 可以使用 execute 變數來取得更多資訊
    execute, err := client.CreateCryptoPayOrder.Execute()
    if err != nil {
        fmt.Print(err)
        return
    }
    empJSON, err := json.MarshalIndent(execute, "", "    ")
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("CreateCryptoPayOrder Response\n%s\n", string(empJSON))
}

```
#### 如果成功會顯示像是以下結果
```json
CreateCryptoPayOrder Response
{
    "OrderId": "53190266162719072",
    "Qrcode": "https://stage-api.bitwin.ai/order/53190266162719072",
    "Amount": "700000000",
    "RealAmount": "700000000",
    "CryptoWallet": "0x54420B5AB4a26Aa294C4F039052Ca3A7069E2C99",
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "474A8E4E3A8FF8FDC5D392BDCB4672D9"
}
```

### 查詢付款訂單
```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    // true 是正式環境, false 是測試環境, 預設為 false
    client.QueryCryptoPayOrder.SetIsProdEnvironment(false)
    // SetMerchantOrderID or SetOrderID 只需要選擇填寫一個，也可以兩者都填
    client.QueryCryptoPayOrder.SetMerchantOrderID("YOZERO_ORDER_0000001")
    client.QueryCryptoPayOrder.SetOrderID("53190266162719072")
    client.QueryCryptoPayOrder.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
    // 可以使用 execute 變數來取得更多資訊
    execute, err := client.QueryCryptoPayOrder.Execute()
    if err != nil {
        fmt.Print(err)
        return
    }
    empJSON, err := json.MarshalIndent(execute, "", "    ")
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("QueryCryptoPayOrder Response\n%s\n", string(empJSON))
}

```
#### 如果使用者尚未付款會顯示
```json
QueryCryptoPayOrder Response
{
    "OrderId": "53190266162719072",
    "MerchantOrderId": "YOZERO_ORDER_0000001",
    "MerchantUserId": "YOZERO_USER_0000001",
    "OrderDescription": "YOZERO_DESC_0000001",
    "Symbol": "USDT_ERC20",
    "Amount": "700000000",
    "RealAmount": "700000000",
    "MerchantRMB": "45.38",
    "ExchangeRMB": "45.85",
    "OrderStatus": "PENDING",
    "CallBackUrl": "https://yourdomain/api/callback",
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "8B7FA6A133EB46FEBBEA21870FE36C4E"
}
```
#### 如果使用者付款成功會顯示
```json
QueryCryptoPayOrder Response
{
    "OrderId": "53190266162719072",
    "MerchantOrderId": "YOZERO_ORDER_0000001",
    "MerchantUserId": "YOZERO_USER_0000001",
    "OrderDescription": "YOZERO_DESC_0000001",
    "Symbol": "USDT_ERC20",
    "Amount": "700000000",
    "RealAmount": "700000000",
    "MerchantRMB": "45.38",
    "ExchangeRMB": "45.85",
    "OrderStatus": "SUCCESS",
    "CallBackUrl": "https://yourdomain/api/callback",
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "1227B59B65769CC8E497228638A54D40"
}
```
#### 如果使用者付款成功，你的回調伺服器將會收到像是以下 JSON 內容
```json
{
  "MerchantId": "<Merchant ID>",
  "MerchantOrderId": "YOZERO_ORDER_0000001",
  "OrderId": "53190266162719072",
  "Symbol": "USDT_ERC20",
  "Amount": "700000000",
  "PayAmount": "700000000",
  "MerchantRMB": "45.38",
  "ExchangeRMB": "45.85",
  "PayUnixTimestamp": 1628563840,
  "Sign": "78AEA2AEA64681F2CE96AE537CDDFADA"
}
```
關於回調詳情請參閱文件

### 建立商戶出款單

```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    // true 是正式環境, false 是測試環境, 預設為 false
    client.MerchantWithdraw.SetIsProdEnvironment(false)
    client.MerchantWithdraw.SetMerchantUserID("YOZERO_USER_0000001")
    client.MerchantWithdraw.SetMerchantWithdrawID("YOZERO_WITHDRAW_0000001")
    client.MerchantWithdraw.SetUserWallet("0x875EDa094F03Ed4c93adb3dbb77913F860dC888f")
    client.MerchantWithdraw.SetAmount("1000000000") // 10 USDT
    client.MerchantWithdraw.SetMerchantRMB("64.81")
    client.MerchantWithdraw.SetSymbol("USDT_ERC20")
    client.MerchantWithdraw.SetCallBackURL("https://yourdomain/api/callback")
    client.MerchantWithdraw.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
    // 可以使用 execute 變數來取得更多資訊
    execute, err := client.MerchantWithdraw.Execute()
    if err != nil {
        fmt.Print(err)
        return
    }
    empJSON, err := json.MarshalIndent(execute, "", "    ")
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("MerchantWithdraw Response\n%s\n", string(empJSON))
}

```
#### 如果成功會顯示像是以下結果
```json
MerchantWithdraw Response
{
    "WithdrawId": "53192080311396704",
    "MerchantWithdrawId": "YOZERO_WITHDRAW_0000001",
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "878DCEE083E966D661DD2070B968850D"
}
```
### 查詢商戶出款單

```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    // true 是正式環境, false 是測試環境, 預設為 false
    client.QueryMerchantWithdraw.SetIsProdEnvironment(false)
    // SetMerchantWithdrawID or SetWithdrawID 只需要選擇填寫一個，也可以兩者都填
    client.QueryMerchantWithdraw.SetMerchantWithdrawID("YOZERO_WITHDRAW_0000001")
    client.QueryMerchantWithdraw.SetWithdrawID("53192080311396704")
    client.QueryMerchantWithdraw.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
    // 可以使用 execute 變數來取得更多資訊
    execute, err := client.QueryMerchantWithdraw.Execute()
    if err != nil {
        fmt.Print(err)
        return
    }
    empJSON, err := json.MarshalIndent(execute, "", "    ")
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("QueryMerchantWithdraw Response\n%s\n", string(empJSON))
}

```
#### 如果提領正在等待中將會顯示像是
```json
QueryMerchantWithdraw Response
{
    "MerchantUserId": "YOZERO_USER_0000001",
    "UserWallet": "0x875EDa094F03Ed4c93adb3dbb77913F860dC888f",
    "MerchantWithdrawId": "YOZERO_WITHDRAW_0000001",
    "WithdrawId": "53192080311396704",
    "Symbol": "USDT_ERC20",
    "Amount": "1000000000",
    "MerchantRMB": "64.81",
    "ExchangeRMB": "68.00",
    "Status": "PENDING",
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "0D7827A91A2DFEC4A37930B7B63CB3D7"
}
```
#### 如果已經提領成功將會顯示像是
```json
QueryMerchantWithdraw Response
{
    "MerchantUserId": "YOZERO_USER_0000001",
    "UserWallet": "0x875EDa094F03Ed4c93adb3dbb77913F860dC888f",
    "MerchantWithdrawId": "YOZERO_WITHDRAW_0000001",
    "WithdrawId": "53192080311396704",
    "Symbol": "USDT_ERC20",
    "Amount": "1000000000",
    "MerchantRMB": "64.81",
    "ExchangeRMB": "68.00",
    "Status": "SUCCESS",
    "WithdrawDateTime": 1628565087,
    "ApprovedDateTime": 1628564721,
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "0031A5895817164033D39618C4C687B1"
}
```
#### 如果提領成功，你的回調伺服器將會收到像是以下 JSON 內容
```json
{
  "MerchantId": "<Merchant ID>",
  "MerchantUserId": "YOZERO_USER_0000001",
  "MerchantWithdrawId": "YOZERO_WITHDRAW_0000001",
  "UserWallet": "0x875EDa094F03Ed4c93adb3dbb77913F860dC888f",
  "WithdrawId": "53192080311396704",
  "WithdrawAmount": "1000000000",
  "MerchantRMB": "64.81",
  "ExchangeRMB": "68.00",
  "Symbol": "USDT_ERC20",
  "ReplyDateTime": 1628565087,
  "Sign": "AA367497C3BC87FC4E544D669C58DDAF"
}
```

關於回調詳情請參閱文件

### 查詢建議匯率

```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    // true 是正式環境, false 是測試環境, 預設為 false
    client.ExchangeRate.SetIsProdEnvironment(false)
    client.ExchangeRate.SetSymbol("USDT_ERC20")
    // 可以使用 execute 變數來取得更多資訊
    execute, err := client.ExchangeRate.Execute()
    if err != nil {
        fmt.Print(err)
        return
    }
    empJSON, err := json.MarshalIndent(execute, "", "    ")
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("ExchangeRate Response\n%s\n", string(empJSON))
}
```
#### 結果將像是
```json
ExchangeRate Response
{
    "RMBRate": "6.55",
    "RMBBuyRate": "6.80",
    "ReturnCode": "200",
    "ReturnMessage": ""
}

```

### BITWIN 會員錢包綁定

```go
import (
    "github.com/xup6m6fu04/bitwin-sdk-go/v3/bitwin"
)

func main() {
    client := bitwin.New("<Merchant ID>", "<Merchant Sign key>")
    // true 是正式環境, false 是測試環境, 預設為 false
    client.BuildRelationUser.SetIsProdEnvironment(false)
    client.BuildRelationUser.SetMerchantUserID("YOZERO_USER_0000001")
    client.BuildRelationUser.SetCallBackURL("https://yourdomain/api/callback")
    client.BuildRelationUser.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
    // 可以使用 execute 變數來取得更多資訊
    execute, err := client.BuildRelationUser.Execute()
    if err != nil {
        fmt.Print(err)
        return
    }
    empJSON, err := json.MarshalIndent(execute, "", "    ")
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("BuildRelationUser Response\n%s\n", string(empJSON))
}
```
#### 結果將像是
```json
BuildRelationUser Response
{
    "QrcodeData": "BITWIN$eyJhY3Rpb24iOiJtb2JpbGUvdjMvdXNlci9iaW5kIiwiZGF0YSI6eyJuYW1lIjoiOTkwODg1MzBEMDkzQzRDRkY0MzEyRjAyRUY3RkI1RjYiLCJjb2RlIjoiRUhrSWoiLCJtZXJjaGFudF9uYW1lIjoieW96ZXJvIn19",
    "QrcodeImageUrl": "https://stage-api.bitwin.ai/web/v3/bind/user/99088530D093C4CFF4312F02EF7FB5F6",
    "ReturnCode": "200",
    "ReturnMessage": "",
    "Sign": "414E783ECC46DB9403916694741DF275"
}
```

#### 當使用者掃描並認證後, 你的回調伺服器將會收到像是以下 JSON 內容
```json
{
  "MerchantId": "<Merchant ID>",
  "MerchantUserId": "YOZERO_USER_0000001",
  "UserName": "48847933077253904",
  "BTC": "2My4ttAncyVKbAQWwAMLsG7JCMif3KkpHBC",
  "ETH": "0xe4f3Ad1005ac2FbD22f7F22871A8Ea1d688866a0",
  "USDT_ERC20": "0x7F8FAe2d400cD767d4184638eD296DBc44F218Bb",
  "USDT_TRC20": "TYGzJX3tyDy81eQGGw92US821LiykuHPFi",
  "USDT_BEP20": "0x84e6B02d0223c004bc350F481038371Cfd7e4512",
  "Sign": "9C516FBC1513EB54CC628C7AB5C105AF"
}
```

## Versioning
This project respects semantic versioning.

See http://semver.org/

## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fxup6m6fu04%2Fbitwin-sdk-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fxup6m6fu04%2Fbitwin-sdk-go?ref=badge_large)
