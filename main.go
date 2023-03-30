package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

func main() {
	// Open the CSV file.
	carPrice, err := os.Open("car_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer carPrice.Close()
	// Create a dataframe from the CSV file.
	carDF := dataframe.ReadCSV(carPrice)
	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	carsSummary := carDF.Describe()
	// Output the summary statist
	fmt.Println(carsSummary)

	fmt.Println("Complete Dataframe:")
	fmt.Println(carDF)

	// Create a new regression instance.
	r := new(regression.Regression)

	// Add the independent variables to the regression.
	for _, colName := range carDF.Names() {
		if colName == "Selling_Price" {
			continue
		}
		colData := carDF.Col(colName).Float()
		dp := make([]regression.DataPoint, len(colData))
		for i, f := range colData {
			dp[i].Observed = f
			dp[i].Variables = []float64{carDF.Col("Age").Float()[i]}
		}
		r.Train(dp)
	}

	// Add the dependent variable to the regression.
	sellingPrice := carDF.Col("Selling_Price").Float()
	dp := make([]regression.DataPoint, len(sellingPrice))
	for i, f := range sellingPrice {
		dp[i].Observed = f
	}
	r.Train(dp)

	// Fit the regression using the normal equation.
	r.Run()

	// Output the regression formula.
	formula := "Selling_Price = "
	for i, coef := range r.Weights {
		if i == 0 {
			formula += fmt.Sprintf("%.2f", coef)
		} else {
			formula += fmt.Sprintf(" + %.2f * %s", coef, carDF.Names()[i-1])
		}
	}
	fmt.Println("Regression Formula:", formula)

	// Output the regression statistics.
	fmt.Println("R²:", r.R2)
	fmt.Println("Coefficients:", r.Weights)
	fmt.Println("Intercept:", r.Intercept)

	fmt.Println("Complete Dataframe:")
	fmt.Println(carDF)
}
