English | [简体中文](README.md)
### 🤷‍ dirserver

dirserver has two command modes:

1. **upload** - Provides files to the share server
2. **share** - Shares files and receives files via HTTP (optional)

The share server command functions similarly to Python's http.server,
launching a web server for the directory file system. It can easily start web services for static resources,
allowing users to:
1. Browse files within the --dir scope
2. View web pages
3. Download zip files
4. Receive files when --open_receive=true is enabled (generates a passcode ID that forms part of the HTTP receive API: {ip:port}/receive/${passcode})

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

### dirserver share api
- Receive files
  ```apidoc
  /**
   * @api {post} /receive/${passcode} File Reception
   * @apiName passcode
   * @apiGroup /receive
   *
   * @apiParam {File} file File
   * @apiParam {String} params like {"target_path": "Target Path", "is_dir": false}
   *
   * @apiHeaderExample {json} Header-Example:
   *     {
   *       "Content-Type": "multipart/form-data"
   *     }
   */
  ```