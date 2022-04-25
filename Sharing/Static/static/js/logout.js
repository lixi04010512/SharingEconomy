
$("#confirm").click(function(){
    confirm('是否退出当前账号？').then(() => {
        $.ajax({
            type: "POST",
            url: "/logoutPost",
            success: function (data) {
                alert("用户退出成功!");
                window.location = "/login";
            }
        })
    }).catch(() => {
        console.log('已取消')
    })
})


