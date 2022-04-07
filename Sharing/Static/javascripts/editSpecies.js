//修改图书
$("#editSpecies").click(function () {
    $.ajax({
        type: "POST",
        url: "/updateStick",
        data: {"id": $("#roleId").val(), "name": $("#roleName").val()},
        success: function () {
            alert("成功修改！");
            window.location = "/species";
        }
    })
})