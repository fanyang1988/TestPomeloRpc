// Generated by CoffeeScript 1.8.0
(function() {
  var Server, paths, port, server;

  Server = require("pomelo-rpc").server;

  paths = [
    {
      namespace: "user",
      path: __dirname + "/remote/test"
    }
  ];

  port = 3333;

  server = Server.create({
    paths: paths,
    port: port
  });

  server.start();

  console.log("rpc server started.");

}).call(this);
