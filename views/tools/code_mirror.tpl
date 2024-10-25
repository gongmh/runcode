<!DOCTYPE html>
<html>
<head>
    <title>在线代码运行</title>
    <link href="/tools/static/js/codemirror-5.41.0/lib/codemirror.css" rel="stylesheet">
    <script src="/tools/static/js/codemirror-5.41.0/lib/codemirror.js"></script>
    <script src="/tools/static/js/codemirror-5.41.0/mode/go/go.js"></script>
    <script src="/tools/static/js/codemirror-5.41.0/mode/python/python.js"></script>
    <script src="/tools/static/js/codemirror-5.41.0/mode/php/php.js"></script>
    <script src="/tools/static/js/codemirror-5.41.0/mode/lua/lua.js"></script>
    <script src="/tools/static/js/codemirror-5.41.0/mode/clike/clike.js"></script>
    <link href="/tools/static/js/codemirror-5.41.0/theme/darcula.css" rel="stylesheet">
    <script src="https://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <script>
        var _hmt = _hmt || [];
        (function() {
            var hm = document.createElement("script");
            hm.src = "https://hm.baidu.com/hm.js?a5aa0951990786a4091fa61d0df87b7f";
            var s = document.getElementsByTagName("script")[0];
            s.parentNode.insertBefore(hm, s);
        })();
    </script>
    <style type="text/css">
        .navdown {
            color: #ff0000;
            border: 1px solid #F00;
        }
        .jinrishici {
            position:absolute;
            right:5px;
            top:30px;
            /*font-family: "Arial","Microsoft YaHei","黑体","宋体",sans-serif;*/
            font: bold 12px/20px arial,sans-serif;
            /*font: 14px/1.5 'Microsoft YaHei',arial,tahoma,\5b8b\4f53,sans-serif;*/
            /*font-family: Helvetica, 'Hiragino Sans GB', 'Microsoft Yahei', '微软雅黑', Arial, sans-serif;*/
            color: #999999;
        }
        .CodeMirror {
            height: 530px;
        }
    </style>
</head>

<body>

<div>
    <div>
        <h1>在线代码运行</h1>
    </div>

    <span id="jinrishici-sentence" class="jinrishici">正在加载今日诗词....</span>
    <script src="https://sdk.jinrishici.com/v2/browser/jinrishici.js" charset="utf-8"></script>

    <div id="button-list"></div>

    <div>
        <div style="float: left;width:80%;height: 550px;">
        <pre class="ace_editor" id="code_text" style="min-height:530px;max-width:98%"><textarea class="ace_text-input"
                                                                                           id="ace_text"></textarea></pre>
        </div>
        <div style="float: left;width:20%;height:550px;position:relative">
            <div>
                <p/>
                <p style="font-family: arial, serif;color: darkblue;font-size:12px;">点击分享也可以将您的优秀代码分享给朋友~</p>
                <p/><p/>
                <button id="code_share" style="font-family: arial, serif;color: darkblue">分享</button>
                <input id="share_url" disabled="disabled" style="background:#e0e0e0;width:98%;border:0"/>
            </div>
            <div style="position:absolute;bottom: 0px">
                <p style="font:italic bold 12px/20px arial,sans-serif;color: #999999">欢迎试用，有问题或建议随时反馈~ <a
                        class="email" href="mailto:{{.email}}">{{.email}}</a></p>
            </div>
        </div>
    </div>

    <div>
        <button id="code_run">运行</button>
    </div>

    <div>
        <p>运行结果:</p>
        <textarea id="run_result" disabled="disabled" style="min-height:200px;width:78.4%;background-color: beige"></textarea>
    </div>
</div>

<script type="text/javascript">

    var codeMap = {{.code_map}}
    var defaultLag = {{.default_lag}}

    var codeJsonObj;

    var editor = CodeMirror.fromTextArea(document.getElementById("ace_text"), {
        lineNumbers: true,
        extraKeys: {"Ctrl-Q": "autocomplete"},  //避免热键冲突
        mode: {name: "text/x-go", globalVars: true},
        theme: "darcula"
    });

    $(document).ready(function () {
        codeJsonObj = $.parseJSON(codeMap);
        for (var key in codeJsonObj) {
            if (key == defaultLag) {
                $("#button-list").append("<input type=\"reset\" id=\"button-\"+key value=" + codeJsonObj[key].ShowName + " class=\"navdown\" />")
            } else {
                $("#button-list").append("<input type=\"reset\" id=\"button-\"+key value=" + codeJsonObj[key].ShowName + " />")
            }
        }

        editor.setValue(codeJsonObj[defaultLag].DemoCode)

        $("#share_url").attr("value", window.location.href);

        $("#code_run").click(function () {
            $("#run_result").text("");
            $.post("/tools/runcode",
                {"code_type": defaultLag, "code_content": editor.getValue()},
                // {"code_type": defaultLag, "code_content": encodeURIComponent(editor.getValue())},
                function (data, status) {
                    if (status != "success") {
                        $("#run_result").text("网络错误，请重试~");
                        return
                    }

                    if (data.errno != 0) {
                        text = (data.data) ? (data.data) : "代码错误，请检查后重试~";
                        $("#run_result").text(text);
                        return
                    }

                    $("#run_result").text(data.data);
                }
            );
        });

        $(":reset").click(function () {
            defaultLag = $(this).val();
            $(this).addClass("navdown").siblings().removeClass("navdown");

            editor.setValue(codeJsonObj[defaultLag].DemoCode);
            editor.setOption(
                "mode", {name: codeJsonObj[defaultLag].JsThemeCM}
            )
        })

        $("#code_share").click(function () {
            $("#share_url").attr("value", "");
            $.post(
                "/tools/s",
                {"code_type": defaultLag, "code_content": editor.getValue()},
                function (data, status) {
                    if (status != "success") {
                        $("#share_url").attr("value", "网络错误，请重试~");
                        return
                    }

                    if (data.errno != 0) {
                        text = (data.data) ? (data.data) : "分享链接生成错误，请反馈~";
                        $("#share_url").attr("value", text);
                        return
                    }
                    var shareUrl = window.location.protocol + "//" + window.location.host + "/tools/s?id=" + data.data
                    $("#share_url").attr("value", shareUrl);
                }
            );
        })
    });
</script>


</body>
</html>