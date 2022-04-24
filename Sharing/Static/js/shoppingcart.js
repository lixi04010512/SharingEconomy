$(function(){
	//加的效果
	$(".product-add").click(function(){
		var n=$(this).prev().val();
		var num=parseInt(n)+1;
		if(num==99){ return;}
			$(this).prev().val(num);
		TotalPrice();
	});
	//减的效果
	$(".product-jian").click(function(){
		var n=$(this).next().val();
		var num=parseInt(n)-1;
		if(num==0){ return;}
			$(this).next().val(num);
		TotalPrice();
	});
	
	$(".product-ckb").click(function(){
		$(this).children("em").toggleClass("product-xz");
		TotalPrice();
		productxz();
	});

//计算产品价格
function TotalPrice(){
	//总价
	var total = 0;
	$(".product-em").each(function(){
		
		if($(this).is(".product-xz")){
			var price = parseInt($(this).parents(".product-ckb").siblings().find(".price").text());//得到产品单价
			var slproice = parseInt($(this).parents(".product-ckb").siblings().find(".product-num").val());//得到产品数量
			var rent = parseInt($(this).parents(".product-ckb").siblings().find(".rent").text());
			var dgtotal = price * slproice + rent;
			total+=dgtotal;
		}
		$(".all-price").text(total.toFixed(2)); //输出全部总价
	});
	
}

	
	TotalPrice();
});

