package service

import (
	"context"
	"errors"
	"fmt"
	gfconfig "gofirst/src/config/webserver"
	gfutil "gofirst/src/util"
	gflog "gofirst/src/util/log"
	gfnet "gofirst/src/util/net"
	"html/template"
	"io"
	"net"
	"net/http"
	"os"
)

// Type defenition to avoid any conflicts
type contextType string

var (
	keyServerAddr contextType = ""
)

func test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	str_message := fmt.Sprintf("got / request %v", ctx.Value(keyServerAddr))
	gflog.Info(str_message)
	io.WriteString(w, "This is my website!\n")
}

type Pagestruct struct {
	PageTitle string
	MovieData []byte
}

func movie(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/config/webserver/html/index.html"))
	address := GetServerAddress()
	tmpl.Execute(w, address)
}

func movie_test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	str_message := fmt.Sprintf("got /movies request %v", ctx.Value(keyServerAddr))
	gflog.Info(str_message)

	// TODO: Have a database of movies and videos
	fd, err := os.Open("/home/asn/small.mp4")

	if err != nil {
		gflog.Error(err.Error())
		io.WriteString(w, "error opening file")
	}
	defer fd.Close()

	// w.Header().Add("Content-Type", "video/mp4")

	buff := make([]byte, gfutil.DEFAULT_BUFFER_SIZE)
	for {
		read, err := fd.Read(buff)
		if err == io.EOF || read == 0 {
			break
		}
		if err != nil {
			w.Header().Del("Content-Type")
			io.WriteString(w, "error writing data")
		}
		w.Write(buff[:read])
	}
}

func getHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", test)
	mux.HandleFunc("/movies", movie)
	mux.HandleFunc("/movie_test", movie_test)
	return mux
}

func startServer(mux *http.ServeMux) error {
	address := GetServerAddress()
	str_msg := fmt.Sprintf("Starting webserver on %s to accept requests...", address)
	gflog.Info(str_msg)

	ctx, cancelCtx := context.WithCancel(context.Background())

	server := &http.Server{
		Addr:    address,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
		ErrorLog: gflog.GetErrorLogger(),
	}

	var err error = nil
	go func(serve *http.Server) {
		err = serve.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			str_message := fmt.Sprintf("server closed; address: %s", serve.Addr)
			gflog.Error(str_message)
		} else if err != nil {
			gflog.Error(err.Error())
		}
		cancelCtx()
	}(server)

	<-ctx.Done()
	return err
}

func GetServerAddress() string {
	server_ip, err := gfnet.PublicIP()
	if err != nil {
		gflog.Error(err.Error())
	}

	address := fmt.Sprintf("%s:%d", server_ip, gfconfig.SERVER_PORT)
	return address
}

func Start() error {
	return startServer(getHandler())
}
