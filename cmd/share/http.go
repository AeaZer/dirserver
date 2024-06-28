package share

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aeazer/dirserver/utils/color"
)

func doWriteJson(w http.ResponseWriter, code int, v any) {
	switch t := v.(type) {
	case error:
		fmt.Println(color.RedDA.Dyeing(t.Error()))
	}
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("marshal json failed, error: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if n, err := w.Write(bs); err != nil {
		if !errors.Is(err, http.ErrHandlerTimeout) {
			fmt.Printf("write response failed, error: %v\n", err)
			return
		}
	} else if n < len(bs) {
		fmt.Printf("actual bytes: %d, written bytes: %d\n", len(bs), n)
	}
}
