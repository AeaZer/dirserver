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
	"path/filepath"

	"github.com/aeazer/dirserver/utils/color"
	"github.com/aeazer/dirserver/utils/file"
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
	targetDir  string

	uploadIsDir bool
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
	flagSet.StringVar(&targetDir, "target_dir", defaultPath, "Remote target dir which relative to share server dir")
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

	if err = beforeUpload(); err != nil {
		fmt.Println("before upload exec failed: ", color.RedDA.Dyeing(err.Error()))
	}

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
		uploadPath, _ = filepath.Abs(uploadPath)
		uploadIsDir = true
	}
}

func beforeUpload() error {
	if uploadIsDir {
		err := file.ZipFolder(uploadPath)
		if err != nil {
			log.Fatalf(color.RedDA.Dyeing("zip folder failed: path: %s", uploadPath))
		}
		uploadPath = uploadPath + ".zip"
	}
	return nil
}

func upload() error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	osFile, err := os.Open(uploadPath)
	if err != nil {
		return fmt.Errorf("error opening osFile: %s", err)
	}
	defer func() {
		_ = osFile.Close()
		if uploadIsDir {
			_ = os.Remove(uploadPath)
		}
	}()

	fileBaseName := filepath.Base(osFile.Name())
	part, err := writer.CreateFormFile("file", fileBaseName)
	if err != nil {
		return fmt.Errorf("error creating form osFile: %v", err)
	}
	_, err = io.Copy(part, osFile)
	if err != nil {
		return fmt.Errorf("error copying osFile data: %v", err)
	}

	type uploadParams struct {
		TargetPath string `json:"target_path"`
		IsDir      bool   `json:"is_dir"`
	}
	p := &uploadParams{TargetPath: filepath.Join(targetDir, fileBaseName), IsDir: uploadIsDir}
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
