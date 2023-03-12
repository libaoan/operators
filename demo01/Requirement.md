
# 通过client-go实现operator

## 需求说明
- 1. 利用 kubernetes v1.25.3 和 client-go v0.25.3 版本实现operator案例
- 2. operator监听k8s集群，获取service和ingress对象，根据变化作出以下相应
- 3. 当新增service 且annotation带ingress/http 属性时，新增ingress对象 
- 4. 当更新service对象，删除annotation ingress/http 属性时，删除ingress对象
- 5. 当更新service对象，新增annotation ingress/http 属性时，新增ingress对象
- 6. 当删除ingress对象时，新增ingress对象
    
# 涉及的依赖     
1、获取nginx controller v1.1.1，并完成集群加载
    [!https://kubernetes.github.io/ingress-nginx/deploy/][yaml配置链接]

# 部署方法

- run ingress-nginx controller
```shell
kubectl apply -f ingress-nginx-deploy-v1.1.1.yaml
```
- build demon01
```shell
go build -o test github.com/operators/demo01
./test
```

# 测试方法

1. 创建或在集群中已经存在的svc添加annotations标记
```shell
metadata:
  annotations:
    ingress/http: "true"
```

2. 通过kubectl get ingress 查看ingress对象自动被创建
```shell
stan@standeMacBook-Pro demo01 % kubectl get ingress
NAME         CLASS   HOSTS         ADDRESS   PORTS   AGE
kubernetes   nginx   example.com             80      6s
```
3. 其他删除ingress对象和删除svc对象场景同理论，不再单独罗列


