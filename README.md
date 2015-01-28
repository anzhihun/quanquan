#QuanQuan

[中文说明](README_cn.md)

QuanQuan is a tool for teamwork. It can help us to work more easy.

##Motivations
I like slack very much when I used it in 2014. It's very good at teamwork. But I have some problems as follows:
>1. It is slow and always broken because of our country network. I feel depressed. Maybe you have the same experience.
>2. And I also worry about the information security. Something we are talked about is secret in our company. Such as new product's information. It should not be known by others at the begining. 
>3. And some companies can not connecte to the internet. But they also need to communicate and work.
>4. There are fewer tools like slack, which is good at teamwork indeed.

So, I decide to make a new open source tool like slack to help us in teamwork. I name it `QuanQuan`. I hope it will be the all works's entrance. All the internal system connect to it. Then all of us can only use `QuanQuan` to deal all the works efficiently.

##How to use
1. Compile and Run by self：Please get [thirdparty library](#thirdparty) before compiling if you didn't get and install them. Then change to directory `build` and execute dist.bat if the os is Windows or dist.sh if the os is Linux or Mac OSX to compile `QuanQuan`. It will start server automaticly after compiled. After the server is started, please open your browser and input `http://localhost:52013` to open login page. Hope you will like it.
2. Run binary file: You can not use this way now, since we didn't release any version before. Maybe later soon. 

##<a name="thirdparty" id="thirdpartyt">Thirdparty Library</a>
1. [gocraft/web](https://github.com/gocraft/web): This is a router used in server. Please execute the command `go get github.com/gocraft/web` to get and install.
2. [websocket](https://code.google.com/p/go.net/websocket): This is websocket library used in server. Please execute the command `go get code.google.com/p/go.net/websocket` to get and install.
3. [tiedot](https://github.com/HouzuoGuo/tiedot):  This is a documented nosql database used in server. Please execute the command `go get github.com/HouzuoGuo/tiedot` to get and install.

##Helping QuanQuan

#### I found a bug!

If you found a bug, please [search existing issues](https://github.com/anzhihun/quanquan/issues) first  to
see if it's already there. If not, please create a new [issue](https://github.com/anzhihun/quanquan/issues), Include steps to consistently reproduce the problem, actual vs. expected results, screenshots, and your OS and
QuanQuan version number. 

#### I have a new suggestion, but don't know how to program!

For feature requests please [search existing feature issues](https://github.com/anzhihun/quanquan/issues) to
see if it's already there; you can comment to upvote it if so. If not, feel free to create an new issue; we'll
change it to the feature issue for you.

#### I want to help with the code!

Awesome! Please feel free to push your request. To make it more easy, please see [Develop Guide](#developGuide)

##<a name="developGuide" id="developGuide">Develop Guide</a>
[Develop Guide]()

##Credits
* [Underscore.JS](http://underscorejs.org/)
* [Backbone.JS](http://backbonejs.org/)
* [Foundation](http://foundation.zurb.com/)
* [Golang](https://golang.org/)
* [gocraft/web](https://github.com/gocraft/web)
* [tiedot](https://github.com/HouzuoGuo/tiedot)

##License
[Apache License Version 2.0](LICENSE)
