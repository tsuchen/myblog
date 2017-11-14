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