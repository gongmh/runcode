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
	{Date: "2024-02-19", EventDesc: "第一次吃辅食：米粉"},
	{Date: "2024-02-28", EventDesc: "体检档案转移"},
	{Date: "2024-03-01", EventDesc: "接种疫苗"},
	{Date: "2024-03-07", EventDesc: "领取医保卡"},
	{Date: "2024-03-15", EventDesc: "感冒流鼻涕🤧"},
	{Date: "2024-03-29", EventDesc: "接种疫苗：五联第三针"},
	{Date: "2024-04-12", EventDesc: "第一次坐着洗澡🛀"},
	{Date: "2024-04-16", EventDesc: "接种疫苗：麻腮风疫苗"},
	{Date: "2024-04-27", EventDesc: "回老家"},
	{Date: "2024-05-04", EventDesc: "回京"},
	{Date: "2024-05-05", EventDesc: "第一次理发"},
	{Date: "2024-05-07", EventDesc: "接种疫苗：13价肺炎球菌多糖结合第三针"},
	{Date: "2024-05-21", EventDesc: "接种疫苗：流脑第二针"},
	{Date: "2024-06-04", EventDesc: "接种疫苗：乙肝第三针"},
	{Date: "2024-07-02", EventDesc: "接种疫苗：吃轮状疫苗"},
	{Date: "2024-07-18", EventDesc: "接种疫苗：手足口疫苗"},
	{Date: "2024-07-20", EventDesc: "认识火火兔，自己能打开"}.cDate: "2024-08-29", EventDesc: "接种疫苗：乙脑，发高烧"},
	{Date: "2024-09-10", EventDesc: "会自己满屋子溜达了😄"},
	{Date: "2024-10-22", EventDesc: "会叫爸爸了~"},
	{Date: "2024-11-26", EventDesc: "今年第一场雪，清晰的叫爸爸了。"},
	{Date: "2025-01-28", EventDesc: "理发"},
	{Date: "2025-02-19", EventDesc: "接种疫苗：甲肝"},
	{Date: "2025-03-07", EventDesc: "一直叫爸爸爸爸。。"},
	{Date: "2025-03-08", EventDesc: "自己知道用手吃饭，不过今天摔了一跤，嘴巴破了"},
}

var lengthList = []lengthConf{
	{StandMaxLength: 54.0, StandMinLength: 46.9, CurrentLength: 50.00, Date: "2023-08-15"},
	{StandMaxLength: 59.0, StandMinLength: 50.7, CurrentLength: 56.00, Date: "2023-09-14"},
	{StandMaxLength: 60.0, StandMinLength: 51.7, CurrentLength: 56.00, Date: "2023-09-19"},
	{StandMaxLength: 61.0, StandMinLength: 52.6, CurrentLength: 56.00, Date: "2023-09-26"},
	{StandMaxLength: 73.6, StandMinLength: 64.0, CurrentLength: 70.00, Date: "2024-02-14"},
	{StandMaxLength: 76.6, StandMinLength: 66.6, CurrentLength: 70.07, Date: "2024-04-16"},
	{StandMaxLength: 80.4, StandMinLength: 69.8, CurrentLength: 73.00, Date: "2024-07-05"},
	{StandMaxLength: 84.8, StandMinLength: 73.3, CurrentLength: 77.00, Date: "2024-10-26"},
	{StandMaxLength: 85.4, StandMinLength: 73.7, CurrentLength: 80.00, Date: "2024-11-08"},
	{StandMaxLength: 87.1, StandMinLength: 75.1, CurrentLength: 81.00, Date: "2024-12-20"},
	{StandMaxLength: 90.1, StandMinLength: 77.2, CurrentLength: 84.00, Date: "2025-03-09"},
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
	{StandMaxWeight: 10.9, StandMinWeight: 7.0, CurrentWeight: 18.3, Date: "2024-03-07"},
	{StandMaxWeight: 10.9, StandMinWeight: 7.0, CurrentWeight: 18.4, Date: "2024-03-14"},
	{StandMaxWeight: 11.2, StandMinWeight: 7.1, CurrentWeight: 19.4, Date: "2024-03-29"},
	{StandMaxWeight: 11.3, StandMinWeight: 7.2, CurrentWeight: 19.4, Date: "2024-04-11"},
	{StandMaxWeight: 11.4, StandMinWeight: 7.3, CurrentWeight: 19.6, Date: "2024-04-16"},
	{StandMaxWeight: 11.8, StandMinWeight: 7.6, CurrentWeight: 19.7, Date: "2024-05-25"},
	{StandMaxWeight: 11.9, StandMinWeight: 7.6, CurrentWeight: 19.9, Date: "2024-06-06"},
	{StandMaxWeight: 12.2, StandMinWeight: 7.8, CurrentWeight: 21.0, Date: "2024-07-05"},
	{StandMaxWeight: 12.3, StandMinWeight: 7.9, CurrentWeight: 21.2, Date: "2024-07-12"},
	{StandMaxWeight: 13.1, StandMinWeight: 8.5, CurrentWeight: 23.0, Date: "2024-10-26"},
	{StandMaxWeight: 13.4, StandMinWeight: 8.6, CurrentWeight: 23.7, Date: "2024-11-26"},
	{StandMaxWeight: 13.5, StandMinWeight: 8.8, CurrentWeight: 23.7, Date: "2024-12-20"},
	{StandMaxWeight: 14.3, StandMinWeight: 9.3, CurrentWeight: 25.5, Date: "2025-03-09"},
}

type LengthInfo struct {
	Date          []string
	MaxLength     []float64
	MinLength     []float64
	CurrentLength []float64
}

func GetChildLengthInfo() LengthInfo {
	var lengthInfo LengthInfo

	for _, info := range lengthList {
		lengthInfo.CurrentLength = append(lengthInfo.CurrentLength, info.CurrentLength)
		lengthInfo.MaxLength = append(lengthInfo.MaxLength, info.StandMaxLength)
		lengthInfo.MinLength = append(lengthInfo.MinLength, info.StandMinLength)
		lengthInfo.Date = append(lengthInfo.Date, handleDay(info.Date))
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
	for _, info := range weightList {
		weightInfo.CurrentWeight = append(weightInfo.CurrentWeight, info.CurrentWeight)
		weightInfo.MaxWeight = append(weightInfo.MaxWeight, info.StandMaxWeight*2)
		weightInfo.MinWeight = append(weightInfo.MinWeight, info.StandMinWeight*2)
		weightInfo.Date = append(weightInfo.Date, handleDay(info.Date))
	}
	return weightInfo
}

func handleDay(date string) string {
	birthDate, _ := time.Parse("2006-01-02", "2023-08-15")
	tmpData, _ := time.Parse("2006-01-02", date)
	day := int64(tmpData.Sub(birthDate).Hours())/24 + 1
	if day > 100 {
		month := day / 30
		day := day % 30
		return fmt.Sprintf("%d月%d天", month, day)
	}

	return fmt.Sprintf("%d天", day)
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
