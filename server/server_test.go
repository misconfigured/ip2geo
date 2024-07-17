package server_test

import (
	"github.com/misconfigured/ip2geo/server"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Server", func() {
	Describe("Auth", func() {
		var (
			handler  http.HandlerFunc
			recorder *httptest.ResponseRecorder
			request  server.IPNet
		)

		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				server.GeoHandler(db, w, r)
			})

			//handler = http.HandleFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			//	server.GeoHandler(db, w, r)
			//})

			request = server.IPNet{IP: "4.16.85.242"}
		})

		JustBeforeEach(func() {
			callJSON(handler, "POST", recorder, "localhost:8081/location/resolve", request)
		})

		Context("when valid data is provided", func() {
			It("creates a token", func() {
				spew.Dump("recorder.Body.String()", recorder.Body.String())
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

func callJSON(handler http.HandlerFunc, callType string, recorder *httptest.ResponseRecorder, urlPath string, request interface{}) {
	b, err := json.Marshal(request)
	if err != nil {
		_, _ = fmt.Fprintf(GinkgoWriter, "error when marshalling to JSON")
		Expect(err).ToNot(HaveOccurred())
	}

	req, err := http.NewRequest(callType, urlPath, bytes.NewReader(b))
	if err != nil {
		_, _ = fmt.Fprintf(GinkgoWriter, "error when creating %s request", callType)
		Expect(err).ToNot(HaveOccurred())
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	handler.ServeHTTP(recorder, req)
}
