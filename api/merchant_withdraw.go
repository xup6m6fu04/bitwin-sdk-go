package api

import "encoding/json"

type MerchantWithdrawService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           MerchantWithdrawRequest
}

type MerchantWithdrawRequest struct {
	MerchantId         string `json:"MerchantId"`
	MerchantUserId     string `json:"MerchantUserId,omitempty"`
	MerchantWithdrawId string `json:"MerchantWithdrawId"`
	UserWallet         string `json:"UserWallet"`
	Symbol             string `json:"Symbol"`
	Amount             string `json:"Amount"`
	MerchantRMB        string `json:"MerchantRMB,omitempty"`
	CallBackUrl        string `json:"CallBackUrl"`
	TimeStamp          string `json:"TimeStamp"`
	Sign               string `json:"Sign"`
}

type MerchantWithdrawResponse struct {
	WithdrawID         string `json:"WithdrawId,omitempty"`
	MerchantWithdrawID string `json:"MerchantWithdrawId,omitempty"`
	ReturnCode         string `json:"ReturnCode"`
	ReturnMessage      string `json:"ReturnMessage"`
	Sign               string `json:"Sign"`
}

func NewMerchantWithdrawService(merchantId string, signKey string) *MerchantWithdrawService {
	return &MerchantWithdrawService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: MerchantWithdrawRequest{
			MerchantId:         merchantId,
			MerchantUserId:     "",
			MerchantWithdrawId: "",
			UserWallet:         "",
			Symbol:             "",
			Amount:             "",
			MerchantRMB:        "",
			CallBackUrl:        "",
			TimeStamp:          "",
			Sign:               "",
		},
	}
}

func (srv *MerchantWithdrawService) SetIsProdEnvironment(IsProdEnvironment bool) *MerchantWithdrawService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

func (srv *MerchantWithdrawService) SetMerchantUserId(merchantUserId string) *MerchantWithdrawService {
	srv.Request.MerchantUserId = merchantUserId
	return srv
}

func (srv *MerchantWithdrawService) SetMerchantWithdrawId(merchantWithdrawId string) *MerchantWithdrawService {
	srv.Request.MerchantWithdrawId = merchantWithdrawId
	return srv
}

func (srv *MerchantWithdrawService) SetUserWallet(userWallet string) *MerchantWithdrawService {
	srv.Request.UserWallet = userWallet
	return srv
}

func (srv *MerchantWithdrawService) SetAmount(amount string) *MerchantWithdrawService {
	srv.Request.Amount = amount
	return srv
}

func (srv *MerchantWithdrawService) SetMerchantRMB(merchantRMB string) *MerchantWithdrawService {
	srv.Request.MerchantRMB = merchantRMB
	return srv
}

func (srv *MerchantWithdrawService) SetSymbol(symbol string) *MerchantWithdrawService {
	srv.Request.Symbol = symbol
	return srv
}

func (srv *MerchantWithdrawService) SetCallBackUrl(callBackUrl string) *MerchantWithdrawService {
	srv.Request.CallBackUrl = callBackUrl
	return srv
}

func (srv *MerchantWithdrawService) SetTimeStamp(timeStamp string) *MerchantWithdrawService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

func (srv *MerchantWithdrawService) setSign(sign string) *MerchantWithdrawService {
	srv.Request.Sign = sign
	return srv
}

func (srv *MerchantWithdrawService) Execute() (*MerchantWithdrawResponse, error) {

	url := "MerchantWithdraw"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	var resp MerchantWithdrawResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
