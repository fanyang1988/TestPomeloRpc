# remote service
module.exports = (context) ->
  echo: (msg, cb) ->
    console.error "msg " + msg

    cb null, "echo: " + msg
    return  


    
