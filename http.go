package httputil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func String(rw http.ResponseWriter, code int, str string) {
	rw.WriteHeader(code)
	rw.Write([]byte(str))
}

func JSON(rw http.ResponseWriter, code int, obj interface{}) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	rw.WriteHeader(code)
	rw.Write(raw)
	return nil
}

func BindJSON(r *http.Request, obj interface{}, dupBody bool) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if dupBody {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}

	return json.Unmarshal(b, obj)
}