package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc.go/gorm"
	pb "grpc.go/proto"
	"log"
	"net"
	"os"
	"reflect"
)

const (
	PORT = ":50051"
)

type Service struct {

}

//获取一条产品信息
func (s *Service) GetProductById(ctx context.Context, req *pb.SbyId) (*pb.Response, error) {
    var res pb.Response
	var pro []*pb.Product
	gorm.DbInit.Conn.Table("hb_product").Where("id=?",req.Id).Find(&pro)
	if reflect.ValueOf(res).IsValid() {
		res.Pro=pro
		res.Code=200
		res.Msg="获取数据成功"
	}else{
		res.Code=404
		res.Msg="未找到数据"
	}

	return &res,nil
}

//分页查询
func (s *Service) GetAllProduct(ctx context.Context,req *pb.SAll)(*pb.Response,error)  {
	var res pb.Response
	var pro []*pb.Product

	var  offset int32
	var  limit int32
	//对传过来page和limit判断赋值
	if req.Limit==0{
		limit=10
	}else{
		limit=req.Limit
	}

	if req.Page==0 {
		offset =0
	}else{
		offset=(req.Page-1)*req.Limit
	}
	fmt.Println("开始查询")
	gorm.DbInit.Conn.Table("hb_product").Offset(offset).Limit(limit).Debug().Find(&pro)
	if reflect.ValueOf(pro).IsValid() {
		for _,v:=range pro{
			res.Pro=append(res.Pro, v)
		}
		res.Code=200
		res.Msg="获取数据成功"
	}else{
		res.Code=404
		res.Msg="未找到数据"
	}
	return &res,nil
}

//插入一条产品信息
func (s *Service) InsertOneProduct(ctx context.Context,req *pb.IPOne)(*pb.ReSimp,error)  {
	//组建sql语句
	sql:=fmt.Sprintf("insert into hb_product(product_name,manual,address,content,state,adminuser_id,sorts,created,site_id)values('%s','%s','%s','%s',%v,%v,%v,'%s',%v)",
		req.Pro.ProductName,req.Pro.Manual,req.Pro.Address,req.Pro.Content,req.Pro.State,req.Pro.AdminuserId,req.Pro.Sorts,req.Pro.Created,req.Pro.SiteId)

	res:=gorm.DbInit.Conn.Table("hb_product").Exec(sql).Debug()

	if res.Error!=nil {
		fmt.Println("错误信息:",res.Error)
		return &pb.ReSimp{Code:500,Msg:"服务器响应错误"},nil
	}
	return &pb.ReSimp{Code:200,Msg:"插入成功"},nil

}

//插入多条产品
func (s *Service) InsertManyProduct(ctx context.Context,req *pb.IPMany)(*pb.ReSimp,error)  {
	//批量插入
	sql:= "insert into `hb_product`(`product_name`,`manual`,`address`,`content`,`state`,`adminuser_id`,`sorts`,`created`,`site_id`)values"
	//循环数据并进行封装
	pros :=req.Pro
	if reflect.ValueOf(pros).IsValid() {
		for k,v:=range pros{
			if len(pros)-1==k{
				sql+=fmt.Sprintf("('%s','%s','%s','%s',%v,%v,%v,'%s',%v);",v.ProductName,v.Manual,v.Address,v.Content,v.State,v.AdminuserId,v.Sorts,v.Created,v.SiteId)
			}else{
				sql+=fmt.Sprintf("('%s','%s','%s','%s',%v,%v,%v,'%s',%v),",v.ProductName,v.Manual,v.Address,v.Content,v.State,v.AdminuserId,v.Sorts,v.Created,v.SiteId)
			}

		}

		res:=gorm.DbInit.Conn.Exec(sql)

		if res.Error!=nil {
			log.Fatalln("批量插入错误",res.Error)
			return &pb.ReSimp{Code:500,Msg:"服务器响应失败"},res.Error
		}
	}else{
		return &pb.ReSimp{Code:404,Msg:"传入空数据"},nil
	}
	return &pb.ReSimp{Code:200,Msg:"插入数据成功"},nil
}


func main() {

	//初始化数据库
	gorm.NewWorkerConnect()
	listen,err:= net.Listen("tcp",PORT)
	if err!=nil {
		fmt.Fprint(os.Stderr,err)
		return
	}

	fmt.Printf("listen on: %s\n",PORT)

	server :=grpc.NewServer()
	pb.RegisterProductServiceServer(server,&Service{})

	if err:=server.Serve(listen); err!=nil{
		fmt.Fprint(os.Stderr,err)
	}
}
