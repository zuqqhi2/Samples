package com.example.hello

import org.scalatest.FunSuite

class HelloTest extends FunSuite {
  test("Hello should run main") {
    expectResult()(Hello.main(Array.empty))
  }
}
