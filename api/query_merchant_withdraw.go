package api

import "encoding/json"

// QueryMerchantWithdrawService type
type QueryMerchantWithdrawService struct {
	IsProdEnvironment bool
	SignKey           string
	AccessKey		  string
	Request           QueryMerchantWithdrawRequest
}

// QueryMerchantWithdrawRequest type
type QueryMerchantWithdrawRequest struct {
	MerchantID         string `json:"MerchantId"`
	WithdrawID         string `json:"WithdrawId,omitempty"`
	MerchantWithdrawID string `json:"MerchantWithdrawId"`
	TimeStamp          string `json:"TimeStamp"`
	Sign               string `json:"Sign"`
}

// QueryMerchantWithdrawResponse type
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

// NewQueryMerchantWithdrawService returns a new service for the given merchant ID and sign key
func NewQueryMerchantWithdrawService(merchantID string, signKey string, accessKey string) *QueryMerchantWithdrawService {
	return &QueryMerchantWithdrawService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		AccessKey:         accessKey,
		Request: QueryMerchantWithdrawRequest{
			MerchantID:         merchantID,
			WithdrawID:         "",
			MerchantWithdrawID: "",
			TimeStamp:          "",
			Sign:               "",
		},
	}
}

// SetIsProdEnvironment sets the environment (test or prod)
func (srv *QueryMerchantWithdrawService) SetIsProdEnvironment(IsProdEnvironment bool) *QueryMerchantWithdrawService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

// SetWithdrawID sets WithdrawID
func (srv *QueryMerchantWithdrawService) SetWithdrawID(withdrawID string) *QueryMerchantWithdrawService {
	srv.Request.WithdrawID = withdrawID
	return srv
}

// SetMerchantWithdrawID sets MerchantWithdrawID
func (srv *QueryMerchantWithdrawService) SetMerchantWithdrawID(merchantWithdrawID string) *QueryMerchantWithdrawService {
	srv.Request.MerchantWithdrawID = merchantWithdrawID
	return srv
}

// SetTimeStamp sets Timestamp
func (srv *QueryMerchantWithdrawService) SetTimeStamp(timeStamp string) *QueryMerchantWithdrawService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

// setSign sets sign key
func (srv *QueryMerchantWithdrawService) setSign(sign string) *QueryMerchantWithdrawService {
	srv.Request.Sign = sign
	return srv
}

// Execute build and send request(*QueryMerchantWithdrawService) return *QueryMerchantWithdrawResponse
func (srv *QueryMerchantWithdrawService) Execute() (*QueryMerchantWithdrawResponse, error) {

	url := "QueryMerchantWithdraw"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.AccessKey, srv.IsProdEnvironment, url)
	if err != nil {
		return nil, err
	}
	var resp QueryMerchantWithdrawResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
