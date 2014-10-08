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
        var received_msg = evt.data;
        console.log("WebSocket msg received: " + received_msg);
    };
    
    ws.onclose = function() {
        console.log("WebSocket connection closed.");
    };
    
    $("#msg_box_form").submit(function(evt) {
        var msg = $(this).find('[name=msg]').val();
        
        if (!msg) return false;
        
        ws.send(msg);
        console.log("Sent msg: "+ msg);
        return false;
    });
});
