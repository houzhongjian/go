$(function(){
	button();
})

function button() {
	
	//当点击了登录页面的登录按钮.
	$('#button').on('click', function(){
		$.post('/login?a=login', $('form').serialize(), function(data){
			
		},'json')
	})
}