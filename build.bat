@echo off
cd %cd%
echo ����ͼ�� ������
echo IDI_ICON1 ICON "main.ico" > main.rc
windres -o main.syso main.rc
echo ���ڱ��� ������
go build -ldflags "-s -w -H windowsgui"
echo ������� ������
pause