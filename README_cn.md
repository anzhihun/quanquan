#QuanQuan

[English Document](README.md)

QuanQuan 是一个用于团队协作的即时通信工具，其目标是为了让大家工作更轻松。

##起源
在2014年第一次使用slack的时候，感觉它太适合团队协作了。很是喜欢。但是由于下面的一些原因，促使我开发这个项目：
>1. 国内的网络环境不是很好，在登录slack的时候，比较慢，还比较容易掉线。这个难以接受。
>2. 既然是团队协作，这里面沟通的事物多与工作有关，从安全的角度而言，并不是很让人放心。特别是公司内部开发新产品，集成内部信息，更是不敢。
>3. 有些厂是不让使用外网的，但并不是说不需要高效的协作和沟通。
>4. 国内这方面的工具太少，以前内网使用的飞秋，在群聊，信息沟通和集成上，与slack相比，差距太大。

鉴于以上原因，自己决定弄一个开源的版本出来，以解决公司内部协作的问题。期望所有的内部系统都能集成到quanquan上来，quanquan成为一个前端的入口，也只用进入这个入口，就可以即时高效的处理各种工作任务。


##如何运行
1. 编译运行：进入build目录，如果是windows平台，请执行dist.bat。如果是linux或者Mac OS X平台，请执行dist.sh。在执行前，请先获取[第三方依赖包](#thirdparty)，若已经安装了，则不用再次获取安装。
2. 可执行文件执行: 目前还没有发布版本，暂时只能自己在本地编译运行。

##<a name="thirdparty" id="thirdpartyt">第三方依赖包</a>
1. [gocraft/web](https://github.com/gocraft/web): 这是一个第三方的web路由，执行命令 go get github.com/gocraft/web 获取并安装到本地
2. [websocket](https://code.google.com/p/go.net/websocket): 这是大家都熟悉的websocket库，国内的不能直接下载，可以通过其他的途径（你懂的）下载，执行命令 go get code.google.com/p/go.net/websocket 获取并安装到本地
3. [tiedot](https://github.com/HouzuoGuo/tiedot): 这是项目使用的数据库，执行命令 go get github.com/HouzuoGuo/tiedot 获取并安装到本地

##帮助我们
1. 发现bug：请优先在已经存在的issue里面查找改bug是否已经有了，如果没有，麻烦提供重现bug的步骤，截图，期望结果，以及运行环境，包括操作系统，quanquan版本号等信息。
2. 新建议：请优先在已经存在的issue里面查找是否已经存在了，如果没有，请新建issue，我们会尽快将确认。
3. 参与开发：非常的欢迎大家提交PR，为了帮助你更容易更好的参与开发，请参见[开发指南](#developGuide)

##<a name="developGuide" id="developGuide">开发指南</a>
详情参见[quanquan开发指南]()
