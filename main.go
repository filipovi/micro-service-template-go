package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"filipovi/go-micro-template/api/mrapi"

	"github.com/urfave/negroni"
	"goji.io"
	"goji.io/pat"
)

// Env is the container
type Env struct {
	client mrapi.Client
	apiURL string
}

func send(content []byte, contentType string, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(content)))
	w.WriteHeader(status)
	w.Write(content)
}

func (env *Env) handleIndexRequest(w http.ResponseWriter, r *http.Request) {
	send([]byte("{}"), "application/json", http.StatusOK, w)
}

func (env *Env) handleRoleRequest(w http.ResponseWriter, r *http.Request) {
	response, err := env.client.Get(env.apiURL + "/roles")
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	send([]byte(string(body)), "application/json", http.StatusOK, w)
}

func initialize() (*Env, error) {
	client, err := mrapi.New(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("CLIENT_URL"),
	)
	if nil != err {
		return nil, err
	}
	log.Println("MR API configured!")

	return &Env{client: *client, apiURL: os.Getenv("CLIENT_URL")}, nil
}

func main() {
	env, err := initialize()
	if nil != err {
		log.Fatalf("FATAL ERROR: %s", err)
	}

	n := negroni.Classic()

	// Routing
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), env.handleIndexRequest)
	mux.HandleFunc(pat.Get("/roles"), env.handleRoleRequest)
	n.UseHandler(mux)

	// Launch the Web Server
	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	srv := &http.Server{
		Handler:      n,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server run on http://" + addr)
	log.Fatal(srv.ListenAndServe())
}
