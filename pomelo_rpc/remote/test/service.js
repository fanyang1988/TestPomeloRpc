// Generated by CoffeeScript 1.7.1
(function() {
  module.exports = function(context) {
    return {
      echo: function(msg, cb) {
        console.error("msg " + msg);
        cb(null, "echo: " + msg);
      }
    };
  };

}).call(this);
