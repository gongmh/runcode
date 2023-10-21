package service

import (
	"fmt"
	"time"
)

var eventMap = []EventInfo{
	{Date: "2023-08-15", EventDesc: "出生"},
	{Date: "2023-08-16", EventDesc: "接种疫苗：卡介苗、乙肝第一针"},
	{Date: "2023-08-18", EventDesc: "胎便排完"},
	{Date: "2023-08-19", EventDesc: "出院"},
	{Date: "2023-09-09", EventDesc: "回家后第一次洗澡"},
	{Date: "2023-09-11", EventDesc: "脐带脱落，办理出生证明"},
	{Date: "2023-09-18", EventDesc: "第一罐奶粉吃完"},
	{Date: "2023-09-19", EventDesc: "满月体检打疫苗，因湿疹严重未打疫苗"},
	{Date: "2023-09-24", EventDesc: "42天儿保，指标正常"},
	{Date: "2023-10-10", EventDesc: "接种疫苗：乙肝第二针"},
}

var lengthList = []lengthConf{
	{StandMaxLength: 54.0, StandMinLength: 46.9, CurrentLength: 50.00, Date: "2023-08-15"},
	{StandMaxLength: 59.0, StandMinLength: 50.7, CurrentLength: 56.00, Date: "2023-09-14"},
	{StandMaxLength: 60.0, StandMinLength: 51.7, CurrentLength: 56.00, Date: "2023-09-19"},
	{StandMaxLength: 61.0, StandMinLength: 52.6, CurrentLength: 56.00, Date: "2023-09-26"},
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
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.00, Date: "2023-08-30"},
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.10, Date: "2023-08-31"},
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.20, Date: "2023-09-01"},
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.30, Date: "2023-09-02"},
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.10, Date: "2023-09-03"},
	{StandMaxWeight: 5.0, StandMinWeight: 2.9, CurrentWeight: 6.30, Date: "2023-09-04"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 6.50, Date: "2023-09-05"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 6.50, Date: "2023-09-06"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 6.90, Date: "2023-09-07"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 7.00, Date: "2023-09-08"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 7.10, Date: "2023-09-09"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 7.10, Date: "2023-09-10"},
	{StandMaxWeight: 5.4, StandMinWeight: 3.2, CurrentWeight: 7.70, Date: "2023-09-11"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 7.30, Date: "2023-09-12"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 7.40, Date: "2023-09-13"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 7.70, Date: "2023-09-14"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 7.60, Date: "2023-09-15"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 7.90, Date: "2023-09-16"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 8.00, Date: "2023-09-17"},
	{StandMaxWeight: 5.7, StandMinWeight: 3.5, CurrentWeight: 8.00, Date: "2023-09-18"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.20, Date: "2023-09-19"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.20, Date: "2023-09-20"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.50, Date: "2023-09-21"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.50, Date: "2023-09-22"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.60, Date: "2023-09-23"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.70, Date: "2023-09-24"},
	{StandMaxWeight: 6.1, StandMinWeight: 3.7, CurrentWeight: 8.90, Date: "2023-09-25"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.00, Date: "2023-09-26"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.20, Date: "2023-09-27"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.10, Date: "2023-09-28"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.30, Date: "2023-09-29"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.40, Date: "2023-09-30"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.60, Date: "2023-10-01"},
	{StandMaxWeight: 6.5, StandMinWeight: 3.9, CurrentWeight: 9.70, Date: "2023-10-02"},
	{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.2, Date: "2023-10-03"},
	{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.2, Date: "2023-10-04"},
	{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.1, Date: "2023-10-05"},
	{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.2, Date: "2023-10-06"},
	{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.4, Date: "2023-10-07"},
	{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.2, Date: "2023-10-08"},
	//{StandMaxWeight: 6.9, StandMinWeight: 4.2, CurrentWeight: 10.2, Date: "2023-10-09"},
	{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.4, Date: "2023-10-10"},
	{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.7, Date: "2023-10-11"},
	{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.8, Date: "2023-10-12"},
	{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.8, Date: "2023-10-13"},
	//{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.2, Date: "2023-10-14"},
	{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.9, Date: "2023-10-15"},
	//{StandMaxWeight: 7.1, StandMinWeight: 4.5, CurrentWeight: 10.2, Date: "2023-10-16"},
	{StandMaxWeight: 7.4, StandMinWeight: 4.7, CurrentWeight: 10.9, Date: "2023-10-17"},
	{StandMaxWeight: 7.4, StandMinWeight: 4.7, CurrentWeight: 11.4, Date: "2023-10-18"},
	//{StandMaxWeight: 7.4, StandMinWeight: 4.7, CurrentWeight: 11.4, Date: "2023-10-19"},
	{StandMaxWeight: 7.4, StandMinWeight: 4.7, CurrentWeight: 11.0, Date: "2023-10-20"},
	{StandMaxWeight: 7.4, StandMinWeight: 4.7, CurrentWeight: 11.6, Date: "2023-10-21"},
}

type LengthInfo struct {
	Date          []string
	MaxLength     []float64
	MinLength     []float64
	CurrentLength []float64
}

func GetChildLengthInfo() LengthInfo {
	var lengthInfo LengthInfo

	birthDate, _ := time.Parse("2006-01-02", "2023-08-15")
	for _, info := range lengthList {
		lengthInfo.CurrentLength = append(lengthInfo.CurrentLength, info.CurrentLength)
		lengthInfo.MaxLength = append(lengthInfo.MaxLength, info.StandMaxLength)
		lengthInfo.MinLength = append(lengthInfo.MinLength, info.StandMinLength)
		tmpData, _ := time.Parse("2006-01-02", info.Date)
		lengthInfo.Date = append(lengthInfo.Date, fmt.Sprintf("%s(%d天)", info.Date, int64(tmpData.Sub(birthDate).Hours())/24+1))
	}

	return lengthInfo
}

type WeightInfo struct {
	Date          []string
	MaxWeight     []float64
	MinWeight     []float64
	CurrentWeight []float64
}

func GetChildWeightInfo() WeightInfo {
	var weightInfo WeightInfo
	birthDate, _ := time.Parse("2006-01-02", "2023-08-15")
	for _, info := range weightList {
		weightInfo.CurrentWeight = append(weightInfo.CurrentWeight, info.CurrentWeight)
		weightInfo.MaxWeight = append(weightInfo.MaxWeight, info.StandMaxWeight*2)
		weightInfo.MinWeight = append(weightInfo.MinWeight, info.StandMinWeight*2)
		tmpData, _ := time.Parse("2006-01-02", info.Date)
		weightInfo.Date = append(weightInfo.Date, fmt.Sprintf("%s(%d天)", info.Date, int64(tmpData.Sub(birthDate).Hours())/24+1))
	}
	return weightInfo
}

type EventInfo struct {
	Date      string
	EventDesc string
}

func GetChildEventInfo() []EventInfo {
	return eventMap
}

type weightConf struct {
	StandMaxWeight float64
	StandMinWeight float64
	CurrentWeight  float64
	Date           string
}

type lengthConf struct {
	StandMaxLength float64
	StandMinLength float64
	CurrentLength  float64
	Date           string
}
