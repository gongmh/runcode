<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script src="/tools/static/js/echarts.min.js"></script>
</head>
<body>
    <!-- 为 ECharts 准备一个定义了宽高的 DOM -->
    <div id="main" style="float: middle;width:100%;height:550px;position:relative"></div>
    <script type="text/javascript">

      var showInfo = {{.show_info}}

      var chartDom = document.getElementById('main');
      var myChart = echarts.init(chartDom);
      var option;

      option = {
        title: {
          text: '儿童体重'
        },
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          data: ['当前体重', '标准最高值', '标准最低值']
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        toolbox: {
          feature: {
            saveAsImage: {}
          }
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: showInfo.Date
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: '当前体重',
            type: 'line',
            //stack: 'Total',
            data: showInfo.CurrentWeight,
            lineStyle: {normal: {color: 'green',width: 4,type: 'dashed'}},
            label: {show: true,position: 'bottom',textStyle: {fontSize: 10}}
          },
          {
            name: '标准最高值',
            type: 'line',
            data: showInfo.MaxWeight,
            areaStyle: {color:'#C0D9D9',opacity:1,origin:"start"},
          },
          {
            name: '标准最低值',
            type: 'line',
            data: showInfo.MinWeight,
            areaStyle: {color:'#ffffff',opacity:1,origin:"start"},
          }
        ]
      };

      option && myChart.setOption(option);
    </script>
  </body>
</html>
