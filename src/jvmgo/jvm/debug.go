package jvm

import (
    "fmt"
    rtc "jvmgo/jvm/rtda/class"
)

func LogJString(jStr *rtc.Obj) {
    valueField := jStr.Class().GetField("value", "[C")
    charArr := valueField.GetValue(jStr).(*rtc.Obj)
    chars := charArr.Fields().([]uint16)
    // todo
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}