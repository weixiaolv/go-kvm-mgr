# KVM Manager tools

## Node

项目依赖 <https://github.com/digitalocean/go-libvirt>，一个纯 go 的框架，通过 RPC 方式调用 libvirt 服务，实现对 kvm 管理。

## Init

go-libvirt 使用前，需要自动生成对应版本的 go 源代码，以 libvirt7.0.0 为例，构建本地运行环境。

* 下载依赖环境
  
  ```sh
  sudo apt install python3-pip
  pip install docutils
  export PATH=$PATH:~/.local/bin
  sudo apt install meson libxml2-utils xsltproc libgnutls28-dev libxml2-dev
  
  mkdir build_libvirt
  cd build_libvirt
  
  git clone https://github.com/libvirt/libvirt.git
  cd libvirt
  git checkout v7.0.0
  meson setup build
  export LIBVIRT_SOURCE=~/projects/golang/go-kvm-mgr/build_libvirt/libvirt
  cd ../go-libvirt/
  go generate ./...
  
  cp $LIBVIRT_SOURCE/*.gen.go ../../libvirt/
  
  
  
  
  cd ..
  git clone https://github.com/digitalocean/go-libvirt.git
  cd go-libvirt
  
  sudo apt install 
  
  
  
  
  ```

* 构建版本运行源码

  ```sh

  ```

* 测试验证

  ```sh

  ```
  