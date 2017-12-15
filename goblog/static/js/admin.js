/**
 * desc: admin.js
 * author: xuchen
 */

var success = 1;
var tipsMessage = "";

//dom init finish
$(function(){
  //加载menu
  addMenuCallFunc();
});

/**
 * 添加菜单事件
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

/**
 * 注销登录
 */
function logout(){
  request("/logout", "get", {}, true, function(resp){
    console.log(resp)
    if (resp.Status === success){
      location.assign(resp.Data);
    }else{
      showTipsModal("登出失败");
    }    
  }); 
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

$("#AlterCategoryModal").on("show.bs.modal", function(event){
  var button = $(event.relatedTarget); 
  var categoryId = button.data("id");

  var modal = $(this);
  modal.find(".modal-title").text("修改分类名称");
  var input = modal.find("#AlterCategoryName")
  var comfirmBtn = modal.find("#ComfirmAlterCategory");
  $(comfirmBtn).click(function(){
    var categoryName = input.val()
    var info = checkCategoryName(categoryName);
    if(info.Legal){
      alterCategory(categoryId, categoryName)
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