package main

import (
	// "fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
	"fmt"
)

var (
	device       string = "enp0s8"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
	buffer 		 gopacket.SerializeBuffer
	options		 gopacket.SerializeOptions
	
	seq			uint32  = 1664505538
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	
	// パケットの書き込み
	// rawBytes := []byte{10, 20, 30, 40, 50, 60, 70, 80} // see ASCII code
	rawBytes := "this is payload"
	// err = handle.WritePacketData(rawBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// ヘッダ情報
	options := gopacket.SerializeOptions{}
	buffer := gopacket.NewSerializeBuffer()
	// gopacket.SerializeLayers(buffer, options, 
	// 	&layers.Ethernet{},
	// 	&layers.IPv4{},
	// 	&layers.TCP{},
	// 	gopacket.Payload(rawBytes),
	// )
	// outgoingPacket := buffer.Bytes()
	// err = handle.WritePacketData(outgoingPacket)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0x08, 0x00, 0x27, 0x5a, 0xc5, 0xbc},
		DstMAC: net.HardwareAddr{0x0a, 0x00, 0x27, 0x00, 0x00, 0x02},
		EthernetType: layers.EthernetTypeIPv4,
	}

	ipLayer := &layers.IPv4{
		SrcIP: net.IP{192, 168, 33, 10},
		DstIP: net.IP{192, 168, 33, 1},
		Protocol: layers.IPProtocolTCP,
		Version: uint8(4),
		IHL    : uint8(5),
		TTL    : uint8(64),
	}
	
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(3000),
		DstPort: layers.TCPPort(3000),
		Seq    : seq,
		DataOffset: uint8(5),
	}
	
	// パケットの生成
	// buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(rawBytes),
	)	
	outgoingPacket := buffer.Bytes()
	fmt.Println(outgoingPacket)	
	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal(err)
	}
	
}