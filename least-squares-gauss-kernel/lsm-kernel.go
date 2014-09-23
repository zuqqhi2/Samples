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

// Make random value with normal distribution by Box-Muller Transform
func normalRand(mu, sigma float64) float64 {
  z := math.Sqrt(-2.0 * math.Log(rand.Float64())) * math.Sin(2.0 * math.Pi * rand.Float64())
  return sigma*z + mu
}

// Make sequence of numbers with common difference
func linspace(start, end float64, n int, x plotter.XYs) {
  for i := 0; i < n; i++ {
    t := float64(i) / float64(n-1)
    x[i].X = (1.0 - t) * start + t * end 
  }
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
  // LSM with Gauss Kernel Model

  band_width := 0.3
  step_size  := 0.1
  t0 := make([]float64, n);
  for i := 0; i < n; i++ {
    t0[i] = rand.Float64()
  }

  for step := 0; step < n*1000; step++ {
    idx := rand.Intn(n)
    kernel := make([]float64, n)
    t := make([]float64, n)
    sumError := 0.0
    for i := 0; i < n; i++ {
      kernel[i] = math.Exp(-math.Pow(answer[i].X - answer[idx].X, 2.0) / (2.0 * band_width * band_width))
      t[i] = t0[i] - step_size * kernel[i] * (kernel[i] * t0[i] - answer[i].Y)
      sumError += (t[i] - t0[i]) * (t[i] - t0[i])
      t0[i] = t[i]
    }
    fmt.Printf(" Step %d, RSS %f\n", step, sumError)
    if sumError < 0.000001 {
      break
    }
  }

  
  // Calculation Error
  sumError := 0.0
  for i := 0; i < n; i++ {
    y := t0[i]
    sumError += math.Pow(answer[i].Y - y, 2.0)
  }
  fmt.Printf("\nRSS(Training Data) = %f\n", sumError)
  
  sumError = 0.0
  for i := 0; i < n; i++ {
    // New Data
    x := rand.Float64() * 6.0 - 3.0
    ansY := math.Sin(math.Pi * x) / (math.Pi * x) + 0.1 * x + normalRand(1.0, 0.05)

    // Estimation
    y := 0.0
    for k := 0; k < n; k++ {
      kernel := math.Exp(-math.Pow(x - answer[k].X, 2.0) / (2.0 * band_width * band_width))
      y += kernel * t0[k] 
    }

    // Error
    sumError += math.Pow(ansY - y, 2.0)
  }
  fmt.Printf("\nRSS(New Data)      = %f\n", sumError)
  

  //====================================================
  // Graph Setting

  // Make result data
  N := 1000
  result := make(plotter.XYs, N)
  linspace(-3, 3, N, result)
  
  for i := 0; i < N; i++ {
    result[i].Y = 0.0
    for j := 0; j < n; j++ {
      kernel := math.Exp(-(result[i].X - answer[j].X) * (result[i].X - answer[j].X) / (2.0 * band_width * band_width))
      result[i].Y += kernel * t0[j] 
    }
  }

  // Create a new plot, set its title and axis labels
  p, err := plot.New()
  if err != nil {
    panic(err)
  }
  p.Title.Text = "LSM with Kernel"
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
  filename := "lsm-kernel" + strconv.Itoa(n) + ".png"
  if err := p.Save(4, 4, filename); err != nil {
    panic(err)
  }
}
