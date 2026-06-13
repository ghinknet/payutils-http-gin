package httpgin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	stderrors "errors"

	"github.com/gin-gonic/gin"

	"go.gh.ink/payutils/v3/errors"
)

func TestNewInstance_RejectsUnsupported(t *testing.T) {
	_, err := Driver{}.NewInstance(123)
	if !stderrors.Is(err, errors.ErrUnsupportedInstance) {
		t.Errorf("err = %v, want ErrUnsupportedInstance", err)
	}
}

func TestNewInstance_AcceptsRouterAndRoutesPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	inst, err := Driver{}.NewInstance(r)
	if err != nil {
		t.Fatalf("NewInstance(*gin.Engine) error: %v", err)
	}

	called := false
	inst.Post("/pay/callback", func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success"))
	})

	req := httptest.NewRequest(http.MethodPost, "/pay/callback", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if !called {
		t.Error("registered handler was not invoked")
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", rec.Code)
	}
}
