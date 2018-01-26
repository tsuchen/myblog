/**
 * desc: admin.js
 * author: xuchen
 */

var success = 1;
var tipsMessage = "";

//dom init finish
// $(function(){
//   //加载menu
//   addMenuCallFunc();
// });

/**
 * cd 添加菜单事件
 */
function addMenuCallFunc(){
  var menuButtons = document.getElementsByClassName("menu-button");
  var listGroups = document.getElementsByClassName("menu-list-group");
  for(let index = 0; index < menuButtons.length; index ++){
    var menuButton = $(menuButtons[index]);
    var icon = $(menuButton.find(".glyphicon"));
    var group = $(listGroups[index]);
    if (group.is(":visible")){
      icon.addClass("rotate-icon");
    }

    var clickFunc = function(n){
      menuButton.click(function(){
        var span = $($(this).find(".glyphicon"));
        var listGroup = listGroups[n];
        if ($(listGroup).is(":visible")){
          $(listGroup).slideUp();
          span.removeClass("rotate-icon");
        }else{
          $(listGroup).slideDown();
          span.addClass("rotate-icon");
        }
      });
    };

    clickFunc(index);
  }
}

$("#AddCategory").click(function(){
  addCategory()
});

/**
 * 检查分类名称输入
 */
function checkCategoryName(str){
  var checkInfo = {
    Legal: false,
    Message: "",
  };

  var nameReg = /.{1,20}/;
  checkInfo.Legal = nameReg.test(str);
  if(!checkInfo.Legal){
    checkInfo.Message = "输入的名称不合法";
  }

  return checkInfo;
}

/**
 * 添加博客分类
 */
function addCategory(){
  var nameInput = $("#InputCategoryName");
  var inputStr = nameInput.val();

  var info = checkCategoryName(inputStr);
  if(info.Legal){
    request("/admin/category", "post", {Type: "add", CategoryName: inputStr}, true, function(resp){
      if (resp.Status === success){
        location.assign(resp.Data);
      }else{
        showTipsModal("添加分类失败");
      }    
    });     
  }else{
    showTipsModal(info.Message);
  }
}

//删除博客分类
function deleteCategory(e){
  var categoryName = e.getAttribute("data-name");
  request("/admin/category", "post", {Type: "delete", CategoryName: categoryName}, true, function(resp){
    if (resp.Status === success){
      location.assign(resp.Data);
    }else{
      showTipsModal("删除分类失败");
    }    
  }); 
}

$("#AlterModal").on("show.bs.modal", function(event){
  var button = $(event.relatedTarget); 
  var id = button.data("id");
  var name = button.data("name")
  var type = button.data("type")

  var modal = $(this);
  modal.find(".modal-title").text("修改名称");
  var input = modal.find("#AlterName")
  input.val(name)
  var comfirmBtn = modal.find("#ComfirmAlter");
  $(comfirmBtn).click(function(){
    var alterName = input.val()
    var info = checkCategoryName(alterName);
    if(info.Legal){
      if (type == "AlterCategory") {
        alterCategory(id, alterName)
      } else if (type == "AlterTag") {
        alterCategory(id, alterName)
      }
    }
  });
})

//修改博客分类
function alterCategory(categoryId, categoryName){
  request("/admin/category", "post", {Type: "alter", CategoryId: categoryId, CategoryName: categoryName}, true, function(resp){
    if (resp.Status === success){
      location.assign(resp.Data)
    }else{
      showTipsModal("修改分类失败");
    }    
  }); 
}

$("#TipsModal").on("show.bs.modal", function(event){
  var modal = $(this);
  modal.find(".tips-modal-body").text(tipsMessage);
})

//提示框
function showTipsModal(message){
  tipsMessage = message;
  $("#TipsModal").modal();
}

//添加标签
$("#AddTag").click(function(){
  var nameInput = $("#InputTagName");
  var inputStr = nameInput.val();

  var info = checkCategoryName(inputStr);
  if(info.Legal){
    request("/admin/tag", "post", {Type: "add", TagName: inputStr}, true, function(resp){
      if (resp.Status === success){
        location.assign(resp.Data);
      }else{
        showTipsModal("添加标签失败");
      }    
    });     
  }else{
    showTipsModal(info.Message);
  }
});

//删除标签
function deleteTag(e) {
  var tagName = e.getAttribute("data-name");
  request("/admin/tag", "post", {Type: "delete", TagName: tagName}, true, function(resp){
    if (resp.Status === success){
      location.assign(resp.Data);
    }else{
      showTipsModal("删除分类失败");
    }    
  }); 
}

//修改博客标签
function alterTag(tagId, tagName){
  request("/admin/tag", "post", {Type: "alter", TagId: tagId, TagName: tagName}, true, function(resp){
    if (resp.Status === success){
      location.assign(resp.Data)
    }else{
      showTipsModal("修改标签失败");
    }    
  }); 
}


///////////////////////////测试代码/////////////////////////////
$(function(){
  //加载menu
  initTriangleIcon();
});

function initTriangleIcon(){
  var li = $(".sidebar-nav>ul>li");
  li.map(function(){
    var span = $(this).first().find("span");
    var ul = $(this).next().find("ul");
    if (ul.hasClass("in")){
      span.addClass("rotate-icon")
    } 
  });
}

$(".sidebar-nav .nav-header").click(function(){
  var icon = $(this).find(".down-icon");
  var target = $(this).parent().next().find("ul");
  var exist = icon.hasClass("rotate-icon");
  if(exist){
    target.slideUp(200, function(){
      icon.removeClass("rotate-icon");
    });
  }else{
    target.slideDown(200, function(){
      icon.addClass("rotate-icon");
    });
  }
});

function updateUserProfile(){
  var nickName = $("#InputNickName").val();
  var sex= $("#InputSex").val();
  var birth = $("#InputBirthday").val();
  var phoneNumber = $("#InputPhoneNumber").val();
  var email = $("#InputEmail").val();
  var desc = $("#InputDesc").val();
  if(email != ""){
    if(!checkEmailInput(email)){
      showTipsModal("邮箱格式不正确。");
      return 
    }
  }
  //请求更新用户信息(type:1)
  request("/admin", "post", {Type: 1, NickName: nickName, Sex: sex, Birth: birth, PhoneNumber: phoneNumber,
    Email: email, Desc: desc}, true, function(resp){
      if(resp.Status === success){
        location.assign(resp.Data);
      }else{
        showTipsModal("更新用户信息失败。");
      }    
  }); 
}

function checkEmailInput(inputStr){
  //正则表达式
  var reg = new RegExp("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"); 
  
  return reg.test(inputStr)
}

$("#UpdateProfile").click(function(){
  updateUserProfile()
});

function updateUserPassword(){
  var oldPass = $("#InputOldPass").val();
  var newPass = $("#InputNewPass").val();
  var comfirmPass = $("#InputConfirmPass").val();
  if(oldPass == ""){
    showTipsModal("原密码不能为空。");
    return ;
  }

  if(newPass == ""){
    showTipsModal("新密码不能为空。");
    return;
  }

  if(oldPass == newPass){
    showTipsModal("原密码和新密码不能相同。");
  }

  if(comfirmPass != newPass){
    showTipsModal("两次输入密码不一致。");
    return;
  }

  var passwordReg = /^\w\w{7,15}$/;
  var legal = passwordReg.test(newPass);
  if (!legal) {
    showTipsModal("密码输入格式不正确，请重新输入。");
    return;
  }


  //请求更新用户密码(type:2)
  request("/admin", "post", {Type: 2, OldPassword: oldPass, Password: newPass}, true, function(resp){
    if(resp.Status === success){
      location.assign(resp.Data);
    }else{
      showTipsModal("更新用户密码失败。");
    }    
  });
}

$("#UpdatePass").click(function(){
  updateUserPassword()
});

$("#PrePageLink").click(function(){

});

$("#NextPageLink").click(function(){

});