# upload-and-download-file
Define Api for Download and Upload (one and multiple) file through gin-gonic.

In this repository, you will find an easy way to write APIs which download a file at the user side on hitting a particular API, in the browser the fill get download and can be discovered later on your downloading path.

Here you will also figure out that one can also upload as many files as he wants at once with the help of /upload API!

The uploaded file will be stored in a specified dir In this case it stores at the `file/save/` directory.

And the file we want the user to download at their side are sent by us from a specified directory and in this case, it's a specified `sample` file which you can discover in the `file/give` directory.

This Project is Independent of GOPATH. As to achieve this and easy to use `Makefile` is present, through which you can make this project run 

follow:-
`git clone git@github.com:02amanag/upload-and-download-file.git`
`cd upload-and-download`
`make build`
`make run`

following the above command in the sequence, your server will be started and visible in the same terminal.

`USE` 
    `make docker`
    to create a docker image for this project and then can be run through docker as well!