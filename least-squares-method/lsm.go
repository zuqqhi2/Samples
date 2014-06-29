package main

import (
  "code.google.com/p/plotinum/plot"
  "code.google.com/p/plotinum/plotter"
  //"code.google.com/p/plotinum/vg"
  "image/color"
  "math/rand"
  "math"
  //"fmt"
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

/*
func gaussElimination(mat [][]float64, height, width int) []float64 {
  // Pivoting
  for x := 0; x < width; x++ {
    pivot := x

    for y := x; y < height; y++ {
      
    }
  }
}
*/

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

  // Prepare X axis of LSM result points
  N := 1000
  result := make(plotter.XYs, N)
  linspace(-3, 3, N, result)

  // Trigonometric Polynomial Basis Function for solving
  numParams := 31
  basisFunc := make([][]float64, n, n)
  for y := 0; y < n; y++ {
    basisFunc[y] = make([]float64, numParams, numParams)
    for x := 0; x <= (numParams-1)/2; x++ {
      if x == 0 {
        basisFunc[y][x] = 1.0
      } else {
        basisFunc[y][2*x-1] = math.Sin(float64(x)/2.0 * answer[y].X)
        basisFunc[y][2*x]   = math.Cos(float64(x)/2.0 * answer[y].X)
      }
    }
  }

  /*
  // Make Normal Equation
  equation := make([][]float64, numParams, numParams)
  for y := 0; y < numParams; y++ {
    equation[y] = make([]float64, numParams+1, numParams+1)
    for x := 0; x < numParams; x++ {
      summation := 0.0
      for k := 0; k < n; k++ {
        summation += basisFunc[k][y] * basisFunc[y][k]
      }
      equation[y][x] = summation
    }
  }

  // Solve y=phi*theta
  //theta := gaussElimination(gaussMat)
  */

  //====================================================
  // Graph Setting

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
  
  // Add data and legend
  p.Add(lpPoints)
  p.Legend.Add("answer", lpPoints)
  
  // Save the plot to a PNG file.
  if err := p.Save(4, 4, "points.png"); err != nil {
    panic(err)
  }


  /*
        // Get some random points
        rand.Seed(int64(0))
        n := 15
        scatterData := randomPoints(n)
        lineData := randomPoints(n)
        linePointsData := randomPoints(n)

        // Create a new plot, set its title and
        // axis labels.
        p, err := plot.New()
        if err != nil {
                panic(err)
        }
        p.Title.Text = "Points Example"
        p.X.Label.Text = "X"
        p.Y.Label.Text = "Y"
        // Draw a grid behind the data
        p.Add(plotter.NewGrid())

        // Make a scatter plotter and set its style.
        s, err := plotter.NewScatter(scatterData)
        if err != nil {
                panic(err)
        }
        s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

        // Make a line plotter and set its style.
        l, err := plotter.NewLine(lineData)
        if err != nil {
                panic(err)
        }
        l.LineStyle.Width = vg.Points(1)
        l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
        l.LineStyle.Color = color.RGBA{B: 255, A: 255}

        // Make a line plotter with points and set its style.
        lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
        if err != nil {
                panic(err)
        }
        lpLine.Color = color.RGBA{G: 255, A: 255}
        lpPoints.Shape = plot.PyramidGlyph{}
        lpPoints.Color = color.RGBA{R: 255, A: 255}

        // Add the plotters to the plot, with a legend
        // entry for each
        p.Add(s, l, lpLine, lpPoints)
        p.Legend.Add("scatter", s)
        p.Legend.Add("line", l)
        p.Legend.Add("line points", lpLine, lpPoints)

        // Save the plot to a PNG file.
        if err := p.Save(4, 4, "points2.png"); err != nil {
                panic(err)
        }
  */
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
        pts := make(plotter.XYs, n)
        for i := range pts {
                if i == 0 {
                        pts[i].X = rand.Float64()
                } else {
                        pts[i].X = pts[i-1].X + rand.Float64()
                }
                pts[i].Y = pts[i].X + 10*rand.Float64()
        }
        return pts
}
