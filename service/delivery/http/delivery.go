package delivery

import (
	"bd_tp_V2/response"
	serviceUsecase "bd_tp_V2/service/usecase"
	"net/http"
)

type ServiceDelivery struct {
	useCase *serviceUsecase.Usecase
}

func NewServiceDelivery(useCase *serviceUsecase.Usecase) *ServiceDelivery {
	return &ServiceDelivery{
		useCase: useCase,
	}
}

func (d *ServiceDelivery) Clear(w http.ResponseWriter, r *http.Request) {
	err := d.useCase.Clear()
	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError,err)
		return
	}
	response.SendResponse(w, http.StatusOK, nil)
}

func (d *ServiceDelivery) GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := d.useCase.GetStatus()
	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError,err)
		return
	}
	response.SendResponse(w, http.StatusOK, status)
}