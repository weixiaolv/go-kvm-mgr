package libs

import (
    "log"

    "github.com/digitalocean/go-libvirt"
)

func DisConn(conn *libvirt.Libvirt)  {
    if err := conn.Disconnect(); err != nil {
        log.Printf("Failed to disconnect: %v", err)
    }
}

