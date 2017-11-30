/**
 * desc: admin.js
 * author: xuchen
 */

//dom init finish
$(function(){
  //加载menu
  hideMenuList()
});

/**
 *  隐藏menu子菜单
 */
function hideMenuList(){
  $(".menu-list-group").hide(0);
};

/**
 * 添加菜单事件
 */
var menuButtons = $(".menu-button");
for(let index = 0; index < menuButtons.length; index++){
  var menuButton = $(menuButtons[index])
  menuButton.click(function(){
    var listGrounp = $($(".menu-list-group")[index]);
    if (listGrounp.is(":visible")){
      listGrounp.hide(300)
    }else{
      listGrounp.show(300);
    }
  });
};

/**
 * 注销登录
 */
var logoutLinks = $(".logout");
logoutLinks.click(function(){
  request("/logout", "get", {}, true, function(resp){
    if (resp.Status === success){
      location.assign(resp.Data);
    }else{
      console.log("登出失败")
    }    
  }); 
});