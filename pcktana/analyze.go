// package pcktana
package main

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	pf := "./data/test.pcapng"

	// Open file instead of device
	h, err := pcap.OpenOffline(pf)
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	flows := make(map[gopacket.Flow]int, 0)

	pSrc := gopacket.NewPacketSource(h, h.LinkType())
	for p := range pSrc.Packets() {
		flows[p.NetworkLayer().NetworkFlow()]++
		// fmt.Println(p.NetworkLayer().NetworkFlow().Src())
		// fmt.Println(p.NetworkLayer().NetworkFlow().Dst())
	}

	// fmt.Println(flows)

}
