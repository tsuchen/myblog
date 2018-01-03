/**
 * desc: editblog.js
 * author: xuchen
 */

 var mdEditor;

 //init finish
 $(function(){
    createMdEditor();
 });

 //创建MdEditor
 function createMdEditor(){
    mdEditor = editormd({
        id      : "BlogEditor",
        width   : "100%",
        height  : 900,
        path    : "/static/js/lib/"
    });
 }

 ////////////////////////编辑博客文//////////////////////////////////
var selectedCategoryId = 0    //当前选中的分类id
function addCategoryForBlog(e){
  var categoryId = e.getAttribute("data-id");
  var categoryName = e.getAttribute("data-name");
  $("#SelectCategoryBtn").text(categoryName);
  $("#SelectCategoryBtn").append("<span class='caret'></span>");
  selectedCategoryId = categoryId;
}

var selectedTagsArr = new Array();
function addTagForBlog(e){
  var tagId = e.getAttribute("data-id");
  var tagName = e.getAttribute("data-name");
  var parent = $("#SelectTag");

  for (let i = 0; i < selectedTagsArr.length; i ++) {
    if (selectedTagsArr[i] === tagId) {
      return;
    }
  }

  selectedTagsArr.push(tagId);

  var tagDiv = $(document.createElement("div"));
  tagDiv.addClass("btn btn-info tag-block");
  parent.append(tagDiv);  
  tagDiv.text(tagName)

   //删除标签
   var deleteBtn = $(document.createElement("button"));
   deleteBtn.addClass("btn btn-link");
   deleteBtn.text("删除");
   tagDiv.append(deleteBtn);
   deleteBtn.hide();
   deleteBtn.click(function(){
    tagDiv.remove();
    var index = selectedTagsArr.indexOf(tagId);
    selectedTagsArr.splice(index,1);
   });
   
   tagDiv.mouseover(function (){
     deleteBtn.show();
   }).mouseout(function (){
     deleteBtn.hide();
   });
}    