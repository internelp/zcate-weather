# 工具说明

此工具可以查询天气，配合 crontab 可以做到每天定时查询天气，若天气晴则无动作，若遇到异常天气（"雨", "雷", "尘", "霾", "雹", "雪", "雾", “沙”），则会调用 zcate-push-tool 通过 zcate 将异常天气信息发送到 zcate。

使用 zcate 发送消息请访问：https://www.qiansw.com/how-to-use-zcate-to-receive-zabbix-alarm-messages.html

# 使用截图

在有异常天气时,会收到下面的告警.

![weather](IMG_3122.png)
