Server = require("pomelo-rpc").server

# remote service path info list
paths = [
  namespace: "user"
  path: __dirname + "/remote/test"
]
port = 3333
server = Server.create(
  paths: paths
  port: port
)
server.start()
console.log "rpc server started."
