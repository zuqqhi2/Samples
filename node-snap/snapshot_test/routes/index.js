var express = require('express');
var router = express.Router();

var request = require('request');
var cheerio = require('cheerio');

/* GET home page. */
router.get('/', function(req, res) {
  var url = 'http://zuqqhi2.com/';
  var result = { status : 'error', contents : '' };

  request(url, function (error, response, body) {
    // Connection success
    if (!error && response.statusCode == 200) {
      // Parse
      var $ = cheerio.load(body);
      var hrefs = [];
     
      // Retrieve all link urls in the page
      $('a').each(function() {
        hrefs.push($(this).attr('href'));
      });

      // return response
      result.status = 'success';
      result.contents = hrefs;
      res.send(JSON.stringify(result));
    // Error result
    } else {
      res.send(JSON.stringify(result));
    }
  })
});

module.exports = router;
