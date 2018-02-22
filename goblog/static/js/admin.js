/**
 * desc: admin.js
 * author: xuchen
 */

var success = 1;
var mdEditor;

//dom init finish
$(function(){
  initUI();
});

function initUI(){
  var ul = $(".sidebar-nav>ul")
  var activeId = ul.data("li-actived");
  if(activeId === "editblog-menu"){
    $(ul).parent().hide();
    var content = $("div.content")
    content.removeClass("content")
    content.addClass("editblog-content")
    createMdEditor();
  }
  initLeftMenu(activeId);
  initTriangleIcon();
}

//创建MdEditor
function createMdEditor(){
  mdEditor = editormd({
      id      : "my-editormd",
      width   : "100%",
      height  : 800,
      path    : "/static/js/lib/",
      saveHTMLToTextarea : true,//注意3：这个配置，方便post提交表单

      /**上传图片相关配置如下*/
      imageUpload : true,
      imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
      imageUploadURL : "/static/img/upload/",//注意你后端的上传图片服务地址
  });
}

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

function initLeftMenu(activeId){
  $('ul[id="' + activeId + '"]').addClass("in");
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

//关闭提示窗口的回调函数
var closeCallbackOfTipsModal = null;

$("#TipsModal").on("show.bs.modal", function(event){
  // var modal = $(this);
  // modal.find(".tips-modal-body").text(tipsMessage);
})

$("#TipsModal").on("hide.bs.modal", function(event){
  if(closeCallbackOfTipsModal != null){
    setTimeout(function(){
      closeCallbackOfTipsModal();
      closeCallbackOfTipsModal = null;
    }, 500);
  }
});

//提示框
function showTipsModal(message, callback){
  $("#tips-modal-body").text(message);
  if(callback != null){
    closeCallbackOfTipsModal = callback
  }
  $("#TipsModal").modal("show");
}

//创建一个新用户
$("#CreateNewUser").click(function(){
  var userName = $("#InputUserName").val();
  var password = $("#InputPassword").val();
  var comfirmPass = $("#ComfirmPassword").val();

  if(comfirmPass != password){
    showTipsModal("密码输入不一致。")
    return;
  }

  var legal = checkUserName(userName);
  if(!legal){
    showTipsModal("用户名格式输入不正确。");
    return;
  }

  legal = checkPassword(password);
  if(!legal){
    showTipsModal("密码格式输入不正确。");
    return;
  }

  var curPageIndex = $("ul.pagination li.active a").text();
  //请求创建新用户
  request("/admin/userlist/p/" + curPageIndex, "post", {Type: "NewUser", UserName: userName, Password: password}, true, function(resp){
      if(resp.Status === success){
        showTipsModal("创建新用户成功。", function(){
          location.assign(resp.Data);
        });
      }else{
        showTipsModal("创建新用户失败。");
      }    
  }); 
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
        showTipsModal("更新用户信息成功。", function(){
          location.assign(resp.Data);
        });     
      }else{
        showTipsModal("更新用户信息失败。");
      }    
  }); 
}

$("#UpdateProfile").click(function(){
  updateUserProfile()
});

function updateUserPassword(){
  var oldPass = $("#InputOldPass").val();
  var newPass = $("#InputNewPass").val();
  var comfirmPass = $("#InputConfirmPass").val();

  if(oldPass === newPass){
    showTipsModal("新密码和原密码不能相同。");
    return;
  }

  if(comfirmPass != newPass){
    showTipsModal("两次输入密码不一致。");
    return;
  }

  var legal = checkPassword(newPass);
  if (!legal) {
    showTipsModal("密码输入格式不正确，请重新输入。");
    return;
  }

  //请求更新用户密码(type:2)
  request("/admin", "post", {Type: 2, OldPassword: oldPass, Password: newPass}, true, function(resp){
    if(resp.Status === success){
      showTipsModal("更新密码成功。", function(){
        location.assign(resp.Data);
      });
    }else{
      showTipsModal("更新密码失败。");
    }    
  });
}

$("#UpdatePass").click(function(){
  updateUserPassword();
});

$("#PrePageLink").click(function(){
  //获取当前页数
  var curPageIndex = $("ul.pagination li.active a").text();
  var prePageIndex = parseInt(curPageIndex) - 1;
  if (prePageIndex >= 1) {
    location.assign(prePageIndex);
  }
});

$("#NextPageLink").click(function(){
  //获取当前页数
  var curPageIndex = $("ul.pagination li.active a").text();
  var totalPages = $("#content-list").data("total-pages");
  var nextPageIndex = parseInt(curPageIndex) + 1;
  if (nextPageIndex <= totalPages) {
    location.assign(nextPageIndex);
  }
});

$("#NewCategory").click(function(){
  addCategory();
});

/**
 * 添加博客分类
 */
function addCategory(){
  var inputStr = $("#InputCategoryName").val();
  var legal = checkCategoryName(inputStr);
  if(legal){
    var curPageIndex = $("ul.pagination li.active a").text();
    request("/admin/categorylist/p/" + curPageIndex, "post", {Type: "add", CategoryName: inputStr}, true, function(resp){
      if (resp.Status === success){
        showTipsModal("添加分类成功", function(){
          location.assign(resp.Data);
        });
      }else{
        showTipsModal("添加分类失败");
      }    
    });     
  }else{
    showTipsModal("分类名称输入不合法。");
  }
}

//删除博客分类
function deleteCategory(e){
  var categoryName = e.getAttribute("data-name");
  var curPageIndex = $("ul.pagination li.active a").text();
  request("/admin/categorylist/p/" + curPageIndex, "post", {Type: "delete", CategoryName: categoryName}, true, function(resp){
    if (resp.Status === success){
      showTipsModal("删除分类成功", function(){
        location.assign(resp.Data);
      });
    }else{
      showTipsModal("删除分类失败");
    }    
  }); 
}

$("#AlterModal").on("show.bs.modal", function(event){
  var button = $(event.relatedTarget); 
  var id = button.data("id");
  var name = button.data("name");
  var type = button.data("type");

  var modal = $(this);
  modal.find(".modal-title").text("修改名称");
  var input = modal.find("#AlterName");
  input.val(name);
  var comfirmBtn = modal.find("#ComfirmAlter");
  $(comfirmBtn).click(function(){
    var alterName = input.val();
    var legal = checkCategoryName(alterName);
    if (type == "AlterCategory") {
      if(checkCategoryName(alterName)){
        alterCategory(id, alterName);
      } 
    }else if (type == "AlterTag"){
      if(checkTagName(alterName)){
        alterTag(id, alterName);
      }
    }
  });
})

function alterCategory(categoryId, categoryName){
  var curPageIndex = $("ul.pagination li.active a").text();
  request("/admin/categorylist/p/" + curPageIndex, "post", {Type: "alter", CategoryId: categoryId, CategoryName: categoryName}, true, function(resp){
    if (resp.Status === success){
      showTipsModal("修改分类成功", function(){
        location.assign(resp.Data)
      });
    }else{
      showTipsModal("修改分类失败");
    }    
  }); 
}

//添加标签
$("#NewTag").click(function(){
  var nameInput = $("#InputTagName");
  var inputStr = nameInput.val();

  var legal = checkTagName(inputStr);
  if(legal){
    var curPageIndex = $("ul.pagination li.active a").text();
    request("/admin/taglist/p/" + curPageIndex, "post", {Type: "add", TagName: inputStr}, true, function(resp){
      if (resp.Status === success){
        showTipsModal("添加标签成功", function(){
          location.assign(resp.Data);
        });
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
  var curPageIndex = $("ul.pagination li.active a").text();
  request("/admin/taglist/p/" + curPageIndex, "post", {Type: "delete", TagName: tagName}, true, function(resp){
    if (resp.Status === success){
      showTipsModal("删除标签成功", function(){
        location.assign(resp.Data);
      });
    }else{
      showTipsModal("删除标签失败");
    }    
  }); 
}

//修改博客标签
function alterTag(tagId, tagName){
  var curPageIndex = $("ul.pagination li.active a").text();
  request("/admin/taglist/p/" + curPageIndex, "post", {Type: "alter", TagId: tagId, TagName: tagName}, true, function(resp){
    if (resp.Status === success){
      showTipsModal("修改标签成功", function(){
        location.assign(resp.Data)
      }); 
    }else{
      showTipsModal("修改标签失败");
    }    
  }); 
}

//暂存博客文章
function saveArticle(){
  var title = $("title-input").val();
  var cate = $("cate-select").val();
  var tag = $("tag-select").val();
  var content = $("my-editormd-markdown-doc").val();
  var blogID = $("editblog-form").data("article-id");
  request("/admin/editblog/blog/" + blogID, "post", {Type: "save", Title: title, Cate: cate, Tags: tag, Content: content}, 
    true, function(resp){
      if (resp.Status === success){
        showTipsModal("暂存文章成功");
      }else{
        showTipsModal("暂存文章失败");
    }    
  }); 
}

//删除博客文章
function deleteArticle(){
  
}

//发表博客
function sendArticle(){
  var title = $("title-input").val();
  var cate = $("cate-select").val();
  var tag = $("tag-select").val();
  var content = $("my-editormd-markdown-doc").val();
  var blogID = $("editblog-form").data("article-id");
  request("/admin/editblog/blog/" + blogID, "post", {Type: "save", Title: title, Cate: cate, Tags: tag, Content: content}, 
    true, function(resp){
      if (resp.Status === success){
        showTipsModal("发表文章成功");
      }else{
        showTipsModal("发表文章失败");
    }    
  });
}