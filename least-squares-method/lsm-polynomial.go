package main

import (
  "code.google.com/p/plotinum/plot"
  "code.google.com/p/plotinum/plotter"
  //"code.google.com/p/plotinum/vg"
  "image/color"
  "math/rand"
  "math"
  "strconv"
  "fmt"
)

func normalRand(mu, sigma float64) float64 {
  z := math.Sqrt(-2.0 * math.Log(rand.Float64())) * math.Sin(2.0 * math.Pi * rand.Float64())
  return sigma*z + mu
}

func linspace(start, end float64, n int, x plotter.XYs) {
  for i := 0; i < n; i++ {
    t := float64(i) / float64(n-1)
    x[i].X = (1.0 - t) * start + t * end 
  }
}


func gaussElimination(mat [][]float64, numParams int) []float64 {
  // Pivoting
  for x := 0; x < numParams; x++ {
    pivot := x
    maxVal := 0.0
    for y := x; y < numParams; y++ {
      if math.Abs(mat[y][x]) > maxVal {
        maxVal = math.Abs(mat[y][x])
        pivot = y
      }
    }

    // switch column when pivot != x
    if pivot != x {
      for k := 0; k < numParams+1; k++ {
        tmp           := mat[x][k]
        mat[x][k]     =  mat[pivot][k]
        mat[pivot][k] =  tmp
      }
    }
  }

  // Make upper triangular matrix with basic deformation of matrix
  for k := 0; k < numParams; k++ {
    // Save diagonal element
    diagElem := mat[k][k]
    // Diagonal element must be 1.0
    mat[k][k] = 1.0
    // Basic deformation
    for x := k+1; x < numParams+1; x++ { mat[k][x] /= diagElem  }
    //
    for y := k+1; y < numParams; y++ {
      tmp := mat[y][k]
      for x := k+1; x < numParams+1; x++ {
        mat[y][x] -= tmp*mat[k][x]
      }
      mat[y][k] = 0.0
    }
  }

  // Solve
  result := make([]float64, numParams, numParams)
  for y := numParams-1; y >= 0; y-- {
    result[y] = mat[y][numParams]
    for x := numParams-1; x > y; x-- {
      result[y] -= mat[y][x] * result[x]
    }
  }

  return result
}


func main() {
  //===================================================
  // Make observed points

  rand.Seed(int64(0))

  // Prepare X axis of observed points
  n := 50
  answer := make(plotter.XYs, n)
  linspace(-3, 3, n, answer)

  // make observed points
  pix := make([]float64, n)
  for i := 0; i < n; i++ {
    pix[i] = math.Pi * answer[i].X
  }
  for i := 0; i < n; i++ {
    answer[i].Y = math.Sin(pix[i]) / pix[i] + 0.1 * answer[i].X + normalRand(1.0, 0.05)
  }

  //====================================================
  // LSM

  // Make a matrix to solve
  numParams := 9
  mat := make([][]float64, numParams, numParams)
  fmt.Printf("[Matrix for Gausss Elimination]\n")
  for y := 0; y < numParams; y++ {
    mat[y] = make([]float64, numParams+1, numParams+1)
    for x := 0; x < numParams+1; x++ {
      mat[y][x] = 0.0
      if x < numParams {
        for k := 0; k < n; k++ {
          mat[y][x] += math.Pow(answer[k].X, float64(y+x))
        }
      } else {
        for k := 0; k < n; k++ {
          mat[y][x] += math.Pow(answer[k].X, float64(y)) * answer[k].Y
        }
      }
      fmt.Printf("%f\t", mat[y][x])
    }
    fmt.Printf("\n")
  }
  fmt.Printf("\n")

  // Solve y=phi*theta
  params := gaussElimination(mat, numParams)
  
  // Output
  fmt.Printf("[Optimized Params]\n")
  for i := 0; i < numParams; i++ {
    fmt.Printf("param %d : %f\n", i, params[i])
  }

  //====================================================
  // Graph Setting

  // Result graph
  N := 1000
  result := make(plotter.XYs, N)
  linspace(-3, 3, N, result)
  for i := 0; i < N; i++ {
    result[i].Y = 0.0
    for k := 0; k < numParams; k++ {
      result[i].Y += params[k]*math.Pow(result[i].X, float64(k))
    }
  }

  // Create a new plot, set its title and axis labels
  p, err := plot.New()
  if err != nil {
    panic(err)
  }
  p.Title.Text = "Least Square Method"
  p.X.Label.Text = "X"
  p.Y.Label.Text = "Y"
  p.Add(plotter.NewGrid())

  // Make a scatter plotter and set its style
  // Make a line plotter with points and set its style.
  lpLine, lpPoints, err := plotter.NewLinePoints(answer)
  if err != nil {
    panic(err)
  }
  lpLine.Color = color.RGBA{G: 255, A: 255}
  lpPoints.Shape = plot.PyramidGlyph{}
  lpPoints.Color = color.RGBA{R: 255, A: 255}
 
  // Result line
  lpResultLine, _, err := plotter.NewLinePoints(result)
  if err != nil {
    panic(err)
  }
  lpResultLine.Color = color.RGBA{B: 255, A: 255}
 
  // Add data and legend
  p.Add(lpPoints, lpResultLine)
  p.Legend.Add("Observed Points", lpPoints)
  p.Legend.Add("Result", lpResultLine)
  
  // Save the plot to a PNG file.
  filename := "lsm-polynomial" + strconv.Itoa(numParams) + ".png"
  if err := p.Save(4, 4, filename); err != nil {
    panic(err)
  }
}
