package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Environment URL
const (
	TestAPIBaseURL = "https://stage-api.bitwin.ai/api/v3"
	ProdAPIBaseURL = "https://api.bitwin.ai/api/v3"
)

// Call server method
func callServer(request interface{}, accessKey string, IsProdEnvironment bool, url string) ([]byte, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	baseURL := TestAPIBaseURL
	if IsProdEnvironment {
		baseURL = ProdAPIBaseURL
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", baseURL, url), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Bitwin-Access", accessKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get sign key
func getSign(data interface{}, appSecret string) string {
	switch v := reflect.ValueOf(data); v.Kind() {
	case reflect.String:
		return getMd5String(v.String() + appSecret)
	case reflect.Struct:
		orderStr := structToMapSing(v.Interface(), appSecret)
		return getMd5String(orderStr)
	case reflect.Ptr:
		originType := v.Elem().Type()
		if originType.Kind() != reflect.Struct {
			return ""
		}
		dataType := reflect.TypeOf(data).Elem()
		dataVal := v.Elem()
		orderStr := buildOrderStr(dataType, dataVal, appSecret)
		return getMd5String(orderStr)
	default:
		return ""
	}
}

// structToMapSing method for "getSign"
func structToMapSing(content interface{}, appSecret string) (returnStr string) {
	t := reflect.TypeOf(content)
	v := reflect.ValueOf(content)
	returnStr = buildOrderStr(t, v, appSecret)
	return returnStr
}

// buildOrderStr method for "getSign"
func buildOrderStr(t reflect.Type, v reflect.Value, appSecret string) (returnStr string) {
	keys := make([]string, 0, t.NumField())
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") == "sign" {
			continue
		}
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
		keys = append(keys, t.Field(i).Tag.Get("json"))
	}
	sort.Sort(sort.StringSlice(keys))
	var buf bytes.Buffer
	for _, k := range keys {
		if data[k] == "" || k == "Sign" {
			continue
		}
		switch vv := data[k].(type) {
		case string:
			buf.WriteString(vv)
		case int:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		case int8:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		case int16:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		case int32:
			buf.WriteString(strconv.FormatInt(int64(vv), 10))
		case int64:
			buf.WriteString(strconv.FormatInt(vv, 10))
		case bool:
			buf.WriteString(strconv.FormatBool(vv))
		default:
			continue
		}

		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
	}

	buf.WriteString(appSecret)
	return buf.String()
}

// getMd5String method for "getSign"
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
