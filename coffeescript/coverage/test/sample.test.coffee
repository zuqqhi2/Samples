main   = require('../src/');
assert = require('chai').assert

describe 'testing', ->

  it 'is pretty nice with CoffeeScript', ->
    assert main() == true
