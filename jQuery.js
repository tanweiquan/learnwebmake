// get和post一样写:
$('button').eq(0).click(function(){
    
    $.get('http://127.0.0.1:8080/api/article_content',{/* 参数k1:v1,k2:v2 */},function(data){
        console.log(data);
    },'json')// 这里json代表响应体是json
})
// jquery通用方法：
$('button').eq(1).click(function(){
    // get和post一样写
    $.ajax({
        // url
        url:'http://127.0.0.1:8080/api/article_content',
        // 参数 k:v形式
        data:{/* 参数k1:v1,k2:v2 */},
        // 请求方法：
        type: 'GET',
        // 超时：
        timeout: 2000,
        // 请求头信息：
       /*  headers:{
            k1:1,
            k2:2,
        }, */
        // 响应体结果类型：
        dataType:'json',
        error: function(){
            console.log('出差了');
        },
        // 成功回调：
        success: function(data){
            console.log(data);
        },
    })
})
// jsonp在jquery中的使用
$('button').eq(2).click(function(){
    $.getJSON("https://www.runoob.com/try/ajax/jsonp.php?jsoncallback=?", function(data) { // url?callback=?
        console.log(data);
    });
})
$('button').eq(3).click(function(){
    $.ajax({
        url: 'https://www.runoob.com/try/ajax/jsonp.php?jsoncallback=?',
        dataType: 'jsonp',
        jsonp: 'jsoncallback', //自定义参数的名称  一般让它默认为callback，不会做修改，当与后端不一样时要修改
        jsonpCallback: '自己定义不影响获取',   //自定义回调函数的名称
        success: function (res) {
        console.log(res);
    }
    })
})