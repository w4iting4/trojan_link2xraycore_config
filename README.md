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
![image](https://github.com/w4iting4/trojan_link2xraycore_config/assets/41547947/a7a7f43d-e21b-4c8e-95cb-78e87de608fe)

最后会在本地生成一个`result.txt`,文件中你需要将domian替换为你运行`xray`的服务器`ip`地址。
随后你可以使用:https://github.com/w4iting4/socks_enable 来检查`socks`的有效性