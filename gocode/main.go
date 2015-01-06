package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	webStaticRoot = "web/app/"
)

type NameValue struct {
	Name, Value string
}

var cfg *config

func main() {
	var err error
	cfg, err = NewConfig(configFilePath)
	if err != nil {
		fmt.Println("main.go:main(): Error reading config file %v: %v", configFilePath, err)
		return
	}

	r := mux.NewRouter()

	servedirs := []string{"bower_components", "components", "css", "filters", "sysinfo", "env", "files"}
	for _, dir := range servedirs {
		r.PathPrefix("/" + dir + "/").Handler(http.StripPrefix("/"+dir+"/", http.FileServer(http.Dir(webStaticRoot+dir))))
	}

	r.HandleFunc("/app.js", func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, webStaticRoot+"app.js")
	})

	r.HandleFunc("/s/sysinfo", sysinfo).Methods("GET")
	r.HandleFunc("/s/env", env).Methods("GET")
	r.HandleFunc("/s/files", fileDetails).Methods("GET")
	r.HandleFunc("/s/files/file", downloadFile).Methods("GET")
	r.HandleFunc("/s/mem", memory).Methods("GET")
	r.HandleFunc("/ws/files/tail", tailFile).Methods("GET")

	r.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, webStaticRoot+"index.html")
	})

	port := 8083
	if cfg.Port != 0 {
		port = cfg.Port
	}

	memInfoSrvc(memDuration)

	fmt.Printf("main.go:main(): Listening at %v\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), r)

}

func sysinfo(rw http.ResponseWriter, req *http.Request) {
	s := []NameValue{NameValue{"CPU", "Intel"}}
	// s := "[]"
	// fmt.Println("main.go:sysinfo(): returning string: ", s)
	writeJSON(rw, http.StatusOK, s)
}

func env(rw http.ResponseWriter, req *http.Request) {
	s := os.Environ()
	writeJSON(rw, http.StatusOK, s)
}

func scanLocalFiles() ([]os.FileInfo, error) {
	return nil, nil
}

func fileDetails(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("main.go:fileDetails(): Starting to look for file details.\n")
	ifDetails, err := getHighLevelFileDetails(cfg)
	if err != nil {
		fmt.Printf("main.go:fileDetails(): Error getting high level file details: %v\n", err)
		return
	}

	writeJSON(rw, http.StatusOK, ifDetails)
}

func memory(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("main.go:memory(): Getting memory.\n")
	writeJSON(rw, http.StatusOK, memHistory)
}

func downloadFile(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Query().Get("path")
	fmt.Printf("main.go:downloadFile(): Request to download file: %v\n", path)
	fi, err := os.Stat(path)
	if err != nil {
		writeJSON(rw, http.StatusInternalServerError, err.Error())
		return
	}
	rw.Header().Set("Content-Disposition", "attachment; filename="+fi.Name())
	http.ServeFile(rw, req, path)
}

func tailFile(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Query().Get("path")
	fmt.Printf("main.go:tailFile(): Request to download file: %v\n", path)
	/*
		fi, err := os.Stat(path)
		if err != nil {
			writeJSON(rw, http.StatusInternalServerError, err.Error())
			return
		}
	*/
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		writeJSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	ticker := time.NewTicker(time.Duration(1) * time.Second)
	go func() {
		s := ""
		for range ticker.C {
			s += "a"
			conn.WriteMessage(websocket.TextMessage, []byte(s))
		}
	}()
}

// func trace() (file string, funcName string, line int, ok bool) {
func trace() string {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	return file + ":" + f.Name() + ":" + strconv.Itoa(line) + ": "
}

func writeJSON(w http.ResponseWriter, httpCode int, in interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	b, err := json.Marshal(in)
	if err != nil {
		log.Printf("Error in marshaling: %s: %v\n", err.Error(), in)
		return err
	}

	max := 500
	last := len(b)
	extra := ""
	if len(b) > max {
		last = max
		extra = fmt.Sprintf(" ... and %d more bytes.", len(b)-max)

	}
	log.Printf("Writing JSON: %s %s\n", string(b[0:last]), extra)

	fmt.Fprintf(w, string(b))
	return nil
}
