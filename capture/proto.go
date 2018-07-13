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

package capture

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stdout",
	}
	return cfg.Build()
}

func Start(i string) {
	logger, _ := NewLogger()

	if handle, err := pcap.OpenLive(i, 65535, true, pcap.BlockForever); err != nil {
		logger.Error("err",
			zap.String("msg", err.Error()),
		)
		panic(err)
	} else if err := handle.SetBPFFilter("ip"); err != nil {
		logger.Error("err",
			zap.String("msg", err.Error()),
		)
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {

			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.LinkLayer() == nil {
				logger.Error("err",
					zap.String("msg", "null layer reference"),
					zap.Int("bytes", len(packet.Data())),
					zap.String("raw", packet.Dump()),
				)
				continue
			}

			logger.Info("packet",
				zap.String("src.ip", packet.NetworkLayer().NetworkFlow().Src().String()),
				zap.String("src.port", packet.TransportLayer().TransportFlow().Src().String()),
				zap.String("src.mac", packet.LinkLayer().LinkFlow().Src().String()),
				zap.String("dst.ip", packet.NetworkLayer().NetworkFlow().Dst().String()),
				zap.String("dst.port", packet.TransportLayer().TransportFlow().Dst().String()),
				zap.String("dst.mac", packet.LinkLayer().LinkFlow().Dst().String()),
				zap.Int("bytes", len(packet.Data())),
			)
		}
	}

}
