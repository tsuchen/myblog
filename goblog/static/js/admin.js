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
    var menuButton = $(menuButtons[index]);

    var clickFunc = function(n){
      menuButton.click(function(){
        var listGroup = listGroups[n];
        if ($(listGroup).is(":visible")){
          $(listGroup).slideUp(300);
        }else{
          $(listGroup).slideDown(300);
        }
      });
    };

    clickFunc(index);
  }
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