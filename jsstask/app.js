function AddingElement(st) {
	var prop = $('<li><span></span> <button class="delete-but">Удалить</button></li>');
	$('span',prop).text(st);
	$('#root ul').append(prop);
	$('.delete-but', el).click(function(deleete) {
		$(this).parent().remove() });
}
$(function(){
	$('#root').append('<ul></ul>');
	$('#root').append('<input id="add_task_input">');
	$('#root').append('<button id ="add_task">Добавить</button>');
	$('#add_task').click(function(){
	AddingElement($('#add_task_input').val())});
	AddingElement("Сделать задание #3 по web-программированию");
})   
