package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	defaultPort    = 2233
	defaultDirPath = "./"
)

var (
	help = flag.Bool("h", false, "Print this help message and exit.")

	webPort = flag.Int("port", defaultPort, "Port to listen on")
	dirPath = flag.String("dir", defaultDirPath, "Http server for dir")
)

func main() {
	flagParse()

	fileServer := http.FileServer(http.Dir(*dirPath))
	http.Handle("/", fileServer)

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

	stat, err := os.Stat(*dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf("failed to start: %v\n", err)
		}
	}

	if !stat.IsDir() {
		log.Fatalf("--dir :%s not a folder path", *dirPath)
	}

	_, err = os.Stat(path.Join(*dirPath, "index.html"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Server failed to start: %v\n", err)
		}
		log.Fatalf("Server failed to start: %v\n", err)
	}
}