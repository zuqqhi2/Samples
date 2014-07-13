# test/sample.test.coffee
should = (require 'chai').should()

describe 'String', ->
  describe '#concat()', ->
    it 'should return "John Doe"', ->
      # assert values
      'foo'.concat('bar').should.equal 'foobar'

  describe '#split()', ->
    it 'should return [1,2,3]', ->
      # assert arrays and objects
      'foo,bar,baz'.split(',').should.deep.equal ['foo', 'bar', 'baz']
