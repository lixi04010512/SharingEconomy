
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