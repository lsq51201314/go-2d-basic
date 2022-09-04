@echo off
cd %cd%
echo 生成图标 。。。
echo IDI_ICON1 ICON "main.ico" > main.rc
windres -o main.syso main.rc
echo 正在编译 。。。
go build -ldflags "-s -w -H windowsgui"
echo 编译完成 。。。
pause