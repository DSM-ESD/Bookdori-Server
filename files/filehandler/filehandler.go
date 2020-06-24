package filehandler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// FileHandler 인터페이스는 파일 핸들러에서 구현해야할 메서드를 가집니다
type FileHandler interface {
	URI() string
	Handler(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
}