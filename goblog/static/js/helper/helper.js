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
