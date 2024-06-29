English | [简体中文](README.md)
### 🤷‍ dirserver

Dirserver has two command-line modes: share and upload. These two modes are decoupled, meaning you can use either mode without conflict; however, they are also coupled because upload serves to provide files to the share server. I'm not sure if there will be a remove mode in the future, but for now, I don't need it. If I have the energy, I will extend it to include a remove mode.

The share server command functions similarly to the Python http.server function, launching a Web server for the folder's file system. This allows you to easily start a Web service for static resources, browse the contents of files and web pages under the specified directory, download zips, etc. This is just a bit of encapsulation on top of the native standard library functionality. If file reception is enabled, i.e., `--open_receive=true`, a key ID (passcode) will be generated, forming part of the formatted http receive API: `{ip:port}/receive/${passcode}`. Detailed API documentation will be provided below.

The upload command enables uploading files (which can be a folder) to the share server. You need to obtain the share server's passcode to execute this successfully, but currently, it seems rather redundant.

### quickstart

1. Download the latest executable file from the releases module.
2. Move the executable file to your Windows environment variables .bashrc or .zshrc and refresh the configuration.
3. In the host, type `dirserver -h`. If installed correctly, it will output the following information:
   ```powershell
   GitHub: https://github.com/aeazer/dirserver
   
   Usage:
     dirserver [command] [-subcommand]
   Usage of share command:
     -dir string
           Http server for dir (default "./")
     -open_receive
           Whether to turn on file receiving (default true)
     -passcode_salt string
           Generate passcode salt for enhance security
     -port int
           Port to listen on (default 2233)
   Usage of upload command:
     -addr string
           Address of share server (default "127.0.0.1:2233")
     -passcode string
           Passcode to upload
    -target_dir string
           Remote target dir which relative to share server dir (default "./")
     -upload_path string
           Upload local file path (default "./")
   ```
4. `dirserver --dir ${your_folder_path}` Next, open the browser and you will see the output displayed on the console, as shown in the image built using dirserver.
![Effect Picture](http://www.areazer.top/static/dirserver/dirserver.png)

### dirserver share api
- Receive files
  ```apidoc
  /**
   * @api {post} /receive/${passcode} File Reception
   * @apiName passcode
   * @apiGroup /receive
   *
   * @apiParam {File} file File
   * @apiParam {String} json_data "{"target_path": "Target Path", "is_dir": false}"
   *
   * @apiHeaderExample {json} Header-Example:
   *     {
   *       "Content-Type": "multipart/form-data"
   *     }
   */
  ```