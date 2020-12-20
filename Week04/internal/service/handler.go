package service

import (
	"net/http"
	"strconv"

	xerrors "github.com/pkg/errors"
	
	"goTraining/Week04/internal/biz"

	"github.com/gorilla/mux"
)

type Handler interface {
	Save(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type HandlerImp struct {
	biz.StringService
}

var _ Handler = (*HandlerImp)(nil)

func NewHandler(service biz.StringService) Handler {
	return &HandlerImp{service}
}

func (h *HandlerImp) Save(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	save, err := h.StringService.Save(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if xerrors.Is(err, biz.ErrEmpty) {
			w.Write([]byte("no data"))
		}
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(save)))
}

func (h *HandlerImp) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	amount, err := h.StringService.GetAmount(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if xerrors.Is(err, biz.ErrEmpty) {
			w.Write([]byte("no data"))
		}
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(amount)))
}
