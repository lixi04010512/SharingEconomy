window.onload=function(){
    var input=document.getElementById("uploadfile");
    var div;
    // 当用户上传时触发事件
    input.onchange=function(){
        readFile(this);
    }
    //处理图片并添加都dom中的函数
    var readFile=function(obj){
        // 获取input里面的文件组
        var fileList=obj.files;
        //对文件组进行遍历，可以到控制台打印出fileList去看看
        for(var i=0;i<fileList.length;i++){
            var reader= new FileReader();
            reader.readAsDataURL(fileList[i]);
            // 当文件读取成功时执行的函数
            reader.onload=function(e){
                div=document.createElement('div');
                div.innerHTML='<img src="'+this.result+'" />';
                document.getElementById("img-box").appendChild(div);
            }
        }
    }
}