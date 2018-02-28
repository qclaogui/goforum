<div align=center><h1>goforum</h1></div>

<p align="center">
<a href="https://github.com/qclaogui/goforum/issues?q=is%3Aopen+is%3Aissue"><img src="https://img.shields.io/github/issues/qclaogui/goforum.svg" alt="issues-open"></a>
<a href="https://github.com/qclaogui/goforum/issues?q=is%3Aissue+is%3Aclosed"><img src="https://img.shields.io/github/issues-closed-raw/qclaogui/goforum.svg" alt="issues-closed"></a>
<a href="https://github.com/qclaogui/goforum/blob/master/LICENSE"><img src="https://img.shields.io/github/license/qclaogui/goforum.svg" alt="License"></a>
<a href="https://goreportcard.com/report/github.com/qclaogui/goforum"><img src="https://goreportcard.com/badge/github.com/qclaogui/goforum?v=1" /></a>
</p>


> This project is under development, If you want, Let's go!

goforum use Gin + Vue . [laravel](https://github.com/laravel/laravel) and [laravel-mix](https://github.com/JeffreyWay/laravel-mix) helped me a lot.

 ## Introduction
 Let's go a forum with TDD

 ## Getting started

   pull down the code with `go get`:

   ```
   go get github.com/qclaogui/goforum
   ```

   make sure you have `dep` installed

   ```
   go get -u github.com/golang/dep/cmd/dep
   ```

   (TODO)Go into the source directory and pull down the project dependencies.:

   ```
   cd $GOPATH/src/github.com/qclaogui/goforum

   dep ensure
   ```

 ## Edit configuration

   Before we set up all the tables in your database, our code depends on a small few configuration files,
   you also need to create your database(now we use mysql)

   ```
   cp app.yml.example app.yml
   ```
   Now run

   ```
   go run main.go
   ```
 ## Development

 To start developing, this project requires NodeJS to build the sources,  you have to execute only a few commands

   ```
   npm install

   npm run watch
   ```

 ## License

 The Laravel framework is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).