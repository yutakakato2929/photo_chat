var socket = null;
    var targetname = location.pathname.slice(6);
    // サーバーに接続する (一連の処理)
    window.onload = function () {
    socket = new WebSocket("/ws/" + targetname);
    console.log('connected at websocket')
    //通信が接続された時に実行される
    socket.onopen = function () {
        //append_message("system", "connect to server");
    };
    //サーバーからデータを受信した時に実行される
    socket.onmessage = function (event) {
        //サーバーからメッセージを受け取る
        var message = JSON.parse(event.data);
        append_message(message.username, message.message);
        window.scrollTo(0, document.body.scrollHeight);
    };
    };

    // メッセージ欄を更新する
    function append_message(username, message) {
        
        let li_name = document.createElement("li");
        if (username == targetname) {
            li_name.className = 'chat you';
        }else{
            li_name.className = 'chat me';
        }

        let p_message = document.createElement("p");
        p_message.innerHTML = message;
        p_message.className = 'msg';
        li_name.appendChild(p_message);
        let ul = document.getElementById("messages");
        ul.appendChild(li_name);
    }

    // サーバーにメッセージを送信する
    function send() {
        let send_msg = document.getElementById("message");
        let msg = send_msg.value;
        if (msg == "") {
            return;
        }
        var username = document.getElementById("username").innerHTML
        socket.send(
            JSON.stringify({
                username: username,
                targetname: targetname,
                message: msg
            })
        );
        send_msg.value = "";
    }