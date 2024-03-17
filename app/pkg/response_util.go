package pkg

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data, responseStatus.GetResponseCode())
}

func BuildResponse_[T any](status string, message string, data T, code int) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
		Code:            code,
	}
}
