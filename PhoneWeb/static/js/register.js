//鼠标移开姓名框开始验证
$("#use_name").blur(function(){	
    var use_name = $("#use_name").val();
    var Username = /^[a-zA-Z\u4E00-\u9FA5\uf900-\ufa2d·s]{2,20}$/;
    if (!Username.test(use_name)) {
        $("#p_name").html("<font color=\"red\" size=\"2\">姓名格式填写错误！</font>");
    } else {
        $("#p_name").html("<font color=\"green\" size=\"2\">姓名格式正确！</font>");
    }
})
//鼠标放置姓名框不验证
$("#use_name").focus(function(){
    $("#p_name").html("");
  })

//鼠标移开邮箱框开始验证
	$("#use_email").blur(function(){
		var use_email = $("#use_email").val();
        var Useremail = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
		if (!Useremail.test(use_email)) {
			$("#p_email").html("<font color=\"red\" size=\"2\">邮箱格式填写错误！</font>");
		} else {
			$("#p_email").html("<font color=\"green\" size=\"2\">邮箱格式正确！</font>");
		}
	})
//鼠标放置邮箱框不验证
	$("#use_email").focus(function(){
		$("#p_email").html("");
	  })

//鼠标移开验证码框开始验证
	$("#use_random").blur(function () {
        var use_random = $("#use_random").val();
		if (number!=use_random) {
			$("#p_random").html("<font color=\"red\" size=\"2\">验证码输入错误！</font>");
		} else {
			$("#p_random").html("<font color=\"green\" size=\"2\">验证码输入正确！</font>");
		}
	})

//鼠标放置验证码框不验证
	$("#use_random").focus(function(){
		$("#p_random").html("");
	  })

//鼠标移开密码框开始验证
$("#use_password").blur(function () {
    var password = $("#use_password").val();
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

//鼠标移开确认密码框开始验证
$("#use_password1").blur(function () {
    var password1 = $("#use_password").val();
    var password2 = $("#use_password1").val();
    if (password1==password2) {
        $("#p_password1").html("<font color=\"green\" size=\"2\">两次密码一致！</font>");
    } else {
        $("#p_password1").html("<font color=\"red\" size=\"2\">输入了不同的密码！</font>");
    }
})

//鼠标放置确认密码框不验证
$("#use_password1").focus(function(){
    $("#p_password1").html("");
  })      
  
//注册确认按钮	  
$("#register").click(function(){
	var use_name = $("#use_name").val();
    var Username = /^[a-zA-Z\u4E00-\u9FA5\uf900-\ufa2d·s]{2,20}$/;
	var use_email = $("#use_email").val();
    var Useremail = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
	var use_random = $("#use_random").val();
	var password = $("#use_password").val();
	var Password=/^(\w){6,20}$/;
	var password1 = $("#use_password").val();
    var password2 = $("#use_password1").val();
	if(!Username.test(use_name)){
        alert("姓名输入格式错误！")
	}else if(!Useremail.test(use_email)){
        alert("邮箱输入格式错误!")
	}else if(number!=use_random){
        alert("验证码输入错误！")
	}else if(!Password.test(password)){
        alert("密码输入格式错误！")
	}else if(password1!=password2){
        alert("两次密码输入不一致！")
	}else{
        $.ajax({
			type:"POST",
			url:"/",
			data:{"use_name":$("#use_name").val(),"use_email":$("#use_email").val(),"use_password":$("#use_password").val()},
			success:function(data){
				alert("用户注册成功!");
		  }
		})
	}
})

//邮箱发送按钮
    let number=null;
    $("#deliver").click(function(){
	  $.ajax({
		  type:"POST",
		  url:"/",
		  data:{"use_email":$("#use_email").val()},
		  success:function(data){
				number=data.random;
                  alert("发送成功！");
		}
	})
})