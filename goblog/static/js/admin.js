/**
 * desc: admin.js
 * author: xuchen
 */

//dom init finish
$(function(){
  //加载menu

});

/**
 *  drop menu
 */

$("div.title-right-block").click(function(){
    alert("drop down")
    $(".dropdown-menu").dropdown('toggle');
});