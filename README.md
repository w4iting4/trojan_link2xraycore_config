# trojan_link2xraycore_config
将trojan链接转换为xraycore中的outbound出口，并绑定对应的入站socks。
## 使用说明
```
waiting4@waiting4deMacBook-Pro trojan_link2xraycore_config % go run main.go struct.go -h
  -c string
        Path to the config.json file
  -p int
        Starting value for tag counter (default 10000)
  -pwd string
        Password for inbound settings
  -t string
        Path to the trojan.txt file
  -u string
        Username for inbound settings
```
`-c 指定config.json的路径`

`-p 指定任务起始的端口`

`-pwd 指定socks的统一密码`

`-u 指定socks的统一用户` 

`-t 指定trojan.txt的位置，一行一个`
## 效果
生成完`config.json`，之后上传到目标服务器中,使用`xray run -c config.json` 即可完成`xray`的`socks`监听。