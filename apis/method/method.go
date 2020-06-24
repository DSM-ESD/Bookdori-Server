package method

import (
	"bookdori-server/apis/api"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	// GetNotSupported 는 해당 Request 메서드를 지원하지 않음을 의미합니다
	GetNotSupported struct{}
	// PostNotSupported 는 해당 Request 메서드를 지원하지 않음을 의미합니다
	PostNotSupported struct{}
	// PutNotSupported 는 해당 Request 메서드를 지원하지 않음을 의미합니다
	PutNotSupported struct{}
	// DeleteNotSupported 는 해당 Request 메서드를 지원하지 않음을 의미합니다
	DeleteNotSupported struct{}
)

// Get 메서드는 해당 Request 메서드를 지원하지 않음을 의미합니다
func (GetNotSupported) Get(w http.ResponseWriter, req *http.Request, ps httprouter.Params) api.Response {
	return api.Response{405, "", nil}
}

// Post 메서드는 해당 Request 메서드를 지원하지 않음을 의미합니다
func (PostNotSupported) Post(w http.ResponseWriter, req *http.Request, ps httprouter.Params) api.Response {
	return api.Response{405, "", nil}
}

// Put 메서드는 해당 Request 메서드를 지원하지 않음을 의미합니다
func (PutNotSupported) Put(w http.ResponseWriter, req *http.Request, ps httprouter.Params) api.Response {
	return api.Response{405, "", nil}
}

// Delete 메서드는 해당 Request 메서드를 지원하지 않음을 의미합니다
func (DeleteNotSupported) Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) api.Response {
	return api.Response{405, "", nil}
}