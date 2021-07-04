<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no" />
    <style type="text/css">
        body, html,#allmap {width: 100%;height: 100%;overflow: hidden;margin:0;font-family:"微软雅黑";}
    </style>
    <script type="text/javascript" src="https://api.map.baidu.com/api?v=3.0&ak=uKiGoT2VOxDUZusPVbW24TSa81UDAftF"></script>
    <title>自习室</title>
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

if(showInfo.length>0){

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

for(var i=0; i<showInfo.length; i++) {
    var lat = showInfo[i].lat;
    var lng = showInfo[i].lng;
    var t = showInfo[i].title;

    var marker = new BMap.Marker(new BMap.Point(lng,lat), {icon:myIcon});        // 创建标注

    var label = new BMap.Label(t,{"offset":new BMap.Size(10,-20)});
    marker.setLabel(label);

    marker.setTitle(t);

    map.addOverlay(marker);                                       // 将标注添加到地图中
}

</script>
