

testGen 一个用来演示通过reflect 来自动生成指定类型的对象的生成器
-src
--APi
----ApiDef.go  定义类型
--Gen
----Generator.go  生成器
testGen.go    主文件

ApiDef 通过init 函数，向Generator  注册自己的类型

注意,需要export GOPATH=$GOPATH:`pwd`  将src目录注册到GOPATH方可编译过去
