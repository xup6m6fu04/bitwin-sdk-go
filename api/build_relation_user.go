package api

import "encoding/json"

type BuildRelationUserService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           BuildRelationUserRequest
}

type BuildRelationUserRequest struct {
	MerchantId     string `json:"MerchantId"`
	MerchantUserId string `json:"MerchantUserId"`
	CallBackUrl    string `json:"CallBackUrl"`
	TimeStamp      string `json:"TimeStamp"`
	Sign           string `json:"Sign"`
}

type BuildRelationUserResponse struct {
	QrcodeData     string `json:"QrcodeData,omitempty"`
	QrcodeImageUrl string `json:"QrcodeImageUrl,omitempty"`
	ReturnCode     string `json:"ReturnCode"`
	ReturnMessage  string `json:"ReturnMessage"`
	Sign           string `json:"Sign"`
}

func NewBuildRelationUserService(merchantId string, signKey string) *BuildRelationUserService {
	return &BuildRelationUserService{
		IsProdEnvironment: false,
		SignKey:           signKey,
		Request: BuildRelationUserRequest{
			MerchantId:     merchantId,
			MerchantUserId: "",
			CallBackUrl:    "",
			TimeStamp:      "",
			Sign:           "",
		},
	}
}

func (srv *BuildRelationUserService) SetIsProdEnvironment(IsProdEnvironment bool) *BuildRelationUserService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

func (srv *BuildRelationUserService) SetMerchantUserId(merchantUserId string) *BuildRelationUserService {
	srv.Request.MerchantUserId = merchantUserId
	return srv
}

func (srv *BuildRelationUserService) SetCallBackUrl(callBackUrl string) *BuildRelationUserService {
	srv.Request.CallBackUrl = callBackUrl
	return srv
}

func (srv *BuildRelationUserService) SetTimeStamp(timeStamp string) *BuildRelationUserService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

func (srv *BuildRelationUserService) setSign(sign string) *BuildRelationUserService {
	srv.Request.Sign = sign
	return srv
}

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
