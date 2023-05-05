# grpc_template 

## 概要
golangのgrpcサーバのテンプレート

以下のDevelopersIOのコードを拝借して、 色々足したサンプルです。
https://dev.classmethod.jp/articles/golang-grpc-sample-project/

説明用にちょっとしたDDD、ざざっとClean Architectureでコードを整理して組んでます。

## 前提

DIのためにwireを使用しています。

https://github.com/google/wire

protoのgenerate

``` 
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

主にドメイン周りのコード整理の説明目的で書き直したので、 現状grpcサーバとしては動きません。

DB周りの構築が必要ですが、現状全然やってません。
