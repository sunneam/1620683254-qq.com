# 1620683254-qq.com

1.使用go mod 环境

2.安装 gorm 
   go get -u github.com/go-sql-driver/mysql
   go get -u github.com/jinzhu/gorm

3.安装 grpc
 go get -u google.golang.org/grpc
 
 4.linux
    下载：protoc-3.3.0-linux-x86_64.zip 或 protoc-3.3.0-linux-x86_32.zip
        解压，把bin目录下的protoc复制到GOPATH/bin下，GOPATH/bin加入环境变量。
   window
    下载: protoc-3.3.0-win32.zip
    解压，把bin目录下的protoc.exe复制到GOPATH/bin下，GOPATH/bin加入环境变量。
    
 5.生成pb.go文件
   cd 项目目录/GRpc/proto
   运行:protoc --go_out=plugins=grpc:. micro.proto
 
