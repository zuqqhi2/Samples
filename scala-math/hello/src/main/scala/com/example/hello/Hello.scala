package com.example.hello

import breeze.linalg._

object Hello {
  def main(args: Array[String]) {
    println("Hello!")

    // #1. 2 x 3行列を作成
    //      1.0 2.0 3.0
    //      4.0 5.0 6.0
    val m23 = DenseMatrix(
      (1.0d, 2.0d, 3.0d),
      (4.0d, 5.0d, 6.0d))
 
    // #2. 3 x 2のゼロ埋め行列を作成
    //      0 0
    //      0 0
    //      0 0
    val m32z = DenseMatrix.zeros[Double](3, 2)
 
    // #3. 2 x 2の1埋め行列を作成
    //      1 1
    //      1 1
    val m22o = DenseMatrix.zeros[Double](2, 2)
 
    // #4. 2 x 3の乱数行列を作成
    //      0.8765... 0.23909... 0.61324...
    //      0.51976... 0.80524... 0.44318...
    val m23r = DenseMatrix.rand(2, 3)
    val plusall = m23 + 0.1d

    println(m23r)
    println(plusall)
  }
}
