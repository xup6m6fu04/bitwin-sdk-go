package api

import (
	"encoding/json"
)

// QueryCryptoPayOrderService type
type QueryCryptoPayOrderService struct {
	IsProdEnvironment bool
	SignKey           string
	AccessKey		  string
	Request           QueryCryptoPayOrderRequest
}

// QueryCryptoPayOrderRequest type
type QueryCryptoPayOrderRequest struct {
	MerchantID      string `json:"MerchantId"`
	MerchantOrderID string `json:"MerchantOrderId,omitempty"`
	OrderID         string `json:"OrderId,omitempty"`
	TimeStamp       string `json:"TimeStamp"`
	Sign            string `json:"Sign"`
}

// QueryCryptoPayOrderResponse type
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

// NewQueryCryptoPayOrderService returns a new service for the given merchant ID and sign key
func NewQueryCryptoPayOrderService(merchantID string, signKey string, accessKey string) *QueryCryptoPayOrderService {
	return &QueryCryptoPayOrderService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		AccessKey:         accessKey,
		Request: QueryCryptoPayOrderRequest{
			MerchantID:      merchantID,
			MerchantOrderID: "",
			OrderID:         "",
			TimeStamp:       "",
			Sign:            "",
		},
	}
}

// SetIsProdEnvironment sets the environment (test or prod)
func (srv *QueryCryptoPayOrderService) SetIsProdEnvironment(IsProdEnvironment bool) *QueryCryptoPayOrderService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

// SetMerchantOrderID sets MerchantOrderID
func (srv *QueryCryptoPayOrderService) SetMerchantOrderID(merchantOrderID string) *QueryCryptoPayOrderService {
	srv.Request.MerchantOrderID = merchantOrderID
	return srv
}

// SetOrderID sets OrderID
func (srv *QueryCryptoPayOrderService) SetOrderID(orderID string) *QueryCryptoPayOrderService {
	srv.Request.OrderID = orderID
	return srv
}

// SetTimeStamp sets Timestamp
func (srv *QueryCryptoPayOrderService) SetTimeStamp(timeStamp string) *QueryCryptoPayOrderService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

// setSign sets sign key
func (srv *QueryCryptoPayOrderService) setSign(sign string) *QueryCryptoPayOrderService {
	srv.Request.Sign = sign
	return srv
}

// Execute build and send request(*QueryCryptoPayOrderService) return *QueryCryptoPayOrderResponse
func (srv *QueryCryptoPayOrderService) Execute() (*QueryCryptoPayOrderResponse, error) {

	url := "QueryCryptoPayOrder"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.AccessKey, srv.IsProdEnvironment, url)
	if err != nil {
		return nil, err
	}
	var resp QueryCryptoPayOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
