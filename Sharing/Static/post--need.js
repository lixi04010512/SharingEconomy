$("#needBtn").click(function (){
    demandKind=$(".form-select").val()
    demandName=$("#addnote").val()
    console.log(demandKind,"-----",demandName)
    $.ajax({
        url:"/addDemandPost",
        type:"post",
        dataType:"json" ,
        data: {'demandKind': demandKind,'demandName':demandName},
        success:function (data){
            console.log(data)
            window.location="/chat";
        },


    })
})