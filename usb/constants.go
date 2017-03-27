// Copyright 2013 Google Inc.  All rights reserved.
// Copyright 2016 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usb

// #include <libusb.h>
import "C"

type Class uint8

const (
	ClassPerInterface Class = C.LIBUSB_CLASS_PER_INTERFACE
	ClassAudio        Class = C.LIBUSB_CLASS_AUDIO
	ClassComm         Class = C.LIBUSB_CLASS_COMM
	ClassHID          Class = C.LIBUSB_CLASS_HID
	ClassPrinter      Class = C.LIBUSB_CLASS_PRINTER
	ClassPTP          Class = C.LIBUSB_CLASS_PTP
	ClassMassStorage  Class = C.LIBUSB_CLASS_MASS_STORAGE
	ClassHub          Class = C.LIBUSB_CLASS_HUB
	ClassData         Class = C.LIBUSB_CLASS_DATA
	ClassWireless     Class = C.LIBUSB_CLASS_WIRELESS
	ClassApplication  Class = C.LIBUSB_CLASS_APPLICATION
	ClassVendorSpec   Class = C.LIBUSB_CLASS_VENDOR_SPEC
)

var classDescription = map[Class]string{
	ClassPerInterface: "per-interface",
	ClassAudio:        "audio",
	ClassComm:         "communications",
	ClassHID:          "human interface device",
	ClassPrinter:      "printer dclass",
	ClassPTP:          "picture transfer protocol",
	ClassMassStorage:  "mass storage",
	ClassHub:          "hub",
	ClassData:         "data",
	ClassWireless:     "wireless",
	ClassApplication:  "application",
	ClassVendorSpec:   "vendor-specific",
}

func (c Class) String() string {
	return classDescription[c]
}

type DescriptorType uint8

const (
	DescriptorTypeDevice    DescriptorType = C.LIBUSB_DT_DEVICE
	DescriptorTypeConfig    DescriptorType = C.LIBUSB_DT_CONFIG
	DescriptorTypeString    DescriptorType = C.LIBUSB_DT_STRING
	DescriptorTypeInterface DescriptorType = C.LIBUSB_DT_INTERFACE
	DescriptorTypeEndpoint  DescriptorType = C.LIBUSB_DT_ENDPOINT
	DescriptorTypeHID       DescriptorType = C.LIBUSB_DT_HID
	DescriptorTypeReport    DescriptorType = C.LIBUSB_DT_REPORT
	DescriptorTypePhysical  DescriptorType = C.LIBUSB_DT_PHYSICAL
	DescriptorTypeHub       DescriptorType = C.LIBUSB_DT_HUB
)

var descriptorTypeDescription = map[DescriptorType]string{
	DescriptorTypeDevice:    "device",
	DescriptorTypeConfig:    "configuration",
	DescriptorTypeString:    "string",
	DescriptorTypeInterface: "interface",
	DescriptorTypeEndpoint:  "endpoint",
	DescriptorTypeHID:       "HID",
	DescriptorTypeReport:    "HID report",
	DescriptorTypePhysical:  "physical",
	DescriptorTypeHub:       "hub",
}

func (dt DescriptorType) String() string {
	return descriptorTypeDescription[dt]
}

type EndpointDirection uint8

const (
	EndpointNumMask                         = 0x0f
	EndpointDirectionMask                   = 0x80
	EndpointDirectionIn   EndpointDirection = C.LIBUSB_ENDPOINT_IN
	EndpointDirectionOut  EndpointDirection = C.LIBUSB_ENDPOINT_OUT
)

var endpointDirectionDescription = map[EndpointDirection]string{
	EndpointDirectionIn:  "IN",
	EndpointDirectionOut: "OUT",
}

func (ed EndpointDirection) String() string {
	return endpointDirectionDescription[ed]
}

type TransferType uint8

const (
	TRANSFER_TYPE_CONTROL     TransferType = C.LIBUSB_TRANSFER_TYPE_CONTROL
	TRANSFER_TYPE_ISOCHRONOUS TransferType = C.LIBUSB_TRANSFER_TYPE_ISOCHRONOUS
	TRANSFER_TYPE_BULK        TransferType = C.LIBUSB_TRANSFER_TYPE_BULK
	TRANSFER_TYPE_INTERRUPT   TransferType = C.LIBUSB_TRANSFER_TYPE_INTERRUPT
	TRANSFER_TYPE_MASK        TransferType = 0x03
)

var transferTypeDescription = map[TransferType]string{
	TRANSFER_TYPE_CONTROL:     "control",
	TRANSFER_TYPE_ISOCHRONOUS: "isochronous",
	TRANSFER_TYPE_BULK:        "bulk",
	TRANSFER_TYPE_INTERRUPT:   "interrupt",
}

func (tt TransferType) String() string {
	return transferTypeDescription[tt]
}

type IsoSyncType uint8

const (
	ISO_SYNC_TYPE_NONE     IsoSyncType = C.LIBUSB_ISO_SYNC_TYPE_NONE << 2
	ISO_SYNC_TYPE_ASYNC    IsoSyncType = C.LIBUSB_ISO_SYNC_TYPE_ASYNC << 2
	ISO_SYNC_TYPE_ADAPTIVE IsoSyncType = C.LIBUSB_ISO_SYNC_TYPE_ADAPTIVE << 2
	ISO_SYNC_TYPE_SYNC     IsoSyncType = C.LIBUSB_ISO_SYNC_TYPE_SYNC << 2
	ISO_SYNC_TYPE_MASK     IsoSyncType = 0x0C
)

var isoSyncTypeDescription = map[IsoSyncType]string{
	ISO_SYNC_TYPE_NONE:     "unsynchronized",
	ISO_SYNC_TYPE_ASYNC:    "asynchronous",
	ISO_SYNC_TYPE_ADAPTIVE: "adaptive",
	ISO_SYNC_TYPE_SYNC:     "synchronous",
}

func (ist IsoSyncType) String() string {
	return isoSyncTypeDescription[ist]
}

type IsoUsageType uint8

const (
	ISO_USAGE_TYPE_DATA     IsoUsageType = C.LIBUSB_ISO_USAGE_TYPE_DATA << 4
	ISO_USAGE_TYPE_FEEDBACK IsoUsageType = C.LIBUSB_ISO_USAGE_TYPE_FEEDBACK << 4
	ISO_USAGE_TYPE_IMPLICIT IsoUsageType = C.LIBUSB_ISO_USAGE_TYPE_IMPLICIT << 4
	ISO_USAGE_TYPE_MASK     IsoUsageType = 0x30
)

var isoUsageTypeDescription = map[IsoUsageType]string{
	ISO_USAGE_TYPE_DATA:     "data",
	ISO_USAGE_TYPE_FEEDBACK: "feedback",
	ISO_USAGE_TYPE_IMPLICIT: "implicit data",
}

func (iut IsoUsageType) String() string {
	return isoUsageTypeDescription[iut]
}

type RequestType uint8

const (
	REQUEST_TYPE_STANDARD = C.LIBUSB_REQUEST_TYPE_STANDARD
	REQUEST_TYPE_CLASS    = C.LIBUSB_REQUEST_TYPE_CLASS
	REQUEST_TYPE_VENDOR   = C.LIBUSB_REQUEST_TYPE_VENDOR
	REQUEST_TYPE_RESERVED = C.LIBUSB_REQUEST_TYPE_RESERVED
)

var requestTypeDescription = map[RequestType]string{
	REQUEST_TYPE_STANDARD: "standard",
	REQUEST_TYPE_CLASS:    "class",
	REQUEST_TYPE_VENDOR:   "vendor",
	REQUEST_TYPE_RESERVED: "reserved",
}

func (rt RequestType) String() string {
	return requestTypeDescription[rt]
}
