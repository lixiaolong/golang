# nmap_conf

## 用途

windows的cmd对命令行有长度限制，太长的参数列表会无法执行，所以从配置文件中读取nmap的参数，然后执行。

程序中调用了exec.Command去执行命令，经测试，也不能支持无限长的命令，长度2w左右会出现执行很快返回，无结果的情况。不过8000个字符可以正常执行。

## 首次运行

首次运行，会生成配置文件，之后就可以修改配置文件来运行了。

配置文件conf.ini，内容如下

```
[default]
nmap_path = nmap        # nmap的可执行文件路径
target = 192.168.0.1    # 扫描目标
port_range = 80,443     # 端口范围
sub_param = -O          # 子命令，多个子命令需要用空格隔开
stdout = stdout.log     # 标准输出文件
stderr = stderr.log     # 错误输出文件
```
