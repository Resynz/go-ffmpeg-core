# go-ffmpeg-core
ffmpeg core for golang

基于`ffmpeg`命令封装简单功能，因此运行环境需事先安装有`ffmpeg`命令。

`ffmpeg`官方下载地址: http://ffmpeg.org/download.html



### 1. 剥离视频文件的音频/视频

```go
package main

import (
  go_ffmepg_core "github.com/Resynz/go-ffmpeg-core"
  "log"
)

func main () {
  resourcePath:="./test.mp4" // 原视频路径
  outputPath:="./test-peel-audio.m4a" // 文件输出路径
  peel:=&go_ffmpeg_core.Peel{
    Command: "ffmpeg",
    ResourcePath: resourcePath,
    OutputPath: outputPath,
    Type: go_ffmpeg_core.PeelTypeAudio, // 指定剥离类型为音频/视频
    Stream: "", // 如果该视频文件中有存在多条音轨，可指定音轨编号
  }
  err:=peel.Execute()
  if err!=nil{
    log.Fatalf("peel failed! error:%s\n",err)
  }
  log.Println("peel done.")
}
```



### 2. 合并音频、视频

```go
package main

import (
  go_ffmepg_core "github.com/Resynz/go-ffmpeg-core"
  "log"
)

func main () {
  merge:=&go_ffmpeg_core.MergeAV{
  	Command:    "ffmpeg",
	VideoPath:  "./test-v.mp4", // 纯视频文件路径
	AudioPath:  "./test-a.m4a", // 纯音频文件路径
	OutputPath: "./merge.mp4",  // 结果输出路径
  }
  if err:=merge.Execute();err!=nil{
	log.Fatalf("merge failed! error:%s\n",err.Error())
  }
  log.Println("merge done.")
}
```



### 3. 转换格式

```go
package main

import (
  go_ffmepg_core "github.com/Resynz/go-ffmpeg-core"
  "log"
)

func main () {
  transform:=&go_ffmpeg_core.Transform{
		Command:      "ffmpeg",
		ResourcePath: "./test.mp4",
		OutputPath:   "./test.mkv",
  }
  if err:=transform.Execute();err!=nil{
	log.Fatalf("tranform failed! error:%s\n",err.Error())
  }
  log.Println("transform done.")
}
```

