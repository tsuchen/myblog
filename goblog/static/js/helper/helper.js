/**
* 通信
* url: 请求地址, method: 请求类型, data：请求数据, async: 异步执行
*/
function request(url, method, data, async, callback){
    var async = async || true;
    $.ajax(url, {
      method: method,
      data: data,
      dataType: 'json',
      async: async
    }).done(function(resp){
      callback(resp);
    });
}

/**
 * 注销登录
 */
function logout(){
  request("/logout", "get", {}, true, function(resp){
    console.log(resp)
    if (resp.Status === 1){
      location.assign(resp.Data);
    }else{
      showTipsModal("登出失败");
    }    
  }); 
}


/**
 * 检查用户名是否输入正确
 */
function checkUserName(name){
  var reg = /^[a-zA-z]\w{7,15}$/;
  return reg.test(name);
}

/**
 * 检查密码是否输入正确
 */
function checkPassword(pass){
  var reg = /^\w\w{7,15}$/;
  return reg.test(pass);
}

/**
 * 检查邮箱是否输入正确 
 */
function checkEmailInput(inputStr){
  var reg = new RegExp("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"); 
  return reg.test(inputStr)
}

/**
 * 检查分类名称是否输入正确
 */
function checkCategoryName(str){
  var nameReg = /.{1,20}/;
  return nameReg.test(str);
}

/**
 * 检查标签名称是否输入正确
 */
function checkTagName(str){
  var nameReg = /.{1,20}/;
  return nameReg.test(str);
}

/**
 * 检查文章题目是否输入正确
 */
function checkActicleTitle(title){
  var reg = /.{1, 30}/;
  return reg.test(title);
}

/**
 * 检查文章内容是否输入正确
 */
function checkActicleContent(title){
  var reg = /.{1, 20000}/;
  return reg.test(title);
}