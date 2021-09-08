package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("befor excecute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("after excecute handler")

}

type ErrorHandler struct {
	Handler http.Handler
}

func (e ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "error : %s", err)
		}
	}()
	e.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("handler execute")
		fmt.Fprint(writer, "hello middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("panic")
		panic("ups")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo execute")
		fmt.Fprint(writer, "hello goo")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr: "Localhost:8080",
		Handler: errorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}