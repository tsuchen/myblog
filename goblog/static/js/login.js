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

    var info = checkInputInfo(userName, password);
    info.Legal = true;
    if (info.Legal){
        $("p.login-error-tips").text("")
        request("/login", "post", {username: userName, password: password}, true, function(resp){
            if (resp.Status === success){
                location.assign(resp.Data);
            }else{
                $("p.login-error-tips").text("用户名或密码错误。")
            }    
        });     
    }else{
        $("p.login-error-tips").text(info.Message)
    }
}

//检查输入是否正确
function checkInputInfo(userName, password){
    var checkInfo = {
        Legal: false,
        Message: "",
    };
    
    var legal = checkUserName(userName);
    if(!legal){
        checkInfo.Message = "用户名格式输入不正确。";
        return checkInfo;
    }

    legal = checkPassword(password)
    if(!legal){
        checkInfo.Message = "密码格式输入不正确。";
        return checkInfo;
    }
    
    checkInfo.Legal = true;
    
    return checkInfo;
}