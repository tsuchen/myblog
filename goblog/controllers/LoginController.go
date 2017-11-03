package controllers

import(
	"github.com/astaxie/beego"
	"encoding/json"
)

type Resp struct{
	Data interface{}
}

type LoginController struct{
	beego.Controller;
}

func (c *LoginController) Get(){
	c.Data["URL"] = "http://localhost:8080";
	c.Data["Name"] = "tsuchen";
	c.TplName = "login.html";
}

func (c *LoginController) Post() {
	resp := &Resp{};
	username := c.GetString("username");
	password := c.GetString("password");
	if(username == "xuchen" && password == "1234"){
		resp.Data = "/";
	}else{
		resp.Data = "/login";
	}
	b, err := json.Marshal(resp)
	if err == nil {
		c.Ctx.ResponseWriter.Write(b)
	}
}