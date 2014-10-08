$(function() {  // main()
    
    if (!"WebSocket" in window) {
        console.log("WebSocket is not supported by your browser. Please use a modern browser.");
        return;
    }
    
    var ws = new WebSocket("ws://localhost:8080/ws");
    $(window).on('beforeunload', function(){ // Defer socket.close() until page is closed or we browse away.
        ws.close();
    });
    
    ws.onopen = function()
    {
        console.log("WebSocket connection opened.");
    };
    
    ws.onmessage = function(evt) {
        var msg = evt.data;
        console.log("WebSocket msg received: " + msg);
        msg = $("<div/>").html(msg).text(); // Strip out html tags.
        
        var match = msg.search(":");
        if (match === -1) {
            var date = new Date();
            msg = date.toLocaleDateString() + " " + date.toLocaleTimeString() + " | " + "Broadcast Msg: " + msg;
        }
        
        $("#msg_box").append('<span class="msg">' + msg + '</span>'); // Add msg to msg box.
    };
    
    ws.onclose = function() {
        console.log("WebSocket connection closed.");
    };
    
    $("#msg_box_form").submit(function(evt) {   // On send new message.
        var date = new Date();
        var usr = $(this).find('[name=username]').val();
        var msg = $(this).find('[name=msg]').val();
        
        usr = $("<div/>").html(usr).text();
        msg = $("<div/>").html(msg).text(); // Strip out html tags.
        
        if (!msg || !usr) return false; // Return if username or msg box is empty.
        
        msg = date.toLocaleDateString() + " " + date.toLocaleTimeString() + " | " + usr + ": " + msg;
        ws.send(msg);   // Send the message over websocket.
        console.log("WebSocket msg sent: "+ msg);
        
        $("[name=msg]").val(''); // Clear new msg input box.
        return false;
    });
});
