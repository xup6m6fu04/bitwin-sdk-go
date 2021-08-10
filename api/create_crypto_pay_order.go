package api

import (
	"encoding/json"
)

// CreateCryptoPayOrderService type
type CreateCryptoPayOrderService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           CreateCryptoPayOrderRequest
}

// CreateCryptoPayOrderRequest type
type CreateCryptoPayOrderRequest struct {
	MerchantID       string `json:"MerchantId"`
	MerchantUserID   string `json:"MerchantUserId,omitempty"`
	MerchantOrderID  string `json:"MerchantOrderId"`
	OrderDescription string `json:"OrderDescription,omitempty"`
	Amount           string `json:"Amount"`
	MerchantRMB      string `json:"MerchantRMB,omitempty"`
	Symbol           string `json:"Symbol"`
	CallBackURL      string `json:"CallBackUrl"`
	TimeStamp        string `json:"TimeStamp"`
	Sign             string `json:"Sign"`
}

// CreateCryptoPayOrderResponse type
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

// NewCreateCryptoPayOrderService returns a new service for the given merchant ID and sign key
func NewCreateCryptoPayOrderService(merchantID string, signKey string) *CreateCryptoPayOrderService {
	return &CreateCryptoPayOrderService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: CreateCryptoPayOrderRequest{
			MerchantID:       merchantID,
			MerchantUserID:   "",
			MerchantOrderID:  "",
			OrderDescription: "",
			Amount:           "",
			MerchantRMB:      "",
			Symbol:           "",
			CallBackURL:      "",
			TimeStamp:        "",
			Sign:             "",
		},
	}
}

// SetIsProdEnvironment sets the environment (test or prod)
func (srv *CreateCryptoPayOrderService) SetIsProdEnvironment(isProdEnvironment bool) *CreateCryptoPayOrderService {
	srv.IsProdEnvironment = isProdEnvironment
	return srv
}

// SetMerchantUserID sets MerchantUserID
func (srv *CreateCryptoPayOrderService) SetMerchantUserID(merchantUserID string) *CreateCryptoPayOrderService {
	srv.Request.MerchantUserID = merchantUserID
	return srv
}

// SetMerchantOrderID sets MerchantOrderID
func (srv *CreateCryptoPayOrderService) SetMerchantOrderID(merchantOrderID string) *CreateCryptoPayOrderService {
	srv.Request.MerchantOrderID = merchantOrderID
	return srv
}

// SetOrderDescription sets OrderDescription
func (srv *CreateCryptoPayOrderService) SetOrderDescription(orderDescription string) *CreateCryptoPayOrderService {
	srv.Request.OrderDescription = orderDescription
	return srv
}

// SetAmount sets Amount
func (srv *CreateCryptoPayOrderService) SetAmount(amount string) *CreateCryptoPayOrderService {
	srv.Request.Amount = amount
	return srv
}

// SetMerchantRMB sets MerchantRMB
func (srv *CreateCryptoPayOrderService) SetMerchantRMB(merchantRMB string) *CreateCryptoPayOrderService {
	srv.Request.MerchantRMB = merchantRMB
	return srv
}

// SetSymbol sets Symbol
func (srv *CreateCryptoPayOrderService) SetSymbol(symbol string) *CreateCryptoPayOrderService {
	srv.Request.Symbol = symbol
	return srv
}

// SetCallBackURL sets CallBackURL
func (srv *CreateCryptoPayOrderService) SetCallBackURL(callBackURL string) *CreateCryptoPayOrderService {
	srv.Request.CallBackURL = callBackURL
	return srv
}

// SetTimeStamp set TimeStamp
func (srv *CreateCryptoPayOrderService) SetTimeStamp(timeStamp string) *CreateCryptoPayOrderService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

// setSign set sign key
func (srv *CreateCryptoPayOrderService) setSign(sign string) *CreateCryptoPayOrderService {
	srv.Request.Sign = sign
	return srv
}

// Execute build and send request(*CreateCryptoPayOrderService) return *CreateCryptoPayOrderResponse
func (srv *CreateCryptoPayOrderService) Execute() (*CreateCryptoPayOrderResponse, error) {

	url := "CreateCryptoPayOrder"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	if err != nil {
		return nil, err
	}
	var resp CreateCryptoPayOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
