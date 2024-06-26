package share

import (
	"fmt"
	"net/http"

	"github.com/aeazer/dirserver/utils/color"
)

type requestLog struct{}

func (*requestLog) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := color.RedDA.Dyeing(r.RemoteAddr)
		method := color.GreenDA.Dyeing(r.Method)
		path := color.BlueDA.Dyeing(r.URL.Path)
		ua := color.YellowDA.Dyeing(r.UserAgent())
		fmt.Printf("Received request: Remote Address: %s, Method: %s, Path: %s, User-Agent: %s\n",
			addr, method, path, ua)
		next.ServeHTTP(w, r)
	})
}
