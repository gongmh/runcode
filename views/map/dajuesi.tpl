<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no" />
    <style type="text/css">
        body, html,#allmap {width: 100%;height: 100%;overflow: hidden;margin:0;font-family:"微软雅黑";}
    </style>
    <script type="text/javascript" src="https://api.map.baidu.com/api?v=3.0&ak=uKiGoT2VOxDUZusPVbW24TSa81UDAftF"></script>
    <title>👣大觉寺</title>
</head>
<body>
<div id="allmap"></div>
</body>
</html>
<script type="text/javascript">

var centerLng = 103.388611
var centerLat = 35.563611
var zoomLevel = 5

var showInfo = {{.show_info}}

if(showInfo.zoomLevel>0){
    zoomLevel = showInfo.zoomLevel
}
if(showInfo.centerPoint.lat>0){
    centerLat = showInfo.centerPoint.lat
}
if(showInfo.centerPoint.lng>0){
    centerLng = showInfo.centerPoint.lng
}

// 百度地图API功能
var map = new BMap.Map("allmap");    // 创建Map实例
map.centerAndZoom(new BMap.Point(centerLng,centerLat), zoomLevel);

// map.setMapStyleV2({styleId: '7f931e749f38960cf27eb5fdbd88de6a'});
// map.setMapStyleV2({styleId: '8526f7a861939934384b083607c56fb0'});
map.setMapStyleV2({styleId: 'eff123a007bb8651ae658ee55010aec1'});

//添加地图类型控件
map.addControl(new BMap.MapTypeControl({
    mapTypes:[
        BMAP_NORMAL_MAP,
        BMAP_HYBRID_MAP
    ]}));

map.enableScrollWheelZoom(true);     //开启鼠标滚轮缩放

//BMap_Symbol_SHAPE_POINT
var myIcon = new BMap.Symbol(BMap_Symbol_SHAPE_CIRCLE, {
    scale: 2,//图标缩放大小
    fillColor: "red",//填充颜色
    fillOpacity: 0.5,//填充透明度
    strokeOpacity: 1,
    strokeWeight: 0,
    strokeColor: "red",
});

for(var i=0; i<showInfo.markList.length; i++) {
    var lat = showInfo.markList[i].lat;
    var lng = showInfo.markList[i].lng;
    var t = showInfo.markList[i].title;
    var marker = new BMap.Marker(new BMap.Point(lng,lat), {icon:myIcon});        // 创建标注

    // var label = new BMap.Label(t,{"offset":new BMap.Size(10,-20)});
    // marker.setLabel(label);

    marker.setTitle(t);

    map.addOverlay(marker);                                       // 将标注添加到地图中
}

for(var i=0; i<showInfo.lineList.length; i++) {
    var lineInfo = showInfo.lineList[i];

    var nodes = new Array(0);
    for(var j=0; j<lineInfo.points.length; j++) {
        nodes.push(new BMap.Point(lineInfo.points[j].lng, lineInfo.points[j].lat));
    }
    var polyline = new BMap.Polyline(nodes,
        {strokeColor: "blue", strokeWeight: 3, strokeOpacity: 0.5}
    );
    map.addOverlay(polyline);
}


// var geolocation = new BMap.Geolocation();
// geolocation.getCurrentPosition(function(r){
// 	if(this.getStatus() == BMAP_STATUS_SUCCESS){
// 		var mk = new BMap.Marker(r.point);
// 		map.addOverlay(mk);
// 		map.panTo(r.point);
// 		alert('您的位置：'+r.point.lng+','+r.point.lat);
// 	}
// 	else {
// 		alert('failed'+this.getStatus());
// 	}
// });

</script>
