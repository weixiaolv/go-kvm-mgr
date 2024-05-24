package utils

import (
    "bytes"
    "log"
    "net/url"

    "github.com/digitalocean/go-libvirt"
)

func GetConn(host string) (*libvirt.Libvirt, error) {
    kvmHost := host
    var stringUrlBuilder bytes.Buffer
    // 把字符串写入缓冲
    stringUrlBuilder.WriteString("qemu+ssh://root@")
    stringUrlBuilder.WriteString(kvmHost)
    stringUrlBuilder.WriteString("/system")

    //uri, _ := url.Parse(string(libvirt.QEMUSystem))
    uri, err := url.Parse(stringUrlBuilder.String())
    if err != nil {
        log.Printf("url error: %v", err)
        return nil, err
    }
    l, err := libvirt.ConnectToURI(uri)
    if err != nil {
        log.Printf("failed to connect: %v", err)
        return nil, err
    }
    return l, nil
}
