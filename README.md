# Go installation Ubuntu 17
This is short instruction about GoLang installtion from https://golang.org/dl/

At first - update your system
```
$ sudo apt update
$ sudo apt upgrade
```
Then  download compile file from golang.org (at this moment currently version is go1.10.2.linux-amd64.tar.gz).
Use that command
```
$  sudo curl -O https://storage.googleapis.com/golang/go1.10.2.linux-amd64.tar.gz
```
Next unpack this archive in local folder
```
$ sudo tar -xvf go1.10.2.linux-amd64.tar.gz
$ sudo mv go /usr/local
```

I use VSCODE and therefore execute command with vscode for path variables setting. Just execute this command in terminal and add export part in the end of /etc/profile file

```
$  vscode /etc/profile
```
```
export PATH=$PATH:/usr/local/go/bin
```

Also add this path to /etc/bashrc with next command (for vscode /etc/bashrc I have warning message and therefore use sudo nano)
```
$ sudo nano /etc/bashrc
```

Save it and update your source
```
$ source /etc/profile
```
Then restart your machine and check the version (shold work)
```
$ go version
```
After that make working directory with src folder and try to start HelloWorld.go project
``` 
$ mkdir workplace/src/helloworld
$ cd workplace/src/helloworld
```
Add GOPATH to /etc/bashrc and /etc/profile (like above) for regular use
```
$  export GOPATH=$HOME/workplace/
```
Make Hello.go in helloworld directory
```
$ vscode Hello.go
```
Type next code in VSCode
```
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
Save it and try to execute (you should be in helloworld directory)
```
$ go run Hello.go
```
If you see the result --- perfect :)

Hope, It'll be useful for GoLang istallation on Ubuntu17.
