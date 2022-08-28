const btn = document.getElementsByTagName('button')[0];
const result =document.getElementById("result");
btn.onclick = function(){
    // 1.创建对象
    const xhr =new XMLHttpRequest();
    // 特别的，设置超时时间：
    xhr.timeout =2000;// 超时不超过2s
    // 写超时回调函数
    xhr.ontimeout = function(){
        alert("网络超时，请重试")
    };
    // 写网络异常回调函数
    xhr.onerror = function(){
        alert('网络异常，请重试')
    };
    // 2.初始化，设置请求方法和url，如果是post请求，则写'POST'
    xhr.open('GET','http://127.0.0.1:8080/api/article_content'); // IE浏览器缓存问题可以使用 'url?t='+ Date.now()
    // 特别的，设置请求体：xhr.setRequestHeader('key','value');
    //  如：xhr.setRequestHeader('Content-Type','application/x-www-form-urlencoded')
    // 3.发送请求
    xhr.send(); // 如果是POST请求则写：xhr.send('参数');
    // 取消请求（避免重复请求可以通过取消上次请求来做：判断状态是否小于4）：调用 xhr.abort();
    // 4.事件绑定 处理服务器返回的结构
    // on 当。。。的时候
    // readystate 是xhr中对象的属性，表示状态为0，1，2，3，4
    // 0：xhr未初始化        1：open刚初始化 2：已经发送 
    // 3：服务端返回部分结果  4：服务端返回所有结果
    // change 改变
    let articledata;
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
            // 手动转换：解析请求获得的json数据
           let obj = JSON.parse(xhr.response);
           articledata = obj.data.article[0].article;
           // 自动转换：返回数据时可以使用 xhr.responseType = 'json';

        }else{
            articledata ="发生错误"
            
        }
    }
    result.innerText= articledata;
}
}
