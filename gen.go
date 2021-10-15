package main

/*
#cgo LDFLAGS: -L./lib/rustgen/target/release/ -lrustgen
#include "./lib/rustgen/rustgen.h"
*/
import "C"

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"unsafe"
)

type RustGen struct{}

var (
	air, _   = world.BlockRuntimeID(block.Air{})
	stone, _ = world.BlockRuntimeID(block.Stone{})
	grass, _ = world.BlockRuntimeID(block.Grass{})
	dirt, _  = world.BlockRuntimeID(block.Dirt{})
	water, _ = world.BlockRuntimeID(block.Water{Still: true, Falling: false, Depth: 8})
	lava, _  = world.BlockRuntimeID(block.Lava{Still: true, Falling: false, Depth: 8})
)

func NewRustGen() RustGen {
	return RustGen{}
}

// GenerateChunk ...
func (RustGen) GenerateChunk(pos world.ChunkPos, chunk *chunk.Chunk) {
	chunkData, free := generateChunkData(1548577427, C.int(pos.X()), C.int(pos.Z()))
	n := 0
	for y := int16(0); y < 256; y++ {
		for z := uint8(0); z < 16; z++ {
			for x := uint8(0); x < 16; func() { x++; n++ }() {
				chunk.SetRuntimeID(x, y, z, 0, mapBlockIDToBlock(uint8(chunkData[n])))
			}
		}
	}
	free() // frees the chunk data
}

func generateChunkData(seed, chunkX, chunkZ C.int) (chunkData []C.uint8_t, free func()) {
	cData := C.gen_chunk(seed, chunkX, chunkZ)
	l := cData.len
	data := (*[1 << 30]C.uint8_t)(unsafe.Pointer(cData.data))[:l:l]
	return data, func() {
		C.free_chunk_data(cData)
	}
}

func mapBlockIDToBlock(blockId uint8) uint32 {
	switch blockId {
	case 0:
		return air
	case 1:
		return stone
	case 2:
		return grass
	case 3:
		return dirt
	case 8:
		return water
	case 9:
		return water
	case 10:
		return lava
	case 11:
		return lava
	default:
		return air
	}
}
