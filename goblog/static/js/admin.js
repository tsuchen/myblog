/**
 * desc: admin.js
 * author: xuchen
 */

//dom init finish
$(function(){
  //加载menu
  hideMenuList()
});

/**
 *  隐藏menu子菜单
 */
function hideMenuList(){
  $("#articleListGroup").hide(0);
  $("userListGroup").hide(0);
};

$("#articleButton").click(function () {
  var div = $("#articleListGroup");
  if(div.is(":visible")){
    div.hide(500);
  }else{
    div.show("slow");
  }
});

$("#userButton").click(function () {
  var node = $("#userListGroup");
  if(node.is(":visible")){
    node.hide(500);
  }else{
    node.show("slow");
  }
});

