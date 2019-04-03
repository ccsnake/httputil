package httputil

import (
	"bytes"
	"io/ioutil"
	"net/http"
	json "github.com/json-iterator/go"
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

func BizMessage(w http.ResponseWriter, code int, message string) error {
	b, err := json.Marshal(map[string]interface{}{
		"code":    code,
		"message": message,
	})
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return nil
}
