//获取编辑对象的信息,显示在页面
function load(){
    $.ajax({
        type:"POST",
        url:"",
        success:function(data){
            $("#edit").empty();
          for( var i of data.data){
            $("#edit").append(`
					<label>
							用户id
				    </label>
					<div >
						<input id="roleId" required="" placeholder="请输入用户id" value="${i.user_id}" name="roleId" type="text">
					</div>
                    <label >
							用户名称
				    </label>
					<div>
						<input id="roleName" required="" placeholder="请输入用户名称" value="${i.user_name}" name="roleName" type="text">
					</div>
                    <label>
                         用户密码
				    </label>
					<div >
						<input id="rolePassword" required="" placeholder="请输入用户密码" value="${i.user_password}" name="rolePassword" type="text">
					</div>
              `
            )}
          
        }
    })
}
load();

//修改用户
$("#editRole").click(function(){
    $.ajax({
        type:"POST",
        url:"todo/updateuser",
        data:{"user_id":$("#roleId").val(),"user_name":$("#roleName").val(),"user_password":$("#rolePassword").val()},
        success:function(){
           alert("成功修改！");
           window.location="./user.html";
        }
    })
})