package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// API 인터페이스는 API에서 구현해야할 메서드를 가집니다
type API interface {
	URI() string
	Get(w http.ResponseWriter, req *http.Request, ps httprouter.Params) Response
	Post(w http.ResponseWriter, req *http.Request, ps httprouter.Params) Response
	Put(w http.ResponseWriter, req *http.Request, ps httprouter.Params) Response
	Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) Response
}

// Response 객체는 API에서 응답하는 Response를 정의합니다
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HTTPResponse 함수는 Response 구조체를 HTTP로 응답합니다
func HTTPResponse(w http.ResponseWriter, req *http.Request, res Response) {
	content, err := json.Marshal(res)

	if err != nil {
		abort(w, 500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(content)
}

func abort(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}