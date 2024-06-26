package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aeazer/dirserver/utils/color"
)

const (
	defaultPasscode = ""
	defaultAddr     = "127.0.0.1:2233"
	defaultPath     = "./"
)

var (
	help       bool
	addr       string
	passcode   string
	uploadPath string
	targetPath string
)

const helpMark = "-h"

type Command struct{}

func (*Command) Name() string {
	return "upload"
}

func (c *Command) Run() error {
	flagParse()
	if help {
		return nil
	}

	return nil
}

func flagParse() {
	c := Command{}
	flagSet := flag.NewFlagSet(fmt.Sprintf("%s command", color.YellowDA.Dyeing(c.Name())), flag.ExitOnError)
	flagSet.StringVar(&addr, "addr", defaultAddr, "Address of share server")
	flagSet.StringVar(&passcode, "passcode", defaultPasscode, "Passcode to upload")
	flagSet.StringVar(&uploadPath, "upload_path", defaultPath, "Upload local file path")
	flagSet.StringVar(&targetPath, "target_path", defaultPath, "Remote target file path")
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

	checkCommand()

	if err = upload(); err != nil {
		fmt.Println("upload failed: ", color.RedDA.Dyeing(err.Error()))
	}
	return
}

func checkCommand() {
	stat, err := os.Stat(uploadPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf(color.RedDA.Dyeing("--path=%s file not exists", uploadPath))
		}
	}
	if stat.IsDir() {
		log.Fatal(color.RedDA.Dyeing("Upload folder is not supported for the time being"))
	}
}

func upload() error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(uploadPath)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("error copying file data: %v", err)
	}

	type uploadParams struct {
		TargetPath string `json:"target_path"`
	}
	p := &uploadParams{TargetPath: targetPath}
	bs, _ := json.Marshal(p)
	err = writer.WriteField("json_data", string(bs))
	if err != nil {
		return fmt.Errorf("writing JSON field error: %v", err)
	}
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("closing writer error: %v", err)
	}

	url := fmt.Sprintf("http://%s/receive/%s", addr, passcode)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("creating request error: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		rbs, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error sending request: %d %s", resp.StatusCode, string(rbs))
	}
	fmt.Println(color.GreenDA.Dyeing("upload file successfully"))
	return nil
}
