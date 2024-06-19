package controller

import "gin-gonic-api/app/service"

type RecordController interface{}

type RecordControllerImpl struct {
	recordService service.RecordService
}

func RecordControllerInit(recordService service.RecordService) *RecordControllerImpl {
	return &RecordControllerImpl{
		recordService: recordService,
	}
}
