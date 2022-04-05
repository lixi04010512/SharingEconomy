var gbOption = null; //必须，固定变量名
    var filesArr=[]; //固定写法 这个就是上传能拿到的 files的集合，
    var fileClass = 'input-file'; //必须，固定变量名，值可自定义，就是input的 class自定义后 需要在这重新声明
    var imgObj2 = []; //可自定义，图片上传功能 时必须给的东西，传递到imgsUpload
    var upoad; //可自定义，上传功能时的  工具人，自定义然后传递到imgsUpload 就行
    var imgMaxNum=3;  //最多上传3张
    var ispre; ////反显  时的工具人属性，是一个对象，里面包含了 当前图片预览步骤，和图片数组。 这个最好固定，你要改就一起要改upload.js里面的

    setTimeout(function () { //setTimeout变成异步加载，不然提示有些函数找不到，可自己优化
        imgsUpload(imgObj2,upoad,imgMaxNum); //上传时调用

        var imgObj = ['./t1.png', './t2.png', './t3.png'] //反显, 模拟反显的数据来源，真实是从服务器获取数据
        gbOption=ispre = new imgPre(imgObj, true); //反显

        createPreDiv(); //创建用于图片预览的DIV结构，上传和返显都要调用
    }, 100)



    function gogoole() {
        //得到files的集合，后你想new Fromdata 然后ajax, 还是怎么给后台 就随你玩了!!!!
        if(filesArr.length>0)alert('得到files的集合'); console.log(filesArr)
    }
    function ieSubmit() {
        // 注意：IE方式上传，会动态生成 多个 input ，通过原始表单提交方式提交到服务器，记得上传前删除最后一个 input 。需要打开 测试IE上传的html代码
        // 如果要统一上传方式，请在 在upload.js 搜索  IE方式 ，能检索到需要注意的地方
        alert('传统表单方式上传'); $("#isIeform").submit();
    }



    
var documentW = $(document).width() - 80; 
var documentH = $(document).height() - 80;

$("#editImgs").click(function() {
	$(this).hide();
	$("#uploadImgs").show();
	$("#showImgs").hide();
})
var imgSrc;

function imgsUpload(imgObj, obj,imgNum) {
    var lastInput = null;
    lastInput = $('.' + fileClass).last(); 
    $(".uploadImgs-add").click(function () {
        lastInput = $('.' + fileClass).last(); 
        $('.' + fileClass).each(function (k, v) {
            $(this).attr('id', '');
        })
        lastInput.attr('id', 'fileUpIpt'); 
        lastInput.unbind('change');
        lastInput.bind("change", function () {
            if (lastInput[0].files) { //璋锋瓕 IE10浠ヤ笂          
                var win = window.URL || window.webkitURL;
                imgSrc = win.createObjectURL(lastInput[0].files[0]);
                imgObj.push(imgSrc);            
                filesArr.push(lastInput[0].files[0]);
                var upImgs = "<div class='show-img-div2'><span><img src='images/close.png' /></span><img src='" + imgSrc + "'/></div>";
                $("#labelId").before(upImgs);
              

            } else { 
                var upImgs = "<div class='show-img-div2'><span><img src='images/close.png' /></span><img class='ie-imgshow' src='./img/null.png' /></div>";
                $("#labelId").before(upImgs);
                lastInput[0].select();
                imgSrc = lastInput[0].value;
                imgObj.push(imgSrc);
                $(".ie-imgshow").last()[0].innerHTML = "";
                $(".ie-imgshow").last()[0].style.filter = "progid:DXImageTransform.Microsoft.AlphaImageLoader(enabled='true',sizingMethod='scale',src='" + imgSrc + "')";
                //娣诲姞input file 缁撴瀯
                $(".type-file-div").append(lastInput.clone());

            }
            lastInput.val(''); 
            if(imgObj.length>=imgNum){$("#labelId").hide()}
           
            obj = new imgPre(imgObj);
            gbOption = obj;
            createPreDiv();
          
            $(".show-img-div2").each(function (k, v) {
                $(this).find('span').unbind('click');
                $(this).find('span').bind("click", function () {
                    deleteImg(imgObj, k, obj)
                })
              
                $(this).unbind('click');
                $(this).bind('click', function () {
                    gbOption = obj;
                    createPreDiv.imgMiddleShow(k)
                });
            })
        })
    })

}

function deleteImg(imgObj, k, obj) {
    console.log(k)
    imgObj.splice(k, 1);
    filesArr.splice(k, 1);

    if(imgObj.length<imgMaxNum){
        $("#labelId").show()
    }
    //鍒犻櫎鐐瑰嚮鐨勮嚜宸 
    $(".show-img-div2:eq(" + k + ")").remove();
    //鍒犻櫎鍚庨噸鏂扮粦瀹氫簨浠 
    $(".show-img-div2").each(function (k, v) {
        $(this).find('span').unbind('click');
        $(this).find('span').bind("click", function () {
            deleteImg(imgObj, k)
        })
        //鍜岄瑙堝浘鐗囨満鍒惰仈鍔 
        $(this).unbind('click');
        $(this).bind('click', function () {
            gbOption = obj;
            createPreDiv.imgMiddleShow(k)
        });
    })

    //鐐瑰嚮鍥剧墖鍒犻櫎瀵瑰簲鐨刬nput type=file
    if (!$('.' + fileClass).first()[0].files) { 
        $('.' + fileClass).eq(k).remove(); //鍏煎IE锛岃胺姝岀瓑濡傛灉闇€瑕佸拰IE鏂瑰紡涓€鏍风殑鏂瑰紡锛屽彇娑堝垽鏂紝input灏变笉浼氬垹闄 
    }

}


function imgPre(imgObj, flag) {
    this.imgKey = 0;
    this.imgObj = imgObj;

    var _this = this;
    //flag 鐪熻鏄庡彧鏄函棰勮鍔熻兘锛屽浘鐗囪繑鏄撅紝   涓婁紶鍔熻兘鏃朵笉缁檉lag 鍗冲彲
    if (flag) {
        //鍔ㄦ€佸湪showImgs閲屾坊鍔犲浘鐗 
        var showImgs = '';
        this.imgObj.forEach(function (v, k) {
            showImgs += "<div class='show-img-div'><img src='" + v + "'/></div>";
        })
        $('#showImgs').append(showImgs);

        //棰勮鍥剧墖缁戝畾浜嬩欢
        $('.show-img-div').each(function (k, v) {
            $(this).bind('click', function () {
                if (ispre) gbOption = ispre;
                createPreDiv.imgMiddleShow(k)
            });
        })
    }

}

//棰勮鍥剧墖鐨凞IV缁撴瀯
function createPreDiv() {
    var divStr = "<div class=\"imgs-bg\">\n <div class=\"middle\">\n <div class=\"middle-img\">\n <img src=\"\">\n</div>\n<div class=\"left\">\n<img src=\"img/left.png\">\n</div>\n<div class=\"right\">\n <img src=\"img/right.png\">\n</div>\n<p>\xD7</p>\n</div>\n </div>";
    $('body').append(divStr);
    var maxKey = gbOption.imgObj.length;
    //棰勮鍥惧乏鍙崇澶寸偣鍑讳簨浠 
    $(".imgs-bg .right").click(function () {
        if (gbOption.imgKey >= maxKey) return
        gbOption.imgKey++;
        createPreDiv.imgMiddleShow(gbOption.imgKey)
    })
    $(".imgs-bg .left").click(function () {
        if (gbOption.imgKey == 0) return
        gbOption.imgKey--;
        createPreDiv.imgMiddleShow(gbOption.imgKey)
    })
    $(".imgs-bg p").click(function () {
        $(".imgs-bg").hide();
    })
}
//棰勮鍥   瀹藉害鏍峰紡鍔ㄦ€佽皟鏁达紝涓轰簡鍝嶅簲寮 
createPreDiv.imgW = function () {
    var imgW = $(".middle-img img")[0].clientWidth;
    var imgH = $(".middle-img img")[0].clientHeight + 20;
    // console.log(imgW+'---'+documentW)
    // console.log(imgH+'---'+documentH)
    if (imgW <= documentW) {
        $('.imgs-bg .middle').css("width", imgW + 'px')
    } else {
        $('.imgs-bg .middle').css("width", documentW-80 + 'px')
    }
    if (imgH <= documentH) {
        $('.imgs-bg .middle').css("height", imgH + 'px')
    } else {
        $('.imgs-bg .middle').css("height", '100%')
    }
}
//棰勮鍥  鍔ㄦ€佹敼鍙楽RC
createPreDiv.imgMiddleShow = function (key) {
    gbOption.imgKey = key;
    $(".middle-img img").attr("src", gbOption.imgObj[key]);
    $(".imgs-bg").show();
    setTimeout(function () {
        createPreDiv.imgW();
        $(".middle-img img").css("opacity", 1)
    },30)
}


//})