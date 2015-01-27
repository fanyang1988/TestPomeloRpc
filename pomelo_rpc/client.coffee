Client = require("pomelo-rpc-zeromq").client

# remote service interface path info list
records = [
  namespace: "user"
  serverType: "test"
  path: __dirname + "/remote/test"
]
context = serverId: "test-server-1"

# server info list
servers = [
  id: "test-server-1"
  serverType: "test"
  host: "127.0.0.1"
  port: 3333
]

# route parameter passed to route function
routeParam = null

# route context passed to route function
routeContext = servers

# route function to caculate the remote server id
routeFunc = (routeParam, msg, routeContext, cb) ->
  cb null, routeContext[0].id
  return

client = Client.create(
  routeContext: routeContext
  router:       routeFunc
  context:      context
)
client.start (err) ->
  console.log "rpc client start ok."
  client.addProxies records
  client.addServers servers

  setInterval () ->
      client.proxies.user.test.service.echo routeParam, "hello", (err, resp) ->
          console.error "err stack " + err  if err
          console.log "resp"
          console.log resp
          return
      return
    , 100


  return

