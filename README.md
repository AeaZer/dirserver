### dirserver
Dirserver acts like the Python http.server function, starting a web server for a folder file system. This allows you to easily start a web server for static resources or launch an H5 static web page.

### quickstart
1. Download the latest executable file in the release module.
2. Move the executable file to an environment variable, .bashrc, or .zshrc and source it.
3. In your console, type `dirserver --h`. If installed correctly, it will output the following information:
    ```powshell
    Usage of dirserver:
    -dir string
        Http server for dir (default "./")
    -h  Print this help message and exit.
    -port int
        Port to listen on (default 2233)
    ```
4. `dirserver --dir ${your_folder_path}` next open in browser will output on console, As you can see, this picture is built using dirserver.
![效果图](http://www.areazer.top/static/dirserver/dirserver.png)