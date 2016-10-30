$(function(){
	init();
})

function init(){
	BindValue();
}

//发送请求获取首页的数据.
function BindValue(){
	$.get('/data?a=index',function(data){
		
	},'json')	
}