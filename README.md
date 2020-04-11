# grpc client

1.使用go mod 环境

2.安装 gorm 

   go get -u github.com/go-sql-driver/mysql
   
   go get -u github.com/jinzhu/gorm
   
3. 关于gorm无法兼容pb文件未进行处理而是用gorm的原生语句进行了数据库的读写
   如果对接收的数据结构体进行重构,有点繁琐

4.安装 grpc

 go get -u google.golang.org/grpc
 
 5.linux
    下载：protoc-3.3.0-linux-x86_64.zip 或 protoc-3.3.0-linux-x86_32.zip
        解压，把bin目录下的protoc复制到GOPATH/bin下，GOPATH/bin加入环境变量。
        
   window
    下载: protoc-3.3.0-win32.zip
    解压，把bin目录下的protoc.exe复制到GOPATH/bin下，GOPATH/bin加入环境变量。
    
 6.生成pb.go文件
   cd 项目目录/GRpc/proto
   运行:protoc --go_out=plugins=grpc:. micro.proto
 
 7.项目构建好后先运行server.go 先开启rpc后端服务,再运行route.go
  route.go直接用原生HandleFunc的监听,rpc客户端go接口可直接接收客户端参数,
  数据库文件和json规则可以看mysql.sql和README
  
  8.有问题还请指正
