# KVM Manager tools

## Node

项目依赖 <https://github.com/digitalocean/go-libvirt>，一个纯 go 的框架，通过 RPC 方式调用 libvirt 服务，实现对 kvm 管理，当前版本 go-libvirt 默认使用 `libvirt v8.0.0`。

## 更改 libvirt 版本

go-libvirt 默认版本较新，一般不需要调整，但有的函数调用可能会不支持。可通过代码生成方式，使用其他版本，以 libvirt7.0.0 为例，构建本地运行环境。

* 程序说明

  * `libvirt`，虚拟化服务端工具，通过 `meson` 构建（meson 依赖 python3），开发环境并不需要完全 build，但需要使用 configure 生成的依赖文件
  * `go-libvirt`，可通过代码 `go generate ./...`（转换 libvirt c 接口为 go 代码），支持其他版本

* 下载依赖环境
  
  基础环境

  ```sh
  # python docutils
  sudo apt install python3-pip
  pip3 install docutils

  # install meson and dependence packkage
  sudo apt install meson libxml2-utils xsltproc libgnutls28-dev libxml2-dev libtirpc-dev
  
  # python execute file
  export PATH=$PATH:~/.local/bin
  ```

  构建目录

  ```sh
  # dir for generate
  mkdir build_libvirt
  cd build_libvirt
  ```

  libvirt 环境
  
  ```sh
  # build libvirt run env
  git clone https://github.com/libvirt/libvirt.git
  cd libvirt
  git checkout v7.0.0
  meson setup build
  export LIBVIRT_SOURCE=`pwd`
  cd ..
  ```

  generate libvirt v7.0.0 rpc code

  ```sh
  # build libvirt source code
  git clone https://github.com/digitalocean/go-libvirt.git
  cd go-libvirt

  # generate code tools (go generate auto install in ../bin)
  # go install golang.org/x/tools/cmd/goyacc@latest
  # go install github.com/xlab/c-for-go@latest
  # export PATH=$PATH:~/go/bin

  export PATH=$PATH:`pwd`/bin
  go generate ./...

  # override pkg version
  cp -r ./* ~/go/pkg/mod/github.com/digitalocean/go-libvirt@v0.0.0-20210615174804-eaff166426e3/

  cd ..
  ```

  > go-libvirt 的运行代码和生成工具放在一起，也可以用 go dep 做成本地依赖

  清理

  ```sh
  cd ..
  rm -rf build_libvirt/
  ```
