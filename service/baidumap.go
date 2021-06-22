package service

type Point struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type MarkPoint struct {
	Point
	Title      string `json:"title"`
	Desc       string `json:"-"`
	CreateTime string `json:"-"`
}

type Line struct {
	Points []MarkPoint `json:"points"`
	Title  string      `json:"title"`
}

type ShowInfo struct {
	CenterPoint Point       `json:"centerPoint"`
	ZoomLevel   int         `json:"zoomLevel"`
	MarkList    []MarkPoint `json:"markList"`
	LineList    []Line      `json:"lineList"`
}

func GenMapShowInfo() (si ShowInfo) {
	si = ShowInfo{
		CenterPoint: Point{
			Lng: 35.563611,
			Lat: 103.388611,
		},
		ZoomLevel: 5,
		MarkList:  []MarkPoint{},
		LineList:  []Line{},
	}

	pointList := getMarkList()
	if len(pointList) == 0 {
		return si
	}

	var maxLng = pointList[0].Lng
	var minLng = pointList[0].Lng
	var maxLat = pointList[0].Lat
	var minLat = pointList[0].Lat
	for _, p := range pointList {
		if p.Lng > maxLng {
			maxLng = p.Lng
		}
		if p.Lng < minLng {
			minLng = p.Lng
		}
		if p.Lat > maxLat {
			maxLat = p.Lat
		}
		if p.Lat < minLat {
			minLat = p.Lat
		}
	}
	si.CenterPoint.Lng = (maxLng + minLng) / 2
	si.CenterPoint.Lat = (maxLat + minLat) / 2
	si.MarkList = pointList
	si.LineList = getLineList()

	return
}

func getMarkList() []MarkPoint {
	return []MarkPoint{
		{Title: "信阳东站", Point: Point{114.165959, 32.149747}, Desc: "", CreateTime: "2021-05-08 22:24:30"},
		{Title: "武汉天河机场", Point: Point{114.217379, 30.776258}, Desc: "", CreateTime: "2021-05-08 18:24:30"},
		{Title: "伊宁汉庭上海城酒店", Point: Point{81.300803, 43.947003}, Desc: "", CreateTime: "2021-05-07 22:00:00"},
		{Title: "伊宁六星街", Point: Point{81.316078, 43.938453}, Desc: "", CreateTime: "2021-05-07 19:24:30"},
		{Title: "霍尔果斯国门", Point: Point{80.411855, 44.219006}, Desc: "", CreateTime: "2021-05-07 10:24:30"},
		{Title: "霍尔果斯全季酒店", Point: Point{80.419168, 44.215933}, Desc: "", CreateTime: "2021-05-06 22:24:30"},
		{Title: "果子沟大桥景区", Point: Point{81.151629, 44.474556}, Desc: "", CreateTime: "2021-05-06 15:24:30"},
		{Title: "赛里木湖景区", Point: Point{81.401139, 44.612755}, Desc: "", CreateTime: "2021-05-06 12:24:30"},
		{Title: "新源县格林豪泰商务酒店（那拉提店）", Point: Point{83.264764, 43.44503}, Desc: "", CreateTime: "2021-05-05 22:24:30"},
		{Title: "那拉提景区", Point: Point{84.032821, 43.325552}, Desc: "", CreateTime: "2021-05-05 12:24:30"},
		{Title: "新源县腾龙云海酒店", Point: Point{83.271604, 43.430503}, Desc: "", CreateTime: "2021-05-04 22:24:30"},
		{Title: "新疆喀拉峻游客服务中心景区（主入口）", Point: Point{81.906539, 43.124788}, Desc: "", CreateTime: "2021-05-04 09:24:30"},
		{Title: "特克斯酒店", Point: Point{81.84237, 43.223632}, Desc: "", CreateTime: "2021-05-03 22:24:30"},
		{Title: "特克斯八卦城", Point: Point{81.846232, 43.221328}, Desc: "", CreateTime: "2021-05-03 19:24:30"},
		{Title: "喀拉峻景区布拉克门票站", Point: Point{82.04804, 43.099923}, Desc: "", CreateTime: "2021-05-03 12:24:30"},
		{Title: "伊宁全季酒店（上海城）", Point: Point{81.307841, 43.950092}, Desc: "", CreateTime: "2021-05-02 21:24:30"},
		{Title: "伊宁喀赞其", Point: Point{81.339165, 43.918632}, Desc: "", CreateTime: "2021-05-02 12:24:30"},
		{Title: "伊宁机场", Point: Point{81.338952, 43.963204}, Desc: "", CreateTime: "2021-05-02 08:10:30"},
		{Title: "乌鲁木齐速8智选酒店（乌鲁木齐机场店）", Point: Point{87.49028, 43.906596}, Desc: "", CreateTime: "2021-05-02 01:10:30"},
		{Title: "乌鲁木齐机场", Point: Point{87.487028, 43.912386}, Desc: "", CreateTime: "2021-05-01 23:10:30"},
		{Title: "乌鲁木齐", Point: Point{87.590787, 43.825756}, Desc: ""},
		{Title: "克拉玛依", Point: Point{84.871732, 45.591439}, Desc: ""},
		{Title: "吐鲁番", Point: Point{89.52114, 42.925261}, Desc: ""},
		{Title: "喀纳斯禾木", Point: Point{87.439506, 48.574346}, Desc: ""},
		{Title: "西宁市塔尔寺", Point: Point{101.574541, 36.493597}, Desc: ""},
		{Title: "西宁市东关清真大寺", Point: Point{101.803817, 36.621431}, Desc: ""},
		{Title: "张掖七彩丹霞地质公园", Point: Point{100.075745, 38.960823}, Desc: ""},
		{Title: "嘉峪关长城", Point: Point{98.229967, 39.813063}, Desc: ""},
		{Title: "敦煌莫高窟", Point: Point{94.815727, 40.048724}, Desc: ""},
		{Title: "鸣沙山月牙泉", Point: Point{94.685194, 40.094898}, Desc: ""},
		{Title: "西千佛洞", Point: Point{94.372369, 39.98145}, Desc: ""},
		{Title: "乌素特(水上)雅丹地质公园", Point: Point{93.781293, 37.629759}, Desc: ""},
		{Title: "海西源来青年旅馆", Point: Point{95.360535, 37.862491}, Desc: ""},
		{Title: "德令哈", Point: Point{97.373571, 37.375542}, Desc: ""},
		{Title: "中山桥", Point: Point{103.823431, 36.070959}, Desc: ""},
		{Title: "拉卜楞寺", Point: Point{102.516155, 35.198777}, Desc: ""},
		{Title: "郭莽湿地", Point: Point{102.319261, 34.316128}, Desc: ""},
		{Title: "郎木寺院", Point: Point{102.639326, 34.096656}, Desc: ""},
		{Title: "扎尕那景区", Point: Point{103.198081, 34.243502}, Desc: ""},
		{Title: "迭部吐蕃秘境宾馆", Point: Point{103.217065, 34.062598}, Desc: ""},
		{Title: "岷县", Point: Point{104.047075, 34.441573}, Desc: ""},
		{Title: "定西分水岭", Point: Point{104.096823, 34.938273}, Desc: ""},
		{Title: "西安", Point: Point{108.909777, 34.158992}, Desc: ""},
		{Title: "西北第一漂-丹江漂流", Point: Point{110.591399, 33.454982}, Desc: ""},
		{Title: "乾陵", Point: Point{108.224607, 34.583649}, Desc: ""},
		{Title: "华山", Point: Point{110.073028, 34.497647}, Desc: ""},
		{Title: "太原", Point: Point{112.617131, 37.797443}, Desc: ""},
		{Title: "云冈石窟", Point: Point{113.144291, 40.119273}, Desc: ""},
		{Title: "悬空寺", Point: Point{113.721979, 39.666616}, Desc: ""},
		{Title: "应县木塔", Point: Point{113.195635, 39.572146}, Desc: ""},
		{Title: "岱海", Point: Point{112.65141, 40.601473}, Desc: ""},
		{Title: "成都", Point: Point{104.051731, 30.509578}, Desc: ""},
		{Title: "都江堰", Point: Point{103.617627, 31.00793}, Desc: ""},
		{Title: "九寨沟", Point: Point{103.921224, 33.163223}, Desc: ""},
		{Title: "北京", Point: Point{116.396797, 40.025231}, Desc: ""},
		{Title: "天津", Point: Point{117.216853, 39.142488}, Desc: ""},
		{Title: "山海关", Point: Point{119.802921, 39.978057}, Desc: ""},
		{Title: "北戴河", Point: Point{119.494898, 39.819804}, Desc: ""},
		{Title: "张家口", Point: Point{114.886893, 40.710106}, Desc: ""},
		{Title: "宝昌", Point: Point{115.283741, 41.900326}, Desc: ""},
		{Title: "闪电湖", Point: Point{115.841109, 41.654533}, Desc: ""},
		{Title: "坝上草原", Point: Point{115.411307, 41.245692}, Desc: ""},
		{Title: "桑根达来", Point: Point{115.9624, 42.694208}, Desc: ""},
		{Title: "锡林浩特市平顶山", Point: Point{116.086725, 43.627594}, Desc: ""},
		{Title: "木兰围场坝上旅游风景区", Point: Point{117.501252, 42.326606}, Desc: ""},
		{Title: "青岛栈桥", Point: Point{120.32659, 36.065436}, Desc: ""},
		{Title: "青岛崂山", Point: Point{120.625651, 36.19066}, Desc: ""},
		{Title: "西平", Point: Point{113.878654, 33.397041}, Desc: ""},
		{Title: "郑州", Point: Point{113.668966, 34.787903}, Desc: ""},
		{Title: "漯河", Point: Point{113.970479, 33.578308}, Desc: ""},
		{Title: "嵖岈山", Point: Point{113.736596, 33.135655}, Desc: ""},
		{Title: "武汉", Point: Point{114.372921, 30.543803}, Desc: ""},
		{Title: "南京", Point: Point{118.859406, 32.069291}, Desc: ""},
		{Title: "杭州", Point: Point{120.071528, 30.27289}, Desc: ""},
		{Title: "婺源", Point: Point{118.124476, 29.327558}, Desc: ""},
		{Title: "九江", Point: Point{116.013443, 29.709888}, Desc: ""},
		{Title: "广州", Point: Point{113.264437, 23.154981}, Desc: ""},
		{Title: "珠海", Point: Point{113.555901, 22.221111}, Desc: ""},
		{Title: "海口", Point: Point{110.470153, 19.942054}, Desc: ""},
		{Title: "三亚", Point: Point{109.421115, 18.309499}, Desc: ""},
		{Title: "曼谷", Point: Point{100.494971, 13.749139}, Desc: ""},
		{Title: "普吉岛芭东海滩", Point: Point{98.297053, 7.893212}, Desc: ""},
		{Title: "皮皮岛", Point: Point{98.778401, 7.751461}, Desc: ""},
		{Title: "皇帝岛", Point: Point{98.365837, 7.600926}, Desc: ""},
	}
}

func getLineList() []Line {
	return []Line{
		{
			Title: "伊宁环线",
			Points: []MarkPoint{
				{Title: "伊宁汉庭上海城酒店", Point: Point{81.300803, 43.947003}, Desc: "", CreateTime: "2021-05-07 22:00:00"},
				{Title: "霍尔果斯国门", Point: Point{80.411855, 44.219006}, Desc: "", CreateTime: "2021-05-07 10:24:30"},
				{Title: "果子沟大桥景区", Point: Point{81.151629, 44.474556}, Desc: "", CreateTime: "2021-05-06 15:24:30"},
				{Title: "赛里木湖景区", Point: Point{81.401139, 44.612755}, Desc: "", CreateTime: "2021-05-06 12:24:30"},
				{Title: "新源县格林豪泰商务酒店（那拉提店）", Point: Point{83.264764, 43.44503}, Desc: "", CreateTime: "2021-05-05 22:24:30"},
				{Title: "那拉提景区", Point: Point{84.032821, 43.325552}, Desc: "", CreateTime: "2021-05-05 12:24:30"},
				{Title: "新源县腾龙云海酒店", Point: Point{83.271604, 43.430503}, Desc: "", CreateTime: "2021-05-04 22:24:30"},
				{Title: "新疆喀拉峻游客服务中心景区（主入口）", Point: Point{81.906539, 43.124788}, Desc: "", CreateTime: "2021-05-04 09:24:30"},
				{Title: "特克斯八卦城", Point: Point{81.846232, 43.221328}, Desc: "", CreateTime: "2021-05-03 19:24:30"},
				{Title: "喀拉峻景区布拉克门票站", Point: Point{82.04804, 43.099923}, Desc: "", CreateTime: "2021-05-03 12:24:30"},
				{Title: "伊宁喀赞其", Point: Point{81.339165, 43.918632}, Desc: "", CreateTime: "2021-05-02 12:24:30"},
				{Title: "伊宁机场", Point: Point{81.338952, 43.963204}, Desc: "", CreateTime: "2021-05-02 08:10:30"},
			},
		},
		{
			Title: "青海大环线",
			Points: []MarkPoint{
				{Title: "西宁市塔尔寺", Point: Point{101.574541, 36.493597}},
				{Title: "西宁市东关清真大寺", Point: Point{101.803817, 36.621431}, Desc: ""},
				{Title: "张掖七彩丹霞地质公园", Point: Point{100.075745, 38.960823}, Desc: ""},
				{Title: "嘉峪关长城", Point: Point{98.229967, 39.813063}, Desc: ""},
				{Title: "敦煌莫高窟", Point: Point{94.815727, 40.048724}, Desc: ""},
				{Title: "鸣沙山月牙泉", Point: Point{94.685194, 40.094898}, Desc: ""},
				{Title: "西千佛洞", Point: Point{94.372369, 39.98145}, Desc: ""},
				{Title: "乌素特(水上)雅丹地质公园", Point: Point{93.781293, 37.629759}, Desc: ""},
				{Title: "海西源来青年旅馆", Point: Point{95.360535, 37.862491}, Desc: ""},
				{Title: "德令哈", Point: Point{97.373571, 37.375542}, Desc: ""},
				{Title: "西宁市塔尔寺", Point: Point{101.574541, 36.493597}},
			},
		},

		{
			Title: "甘南环线",
			Points: []MarkPoint{
				{Title: "中山桥", Point: Point{103.823431, 36.070959}, Desc: ""},
				{Title: "拉卜楞寺", Point: Point{102.516155, 35.198777}, Desc: ""},
				{Title: "郭莽湿地", Point: Point{102.319261, 34.316128}, Desc: ""},
				{Title: "郎木寺院", Point: Point{102.639326, 34.096656}, Desc: ""},
				{Title: "扎尕那景区", Point: Point{103.198081, 34.243502}, Desc: ""},
				{Title: "迭部吐蕃秘境宾馆", Point: Point{103.217065, 34.062598}, Desc: ""},
				{Title: "岷县", Point: Point{104.047075, 34.441573}, Desc: ""},
				{Title: "定西分水岭", Point: Point{104.096823, 34.938273}, Desc: ""},
				{Title: "中山桥", Point: Point{103.823431, 36.070959}, Desc: ""},
			},
		},

		{
			Title: "北疆环线",
			Points: []MarkPoint{
				{Title: "乌鲁木齐", Point: Point{87.590787, 43.825756}, Desc: ""},
				{Title: "吐鲁番", Point: Point{89.52114, 42.925261}, Desc: ""},
				{Title: "喀纳斯禾木", Point: Point{87.439506, 48.574346}, Desc: ""},
				{Title: "克拉玛依", Point: Point{84.871732, 45.591439}, Desc: ""},
				{Title: "乌鲁木齐", Point: Point{87.590787, 43.825756}, Desc: ""},
			},
		},

		{
			Title: "大同环线",
			Points: []MarkPoint{
				{Title: "云冈石窟", Point: Point{113.144291, 40.119273}, Desc: ""},
				{Title: "悬空寺", Point: Point{113.721979, 39.666616}, Desc: ""},
				{Title: "应县木塔", Point: Point{113.195635, 39.572146}, Desc: ""},
				{Title: "岱海", Point: Point{112.65141, 40.601473}, Desc: ""},
				{Title: "云冈石窟", Point: Point{113.144291, 40.119273}, Desc: ""},
			},
		},

		{
			Title: "回乡",
			Points: []MarkPoint{
				{Title: "北京", Point: Point{116.396797, 40.025231}, Desc: ""},
				{Title: "保定", Point: Point{115.487699, 38.868475}, Desc: ""},
				{Title: "石家庄", Point: Point{114.490825, 38.016821}, Desc: ""},
				{Title: "邯郸", Point: Point{114.481854, 36.607487}, Desc: ""},
				{Title: "安阳", Point: Point{114.34631, 36.110865}, Desc: ""},
				{Title: "新乡", Point: Point{113.865742, 35.31287}, Desc: ""},
				{Title: "郑州", Point: Point{113.668966, 34.787903}, Desc: ""},
				{Title: "许昌", Point: Point{113.893707, 34.054566}, Desc: ""},
				{Title: "漯河", Point: Point{113.970479, 33.578308}, Desc: ""},
				{Title: "西平", Point: Point{113.878654, 33.397041}, Desc: ""},
			},
		},

		{
			Title: "锡盟",
			Points: []MarkPoint{
				{Title: "北京", Point: Point{116.396797, 40.025231}, Desc: ""},
				{Title: "张家口", Point: Point{114.886893, 40.710106}, Desc: ""},
				{Title: "坝上草原", Point: Point{115.411307, 41.245692}, Desc: ""},
				{Title: "宝昌", Point: Point{115.283741, 41.900326}, Desc: ""},
				{Title: "闪电湖", Point: Point{115.841109, 41.654533}, Desc: ""},
				{Title: "桑根达来", Point: Point{115.9624, 42.694208}, Desc: ""},
				{Title: "锡林浩特市平顶山", Point: Point{116.086725, 43.627594}, Desc: ""},
			},
		},
	}
}
