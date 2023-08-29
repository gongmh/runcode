package service

import "fmt"

type WeightInfo struct {
	Date          []string
	MaxWeight     []float64
	MinWeight     []float64
	CurrentWeight []float64
}

func GetChildWeightInfo() WeightInfo {
	var weightInfo WeightInfo
	for idx, info := range weightList {
		weightInfo.CurrentWeight = append(weightInfo.CurrentWeight, info.CurrentWeight)
		weightInfo.MaxWeight = append(weightInfo.MaxWeight, info.StandMaxWeight*2)
		weightInfo.MinWeight = append(weightInfo.MinWeight, info.StandMinWeight*2)
		weightInfo.Date = append(weightInfo.Date, fmt.Sprintf("%s(%d天)", info.Date, idx))
	}
	return weightInfo
}

var weightList = []weightConf{
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.82, Date: "2023-08-15"},
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.70, Date: "2023-08-16"},
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.60, Date: "2023-08-17"},
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.40, Date: "2023-08-18"},
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.30, Date: "2023-08-19"},
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.20, Date: "2023-08-20"},
	{StandMaxWeight: 4.2, StandMinWeight: 2.6, CurrentWeight: 5.40, Date: "2023-08-21"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 5.40, Date: "2023-08-22"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 5.50, Date: "2023-08-23"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 5.70, Date: "2023-08-24"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 6.00, Date: "2023-08-25"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 5.80, Date: "2023-08-26"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 5.70, Date: "2023-08-27"},
	{StandMaxWeight: 4.6, StandMinWeight: 2.6, CurrentWeight: 5.70, Date: "2023-08-28"},
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.10, Date: "2023-08-29"},
}

type weightConf struct {
	StandMaxWeight float64
	StandMinWeight float64
	CurrentWeight  float64
	Date           string
}
