package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"grpc.go/proto"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const(
	//grpc后台端口
	ADDRESS    = "localhost:50051"
)

//定义全局rpc连接结构体
var  client proto.ProductServiceClient

//初始化rpc服务器
func ConnRpcClient()  {
	conn ,err :=grpc.Dial(ADDRESS,grpc.WithInsecure())
	if err!=nil {
		fmt.Fprint(os.Stderr,err)
	}
	client =proto.NewProductServiceClient(conn)
}


//通过id查询产品接口
func GetProductById(w http.ResponseWriter,req *http.Request)  {
	//解析url
	query:=req.URL.Query()
	id:=query.Get("id")

	//string转换成int32
	ID,err:=strconv.ParseInt(id,10,32)
	if err!=nil {
		w.Write([]byte(string("id不是整形,校验参数失败")))
	}

	//传入参数结构体
	S:= proto.SbyId{Id:int32(ID)}

	//开始调用服务的rpc服务
	res,err :=client.GetProductById(context.Background(),&S)
	if err!=nil {
		w.Write([]byte("获取数据失败"))
		return
	}
	r,_:=json.Marshal(res)
	w.Write(r)
}

//分页查询所有产品
func GetAllProduct(w http.ResponseWriter,req *http.Request)  {

	query:=req.URL.Query()
	//获取第几页
	page:=query.Get("page")
	//分页条数
	limit:=query.Get("limit")

	//string转换成int32
	pg,err1:=strconv.ParseInt(page,10,32)
	lt,err2:=strconv.ParseInt(limit,10,32)

	if err1!=nil||err2!=nil {
		w.Write([]byte(string("校验参数失败")))
		return
	}

	//开始调用服务端的rpc服务
	R,err :=client.GetAllProduct(context.Background(),&proto.SAll{Page:int32(pg),Limit:int32(lt)})
	if err!=nil {
		fmt.Fprint(os.Stderr,err)
		w.Write([]byte(string("获取数据失败")))
	}
	//json序列化
	r,_:=json.Marshal(R)

	w.Write(r)

}

//插入单条产品
func InsertOneProduct(w http.ResponseWriter,req *http.Request)  {
	//从前端获取json数组
	query:=req.URL.Query()
	product:=query.Get("product")

	//定义接口体接收
	 p:=proto.Product{}
	if len(product)>0 {
		err:=json.Unmarshal([]byte(product),&p)
		fmt.Println(product)
		if err!=nil {
			w.Write([]byte("传入的数组不正确"))
			return
		}
	}
	//修改创建时间
	time:=time.Now().Format("2006-01-02 15:04:05")
	p.Created=time


	//传入参数结构体
	ipone:=proto.IPOne{
		Pro:&p,
	}

	//开始调用服务端的rpc服务
	res,err:=client.InsertOneProduct(context.Background(),&ipone)
	if err!=nil{
		w.Write([]byte("插入失败"))
		return
	}
	//转json
	r,_ :=json.Marshal(res)
	w.Write(r)
}

//插入多条产品
func InsertManyProduct(w http.ResponseWriter,req *http.Request)  {

	//获取当前时间
	time:=time.Now().Format("2006-01-02 15:04:05")

	//获取products json数组
	query:=req.URL.Query()
	ps:=query.Get("products")

	//定义接收结构体
	var pros []*proto.Product

	//解码
	err :=json.Unmarshal([]byte(ps),&pros)
	if err!=nil {
		w.Write([]byte("传入的product参数错误"))
		return
	}

	//更改时间
	for k,_:=range pros{
		pros[k].Created=time
	}

	//服务器参数
	Pmany:=proto.IPMany{Pro:pros}

	//调用服务器端rpc服务器接口
	res,err:=client.InsertManyProduct(context.Background(),&Pmany)

	if err!=nil {
		w.Write([]byte("获取结果失败"))
		return
	}

	//序列化
	r,_:=json.Marshal(res)

	w.Write(r)
}

func main() {
	//连接服务端rpc
	ConnRpcClient()

	//通过id获取产品rpc
	http.HandleFunc("/getProductById",GetProductById)

	//分页获取所有产品
	http.HandleFunc("/getAllProduct",GetAllProduct)

	//插入单条产品数据
	http.HandleFunc("/insertOneProduct",InsertOneProduct)

	//插入多条产品数据
	http.HandleFunc("/insertManyProduct",InsertManyProduct)

	//拉起8080端口
	err :=http.ListenAndServe(":8080",nil)

	if err!=nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
