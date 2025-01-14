package apiutils_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/apiutils"
	"github.com/Goldziher/go-monorepo/lib/testutils"
	"github.com/go-chi/render"
	"github.com/stretchr/testify/assert"
)

func TestApiError(t *testing.T) {
	for _, testCase := range []struct {
		ApiErrType     render.Renderer
		ExpectedStatus int
	}{
		{
			ApiErrType:     apiutils.BadRequest("err"),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			ApiErrType:     apiutils.Unauthorized("err"),
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			ApiErrType:     apiutils.UnprocessableContent("err"),
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			ApiErrType:     apiutils.InternalServerError("err"),
			ExpectedStatus: http.StatusInternalServerError,
		},
		{
			ApiErrType:     apiutils.ServiceUnavailable("err"),
			ExpectedStatus: http.StatusServiceUnavailable,
		},
	} {
		client := testutils.CreateTestClient(t, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			_ = render.Render(writer, request, testCase.ApiErrType)
		}))
		res, err := client.Get(context.TODO(), "/")
		assert.Nil(t, err)
		assert.Equal(t, res.StatusCode, testCase.ExpectedStatus)
	}
}
