<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script src="/tools/static/js/echarts.min.js"></script>
</head>
<body>
    <!-- 为 ECharts 准备一个定义了宽高的 DOM -->
    <div id="weight" style="float: middle;width:100%;height:550px;position:relative"></div>
    <div id="event" style="float: middle;width:100%;height:200px;position:relative"></div>
    <script type="text/javascript">
        var eventInfo = {{.event_info}};
        const eventMap = new Map();
        for (i = 0; i < eventInfo.length; i++) {
          eventMap.set(eventInfo[i].Date, eventInfo[i].EventDesc);
        }
        var chartDom = document.getElementById('event');
        var myChart = echarts.init(chartDom,);
        var option;

        function getDailyData(year) {
          const date = +echarts.time.parse(year + '-01-01');
          const end = +echarts.time.parse(+year + 1 + '-01-01');
          const dayTime = 3600 * 24 * 1000;
          const data = [];
          for (let time = date; time < end; time += dayTime) {
            var timeFormat = echarts.time.format(time, '{yyyy}-{MM}-{dd}', false);
            if (eventMap.has(timeFormat)) {
                data.push([
                    timeFormat,
                    1,
                    eventMap.get(timeFormat)
                ]);
            }
          }
          return data;
        }
        option = {
          title: {
            top: 30,
            //left: 'center',
            text: '成长事件-2023'
          },
          tooltip: {
            formatter: function (params) {
              return params.value[0] + ' 事件: ' + params.value[2];
            }
          },
          calendar: {
            top: 80,
            left: 30,
            right: 30,
            cellSize: ['auto', 13],
            range: '2023',
            itemStyle: {
              borderWidth: 0.5
            },
            yearLabel: { show: true }
          },
          series: {
            type: 'heatmap',
            coordinateSystem: 'calendar',
            data: getDailyData('2023')
          }
        };

        option && myChart.setOption(option);
    </script>

    <script type="text/javascript">
      var weightInfo = {{.weight_info}}
      var chartDom = document.getElementById('weight');
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
          name: '日期',
          boundaryGap: false,
          data: weightInfo.Date
        },
        yAxis: {
          type: 'value',
          name: '斤'
        },
        series: [
          {
            name: '当前体重',
            type: 'line',
            //stack: 'Total',
            data: weightInfo.CurrentWeight,
            lineStyle: {normal: {color: 'green',width: 4,type: 'dashed'}},
            label: {show: true,position: 'bottom',textStyle: {fontSize: 10}}
          },
          {
            name: '标准最高值',
            type: 'line',
            data: weightInfo.MaxWeight,
            areaStyle: {color:'#C0D9D9',opacity:1,origin:"start"},
          },
          {
            name: '标准最低值',
            type: 'line',
            data: weightInfo.MinWeight,
            areaStyle: {color:'#ffffff',opacity:1,origin:"start"},
          }
        ]
      };

      option && myChart.setOption(option);
    </script>
  </body>
</html>
