$('#addSpecies').click(function(){
    $.ajax({
        type:"POST",
        url:"/addSticks",
        data:{"species_name":$("#roleName").val()},
        success:function(){
            alert("添加成功！");
            window.location="/species"
        }
    })
})