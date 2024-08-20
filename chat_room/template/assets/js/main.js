const LEFT = "left";
const RIGHT = "right";

const EVENT_MESSAGE = "message"
const EVENT_OTHER = "other"

const userPhotos = [
    "https://cdn-icons-png.flaticon.com/512/4825/4825038.png",
    "https://cdn-icons-png.flaticon.com/512/4825/4825112.png",
    "https://cdn-icons-png.flaticon.com/512/4825/4825015.png",
    "https://cdn-icons-png.flaticon.com/512/4825/4825044.png",
    "https://cdn-icons-png.flaticon.com/512/4825/4825082.png",
    "https://cdn-icons-png.flaticon.com/512/4825/4825087.png",
]
var PERSON_IMG = userPhotos[getRandomNum(0, userPhotos.length - 1)];
var PERSON_NAME = "Guest" + Math.floor(Math.random() * 1000); //隨機分配名稱

var url = "ws://" + window.location.host + "/ws?id=" + PERSON_NAME; // ws://localhost:8080/ws?id=Guest123
var ws = new WebSocket(url); // 建立 WebSocket 連線，並指定 URL
var name = "Guest" + Math.floor(Math.random() * 1000);
var chatroom = document.getElementsByClassName("msger-chat")
var text = document.getElementById("msg"); // 讀取使用者輸入框
var send = document.getElementById("send")

// send 按鈕按下 or 在輸入框按下 enter => 觸發 handleMessageEvent 送出訊息
send.onclick = function (e) {
    handleMessageEvent()
}

text.onkeydown = function (e) {
    if (e.keyCode === 13 && text.value !== "") { // 如果是 enter 且輸入框有值
        handleMessageEvent()
    }
};

function handleMessageEvent() {
    // encode html tag
    content = text.value.replace(/</g, "&lt;").replace(/>/g, "&gt;"); //將 text 元素中的 '<' 與 '>' 符號替換為 '&lt;' 與 '&gt;'，避免 HTML injection

    if (text.value != "") { //如果 text(輸入框) 有值
        ws.send(JSON.stringify({ // 透過 WebSocket 傳送訊息發送訊息
            "event": "message",
            "photo": PERSON_IMG,
            "name": PERSON_NAME,
            "content": content,
        }));
    }
    text.value = ""; // 清空輸入框
}


// 接收來自後端 websocket 的訊息
// 後端會在訊息中告知誰發送的訊息，以及訊息內容
ws.onmessage = function (e) {
    var m = JSON.parse(e.data) 
    var msg = ""
    switch (m.event) {
        case EVENT_MESSAGE:
            if (m.name == PERSON_NAME) { // 如果是自己發送的訊息 => 右邊
                msg = getMessage(m.name, m.photo, RIGHT, m.content);
            } else { // 如果是別人發送的訊息 => 左邊
                msg = getMessage(m.name, m.photo, LEFT, m.content);
            }
            break;
        case EVENT_OTHER: // 其他事件，例如有其他人進入 or 離開聊天室
            if (m.name != PERSON_NAME) {
                msg = getEventMessage(m.name + " " + m.content)
            } else {
                msg = getDateMessage(formatDate(m.timestamp))
                msg += getEventMessage("您已" + m.content)
            }
            break;
    }
    insertMsg(msg, chatroom[0]); //放在chatroom中這個陣列中
};

ws.onclose = function (e) {
    console.log(e)
}



function getEventMessage(msg) { // 在HTML中插入一則事件訊息
    var msg = `<div class="msg-notify">${msg}</div>`
    return msg
}

function getDateMessage(msg) { // 在HTML中插入一則日期訊息
    var msg = `<div class="msg-date"><span class="time-tag">${msg}</span></div>`
    return msg
}

function formatDate(d) {
    return d.split('T')[0];
}

function formatTime(d) {
    return d.toLocaleString('zh-TW', {
        timeZone: 'Asia/Taipei',
        hour12: true,
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
    }).replaceAll("/", "-");
}

function getMessage(name, img, side, text) { // 在HTML中插入一則"訊息框框"
    const d = new Date();
    //   Simple solution for small apps
    var msg = `
    <div class="msg ${side}-msg">
      <div class="msg-img" style="background-image: url(${img})"></div>

      <div class="msg-bubble">
        <div class="msg-info">
          <div class="msg-info-name">${name}</div>
          <div class="msg-info-time">${formatTime(d)}</div>
        </div>

        <div class="msg-text">${text}</div>
      </div>
    </div>
  `
    return msg;
}

function insertMsg(msg, domObj) {
    domObj.insertAdjacentHTML("beforeend", msg);
    domObj.scrollTop += 500;
}

function getRandomNum(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}