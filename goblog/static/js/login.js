/**
 *  desc: login.js
 *  author: xuchen
 */

var success = 1

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

    var errorMessage = checkInputInfo(userName, password);
    errorMessage = "";
    if (errorMessage === ""){
        $("p.login-error-tips").text("")
        request("/login", "post", {username: userName, password: password}, true, function(resp){
            if (resp.Status === success){
                location.assign(resp.Data);
            }else{
                $("p.login-error-tips").text("用户名或密码错误。")
            }    
        });     
    }else{
        $("p.login-error-tips").text(errorMessage)
    }
}

//检查输入是否正确
function checkInputInfo(userName, password){
    var errorMessage = ""

    var nameReg = /^[a-zA-z]\w{7,15}$/
    var isLegal = nameReg.test(userName)
    if (!isLegal) {
        errorMessage = "用户名输入格式不正确，请重新输入。"
        return errorMessage;
    }

    var passwordReg = /^\w\w{7,15}$/
    isLegal = passwordReg.test(password)
    if (!isLegal) {
        errorMessage = "密码输入格式不正确，请重新输入。"
        return errorMessage;
    }

    return errorMessage;
}