(function(doc, win) {
    var docEl = doc.documentElement,
        resizeEvt = 'orientationchange' in window ? 'orientationchange' : 'resize',
        recalc = function() {
            var clientWidth = docEl.clientWidth;
            if (!clientWidth) return;
            if (clientWidth >= 640) {
                docEl.style.fontSize = '100px';
            } else {
                docEl.style.fontSize = 100 * (clientWidth / 640) + 'px';
            }
        };

    if (!doc.addEventListener) return;
    win.addEventListener(resizeEvt, recalc, false);
    doc.addEventListener('DOMContentLoaded', recalc, false);
})(document, window);

function imgChange(obj1, obj2) {
    
    var file = document.getElementById("file");
  
    var imgContainer = document.getElementsByClassName(obj1)[0];

    var fileList = file.files;

    var input = document.getElementsByClassName(obj2)[0];
    var imgArr = [];

    for (var i = 0; i < fileList.length; i++) {
        var imgUrl = window.URL.createObjectURL(file.files[i]);
        imgArr.push(imgUrl);
        var img = document.createElement("img");
        img.setAttribute("src", imgArr[i]);
        var imgAdd = document.createElement("div");
        imgAdd.setAttribute("class", "z_addImg");
        imgAdd.appendChild(img);
        imgContainer.appendChild(imgAdd);
    };
    imgRemove(imgContainer);
};

function imgRemove(imgContainer) {
    var imgList = document.getElementsByClassName("z_addImg");
    
    var pmask = document.getElementsByClassName("p_mask")[0];
    var palert = document.getElementsByClassName("p_alert")[0];
    
    var mask = document.getElementsByClassName("z_mask")[0];
    var cancel = document.getElementsByClassName("z_cancel")[0];
    var sure = document.getElementsByClassName("z_sure")[0];
    for (var j = 0; j < imgList.length; j++) {
        imgList[j].index = j;
        imgList[j].onclick = function() {
        	debugger;
            var t = this;
            pmask.style.display = "block";
            var innerImg = t.innerHTML;
            palert.innerHTML = innerImg;
            palert.onclick = function(){
            	mask.style.display = "block";
	            cancel.onclick = function() {
	                mask.style.display = "none";
	                pmask.style.display = "none";
	            };
	            sure.onclick = function() {
	                mask.style.display = "none";
	                imgContainer.removeChild(t);
	                pmask.style.display = "none";
	            };
            }
            
        }
    };
};