package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"newsPublish/models"
	"math"
	"fmt"
)

//import (
//	"github.com/astaxie/beego"
//	"path"
//	"time"
//	"github.com/astaxie/beego/orm"
//	"newsPublish/models"
//	"math"
//)
//
//type ArticleController struct {
//	beego.Controller
//}
//
//
////展示首页内容
//func(this*ArticleController)ShowIndex(){
//
//	//获取所有文章数据
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取所有文章    select * from article;  queryseter
//	qs :=o.QueryTable("Article")
//	var articles []models.Article
//	//qs.All(&articles)
//	//实现分页
//	//获取总记录数和总页数
//	count,_ := qs.Count()
//	pageSize := 2
//	pageCount := math.Ceil(float64(count) / float64(pageSize))
//	this.Data["count"] = count
//	this.Data["pageCount"] = pageCount
//
//	//处理首页末页内容
//	pageIndex,err :=this.GetInt("pageIndex")
//	if err != nil{
//		pageIndex = 1
//	}
//
//	//获取数据库部分数据  获取几条数据    从哪里开始
//	start := pageSize * (pageIndex - 1)
//	qs.Limit(pageSize,start).All(&articles)
//	this.Data["pageIndex"] = pageIndex
//
//	//传递数据给前端
//	this.Data["articles"] = articles
//
//
//	//指定视图
//	this.TplName = "index.html"
//}
//
////展示添加文章页面
//func(this*ArticleController)ShowAdd(){
//	//获取所有类型,传递给前端显示
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取所有
//	var articleTypes []models.ArticleType
//	o.QueryTable("ArticleType").All(&articleTypes)
//	//传递给前段
//	this.Data["articleTypes"] = articleTypes
//
//	//指定视图
//	this.TplName = "add.html"
//}
//
////处理添加文章数据
//func(this*ArticleController)HandleAdd(){
//	//获取数据
//	articleName :=this.GetString("articleName")
//	content :=this.GetString("content")
//	//file
//	file,head,err :=this.GetFile("uploadname")
//
//	//获取数据
//	if articleName == "" || content == "" || err != nil{
//		beego.Error("获取用户添加数据失败",err)
//		this.TplName = "add.html"
//		return
//	}
//	defer file.Close()
//	//需要判断大小
//	if head.Size > 5000000{
//		beego.Error("图片太大，我不收")
//		this.TplName = "add.html"
//		return
//	}
//	//需要校验格式
//	ext :=path.Ext(head.Filename)
//	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
//		beego.Error("文件格式不正确")
//		this.TplName = "add.html"
//		return
//	}
//
//	//防止重名
//	//beego.Info("time.now = ",time.Now().Format("2006-01-02 15:04:05"))
//	fileName := time.Now().Format("20060102150405")
//	//操作数据
//	this.SaveToFile("uploadname","./static/img/"+fileName+ext)
//
//	//把数据插入到数据库
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取插入对象
//	var article models.Article
//	//给插入对象赋值
//	article.Title = articleName
//	article.Content = content
//	article.Img = "/static/img/"+fileName+ext
//
//	//获取类型数据
//	typeName := this.GetString("select")
//
//	var articleType models.ArticleType
//	articleType.TypeName = typeName
//	o.Read(&articleType,"TypeName")
//
//	article.ArticleType = &articleType
//
//
//	//插入
//	o.Insert(&article)
//
//	//返回数据
//	this.Redirect("/index",302)
//}
//
////展示文章详情页
//func(this*ArticleController)ShowContent(){
//	//获取数据
//	id,err := this.GetInt("id")
//	//校验数据
//	if err != nil{
//		beego.Error("获取数据错误",err)
//		this.TplName = "index.html"
//		return
//	}
//	//处理数据
//	//查询数据库，获取文章信息
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取查询对象
//	var article models.Article
//	//指定查询条件
//	article.Id2 = id
//	//查询
//	o.Read(&article)
//
//	//阅读次数加一
//	//更新操作
//	article.ReadCount += 1
//	o.Update(&article)
//
//	//返回数据
//	this.Data["article"] = article
//
//	//指定视图
//	this.TplName = "content.html"
//}
//
////展示编辑文章页面
//func(this*ArticleController)ShowEditArticle(){
//	//填充的文章原来的数据
//	//获取数据
//	id ,err :=this.GetInt("id")
//	//数据校验
//	if err != nil{
//		beego.Error("获取数据错误",err)
//		this.TplName = "index.html"
//		return
//	}
//	//处理数据
//	//查询
//	o := orm.NewOrm()
//	//获取查询对象
//	var article models.Article
//	//指定查询条件
//	article.Id2 = id
//	//查询
//	o.Read(&article)
//
//
//	//返回数据
//	this.Data["article"] = article
//
//
//	//指定视图
//	this.TplName = "update.html"
//}
//
////封装上传文件函数  重复使用相同代码   函数参数   返回值   写接口
//func UploadFunc(this*ArticleController,fileName string)string{
//	//file
//	file,head,err :=this.GetFile(fileName)
//
//	//获取数据
//	if  err != nil{
//		beego.Error("获取用户添加数据失败",err)
//		this.TplName = "add.html"
//		return ""
//	}
//	defer file.Close()
//	//需要判断大小
//	if head.Size > 5000000{
//		beego.Error("图片太大，我不收")
//		this.TplName = "add.html"
//		return ""
//	}
//	//需要校验格式
//	ext :=path.Ext(head.Filename)
//	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
//		beego.Error("文件格式不正确")
//		this.TplName = "add.html"
//		return ""
//	}
//
//	//防止重名
//	//beego.Info("time.now = ",time.Now().Format("2006-01-02 15:04:05"))
//	filePath := time.Now().Format("20060102150405")
//	//操作数据
//	this.SaveToFile(fileName,"./static/img/"+filePath+ext)
//	return "/static/img/"+filePath+ext
//}
//
////处理编辑数据
//func(this*ArticleController)HandleEditArticle(){
//	//获取数据
//	id,err :=this.GetInt("id")
//	articleName := this.GetString("articleName")
//	content := this.GetString("content")
//	filePath := UploadFunc(this,"uploadname")
//	//校验数据
//	if err != nil || articleName == "" || content == "" || filePath == ""{
//		beego.Error("获取数据错误")
//		this.TplName = "update.html"
//		return
//	}
//
//	//处理数据
//	//更新
//	//获取orm对象
//	o := orm.NewOrm()
//	//获取更新对象
//	var article models.Article
//	//给更新条件赋值
//	article.Id2 = id
//	//先read一下，判断要更新的数据
//	err = o.Read(&article)
//	//更新
//	if err != nil{
//		beego.Error("更新数据不存在")
//		this.TplName = "update.html"
//		return
//	}
//	article.Title = articleName
//	article.Content = content
//	article.Img = filePath
//	o.Update(&article)
//
//	//返回数据
//	this.Redirect("/index",302)
//}
//
////处理删除文章
//func(this*ArticleController)HandleDelete(){
//	//获取数据
//	id,err :=this.GetInt("id")
//	//校验数据
//	if err != nil{
//		beego.Error("删除请求数据错误")
//		this.TplName = "index.html"
//		return
//	}
//	//处理数据
//	//删除操作
//	o := orm.NewOrm()
//	//定义一个删除对象
//	var article models.Article
//	//指定删除条件
//	article.Id2 = id
//	//删除
//	_,err =o.Delete(&article)
//	if err != nil{
//		beego.Error("删除失败")
//		this.TplName = "index.html"
//		return
//	}
//	//返回数据
//	this.Redirect("/index",302)
//}
//
////展示添加文章类型页面
//func(this*ArticleController)ShowAddType(){
//	//获取所有类型
//	o := orm.NewOrm()
//
//	qs := o.QueryTable("ArticleType")
//	var articleTypes []models.ArticleType
//	qs.All(&articleTypes)
//
//	//传递数据给前段
//	this.Data["articleTypes"] = articleTypes
//
//	this.TplName = "addType.html"
//}
//
////处理添加类型数据
//func(this*ArticleController)HandleAddType(){
//	//获取数据
//	typeName := this.GetString("typeName")
//	//校验数据
//	if typeName == ""{
//		beego.Error("类型数据不能为空")
//		this.TplName = "addType.html"
//		return
//	}
//	//处理数据
//	o := orm.NewOrm()
//	//获取插入对象
//	var articleType models.ArticleType
//	//给插入对象赋值
//	articleType.TypeName = typeName
//	//插入
//	o.Insert(&articleType)
//	//返回数据
//	this.Redirect("/addType",302)
//}
type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowIndex() {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	qs.All(&articles)
	count,_:=qs.Count()
	pageSize:=2
	pageCount:=math.Ceil(float64(count)/float64(pageSize))
	this.Data["Count"]=count
	this.Data["pageCount"]=pageCount
	pageIndex,err:=this.GetInt("pageIndex")
	if err!=nil{
		pageIndex=1
	}
	start:=pageSize*(pageIndex-1)
    qs.Limit(pageSize,start).All(&articles)
    this.Data["pageIndex"]=pageIndex

	this.Data["articles"] = articles

	this.TplName = "index.html"
}
func (this *ArticleController) ShowAdd() {
	this.TplName = "add.html"
}
func (this *ArticleController) HandleAdd() {
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	file, head, err := this.GetFile("uploadname")
	fmt.Println(articleName)
	fmt.Println(content)

	if articleName == "" || content == "" || err != nil {
		beego.Error("错误", err)
		this.TplName = "add.html"
		return
	}
	defer file.Close()
	if head.Size > 50000 {
		beego.Error("台大", err)
		this.TplName = "add.html"
		return
	}
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Error("不政权", err)
		this.TplName = "add.html"
		return
	}
	Filename := time.Now().Format("20060102150405")
	err = this.SaveToFile("uploadname", "./static/img/"+Filename+ext)
	if err !=nil{
		panic(err)
	}
	o := orm.NewOrm()
	var article models.Article
	article.Title = articleName
	article.Content = content
	article.Img = "/static/img/" + Filename + ext
	//article.ArticleType = new()
	_,err = o.Insert(&article)
	if err !=nil{
		panic(err)
	}
	this.Redirect("/index", 302)

}
func (this *ArticleController) ShowContent() {
	id, err := this.GetInt("id")
	if err != nil {
		beego.Error("错误", err)
		this.TplName = "index.html"
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id2 = id
	o.Read(&article)
	article.ReadCount += 1

	this.Data["article"] = article
	o.Update(&article)
	this.TplName = "content.html"
}
func(this*ArticleController)ShowEditArticle(){
	//填充的文章原来的数据
	//获取数据
	id ,err :=this.GetInt("id")
	//数据校验
	if err != nil{
		beego.Error("获取数据错误",err)
		this.TplName = "index.html"
		return
	}
	//处理数据
	//查询
	o := orm.NewOrm()
	//获取查询对象
	var article models.Article
	//指定查询条件
	article.Id2 = id
	//查询
	o.Read(&article)


	//返回数据
	this.Data["article"] = article


	//指定视图
	this.TplName = "update.html"
}
func UploadFunc(this*ArticleController,fileName string)string{
	//file
	file,head,err :=this.GetFile(fileName)

	//获取数据
	if  err != nil{
		beego.Error("获取用户添加数据失败",err)
		this.TplName = "add.html"
		return ""
	}
	defer file.Close()
	//需要判断大小
	if head.Size > 5000000{
		beego.Error("图片太大，我不收")
		this.TplName = "add.html"
		return ""
	}
	//需要校验格式
	ext :=path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		beego.Error("文件格式不正确")
		this.TplName = "add.html"
		return ""
	}

	//防止重名
	//beego.Info("time.now = ",time.Now().Format("2006-01-02 15:04:05"))
	filePath := time.Now().Format("20060102150405")
	//操作数据
	this.SaveToFile(fileName,"./static/img/"+filePath+ext)
	return "/static/img/"+filePath+ext
}
func(this*ArticleController)HandleEditArticle(){
	//获取数据
	id,err :=this.GetInt("id")
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	filePath := UploadFunc(this,"uploadname")
	//校验数据
	if err != nil || articleName == "" || content == "" || filePath == ""{
		beego.Error("获取数据错误")
		this.TplName = "update.html"
		return
	}
	//处理数据
	//更新
	//获取orm对象
	o := orm.NewOrm()
	//获取更新对象
	var article models.Article
	//给更新条件赋值
	article.Id2 = id
	//先read一下，判断要更新的数据
	err = o.Read(&article)
	//更新
	if err != nil{
		beego.Error("更新数据不存在")
		this.TplName = "update.html"
		return
	}
	article.Title = articleName
	article.Content = content
	article.Img = filePath
	o.Update(&article)

	//返回数据
	this.Redirect("/index",302)
}
func (this*ArticleController)HandleDelete(){
	id,err:=this.GetInt("id")
	if err!=nil{
		beego.Error("错误")
		this.TplName="index.html"
		return
	}
	o:=orm.NewOrm()
	var article models.Article
	article.Id2=id
	_,err=o.Delete(&article)
	if err!=nil{
		beego.Error("错误1")
		this.TplName="index.html"
		return
	}
	this.Redirect("/index",302)
}
