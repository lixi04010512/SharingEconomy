let canSend = false;
$(function() {
    $('#footer').on('keyup', 'input', function() {
        if ($(this).val().length > 0) {
            $(this).next().css('background', '#114F8E').prop('disabled', true);
            canSend = true;
        } else {
            $(this).next().css('background', '#ddd').prop('disabled', false);
            canSend = false;
        }
    })
    $('#footer .send').click(send)
    $("#footer .my-input").keydown(function(e){
        if(e.keyCode == 13){
            return send();
        }
    });
})



function reply(headSrc, str) {
    var html = "<div class='reply'><div class='dropdown d-inline-block'><a class=' dropdown-toggle arrow-none' id='dLabel8' data-bs-toggle='dropdown' href='#' role='button' aria-haspopup='false' aria-expanded='false'><div class='reply'><div class='msg'><img src=" + headSrc + " /><p><i class='msg_input'></i>" + str + "</p></div></div></a><div class='dropdown-menu dropdown-menu-right' aria-labelledby='dLabel8'><a class='dropdown-item' href='#'>删除</a></div></div></div>";
    return upView(html);
}
function ask(headSrc, str) {
    var html = "<div class='reply' ><div class='dropdown d-inline-block'><a class=' dropdown-toggle arrow-none' id='dLabel8' data-bs-toggle='dropdown' href='#' role='button' aria-haspopup='false' aria-expanded='false'><div class='ask'><div class='msg'><img src=" + headSrc + " />" + "<p><i class='msg_input'></i>" + str + "</p></div></div></a><div class='dropdown-menu dropdown-menu-right' aria-labelledby='dLabel8'><a class='dropdown-item' href='#'>删除</a></div></div></div>";
    return upView(html);
}
function upView(html) {
    let message = $('#message');
    message.append(html);
    return $('html,body').animate({
        scrollTop: message.outerHeight() - window.innerHeight
    }, 200);
}
function send(msg){
    if(canSend){
        let input = $("#footer .my-input");
        ask("assets/img/profile.jpg", input.val());
        input.val('');
        test();
    }
}
function sj() {
    return parseInt(Math.random() * 10)
}
