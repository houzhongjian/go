$(function(){
	init();
})

function init(){
	register();
}

function register(){
	$('#button').on('click',function(){
		$('input').removeClass('border_color')
		$.post('/register/register', $('form').serialize(), function(data){
			if (data.status != true) {
				$.each(data, function(k, v){
					if (v != ""){						
						$('#'+k).addClass('border_color')
					}
				})
			} else {
				//提示层
				layer.msg('注册成功');
				setTimeout(function(){					
					window.location.href="/login"
				},2000)
			}
		},'JSON')
	})
}