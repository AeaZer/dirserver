简体中文 | [English](README-EN.md)

### 🤷‍ dirserver

Dirserver 有两种命令行模式，分别是 share 和 upload，这两种模式是解耦的，随便你使用什么模式，不冲突；这两种模式又是耦合的，upload 是为 share server 提供文件，不确定后面会不会再有 remove 模式，确定的是现在我是不需要的，如果我有精力的话我会去扩展出 remove 

share server command 的作用类似于 Python http.server 函数，为文件夹文件系统启动 Web 服务器。这使您可以轻松地启动静态资源的 Web 服务，可以浏览 --dir 范围下的文件内容、网页，下载 zip 等等，这是原生标准库就提供的，我只是做了一点封装而已，如果开启了文件接收即`--open_receive=true`会生成密钥 ID(passcode) 组成 http receive api 的一部分格式化之后的 api 是`{ip:port}/receive/${passcode}`下面会有详细的接口文档 

upload command 是开启了向 share server 的服务器上传文件（**可以是一个文件夹**），需要拿到 share server 的 passcode 才可以执行成功，但是目前好像是比较鸡肋的。

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
4. `dirserver --dir ${your_folder_path}` 接下来在浏览器中打开将在控制台上输出，如您所见，这张图片是使用 dirserver 构建的。
![效果图](http://www.areazer.top/static/dirserver/dirserver.png)

### dirserver share api
- 接收文件
  ```apidoc
  /**
   * @api {post} /receive/${passcode} 文件接收
   * @apiName passcode
   * @apiGroup /receive
   *
   * @apiParam {File} file 文件
   * @apiParam {String} json_data "{"target_path": "目标路径", "is_dir": false}"
   *
   * @apiHeaderExample {json} Header-Example:
   *     {
   *       "Content-Type": "multipart/form-data"
   *     }
   */
  ```
      
