<!DOCTYPE html>
<html>
<head>
    <title>JSON解析工具</title>
    <script src="https://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <link href="/static/js/codemirror-5.41.0/lib/codemirror.css" rel="stylesheet">
    <script src="/static/js/codemirror-5.41.0/lib/codemirror.js"></script>
    <link href="/static/js/codemirror-5.41.0/theme/darcula.css" rel="stylesheet">
    <script src="/static/js/codemirror-5.41.0/mode/javascript/javascript.js"></script>
</head>

<body>

<div>
<div style="min-height:530px; width:49%;display:inline-block">
    <input class="ace_text-input" id="ace_input"></input>
</div>

<div style="min-height:530px; width:49%;display:inline-block">
    <textarea class="ace_text-input" id="ace_output"></textarea>
</div>

</div>



<script>

    var editorInput = CodeMirror.fromTextArea(document.getElementById("ace_input"), {
        lineNumbers: true,
        extraKeys: {"Ctrl-Q": "autocomplete"},  //避免热键冲突
        mode: {name: "text/javascript", globalVars: true},
        // theme: "darcula"
    });

    var editorOut = CodeMirror.fromTextArea(document.getElementById("ace_output"), {
        lineNumbers: true,
        extraKeys: {"Ctrl-Q": "autocomplete"},  //避免热键冲突
        mode: {name: "text/javascript", globalVars: true},
        // theme: "darcula"
    });

    editorInput.on("change", function () {
        //console.log(1)
        var tmp = JSON.parse(editorInput.getValue());
        var output = JSON.stringify(tmp, undefined, 4);
        editorOut.setValue(output);
    })

</script>
</body>
</html>