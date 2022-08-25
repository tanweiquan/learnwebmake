const btn = document.getElementsByTagName('button')[0];
const result =document.getElementById("result");
btn.onclick = function(){
    // 1.创建对象
    const xhr =new XMLHttpRequest();
    // 2.初始化，设置请求方法和url
    xhr.open('GET','http://127.0.0.1:8080/api/article_content');
    // 3.发送请求
    xhr.send();
    // 4.事件绑定 处理服务器返回的结构
    // on 当。。。的时候
    // readystate 是xhr中对象的属性，表示状态为0，1，2，3，4
    // 0：xhr未初始化        1：open刚初始化 2：已经发送 
    // 3：服务端返回部分结果  4：服务端返回所有结果
    // change 改变
xhr.onreadystatechange = function(){
    // 判断服务端是否返回所有结果
    if(xhr.readyState ===4){
        // 判断响应状态码 200 301 404 500...
        // 2xx：都是成功
        if(xhr.status >=200 &&  xhr.status<300){
            // 处理结果 行 头 空行 体
            // 1、响应行
            // console.log(xhr.status); // 响应状态码
            // console.log(xhr.statusText); // 响应状态字符串
            // console.log(xhr.getAllResponseHeaders()); // 所有响应头
            // console.log(xhr.response);//响应体

            // 设置result的文本
            result.innerHTML= xhr.response;

        }else{

        }
    }
}
}