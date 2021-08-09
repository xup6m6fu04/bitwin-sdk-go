package api

import (
	"encoding/json"
)

type CreateCryptoPayOrderService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           CreateCryptoPayOrderRequest
}

type CreateCryptoPayOrderRequest struct {
	MerchantId       string `json:"MerchantId"`
	MerchantUserId   string `json:"MerchantUserId,omitempty"`
	MerchantOrderId  string `json:"MerchantOrderId"`
	OrderDescription string `json:"OrderDescription,omitempty"`
	Amount           string `json:"Amount"`
	MerchantRMB      string `json:"MerchantRMB,omitempty"`
	Symbol           string `json:"Symbol"`
	CallBackUrl      string `json:"CallBackUrl"`
	TimeStamp        string `json:"TimeStamp"`
	Sign             string `json:"Sign"`
}

type CreateCryptoPayOrderResponse struct {
	OrderID       string `json:"OrderId,omitempty"`
	Qrcode        string `json:"Qrcode,omitempty"`
	Amount        string `json:"Amount,omitempty"`
	RealAmount    string `json:"RealAmount,omitempty"`
	CryptoWallet  string `json:"CryptoWallet,omitempty"`
	ReturnCode    string `json:"ReturnCode"`
	ReturnMessage string `json:"ReturnMessage"`
	Sign          string `json:"Sign"`
}

func NewCreateCryptoPayOrderService(merchantId string, signKey string) *CreateCryptoPayOrderService {
	return &CreateCryptoPayOrderService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: CreateCryptoPayOrderRequest{
			MerchantId:       merchantId,
			MerchantUserId:   "",
			MerchantOrderId:  "",
			OrderDescription: "",
			Amount:           "",
			MerchantRMB:      "",
			Symbol:           "",
			CallBackUrl:      "",
			TimeStamp:        "",
			Sign:             "",
		},
	}
}

func (srv *CreateCryptoPayOrderService) SetIsProdEnvironment(isProdEnvironment bool) *CreateCryptoPayOrderService {
	srv.IsProdEnvironment = isProdEnvironment
	return srv
}

func (srv *CreateCryptoPayOrderService) SetMerchantUserId(merchantUserId string) *CreateCryptoPayOrderService {
	srv.Request.MerchantUserId = merchantUserId
	return srv
}

func (srv *CreateCryptoPayOrderService) SetMerchantOrderId(merchantOrderId string) *CreateCryptoPayOrderService {
	srv.Request.MerchantOrderId = merchantOrderId
	return srv
}

func (srv *CreateCryptoPayOrderService) SetOrderDescription(orderDescription string) *CreateCryptoPayOrderService {
	srv.Request.OrderDescription = orderDescription
	return srv
}

func (srv *CreateCryptoPayOrderService) SetAmount(amount string) *CreateCryptoPayOrderService {
	srv.Request.Amount = amount
	return srv
}

func (srv *CreateCryptoPayOrderService) SetMerchantRMB(merchantRMB string) *CreateCryptoPayOrderService {
	srv.Request.MerchantRMB = merchantRMB
	return srv
}

func (srv *CreateCryptoPayOrderService) SetSymbol(symbol string) *CreateCryptoPayOrderService {
	srv.Request.Symbol = symbol
	return srv
}

func (srv *CreateCryptoPayOrderService) SetCallBackUrl(callBackUrl string) *CreateCryptoPayOrderService {
	srv.Request.CallBackUrl = callBackUrl
	return srv
}

func (srv *CreateCryptoPayOrderService) SetTimeStamp(timeStamp string) *CreateCryptoPayOrderService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

func (srv *CreateCryptoPayOrderService) setSign(sign string) *CreateCryptoPayOrderService {
	srv.Request.Sign = sign
	return srv
}

func (srv *CreateCryptoPayOrderService) Execute() (*CreateCryptoPayOrderResponse, error) {

	url := "CreateCryptoPayOrder"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	var resp CreateCryptoPayOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
