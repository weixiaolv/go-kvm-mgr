package libs

import (
    "bytes"
)

func VolNameBuilder(vmName string, suffix string)  string {
    var stringVolBuilder bytes.Buffer
    stringVolBuilder.WriteString(vmName)
    stringVolBuilder.WriteString(suffix)
    volName := stringVolBuilder.String()
    return volName
}
