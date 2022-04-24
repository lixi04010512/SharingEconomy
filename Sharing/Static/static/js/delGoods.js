//下架商品按钮
$(".delMyDoods").click(function(){

    var delid = $(".delId").val();

        $.ajax({
            type:"POST",
            url:"/delGoodsPost",
            data:{"delId":delid},
            success:function(data){
                console.log(data)
                alert("商品下架成功!");
                window.location="/cart";
            }
        })
    })