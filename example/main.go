// author: liuaifu
// laf163@gmail.com
// 2011-3-4

package main

import (
	"time"
	"winservice"
)

func main() {
	/**
	* 这个GoService1替换为你创建的服务名，创建服务的命令如下：
	* 注意：等号后的空格必不可少
	* sc.exe create GoService1 binPath= x:\...\test.exe
	* sc start GoService1
	*/
	winservice.Start("GoService1")
	for{
		//在这里做一些事情
		//do something here
		//...
		time.Sleep(20*1e9)
	}
	winservice.Stop()
	return
}
