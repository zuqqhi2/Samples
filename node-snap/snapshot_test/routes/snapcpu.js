var express = require('express');
var router = express.Router();

var fs = require('fs');
var profiler = require('nodegrind');

router.get('/', function(req, res) {
  // filename
  var name =  (new Date().getTime()) + '.cpuprofile';

  // Start to measure
  profiler.startCPU(name);

  // Stop to measure after 5 sec
  setTimeout(function(){
    var cpuProfile = profiler.stopCPU(name, 'cpuprofile');
    
    //Write profile
    fs.writeFile(name, cpuProfile, function(err){
      if (err) res.send(err);
      else     res.send('OK:' + name);
    });
  }, 5000);
});

module.exports = router;
