package Chunk

import (
	"PerlinNoise"
	"math"
)

var TILE_SIZE = 16
var CHUNK_SIZE = 32 * 32
var PERLIN_SEED float32 = 600

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Chunk struct {
	ChunkID [2]int
	Map     map[Coordinate]Tile
}
type Tile struct {
	Key string
	X   int
	Y   int
}

func GetChunkID(x, y int) Coordinate {
	tileX := float64(x)
	tileY := float64(y)

	var ChunkID Coordinate
	if tileX < 0 {
		ChunkID.X = int(math.Floor(tileX / float64(TILE_SIZE)))
	} else {
		ChunkID.X = int(math.Ceil(tileX / float64(TILE_SIZE)))
	}
	if tileY < 0 {
		ChunkID.Y = int(math.Floor(tileY / float64(TILE_SIZE)))
	} else {
		ChunkID.Y = int(math.Ceil(tileY / float64(TILE_SIZE)))
	}
	if tileX == 0 {
		ChunkID.X = 1
	}
	if tileY == 0 {
		ChunkID.Y = 1
	}
	return ChunkID

}
func NewChunk(idChunk Coordinate) Chunk {
	// Помечаем чанк уникальным ИД
	chunk := Chunk{ChunkID: [2]int{idChunk.X, idChunk.Y}}
	// Максимальный координаты чанка
	var chunkXMax, chunkYMax int
	// Инициируем карту для тайлов
	var chunkMap map[Coordinate]Tile
	chunkMap = make(map[Coordinate]Tile)
	// Определяем максимальные координаты для чанка
	chunkXMax = idChunk.X * CHUNK_SIZE
	chunkYMax = idChunk.Y * CHUNK_SIZE
	switch {
	// логика генерации чанка для отрицательных координат
	case chunkXMax < 0 && chunkYMax < 0:
		{
			for x := chunkXMax + CHUNK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax + CHUNK_SIZE; y > chunkYMax; y -= TILE_SIZE {
					//Координаты для изображения тайла на клиенте
					posX := float32(x - (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)

					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.01:
						tile.Key = "~" //Вода
					case perlinValue >= -0.01 && perlinValue <= 0.5:
						tile.Key = "1" //Земля

					case perlinValue > 0.5:
						tile.Key = "^" // Горы
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
		//Для отрицательной X координаты
	case chunkXMax < 0:
		{
			for x := chunkXMax + CHUNK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax - CHUNK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x - (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))

					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)

					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "~"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "1"

					case perlinValue > 0.5:
						tile.Key = "^"
					}

					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
		// отрицательной Y координаты
	case chunkYMax < 0:
		{
			for x := chunkXMax - CHUNK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax + CHUNK_SIZE; y > chunkYMax; y -= TILE_SIZE {
					posX := float32(x + (TILE_SIZE / 2))
					posY := float32(y - (TILE_SIZE / 2))
					tile := Tile{}
					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "~"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "1"
					case perlinValue > 0.5:
						tile.Key = "^"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
		// Для положительных координат
	default:
		{
			for x := chunkXMax - CHUNK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax - CHUNK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x + (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}
					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "~"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "1"
					case perlinValue > 0.5:
						tile.Key = "^"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}

	}

	chunk.Map = chunkMap
	return chunk
}
