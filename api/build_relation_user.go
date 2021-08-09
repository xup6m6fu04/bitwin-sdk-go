package api

import "encoding/json"

// BuildRelationUserService type
type BuildRelationUserService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           BuildRelationUserRequest
}

// BuildRelationUserRequest type
type BuildRelationUserRequest struct {
	MerchantID     string `json:"MerchantId"`
	MerchantUserID string `json:"MerchantUserId"`
	CallBackURL    string `json:"CallBackUrl"`
	TimeStamp      string `json:"TimeStamp"`
	Sign           string `json:"Sign"`
}

// BuildRelationUserResponse type
type BuildRelationUserResponse struct {
	QrcodeData     string `json:"QrcodeData,omitempty"`
	QrcodeImageURL string `json:"QrcodeImageUrl,omitempty"`
	ReturnCode     string `json:"ReturnCode"`
	ReturnMessage  string `json:"ReturnMessage"`
	Sign           string `json:"Sign"`
}

// NewBuildRelationUserService returns a new service for the given merchant ID and sign key
func NewBuildRelationUserService(merchantID string, signKey string) *BuildRelationUserService {
	return &BuildRelationUserService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: BuildRelationUserRequest{
			MerchantID:     merchantID,
			MerchantUserID: "",
			CallBackURL:    "",
			TimeStamp:      "",
			Sign:           "",
		},
	}
}

// SetIsProdEnvironment sets the environment (test or prod)
func (srv *BuildRelationUserService) SetIsProdEnvironment(IsProdEnvironment bool) *BuildRelationUserService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

// SetMerchantUserID sets MerchantUserId
func (srv *BuildRelationUserService) SetMerchantUserID(merchantUserID string) *BuildRelationUserService {
	srv.Request.MerchantUserID = merchantUserID
	return srv
}

// SetCallBackURL sets CallBackURL
func (srv *BuildRelationUserService) SetCallBackURL(callBackURL string) *BuildRelationUserService {
	srv.Request.CallBackURL = callBackURL
	return srv
}

// SetTimeStamp sets Timestamp
func (srv *BuildRelationUserService) SetTimeStamp(timeStamp string) *BuildRelationUserService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

// setSign sets sign key
func (srv *BuildRelationUserService) setSign(sign string) *BuildRelationUserService {
	srv.Request.Sign = sign
	return srv
}

// Execute build and send request(*BuildRelationUserService) return *BuildRelationUserResponse
func (srv *BuildRelationUserService) Execute() (*BuildRelationUserResponse, error) {

	url := "BuildRelationUser"

	// Set Sign Column
	srv.setSign(getSign(srv.Request, srv.SignKey))

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	var resp BuildRelationUserResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
