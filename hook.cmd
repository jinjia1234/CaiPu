@echo off
ECHO ��ȫ·����%0
ECHO ȥ�����ţ�%~0
ECHO ���ڷ�����%~d0
ECHO ����·����%~p0
ECHO �ļ�����%~n0
ECHO ��չ����%~x0
ECHO �ļ����ԣ�%~a0
ECHO �޸�ʱ�䣺%~t0
ECHO �ļ���С��%~z0
echo ��ǰ���̷�:%~d0
echo ��ǰ���̷���·��:%~dp0
echo ��ǰ���̷���·���Ķ��ļ�����ʽ:%~sdp0
echo ��ǰ�������ļ���:%~n0
echo.
cls

git fetch --all
git reset --hard origin/master
git pull origin master
go run main.go
