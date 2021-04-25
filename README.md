## `DroneCI` with `Gitea` for `golang` project

使用`Gitea`和`DroneCI`为 `golang` 项目打造的 `devops` 项目模版。

特点：
1. `Gitea` 开源的 git 仓库系统，使用 `ssh` 方式进行 `clone`，需要在 `gitea` 的仓库配置 `ssh` 部署密钥；
2. `DroneCI` 是用 `golang` 开发轻量级 `CI/CD` 工具;
3. 已经为 `golang` 项目进行流程定制，方便扩展使用;
4. 提供基础 `Dockerfile` 对 `golang` 程序打包发布 `Docker` 镜像;
5. `DroneCI` build 结束之后，自动发送钉钉提醒。

### `Dockerfile` 配置参考
> 我们将 `alpine` linux 系统的软件源替换为清华大学的镜像地址。（！！阿里云的 alpine 镜像测试出现404问题）。

```dockerfile
FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
COPY ./test /bin

ENTRYPOINT ["/bin/test"]
```

### DroneCI 的 `.drone.yml` 配置解析
仓库中 `.drone.yml` 提供了3个 `pipeline`，分别是 `default`, `tag`, `promote`，所有的 `pipeline`, 均需要运行在 `docker` 中。
3个 `pipeline` 对应场景:
1. `default` 为默认场景，适合 `master` 和 `develop` 分支在被 `push` 之后进行 `golangci-lint`、 `test` (单元测试) 和 `build`;
2. `tag` 应用在 `git` 仓库创建 `tag` 的场景，`git` 创建 `tag` 之后，将会依次自动进行 `golangci-lint`, `test`, `build`, `publish`(发布 `docker` 镜像);
3. `promote` 是在`DroneCI`的后台`promote`之后，将会依次自动进行 `golangci-lint`, `test`, `build`, `publish`(发布 `docker` 镜像)，你可以在 `publish` 后面增加 `step` 进行 `deploy`;

#### `pipeline` 流程结构

```text
.drone-pipeline
├── default
│   ├── clone
│   ├── linter
│   ├── test
│   ├── build
│   └── notification
├── tag
│   ├── clone
│   ├── linter
│   ├── test
│   ├── build
│   ├── publish
│   └── notification
└── promote
    ├── clone
    ├── linter
    ├── test
    ├── build
    ├── publish
    ├── deploy(complete this step by your self)
    └── notification
```

### `DroneCI` 配置 `secrets` 变量说明
> 此仓库中 `.drone.yml` 文件配置的 `pipeline` 正确运行，需要配置以下 `secrets` 变量。

1. `SSH_HOST`, `git` 仓库的 `ssh` `clone` 地址的 `IP` 或者域名, 例如: `1.2.3.4`, `gitea.example.com`, `github.com`;(`clone step` 使用)
2. `SSH_PORT`, `git` 仓库的默认端口为 `22` 的时候可以不设置，否则需要配置 `git` server 的 `ssh` 端口;(`clone step` 使用)
3. `SSH_KEY`, `git` 仓库中设置用户部署的时候进行 `ssh`登录`private key`，需要对应 `git` 仓库中设置的用于部署 `public key`;(`clone step` 使用)
4. `DOCKER_REPO`, `docker` 镜像的地址，需要包含域名，例如: `registry.example.com/user/repo`;(`publish step` 使用)
5. `DOCKER_REGISTRY`, `docker` 镜像的注册中心地址, 例如: `registry.example.com`;(`publish step` 使用)
6. `DOCKER_REGISTRY_USERNAME`, 用户登录 `docker` 镜像注册中心的账号；(`publish step` 使用)
7. `DOCKER_REGISTRY_PASSWORD`, 用户登录 `docker` 镜像注册中心的密码;(`publish step` 使用)
8. `GOPROXY`, `Golang` 的依赖代理配置;
9. `DINGTALK_TOKEN`, `Dingtalk` 的群 `bot` 消息 `token`;(`notification step` 使用)
10. `DINGTALK_SECRET`, `Dingtalk` 的群 `bot` 消息 `secret`;(`notification step` 使用)