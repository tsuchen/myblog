/**
 * desc: admin.js
 * author: xuchen
 */

var success = 1

//dom init finish
$(function(){
  //加载menu
  hideMenuList()
});

/**
 *  隐藏menu子菜单
 */
function hideMenuList(){
  $(".menu-list-group").hide();
  addMenuCallFunc()
};

/**
 * 添加菜单事件
 */
function addMenuCallFunc(){
  var menuButtons = document.getElementsByClassName("menu-button");
  var listGroups = document.getElementsByClassName("menu-list-group");
  for(let index = 0; index < menuButtons.length; index ++){
   var click = function(index){
      var menuButton = $(menuButtons[index]);
      menuButton.click(function(){
        var listGroup = listGroups[index];
        if ($(listGroup).is(":visible")){
          $(listGroup).hide(300);
          console.log("hide: ", index)
        }else{
          $(listGroup).show(300);
          console.log("show: ", index)
        }
      });
   }
   click(index);
  };
}

/**
 * 注销登录
 */
function logout(){
  request("/logout", "get", {}, true, function(resp){
    console.log(resp)
    if (resp.Status === success){
      location.assign(resp.Data);
    }else{
      console.log("登出失败");
    }    
  }); 
}