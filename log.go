package basm

import "runtime"

//go:wasmimport env bufferLog
func _hostFuncBufferLog(ptr, size uint32)

func hostFuncBufferLog(msg string) {
	msgData := []byte(msg)
	inOffset, inLen := bytesToOffsetSize(msgData)
	_hostFuncBufferLog(inOffset, inLen)
	runtime.KeepAlive(msgData)
}
