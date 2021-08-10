package api

import "encoding/json"

// MerchantWithdrawService type
type MerchantWithdrawService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           MerchantWithdrawRequest
}

// MerchantWithdrawRequest type
type MerchantWithdrawRequest struct {
	MerchantID         string `json:"MerchantId"`
	MerchantUserID     string `json:"MerchantUserId,omitempty"`
	MerchantWithdrawID string `json:"MerchantWithdrawId"`
	UserWallet         string `json:"UserWallet"`
	Symbol             string `json:"Symbol"`
	Amount             string `json:"Amount"`
	MerchantRMB        string `json:"MerchantRMB,omitempty"`
	CallBackURL        string `json:"CallBackUrl"`
	TimeStamp          string `json:"TimeStamp"`
	Sign               string `json:"Sign"`
}

// MerchantWithdrawResponse type
type MerchantWithdrawResponse struct {
	WithdrawID         string `json:"WithdrawId,omitempty"`
	MerchantWithdrawID string `json:"MerchantWithdrawId,omitempty"`
	ReturnCode         string `json:"ReturnCode"`
	ReturnMessage      string `json:"ReturnMessage"`
	Sign               string `json:"Sign"`
}

// NewMerchantWithdrawService returns a new service for the given merchant ID and sign key
func NewMerchantWithdrawService(merchantID string, signKey string) *MerchantWithdrawService {
	return &MerchantWithdrawService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: MerchantWithdrawRequest{
			MerchantID:         merchantID,
			MerchantUserID:     "",
			MerchantWithdrawID: "",
			UserWallet:         "",
			Symbol:             "",
			Amount:             "",
			MerchantRMB:        "",
			CallBackURL:        "",
			TimeStamp:          "",
			Sign:               "",
		},
	}
}

// SetIsProdEnvironment sets the environment (test or prod)
func (srv *MerchantWithdrawService) SetIsProdEnvironment(IsProdEnvironment bool) *MerchantWithdrawService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

// SetMerchantUserID sets MerchantUserID
func (srv *MerchantWithdrawService) SetMerchantUserID(merchantUserID string) *MerchantWithdrawService {
	srv.Request.MerchantUserID = merchantUserID
	return srv
}

// SetMerchantWithdrawID sets MerchantWithdrawID
func (srv *MerchantWithdrawService) SetMerchantWithdrawID(merchantWithdrawID string) *MerchantWithdrawService {
	srv.Request.MerchantWithdrawID = merchantWithdrawID
	return srv
}

// SetUserWallet sets UserWallet
func (srv *MerchantWithdrawService) SetUserWallet(userWallet string) *MerchantWithdrawService {
	srv.Request.UserWallet = userWallet
	return srv
}

// SetAmount sets Amount
func (srv *MerchantWithdrawService) SetAmount(amount string) *MerchantWithdrawService {
	srv.Request.Amount = amount
	return srv
}

// SetMerchantRMB sets MerchantRMB
func (srv *MerchantWithdrawService) SetMerchantRMB(merchantRMB string) *MerchantWithdrawService {
	srv.Request.MerchantRMB = merchantRMB
	return srv
}

// SetSymbol sets Symbol
func (srv *MerchantWithdrawService) SetSymbol(symbol string) *MerchantWithdrawService {
	srv.Request.Symbol = symbol
	return srv
}

// SetCallBackURL sets CallBackURL
func (srv *MerchantWithdrawService) SetCallBackURL(callBackURL string) *MerchantWithdrawService {
	srv.Request.CallBackURL = callBackURL
	return srv
}

// SetTimeStamp sets Timestamp
func (srv *MerchantWithdrawService) SetTimeStamp(timeStamp string) *MerchantWithdrawService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

// setSign sets sign key
func (srv *MerchantWithdrawService) setSign(sign string) *MerchantWithdrawService {
	srv.Request.Sign = sign
	return srv
}

// Execute build and send request(*MerchantWithdrawService) return *MerchantWithdrawResponse
func (srv *MerchantWithdrawService) Execute() (*MerchantWithdrawResponse, error) {

	url := "MerchantWithdraw"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	if err != nil {
		return nil, err
	}
	var resp MerchantWithdrawResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
