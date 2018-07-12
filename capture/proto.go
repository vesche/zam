// Copyright © 2018 Austin Jackson <vesche@protonmail.com>
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

package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/Users/vesche/tmp",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   true, // disabled by default
	})

	listen()
}

func listen() {
	if handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			// this is mostly testing
			//if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.LinkLayer() == nil {
			//	fmt.Printf("%s\n", "error")
			//}
			fmt.Printf("%d - %s\n", len(packet.Data()), packet.LinkLayer().LinkFlow().Src().String() )
			// handlePacket(packet)  // Do something with a packet here.
		}
	}

}

	/*
	
	// defer packetHandle.Close()
	if err != nil {
		fmt.Println("handle error")
	}
	*/

		/*
	packetSource := gopacket.NewPacketSource(packetHandle, packetHandle.LinkType())

	for packet := range packetSource.Packets() {

		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.LinkLayer() == nil {
			//logger.Error("err",
			//	zap.String("msg", "null layer reference"),
			//	zap.Int("bytes", len(packet.Data())),
			//	zap.String("raw", packet.Dump()),
			//)

			// i might get rid of this
			fmt.Printf("%s: %s, %d", "null layer ref", len(packet.Data()), packet.Dump())
			continue
		}

		printf("TEST dst.ip %s - dst.port %s", packet.NetworkLayer().NetworkFlow().Dst().String(), packet.TransportLayer().TransportFlow().Dst().String())

		//logger.Info("packet",
		//	zap.String("src.ip", packet.NetworkLayer().NetworkFlow().Src().String()),
		//	zap.String("src.port", packet.TransportLayer().TransportFlow().Src().String()),
		//	zap.String("src.mac", packet.LinkLayer().LinkFlow().Src().String()),
		//	zap.String("dst.ip", packet.NetworkLayer().NetworkFlow().Dst().String()),
		//	zap.String("dst.port", packet.TransportLayer().TransportFlow().Dst().String()),
		//	zap.String("dst.mac", packet.LinkLayer().LinkFlow().Dst().String()),
		//	zap.Int("bytes", len(packet.Data())),
		//)
		*/


