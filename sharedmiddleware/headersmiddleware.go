package sharedmiddleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HeadersMiddleware struct {
}

func NewHeadersMiddleware() *HeadersMiddleware {
	return &HeadersMiddleware{}
}

func (m *HeadersMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := map[string]interface{}{}
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		body["device"] = r.Header.Get("device")
		body["system"] = r.Header.Get("system")
		body["app_id"] = r.Header.Get("app_id")
		body["packet_name"] = r.Header.Get("packet_name")
		body["ifa"] = r.Header.Get("ifa")
		body["app_version"] = r.Header.Get("app_version")
		body["geo"] = r.Header.Get("geo")
		body["k2_id"] = r.Header.Get("k2_id")

		geo := strings.Split(body["geo"].(string), ",")
		if len(geo) > 1 {
			body["latitude"], _ = strconv.ParseFloat(geo[0], 64)
			body["longitude"], _ = strconv.ParseFloat(geo[1], 64)
		}

		json.Unmarshal(bodyBytes, &body)
		bodyBytes, _ = json.Marshal(body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		next(w, r)
	}
}
