$(function(){
	nav();
	click();
})

function nav(){
	$('.nav_li').hover(function(){
		$(this).css('background', '#232c32')
		$(this).css('cursor', 'pointer')
	}, function(){
		$(this).css('background', '#5f6d7e')
	})
}

function click() {
	$('.nav_li').on('click', function(){
		var path = $(this).attr('path')
		console.debug(path)
		window.location.href="http://"+path
	})
}