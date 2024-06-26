package share

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/aeazer/dirserver/utils/color"
	utilsfile "github.com/aeazer/dirserver/utils/file"
	"github.com/aeazer/dirserver/utils/math"
)

const (
	defaultPort    = 2233
	defaultDirPath = "./"
)

const helpMark = "-h"

const maxFileSize = 100 << 20 // 100 MB

var (
	help         bool
	webPort      int
	dirPath      string
	openReceive  bool
	passcode     string
	passcodeSalt string
)

type Command struct{}

func (*Command) Name() string {
	return "share"
}

func (c *Command) Run() error {
	flagParse()
	if help {
		return nil
	}
	checkCommand()

	fs()
	if openReceive {
		receive()
	}

	log.Printf("Starting server on port %d, passcode: %s\n", webPort, color.YellowDA.Dyeing(passcode))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", webPort), nil); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
	return nil
}

func flagParse() {
	c := Command{}
	flagSet := flag.NewFlagSet(fmt.Sprintf("%s command", color.RedDA.Dyeing(c.Name())), flag.ExitOnError)
	flagSet.IntVar(&webPort, "port", defaultPort, "Port to listen on")
	flagSet.StringVar(&dirPath, "dir", defaultDirPath, "Http server for dir")
	flagSet.BoolVar(&openReceive, "open_receive", true, "Whether to turn on file receiving")
	flagSet.StringVar(&passcodeSalt, "passcode_salt", "", "Generate passcode salt for enhance security")
	if os.Args[1] == helpMark {
		flagSet.Usage()
		help = true
		return
	}
	err := flagSet.Parse(os.Args[2:])
	if err != nil {
		flagSet.Usage()
		log.Fatalf("unknown %s mod command: %v", c.Name(), err)
	}
	passcode = math.MD5Time(passcodeSalt)
}

func checkCommand() {
	stat, err := os.Stat(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Failed to start: %v\n", err)
		}
	}
	if !stat.IsDir() {
		log.Fatalf("--dir :%s not a folder path", dirPath)
	}
}

func fs() {
	fileServer := http.FileServer(http.Dir(dirPath))
	http.Handle("/", mountParent(fileServer))
}

func receive() {
	http.Handle(fmt.Sprintf("/receive/%s", passcode), mountParent(receiveHandler()))
}

func receiveHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(maxFileSize); err != nil {
			doWriteJson(w, http.StatusBadRequest, fmt.Errorf("error parsing form: %v", err))
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			doWriteJson(w, http.StatusBadRequest, fmt.Errorf("error retrieving file: %v", err))
			return
		}
		defer file.Close()
		jss := r.FormValue("json_data")
		type receiveParams struct {
			TargetPath string `json:"target_path"`
		}
		var p receiveParams
		err = json.Unmarshal([]byte(jss), &p)
		if err != nil {
			doWriteJson(w, http.StatusBadRequest, fmt.Errorf("unmarshal json body data occur error: %v", err))
			return
		}
		if p.TargetPath == "" {
			doWriteJson(w, http.StatusBadRequest, errors.New("target path is empty"))
			return
		}
		err = utilsfile.Save(file, path.Join(dirPath, p.TargetPath))
		if err != nil {
			doWriteJson(w, http.StatusBadRequest, fmt.Errorf("save file occur error: %v", err))
			return
		}
		doWriteJson(w, http.StatusAccepted, "file save success!")
	})
}
