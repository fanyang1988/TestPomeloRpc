// Generated by CoffeeScript 1.7.1
var Client, client, context, records, routeContext, routeFunc, routeParam, servers;

Client = require("pomelo-rpc-zeromq").client;

records = [
  {
    namespace: "user",
    serverType: "test",
    path: __dirname + "/remote/test"
  }
];

context = {
  serverId: "test-server-1"
};

servers = [
  {
    id: "test-server-1",
    serverType: "test",
    host: "127.0.0.1",
    port: 3333
  }
];

routeParam = null;

routeContext = servers;

routeFunc = function(routeParam, msg, routeContext, cb) {
  cb(null, routeContext[0].id);
};

client = Client.create({
  routeContext: routeContext,
  router: routeFunc,
  context: context
});

client.start(function(err) {
  console.log("rpc client start ok.");
  client.addProxies(records);
  client.addServers(servers);
  setInterval(function() {
    client.proxies.user.test.service.echo(routeParam, "hello", function(err, resp) {
      if (err) {
        console.error("err stack " + err);
      }
      console.log("resp");
      console.log(resp);
    });
  }, 100);
});