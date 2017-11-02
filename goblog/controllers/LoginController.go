package controllers

import(
	"github.com/astaxie/beego"
)

type LoginController struct{
	beego.Controller;
}

func (c *LoginController) Get(){
	c.Data["URL"] = "http://localhost:8080";
	c.Data["Name"] = "tsuchen";
	c.TplName = "login.html";
}

func (c *LoginController) Post() {
	inputs := c.Input();
	username := inputs.Get("UserName");
	password := inputs.Get("Password");
	if(username == "xuchen" && password == "1234"){
		c.TplName = "index.tpl";
	}else{
		c.Ctx.WriteString("登录失败");
	}
}