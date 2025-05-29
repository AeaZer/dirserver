简体中文 | [English](README-EN.md)

### 🤷‍ dirserver

Dirserver 有两种命令模式。

1. **upload** 是为 share server 提供文件
2. **share** 是共享文件并通过 http 的方式接收文件（这是可选的）
share server 命令的功能类似于 Python 的 http.server，为目录文件系统启动 Web 服务器。它可以轻松启动静态资源的 Web 服务，允许用户：

1. 浏览 --dir 范围内的文件
2. 查看网页
3. 下载 zip 文件
4. 启用 --open_receive=true 时接收文件（生成构成 HTTP 接收 API 一部分的密码 ID：{ip：port}receive{passcode}）

### quickstart

1. 在 releases 模块中下载最新的可执行文件。
2. 将可执行文件移动到 windows 环境变量 .bashrc 或 .zshrc 并刷新配置。
3. 在主机中，键入 dirserver -h 。如果安装正确，它将输出以下信息：
   ```powshell
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
- 接收文件
  ```apidoc
  /**
   * @api {post} /receive/${passcode} 文件接收
   * @apiName passcode
   * @apiGroup /receive
   *
   * @apiParam {File} file 文件
   * @apiParam {String} params {"target_path": "目标路径", "is_dir": false}
   *
   * @apiHeaderExample {json} Header-Example:
   *     {
   *       "Content-Type": "multipart/form-data"
   *     }
   */
  ```
      
