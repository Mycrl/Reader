# Reader

`TXT`文件阅读器后端，自动抽离`TXT`文件章节信息，并可以自动扫描目录下面的文件，具备监听文件变动自动更新缓存数据的能力.


### 生成项目

从Git仓库获取源代码并编译:
```bash
git clone https://github.com/quasipaa/Reader
cd Reader
go build
```

编译完成之后当前目录下面会出现可执行文件:
```bash
chmod +x ./Reader
./Reader
```

### 启动

可通过命名行参数启动服务器.
获取帮助:
```bash
./Reader -h
```

* `-a` `string` 文件存储目录 (default "./data")</br>
* `-d` `string` 数据库存储目录 (default "./kv")</br>
* `-p` `number`  绑定端口 (default "8080")</br>

启动示例:
```bash
./Reader -a /data/AppData -d /data/databse -p 80
```
这将启动服务并监听`0.0.0.0:80`端口.
以`/data/AppData`作为文件存储目录，服务将自动扫描此目录内部变更.
数据库文件存储在`/data/databse`目录.

### License
[MIT](./LICENSE)
Copyright (c) 2020 Mr.Panda.