package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (i Instance) Get(path string, handler http.HandlerFunc) {
	i.Router.GET(path, gin.WrapH(handler))
}

func (i Instance) Post(path string, handler http.HandlerFunc) {
	i.Router.POST(path, gin.WrapH(handler))
}

func (i Instance) Put(path string, handler http.HandlerFunc) {
	i.Router.PUT(path, gin.WrapH(handler))
}

func (i Instance) Patch(path string, handler http.HandlerFunc) {
	i.Router.PATCH(path, gin.WrapH(handler))
}

func (i Instance) Delete(path string, handler http.HandlerFunc) {
	i.Router.DELETE(path, gin.WrapH(handler))
}

func (i Instance) Head(path string, handler http.HandlerFunc) {
	i.Router.HEAD(path, gin.WrapH(handler))
}

func (i Instance) Options(path string, handler http.HandlerFunc) {
	i.Router.OPTIONS(path, gin.WrapH(handler))
}

func (i Instance) Any(path string, handler http.HandlerFunc) {
	i.Router.Any(path, gin.WrapH(handler))
}
