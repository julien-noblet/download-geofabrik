package download

import "sync"

const (
	// bufferSize is 128KB, optimal for high-throughput network and disk I/O.
	bufferSize = 128 * 1024
)

var bufferPool = sync.Pool{
	New: func() any {
		buf := make([]byte, bufferSize)

		return &buf
	},
}

// getBuffer acquires a 128KB buffer from the pool.
func getBuffer() *[]byte {
	val := bufferPool.Get()
	if buf, ok := val.(*[]byte); ok {
		return buf
	}

	buf := make([]byte, bufferSize)

	return &buf
}

// putBuffer returns a 128KB buffer back to the pool.
func putBuffer(buf *[]byte) {
	bufferPool.Put(buf)
}
