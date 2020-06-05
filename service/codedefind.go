package service

//插入类型
const (
	CODE_INSERT_TYPE_BY_RUN   = "run_code"
	CODE_INSERT_TYPE_BY_SHARE = "share_code"
)

//
type CodeStruct struct {
	ShowName  string
	JsTheme   string
	DemoCode  string
	JsThemeCM string
	dockerCmd string
	postfix   string
}

//示例程序
var phpDemoCode = "<?php\nprint(\"hello php\");"
var goDemoCode = "package main\n\nfunc main() { \n    println(\"hello golang\")\n}"
var javaDemoCode = "import java.io.*;\nclass test {\n    public static void main (String[] args) throws java.lang.Exception {\n        System.out.println(\"hi java\");\n    }\n}"
var pyDemoCode = "#!/usr/bin/python\n\nprint(\"hello python\")"
var cppDemoCode = "#include <iostream>\n\nint main(){\n    std::cout << \"hello cpp\" << std::endl;\n}"
var luaDemoCode = "print(\"hello world\")"
var py3DemoCode = "#!/usr/bin/python3\n\nprint(\"hello python3\")"

//docker命令，格式：1. 先创建容器；2.copy文件到容器；3. 运行容器执行命令
var golangCmdV2 = `cid=$(docker create -c 512 -m 512M --rm golang:alpine sh -c "go build /tmp/1.go && ./1 2>> tmp && cat tmp") && docker cp %s $cid:/tmp/1.go && docker start $cid -a`
var phpCmdV2 = `cid=$(docker create -c 512 -m 512M --rm php sh -c "/usr/local/bin/php /tmp/1.php") && docker cp %s $cid:/tmp/1.php && docker start $cid -a`
var javaCmdV2 = `cid=$(docker create -c 512 -m 512M --rm java sh -c "javac /tmp/1.java && sh /tmp/javaRunner.sh") && docker cp ./scripts/javaRunner.sh $cid:/tmp/javaRunner.sh && docker cp %s $cid:/tmp/1.java && docker start $cid -a`
var cppCmdV2 = `cid=$(docker create -c 512 -m 512M --rm -w /tmp gcc:4.9 sh -c "g++ -o myapp 1.cpp && ./myapp") && docker cp %s $cid:/tmp/1.cpp && docker start $cid -a`
var pythonCmdV2 = `cid=$(docker create -c 512 -m 512M --rm python:2.7 sh -c "python /tmp/1.py") && docker cp %s $cid:/tmp/1.py && docker start $cid -a`
var luaCmdV2 = `cid=$(docker create -c 512 -m 512M --rm luafan/luafan-alpine sh -c "lua /tmp/1.lua") && docker cp %s $cid:/tmp/1.lua && docker start $cid -a`
var python3CmdV2 = `cid=$(docker create -c 512 -m 512M --rm python sh -c "python3 /tmp/1.py") && docker cp %s $cid:/tmp/1.py && docker start $cid -a`

//代码整体配置
//codemirror mode: https://codemirror.net/mode/index.html
var codeShowMap = map[string]CodeStruct{
	"php7": {
		"php7",
		"php",
		phpDemoCode,
		"text/x-php",
		phpCmdV2,
		".php",
	},
	"java": {
		"java",
		"java",
		javaDemoCode,
		"text/x-java",
		javaCmdV2,
		".java",
	},
	"golang": {
		"golang",
		"golang",
		goDemoCode,
		"text/x-go",
		golangCmdV2,
		".go",
	},
	"python2.7": {
		"python2.7",
		"python",
		pyDemoCode,
		"text/x-cython",
		pythonCmdV2,
		".py",
	},
	"cpp": {
		"cpp",
		"c_cpp",
		cppDemoCode,
		"text/x-c++src",
		cppCmdV2,
		".cpp",
	},
	"lua": {
		"lua",
		"lua",
		luaDemoCode,
		"text/x-lua",
		luaCmdV2,
		".lua",
	},
	"python3": {
		"python3",
		"python",
		py3DemoCode,
		"text/x-cython",
		python3CmdV2,
		".py",
	},
}
