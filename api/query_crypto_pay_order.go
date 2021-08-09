package api

import (
	"encoding/json"
)

type QueryCryptoPayOrderService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           QueryCryptoPayOrderRequest
}

type QueryCryptoPayOrderRequest struct {
	MerchantId      string `json:"MerchantId"`
	MerchantOrderId string `json:"MerchantOrderId,omitempty"`
	OrderId         string `json:"OrderId,omitempty"`
	TimeStamp       string `json:"TimeStamp"`
	Sign            string `json:"Sign"`
}

type QueryCryptoPayOrderResponse struct {
	OrderID          string `json:"OrderId,omitempty"`
	MerchantOrderID  string `json:"MerchantOrderId,omitempty"`
	MerchantUserID   string `json:"MerchantUserId,omitempty"`
	OrderDescription string `json:"OrderDescription,omitempty"`
	Symbol           string `json:"Symbol,omitempty"`
	Amount           string `json:"Amount,omitempty"`
	RealAmount       string `json:"RealAmount,omitempty"`
	MerchantRMB      string `json:"MerchantRMB,omitempty"`
	ExchangeRMB      string `json:"ExchangeRMB,omitempty"`
	OrderStatus      string `json:"OrderStatus,omitempty"`
	CallBackURL      string `json:"CallBackUrl,omitempty"`
	ReturnCode       string `json:"ReturnCode"`
	ReturnMessage    string `json:"ReturnMessage"`
	Sign             string `json:"Sign"`
}

func NewQueryCryptoPayOrderService(merchantId string, signKey string) *QueryCryptoPayOrderService {
	return &QueryCryptoPayOrderService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: QueryCryptoPayOrderRequest{
			MerchantId:      merchantId,
			MerchantOrderId: "",
			OrderId:         "",
			TimeStamp:       "",
			Sign:            "",
		},
	}
}

func (srv *QueryCryptoPayOrderService) SetIsProdEnvironment(IsProdEnvironment bool) *QueryCryptoPayOrderService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

func (srv *QueryCryptoPayOrderService) SetMerchantOrderId(merchantOrderId string) *QueryCryptoPayOrderService {
	srv.Request.MerchantOrderId = merchantOrderId
	return srv
}

func (srv *QueryCryptoPayOrderService) SetOrderId(orderId string) *QueryCryptoPayOrderService {
	srv.Request.OrderId = orderId
	return srv
}

func (srv *QueryCryptoPayOrderService) SetTimeStamp(timeStamp string) *QueryCryptoPayOrderService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

func (srv *QueryCryptoPayOrderService) setSign(sign string) *QueryCryptoPayOrderService {
	srv.Request.Sign = sign
	return srv
}

func (srv *QueryCryptoPayOrderService) Execute() (*QueryCryptoPayOrderResponse, error) {

	url := "QueryCryptoPayOrder"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	var resp QueryCryptoPayOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
