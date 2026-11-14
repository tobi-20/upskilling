package snowflake

import "errors"

const (
	sequenceBits        = 12
	workerBits          = 10
	customEpoch  uint64 = 1794508485000
	maxWorker           = 1<<workerBits - 1
	maxSequence         = 1<<sequenceBits - 1
)

func GenerateSnowflake(worker, currentTime, sequence uint64) int64 {
	timestamp := currentTime - customEpoch
	snowflake := timestamp<<(sequenceBits+workerBits) | worker<<sequenceBits | sequence
	return int64(snowflake)
}

type Generator struct {
	worker   uint64
	lastTime uint64
	sequence uint64
}

func NewGenerator(worker uint64) (*Generator, error) {
	if worker > maxWorker {
		return nil, errors.New("worker ID is out of range, it must be between 0-1023")
	}
	return &Generator{worker: worker}, nil
}

func (g *Generator) NextID(nowMs uint64) (int64, error) {
	if nowMs < g.lastTime {
		return 0, errors.New("clock moved backwards")
	}
	if nowMs == g.lastTime {
		if g.sequence == maxSequence {
			return 0, errors.New("sequence exhausted for this millisecond (max 4096 IDs per ms)")
		}
		g.sequence++
	} else {
		g.sequence = 0
	}
	g.lastTime = nowMs

	return GenerateSnowflake(g.worker, nowMs, g.sequence), nil
}
