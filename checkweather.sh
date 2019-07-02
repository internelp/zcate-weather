#!/bin/bash
# 将此脚本添加到 crontab
# 可以在有异常天气时将天气信息发送到 zcate
# 异常天气为: "雨", "雷", "尘", "霾", "雹", "雪", "雾", "沙"

WEA=`/usr/sbin/weather`  #这里修改成正确的路径
success=$?
if [ $success -ge "1" ];then
	echo 1
else
	echo 0
	/usr/lib/zabbix/alertscripts/zcate -token="xxx" -title="异常天气预警" -platform=ios -body="$WEA"
fi