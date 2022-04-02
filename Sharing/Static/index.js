$("#sbFile").click(function (){

    ImgName = $('#file').val()
    $.ajax({
        url:"http://127.0.0.1:8080/upload",
        type:"post",
        dataType:"json",
        data:{imgname:ImgName},
        success:function (data) {
            console.log(data)
            // $("#balance").val(data.data2/1000000000000000000)
        },
    })
    console.log(ImgName)
})