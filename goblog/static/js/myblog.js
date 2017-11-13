/**
 * myblog.js
 */

 var MenuJsonUrl = "../static/home_menu.json"

 $(function(){
    //dom init finish
    //加载menu
    addMenuList()
 });

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
   *  加载主页menu
   */
  function addMenuList(){
    $.getJSON(MenuJsonUrl, function(data){
      var menuCtrl = $("#menu");
      var menuList = data.menu_list;
      for (var index in menuList){
        var ps = document.createElement('li');
        var link = document.createElement('a');
        link.innerHTML = menuList[index].title;
        ps.append(link);
        menuCtrl.append(ps);
      }
    });
  }