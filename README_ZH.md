# goforum

[English.md](README.md)

> 项目还处于开发中...

## 开始安装

```
go get github.com/qclaogui/goforum
```

> 有关`go get timeout`的问题处理，[推荐使用代理](https://github.com/golang/go/wiki/GoGetProxyConfig)。

**或者:** 不使用`go get`获取,项目采用`dep`作为依赖管理,`dep`安装:

```
go get -u github.com/golang/dep/cmd/dep
```

克隆代码并获取项目依赖:

```
git clone https://github.com/qclaogui/goforum.git $GOPATH/src/github.com/qclaogui/goforum

cd $GOPATH/src/github.com/qclaogui/goforum && dep ensure
```

## 修改配置文件

新建数据库,目前项目使用mysql

```
cd $GOPATH/src/github.com/qclaogui/goforum

cp app.yml.example app.yml
```

然后运行

```
cd $GOPATH/src/github.com/qclaogui/goforum/cmd/web

go run main.go
```

浏览器访问: [localhost:8321](http://localhost:8321)

## 开发

安装前端依赖文件，项目根执行:

```
cd $GOPATH/src/github.com/qclaogui/goforum

npm install

npm run watch
```
