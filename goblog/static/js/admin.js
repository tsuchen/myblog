/**
 * desc: admin.js
 * author: xuchen
 */

var success = 1;

//dom init finish
$(function(){
  //加载menu
  addMenuCallFunc();
  addMenuCallLink();
});

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
        var span = $($(this).find(".glyphicon"));
        if(span.hasClass("rotate-icon")){
          span.removeClass("rotate-icon")
        }else{
          span.addClass("rotate-icon")
        }
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

/**
 * 添加博客分类
 */
function addCategory(){
  var nameInput = $("#InputCategoryName")
  var inputStr = nameInput.val()

  var isLegal, errorStr = checkCategoryName(inputStr)
  if(isLegal){
    request("/admin/category", "post", {Type: "add", CatgoryName: inputStr}, true, function(resp){
      if (resp.Status === success){
          location.assign(resp.Data);
      }else{
          $("p.login-error-tips").text("用户名或密码错误。")
      }    
    });     
  }else{
    console.log(errorStr)
  }
}

/**
 * 检查分类名称输入
 */
function checkCategoryName(str){
  var isLegal = false
  var errorStr = ""

  var nameReg = /^[a-zA-z]\w{1,15}$/
  isLegal = nameReg.test(str)
  if(!isLegal){
    errorStr = "输入的名称不合法";
  }

  return isLegal, errorStr
}