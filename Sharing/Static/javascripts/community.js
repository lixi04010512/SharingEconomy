//新增公益
$(".btnAdd").click(function () {
    window.location = "/addCommunity"
})

//弹出修改页面
$("#community_body").delegate(".btnedit", "click", function () {
    $.ajax({
        type: "POST",
        url: "/editCommunities",
        data: {"id": $(this).data("id")},
        success: function () {
            window.location = "/editCommunity"
        }
    })
})

//删除公益
$("#community_body").delegate(".btndel", "click", function () {
    $.ajax({
        type: "POST",
        url: "/delCommunity",
        data: {"id": $(this).data("id")},
        success: function () {
            alert("成功删除！");
            location.reload();
        }
    })
})

//查询单个借阅
$(".btnsearch").click(function () {
    $.ajax({
        type: "POST",
        url: "/",
        data: {"user_name": $("#username").val()},
        success: function (data) {
            $("#borrow_body").empty();
            for (var i of data.data) {
                $("#borrow_body").append(`
            <tr>
            <td>${i.borrow_id}</td>
            <td>${i.book_name}</td>
            <td>${i.user_name}</td>
            <td>
                <div class="am-btn-toolbar">
                    <div class="am-btn-group am-btn-group-xs">
                    <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.borrow_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                    <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.borrow_id}"><span class="am-icon-trash-o"></span> 删除</button>
                    </div>
                </div>
            </td>
        </tr>    
       `)
            }
        }
    })
})