# Build
```
go build -ldflags="-s -w" -o uestclogin main.go
```
# run
```
./uestclogin
```
> 账号密码错误（注意 是在代码内填写账号）
```
./uestclogin
2020/12/10 12:22:37 status code: 200 
2020/12/10 12:22:37 IP has been online, please logout.
```
> 已在线 可触发账号下线
```
2020/12/10 12:23:02 status code: 200 
2020/12/10 12:23:02 E2532: The two authentication interval cannot be less than 10 seconds.(两次认证的间隔太短)
```
登陆成功
```
2020/12/10 12:23:15 status code: 200 
2020/12/10 12:23:15 login ok
```