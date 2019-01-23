package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//"newsPublish/models"
	"github.com/astaxie/beego/orm"
	"newsPublish/models"
)
//
//type UserController struct {
//	beego.Controller
//}
//
////展示注册页面
//func(this*UserController)ShowRegister(){
//	//指定视图页面
//	this.TplName = "register.html"
//}
//
////处理注册数据
//func(this*UserController)HandleRegister(){
////把数据插入到数据库中
//	//获取数据
//	userName :=this.GetString("userName")
//	pwd :=this.GetString("password")
//	//beego.Info(userName,pwd)
//	//校验数据
//	if userName == "" || pwd == "" {
//		beego.Error("用户名或者密码不能为空")
//		this.TplName = "register.html"
//		return
//	}
//	//操作数据
//	//插入数据
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取插入对象
//	var user models.User
//	//给插入对象赋值
//	user.Name = userName
//	user.Pwd = pwd
//	//插入
//	_,err := o.Insert(&user)
//	if err != nil{
//		beego.Error("用户注册失败")
//		this.TplName = "register.html"
//		return
//	}
//	//返回数据
//	//this.Ctx.WriteString("用户注册成功")
//	this.TplName = "login.html"
//	//this.Redirect("/login",302)
//}
//
////展示登录页面
//func(this*UserController)ShowLogin(){
//	//指定视图
//	this.TplName = "login.html"
//}
//
////处理登录数据
//func(this*UserController)HandleLogin(){
//	//获取数据
//	userName :=this.GetString("userName")
//	pwd :=this.GetString("password")
//	//校验数据
//	if userName == "" || pwd == ""{
//		beego.Error("用户名或者密码不能为空")
//		this.TplName = "login.html"
//		return
//	}
//	//处理数据
//	//查询
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取查询对象
//	var user models.User
//	//指定查询条件
//	user.Name = userName
//	//查询
//	err := o.Read(&user,"Name")
//	if err != nil{
//		beego.Error("用户不存在")
//		this.TplName = "login.html"
//		return
//	}
//
//	//校验密码是否正确
//	if user.Pwd != pwd{
//		beego.Error("输入的密码错误")
//		this.TplName = "login.html"
//		return
//	}
//	//返回数据
//	this.Ctx.WriteString("登录成功")
//	remeber:=this.GetString("remeber")
//	if remeber=="on"{
//		this.Ctx.SetCookie("userName",userName,60*60*24)
//	}else {
//		this.Ctx.SetCookie("userName",userName,-1)
//	}
//
//
//
//	this.Redirect("/index",302)
//}
//
type UserController struct {
	beego.Controller
}

func (this*UserController)ShowRegister(){
	this.TplName="register.html"
}
func (this*UserController)HandleRegister(){
	userName:=this.GetString("userName")
	Pwd:=this.GetString("password")
	//beego.Info(userName,Pwd)
	if userName==""||Pwd==""{
		beego.Error("错误")
		this.TplName="register.html"
		return
	}
	o:=orm.NewOrm()
	var user models.User
	user.Name=userName
	user.Pwd=Pwd
	   _,err:=o.Insert(&user)
	if err!=nil{
		beego.Error("1错误",err)
		this.TplName="register.html"
		return
	}
	this.Ctx.WriteString("用户注册成功")

}
func (this*UserController)ShowLogin(){
	this.TplName="login.html"
}
func (this*UserController)HandleLogin(){
	userName:=this.GetString("userName")
	Pwd:=this.GetString("password")
	if userName==""||Pwd=="" {
		beego.Error("错误")
		this.TplName="login.html"
		return
	}
	o:=orm.NewOrm()
	var user models.User
	user.Name=userName
	err:=o.Read(&user,"Name")
	if err!=nil {
		beego.Error("错误1")
		this.TplName="login.html"
		return
	}
	if user.Pwd!=Pwd{
		beego.Error("错误2")
		this.TplName="login.html"
		return
	}
	//this.Ctx.WriteString("成功")
	this.Redirect("/index",302)

}