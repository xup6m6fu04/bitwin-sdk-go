package api

import "encoding/json"

type QueryMerchantWithdrawService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           QueryMerchantWithdrawRequest
}

type QueryMerchantWithdrawRequest struct {
	MerchantId         string `json:"MerchantId"`
	WithdrawId         string `json:"WithdrawId,omitempty"`
	MerchantWithdrawId string `json:"MerchantWithdrawId"`
	TimeStamp          string `json:"TimeStamp"`
	Sign               string `json:"Sign"`
}

type QueryMerchantWithdrawResponse struct {
	MerchantUserID     string `json:"MerchantUserId,omitempty"`
	UserWallet         string `json:"UserWallet,omitempty"`
	MerchantWithdrawID string `json:"MerchantWithdrawId,omitempty"`
	WithdrawID         string `json:"WithdrawId,omitempty"`
	Symbol             string `json:"Symbol,omitempty"`
	Amount             string `json:"Amount,omitempty"`
	MerchantRMB        string `json:"MerchantRMB,omitempty"`
	ExchangeRMB        string `json:"ExchangeRMB,omitempty"`
	Status             string `json:"Status,omitempty"`
	WithdrawDateTime   int64  `json:"WithdrawDateTime,omitempty"`
	ApprovedDateTime   int64  `json:"ApprovedDateTime,omitempty"`
	ReturnCode         string `json:"ReturnCode"`
	ReturnMessage      string `json:"ReturnMessage"`
	Sign               string `json:"Sign"`
}

func NewQueryMerchantWithdrawService(merchantId string, signKey string) *QueryMerchantWithdrawService {
	return &QueryMerchantWithdrawService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: QueryMerchantWithdrawRequest{
			MerchantId:         merchantId,
			WithdrawId:         "",
			MerchantWithdrawId: "",
			TimeStamp:          "",
			Sign:               "",
		},
	}
}

func (srv *QueryMerchantWithdrawService) SetIsProdEnvironment(IsProdEnvironment bool) *QueryMerchantWithdrawService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

func (srv *QueryMerchantWithdrawService) SetWithdrawId(withdrawId string) *QueryMerchantWithdrawService {
	srv.Request.WithdrawId = withdrawId
	return srv
}

func (srv *QueryMerchantWithdrawService) SetMerchantWithdrawId(merchantWithdrawId string) *QueryMerchantWithdrawService {
	srv.Request.MerchantWithdrawId = merchantWithdrawId
	return srv
}

func (srv *QueryMerchantWithdrawService) SetTimeStamp(timeStamp string) *QueryMerchantWithdrawService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

func (srv *QueryMerchantWithdrawService) setSign(sign string) *QueryMerchantWithdrawService {
	srv.Request.Sign = sign
	return srv
}

func (srv *QueryMerchantWithdrawService) Execute() (*QueryMerchantWithdrawResponse, error) {

	url := "QueryMerchantWithdraw"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	var resp QueryMerchantWithdrawResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
