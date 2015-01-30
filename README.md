# Setting up dev
* Add dependent libraries 
```go get github.com/cloudfouny/gosigar github.com/gorilla/mux github.com/gorilla/websocket ```

# Running the project

* Install CompileDaemon ```go get github.com/githubnemo/CompileDaemon```
* From root directory run ```CompileDaemon -directory="gocode" -build="go build -o ../server-inspect" -command="./server-inspect"```
* You should be able to access the page at localhost:8083


# Tips and Tricks

* if you are using vim mode in SublimeText, you might see that you have to repeat press for direction commands.  You can solve it via this: https://gist.github.com/kconragan/2510186

* This part didn't work but retaining notes for later:
	** Install gin ```go get github.com/codegangsta/gin```
	** From root directory run ```gin -t gocode -a 8083 -b ../server-inspect run```

# Building for different platforms
It's likely that some of you are writing the code on Windows, Mac OSX, Linux, etc. while the server you deploy it on might be another.  In that case, you can cross compile for multiple platforms.

* First setup the environment to build for multiple platforms as I've demonstrated in this video https://www.youtube.com/watch?v=KLh1pOz4y_Q
* As shown in the video, you can then cross compile this app for another platform:
	** Go into the _gocode_ directory
	** ```GOOS=linux GOARCH=amd64 go build -o server-inspect-linux```
