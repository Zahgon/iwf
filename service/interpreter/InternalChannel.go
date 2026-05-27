package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
)

type InternalChannel struct {
	// key is channel name
	receivedData map[string][]*iwfidl.EncodedObject
}

func NewInternalChannel() *InternalChannel { _ = "STUB: not implemented"; return nil }

func RebuildInternalChannel(refill map[string][]*iwfidl.EncodedObject) *InternalChannel {
	_ = "STUB: not implemented"
	return nil
}

func (i *InternalChannel) GetAllReceived() map[string][]*iwfidl.EncodedObject {
	_ = "STUB: not implemented"
	return nil
}

func (i *InternalChannel) GetInfos() map[string]iwfidl.ChannelInfo {
	_ = "STUB: not implemented"
	return nil
}

func (i *InternalChannel) HasData(channelName string) bool { _ = "STUB: not implemented"; return false }

func (i *InternalChannel) ProcessPublishing(publishes []iwfidl.InterStateChannelPublishing) {
	_ = "STUB: not implemented"
	return
}

func (i *InternalChannel) receive(channelName string, data *iwfidl.EncodedObject) {
	_ = "STUB: not implemented"
	return
}

func (i *InternalChannel) Retrieve(channelName string) *iwfidl.EncodedObject {
	_ = "STUB: not implemented"
	return nil
}
