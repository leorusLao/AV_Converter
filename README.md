# AV_Converter

Audio and Video Converter，音视频转换器

有能力最好还是自行编写 ffmpeg 的编解码。

# 指南

- make
```
// 打包
make

// debug
make debug

// 清理
make clean
```

# 快速开始

```
./av_converter -i xxx.mp3 -o xxx.wav
```

# 一、依赖工具说明

## 1. 音视频转换工具

### 1.1 常规格式转换 FFmpeg

- 项目地址：https://github.com/FFmpeg/FFmpeg
- 位置：tools/ffmpeg.exe

### 1.2 silk格式转换

- 项目地址：https://github.com/kn007/silk-v3-decoder
- 位置：tools/silk_v3_decoder.exe

### 1.3 adu与hzmv格式转换

- 项目地址：https://github.com/leorusLao/hzmv-adu-convert
- 位置：tools/adu_convert.exe

### 1.4 spx 格式转换

- 项目地址：
- 位置：tools/speexdec.exe