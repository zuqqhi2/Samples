var express = require('express');
var router = express.Router();

var fs = require('fs');
var heapdump = require('heapdump');

router.get('/', function(req, res) {
  // filename
  var name = (new Date().getTime()) + '.heapsnapshot';

  // Take a snapshot of heap
  heapdump.writeSnapshot(name, function() {
    res.send('OK:' + name);
  });
});

module.exports = router;
