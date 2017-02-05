function http(url){
        var xhttp = new XMLHttpRequest();
        var method = "GET";
        var sendBody = null;
        var jsonHeader = false;
        function get(params){
            method = "GET";
            if(params != undefined){
                url += "?" + params;
            }
            return {send:send};
        }
        
        function post(data){
            method = "POST";
            if(data != undefined){
                jsonHeader = true;
                sendBody = JSON.stringify(data);
            }
            return {send:send};
        }
        
        function send(status, success, fail){
            xhttp.open(method, url, true);
            if(jsonHeader){
                 xhttp.setRequestHeader('Content-type', 'application/json; charset=utf-8');
            }
            xhttp.onreadystatechange = function() {
                if (this.readyState == 4 ) {
                    if(status == this.status){
                         if(success) {success(this);}
                    }else{
                        if(fail) {fail(this);}
                    }
                }
            };
            xhttp.send(sendBody);
        }
        
        return {
            get:get,
            post:post,
            send:send
        }
    }