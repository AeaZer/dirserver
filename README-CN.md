简体中文 | [English](README.md)
### 🤷‍ dirserver
Dirserver 的作用类似于 Python http.server 函数，为文件夹文件系统启动 Web 服务器。这使您可以轻松地启动静态资源的 Web 服务器或启动 H5 静态网页。

### quickstart
1. 在 releases 模块中下载最新的可执行文件。
2. 将可执行文件移动到 windows 环境变量 .bashrc 或 .zshrc 并刷新配置。
3. 在主机中，键入 dirserver --h 。如果安装正确，它将输出以下信息：
    ```powshell
    Usage of dirserver:
    -dir string
        Http server for dir (default "./")
    -h  Print this help message and exit.
    -port int
        Port to listen on (default 2233)
    ```
4. `dirserver --dir ${your_folder_path}` 接下来在浏览器中打开将在控制台上输出，如您所见，这张图片是使用 dirserver 构建的。
![效果图](http://www.areazer.top/static/dirserver/dirserver.png)