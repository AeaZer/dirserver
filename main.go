package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort     = 2233
	defaultDirPath  = "./"
	defaultMainFile = "index.html"
)

var (
	help = flag.Bool("h", false, "Print this help message and exit.")

	webPort = flag.Int("port", defaultPort, fmt.Sprintf("Port to listen on, default is %d", defaultPort))
	dirPath = flag.String("dir", defaultDirPath,
		fmt.Sprintf("Http server for dir, Default %s", defaultDirPath))
	mainFile = flag.String("mf", defaultMainFile,
		fmt.Sprintf("Main file to open, Default %s", defaultMainFile))
)

func main() {
	flagParse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serverMainFile(w, fmt.Sprintf("%s%s", *dirPath, *mainFile))
	})

	log.Printf("Starting server on port %d\n", *webPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *webPort), nil); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func flagParse() {
	flag.Parse()

	if *help {
		flag.CommandLine.Usage()
		os.Exit(0)
	}
}

func serverMainFile(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(fileBytes)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
	}
}
