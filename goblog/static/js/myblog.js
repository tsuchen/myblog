/**
 * desc: myblog.js
 * author: xuchen
 */

var MenuJsonUrl = "../static/home_menu.json"

//dom init finish
$(function(){
  //加载menu
  addMenuList()
});

/**
 *  加载主页menu
 */
function addMenuList(){
  $.getJSON(MenuJsonUrl, function(data){
    console.log(data)
    var menuCtrl = $("#menu-ul");
    var smallMenuCtrl = $("#small-menu-ul");
    var menuList = data.menu_list;
    for (var index in menuList){
      //add menu
      var ps = $(document.createElement('li'));
      menuCtrl.append(ps); 
      var alink = document.createElement('a');
      alink.innerHTML = menuList[index].title;
      ps.append(alink);
     
      //add smallmenu
      var smallps =$(document.createElement('li'));
      smallMenuCtrl.append(smallps);
      var smallLink = document.createElement('a');
      smallLink.innerHTML = menuList[index].title;
      smallps.append(smallLink);
    }
  });
}