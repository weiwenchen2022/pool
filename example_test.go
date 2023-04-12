package pool_test

import (
	"io"
	"os"
	"time"

	"github.com/weiwenchen2022/pool"
)

type buf struct {
	data []byte
}

func (b *buf) WriteString(s string) (int, error) {
	b.data = append(b.data, s...)
	return len(s), nil
}

func (b *buf) WriteByte(c byte) error {
	b.data = append(b.data, c)
	return nil
}

func (b *buf) Bytes() []byte {
	return b.data
}

func (b *buf) Reset() {
	b.data = b.data[:0]
}

var bufPool pool.Pool[buf]

func ExamplePool() {
	Log(os.Stdout, "path", "/search?q=flowers")
	// Output: 2006-01-02T15:04:05Z path=/search?q=flowers
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get()
	defer bufPool.Put(b)

	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}
