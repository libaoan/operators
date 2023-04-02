# operators 开发实例

## 开发环境(on mac m1)

### Docker Desktop + Kind

- 通过docker 官网下载 Docker Desktop App， 通过拖拽方式安装并启动
- 通过kubernetes下载Kind
- 通过kind命令创建集群（可以创建多节点集群）

#### 优点
- 兼容M1，安装快速
- Kind支持多节点集群

#### 缺点
- docker desktop 耗费资源多
- Kind 创建的集群，无法直接通过local主机访问（kubernetes in Docker）

### Docker + Minikube

为什么要替换Docker Desktop 可以参考这边文章
[我如何从 Docker Desktop 切换到 Colima](https://t.cj.sina.com.cn/articles/view/1772191555/69a17f43019014old)

- 通过 brew install colima 安装(可能会找不到依赖，可以通过brew install 手动安装依赖)
- 可以通过本地编译方式安装colima [Mac M1 安装colima，替代Docker for mac Desktop](https://www.jianshu.com/p/963392b3eb4b])
- [关于Clima 虚拟机](https://zhuanlan.zhihu.com/p/466229156)  
- brew install minikube
  或者直接通过官方链接下载
  ```shell
    curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-arm64
    sudo install minikube-darwin-arm64 /usr/local/bin/minikube
  ```
- brew install kubectl



