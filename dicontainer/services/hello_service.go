package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UpperHandler struct {
	UpperUsecase UpperUsecaseInterface
}

func NewUpperHandler(u UpperUsecaseInterface) *UpperHandler {
	return &UpperHandler{
		UpperUsecase: u,
	}
}

func (h UpperHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := &UpperRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	h.UpperUsecase.Do(w, req)
}

type UpperRequest struct {
	Name string `json:"name"`
}

func (req *UpperRequest) GetName() string {
	return req.Name
}

type UpperRequestInterface interface {
	GetName() string
}

type UpperPresenter struct {
	Name string `json:"name"`
}

func NewUpperPresenter() UpperPresenterInterface {
	return &UpperPresenter{}
}

func (p *UpperPresenter) Output(w io.Writer, res UpperResponseInterface) {
	p.Name = res.GetName()
	json.NewEncoder(w).Encode(p)
}

type UpperUsecaseInterface interface {
	Do(io.Writer, UpperRequestInterface)
}

type UpperPresenterInterface interface {
	Output(w io.Writer, res UpperResponseInterface)
}

type UpperResponseInterface interface {
	GetName() string
}

type UpperResponse struct {
	Name string
}

func (r *UpperResponse) GetName() string {
	return r.Name
}

type UpperUsecase struct {
	UpperPresenter UpperPresenterInterface
}

func NewUpperUsecase(p UpperPresenterInterface) UpperUsecaseInterface {
	return &UpperUsecase{
		UpperPresenter: p,
	}
}

func (u *UpperUsecase) Do(w io.Writer, req UpperRequestInterface) {
	name := req.GetName()
	upperName := strings.ToUpper(name)
	res := &UpperResponse{Name: upperName}
	u.UpperPresenter.Output(w, res)
}
