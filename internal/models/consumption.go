// package name
package models

// lirary imported
import (	
) 

// consumption user (ECOLYO)
type Consumption struct {
	Pseudo                    string
	Latitude                  float64
	Longitude				  float64
	AverageDailyConsumption   float64
	AverageMonthlyConsumption float64
	AverageAnnualConsumption  float64
}