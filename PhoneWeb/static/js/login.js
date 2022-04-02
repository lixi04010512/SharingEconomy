//鼠标移开密码框开始验证
$("#use_password").blur(function () {
    var password = $("#use_password").val();
    //密码验证
    var Password=/^(\w){6,20}$/;
    if (!Password.test(password)) {
        $("#p_password").html("<font color=\"red\" size=\"2\">密码格式填写错误！</font>");
    } else {
        $("#p_password").html("<font color=\"green\" size=\"2\">密码格式正确！</font>");
    }
})

//鼠标放置密码框不验证
$("#use_password").focus(function(){
    $("#p_password").html("");
  })
