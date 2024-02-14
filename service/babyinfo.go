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
	{Date: "2023-10-10", EventDesc: "接种疫苗：重组乙型肝炎疫苗二针"},
	{Date: "2023-11-07", EventDesc: "接种疫苗：13价肺炎球菌多糖结合第一针"},
	{Date: "2023-11-22", EventDesc: "百天"},
	{Date: "2023-11-28", EventDesc: "接种疫苗：13价肺炎球菌多糖结合第二针"},
	{Date: "2024-01-03", EventDesc: "接种疫苗"},
	{Date: "2024-01-03", EventDesc: "办理户口"},
	{Date: "2024-01-21", EventDesc: "第一张银行卡"},
	{Date: "2024-01-26", EventDesc: "生病：高烧、咳嗽"},
	{Date: "2024-01-26", EventDesc: "办理医保"},
	{Date: "2024-02-13", EventDesc: "病愈，能蠕动爬行。。"},
	{Date: "2024-02-15", EventDesc: "半岁"},
}

var lengthList = []lengthConf{
	{StandMaxLength: 54.0, StandMinLength: 46.9, CurrentLength: 50.00, Date: "2023-08-15"},
	{StandMaxLength: 59.0, StandMinLength: 50.7, CurrentLength: 56.00, Date: "2023-09-14"},
	{StandMaxLength: 60.0, StandMinLength: 51.7, CurrentLength: 56.00, Date: "2023-09-19"},
	{StandMaxLength: 61.0, StandMinLength: 52.6, CurrentLength: 56.00, Date: "2023-09-26"},
	{StandMaxLength: 73.6, StandMinLength: 64.0, CurrentLength: 70.00, Date: "2024-02-14"},
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
	{StandMaxWeight: 7.4, StandMinWeight: 4.7, CurrentWeight: 11.7, Date: "2023-10-22"},
	{StandMaxWeight: 7.7, StandMinWeight: 4.8, CurrentWeight: 12, Date: "2023-10-24"},
	{StandMaxWeight: 7.8, StandMinWeight: 5.0, CurrentWeight: 12, Date: "2023-10-31"},
	{StandMaxWeight: 7.8, StandMinWeight: 5.0, CurrentWeight: 12.2, Date: "2023-11-01"},
	{StandMaxWeight: 8.2, StandMinWeight: 5.1, CurrentWeight: 12.4, Date: "2023-11-07"},
	{StandMaxWeight: 8.2, StandMinWeight: 5.1, CurrentWeight: 13.2, Date: "2023-11-09"},
	{StandMaxWeight: 8.2, StandMinWeight: 5.1, CurrentWeight: 12.9, Date: "2023-11-13"},
	{StandMaxWeight: 8.6, StandMinWeight: 5.5, CurrentWeight: 13.4, Date: "2023-11-22"},
	{StandMaxWeight: 8.6, StandMinWeight: 5.5, CurrentWeight: 13.5, Date: "2023-11-25"},
	{StandMaxWeight: 8.8, StandMinWeight: 5.6, CurrentWeight: 14.5, Date: "2023-12-02"},
	{StandMaxWeight: 9.1, StandMinWeight: 5.8, CurrentWeight: 14.2, Date: "2023-12-05"},
	{StandMaxWeight: 9.3, StandMinWeight: 5.9, CurrentWeight: 15.5, Date: "2023-12-15"},
	{StandMaxWeight: 9.7, StandMinWeight: 6.1, CurrentWeight: 15.3, Date: "2023-12-26"},
	{StandMaxWeight: 9.7, StandMinWeight: 6.1, CurrentWeight: 15.7, Date: "2024-01-01"},
	{StandMaxWeight: 10.0, StandMinWeight: 6.4, CurrentWeight: 16.3, Date: "2024-01-11"},
	{StandMaxWeight: 10.1, StandMinWeight: 6.5, CurrentWeight: 16.9, Date: "2024-01-21"},
	{StandMaxWeight: 10.2, StandMinWeight: 6.6, CurrentWeight: 15.8, Date: "2024-01-29"},
	{StandMaxWeight: 10.4, StandMinWeight: 6.7, CurrentWeight: 17.3, Date: "2024-02-05"},
	{StandMaxWeight: 10.5, StandMinWeight: 6.7, CurrentWeight: 16.6, Date: "2024-02-09"},
	{StandMaxWeight: 10.6, StandMinWeight: 6.8, CurrentWeight: 16.5, Date: "2024-02-14"},
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
