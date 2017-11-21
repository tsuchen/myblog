/**
 *  desc: login.js
 *  author: xuchen
 */

//dom init finish
$(function(){

});

//登录处理
$("#login").on("click", function(){
    loginBlog();
});

function loginBlog(){
    var userName = $("#InputUserName").val();
    var password = $("#InputPassword").val();
    // var md5PasswordInput = $("#md5-password").val();
    // md5PasswordInput.value = toMD5(password);
    // md5PasswordInput.value = password;

    var errorMessage = checkInputInfo(userName, password)
    if (errorMessage === ""){
        request("/login", "post", {username: userName, password: password}, true, function(resp){
            location.assign(resp.Data);
        });
    }else{
        $("p.login-error-tips").text(errorMessage)
    }
}

//检查输入是否正确
function checkInputInfo(userName, password){
    var errorMessage = ""

    var nameReg = new RegExp('^[\w][\w|\_]{5, 12}');
    var isLegal = nameReg.test(userName)
    console.log(isLegal)
    if (!isLegal) {
        errorMessage = "用户名输入格式不正确，请重新输入。"
        return errorMessage;
    }

    var passwordReg = new RegExp('[^\w][\w]{5, 14}');
    isLegal = passwordReg.test(password)
    if (!isLegal) {
        errorMessage = "密码输入格式不正确，请重新输入。"
        return errorMessage;
    }

    return errorMessage;
}