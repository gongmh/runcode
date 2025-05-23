<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script src="/tools/static/js/echarts.min.js"></script>
<title>儿童信息</title>
</head>
<body>
    <!-- 为 ECharts 准备一个定义了宽高的 DOM -->
    <div id="chart" style="float: middle;width:100%;height:550px;position:relative">
      <div id="weight" style="float: left;width:60%;height:550px;position:relative"></div>
      <div id="length" style="float: right;width:40%;height:550px;position:relative"></div>
    </div>
    <div id="event" style="float: middle;width:100%;height:400px;position:relative"></div>
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
            text: '成长事件'
          },
          tooltip: {
            formatter: function (params) {
              return params.value[0] + ' 事件: ' + params.value[2];
            }
          },
          calendar: [
            {
                top: 80,
                cellSize: ['auto', 13],
                range: '2025',
                itemStyle: {
                    borderWidth: 0.5
                }
            },
            {
                top: 200,
                cellSize: ['auto', 13],
                range: '2024',
                itemStyle: {
                    borderWidth: 0.5
                }
            },
            {
                top: 320,
                cellSize: ['auto', 13],
                range: '2023',
                itemStyle: {
                    borderWidth: 0.5
                }
            }
          ],
          series: [
            {
                type: 'heatmap',
                coordinateSystem: 'calendar',
                calendarIndex: 0,
                data: getDailyData('2025')
            },
            {
                type: 'heatmap',
                coordinateSystem: 'calendar',
                calendarIndex: 1,
                data: getDailyData('2024')
            },
            {
                type: 'heatmap',
                coordinateSystem: 'calendar',
                calendarIndex: 2,
                data: getDailyData('2023')
            }
          ]
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

    <script type="text/javascript">
          var lengthInfo = {{.length_info}}
          var chartDom = document.getElementById('length');
          var myChart = echarts.init(chartDom);
          var option;

          option = {
            title: {
              text: '儿童身高'
            },
            tooltip: {
              trigger: 'axis'
            },
            legend: {
              data: ['当前身高', '标准最高值', '标准最低值']
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
              data: lengthInfo.Date
            },
            yAxis: {
              type: 'value',
              name: '厘米'
            },
            series: [
              {
                name: '当前身高',
                type: 'line',
                //stack: 'Total',
                data: lengthInfo.CurrentLength,
                lineStyle: {normal: {color: 'green',width: 4,type: 'dashed'}},
                label: {show: true,position: 'bottom',textStyle: {fontSize: 10}}
              },
              {
                name: '标准最高值',
                type: 'line',
                data: lengthInfo.MaxLength,
                areaStyle: {color:'#C0D9D9',opacity:1,origin:"start"},
              },
              {
                name: '标准最低值',
                type: 'line',
                data: lengthInfo.MinLength,
                areaStyle: {color:'#ffffff',opacity:1,origin:"start"},
              }
            ]
          };

          option && myChart.setOption(option);
        </script>
  </body>
</html>
