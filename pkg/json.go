package pkg

import (
	log "github.com/sirupsen/logrus"
	"reflect"
)

const (
	CodeOk       = 0  // 成功
	CodeErr      = -1 // 失败
	CodeErrToken = -2 // token相关的异常
)

// JSONResponse represents an HTTP response which contains a JSON body.
type JSONResponse struct {
	// HTTP status code.
	Code int `json:"code"`
	// JSON represents the JSON that should be serialized and sent to the client
	Data interface{} `json:"data"`
}

type JSONQueryResponse struct {
	// HTTP status code.
	Code int `json:"code"`
	// JSON represents the JSON that should be serialized and sent to the client
	Data      interface{} `json:"data"`
	QueryType interface{} `json:"query_type"`
}

func SuccessResponse(data interface{}) JSONResponse {
	return JSONResponse{
		Code: 0,
		Data: data,
	}
}

func QueryResponse(data interface{}, queryType interface{}) JSONQueryResponse {
	return JSONQueryResponse{
		Code:      0,
		Data:      data,
		QueryType: queryType,
	}
}

// MessageResponse returns a JSONResponse with a 'message' key containing the given text.
func MessageResponse(code int, msg, msgZh string) JSONResponse {
	log.WithFields(log.Fields{
		"code":   code,
		"msg_zh": msgZh,
	}).Warnf(msg)
	return JSONResponse{
		Code: code,
		Data: struct {
			Message   string `json:"message"`
			MessageZh string `json:"message_zh"`
		}{msg, msgZh},
	}
}

func PrintStruct(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldValue := val.Field(i)
			log.Debugf("%s: %v\n", field.Name, fieldValue.Interface())
		}
	} else if val.Kind() == reflect.Ptr {
		val = val.Elem()
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldValue := val.Field(i)
			log.Debugf("%s: %v\n", field.Name, fieldValue.Interface())
		}
	} else {
		log.Warningln("[PrintStruct] Not a struct or pointer.")
	}
}
