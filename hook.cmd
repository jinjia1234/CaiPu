@echo off
ECHO 完全路径：%0
ECHO 去掉引号：%~0
ECHO 所在分区：%~d0
ECHO 所处路径：%~p0
ECHO 文件名：%~n0
ECHO 扩展名：%~x0
ECHO 文件属性：%~a0
ECHO 修改时间：%~t0
ECHO 文件大小：%~z0
echo 当前的盘符:%~d0
echo 当前的盘符及路径:%~dp0
echo 当前的盘符及路径的短文件名格式:%~sdp0
echo 当前批处理文件名:%~n0
echo.
cls

git fetch --all
git reset --hard origin/master
git pull origin master
go run main.go
