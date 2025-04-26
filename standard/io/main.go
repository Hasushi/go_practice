package main

import (
	"fmt"
	"io"
	"strings"
)

// 1. io.Writerインターフェースを実装する構造体を定義
type MyWriter struct{}


// 2. Writeメソッドを実装
func (w *MyWriter) Write(p []byte) (n int, err error) {
	// 受け取ったバイトスライスを文字列に変換して表示
	str := string(p)
	fmt.Println("MyWriter:", str)
	return len(p), nil // 書き込んだバイト数を返す
}


// 3. io.Readerインターフェースを実装する構造体を定義
type MyReader struct {
	data string
	index int
}

// 4. Readメソッドを実装
func (r *MyReader) Read(p []byte) (n int, err error) {
	if r.index >= len(r.data) {
		return 0, io.EOF // 読み込むデータがない場合はEOFを返す
	}
	n = copy(p, r.data[r.index:])
	r.index += n
	return n, nil // 読み込んだバイト数を返す
}

// 5. io.Closerインターフェースを実装する構造体を定義
type MyCloser struct {
	data string
}

// 6. Closeメソッドを実装
func (c *MyCloser) Close() error {
	fmt.Println("Closing MyCloser")
	return nil // クローズ処理が成功した場合はnilを返す
}

// 7. io.Seekerインターフェースを実装する構造体を定義
type MySeeker struct {
	data string
	index int
}

// 8. Seekメソッドを実装
func (s *MySeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		s.index = int(offset)
	case io.SeekCurrent:
		s.index += int(offset)
	case io.SeekEnd:
		s.index = len(s.data) + int(offset)
	default:
		return 0, fmt.Errorf("invalid whence")
	}
	if s.index < 0 {
		s.index = 0
	} else if s.index > len(s.data) {
		s.index = len(s.data)
	}
	return int64(s.index), nil // 新しいインデックスを返す
}


// 9. io.StringWriterインターフェースを実装する構造体を定義
type MyStringWriter struct {
	data strings.Builder
}

// 10. WriteStringメソッドを実装
func (w *MyStringWriter) WriteString(s string) (n int, err error) {

	w.data.WriteString(s) // 文字列を追加
	return len(s), nil // 書き込んだバイト数を返す
}


// 11. io.StringReaderインターフェースを実装する構造体を定義
type MyStringReader struct {
	data string
	index int
}

// 12. ReadStringメソッドを実装
func (r *MyStringReader) ReadString(delim byte) (string, error) {
	if r.index >= len(r.data) {
		return "", io.EOF // 読み込むデータがない場合はEOFを返す
	}
	end := strings.IndexByte(r.data[r.index:], delim)
	if end == -1 {
		end = len(r.data) - r.index
	} else {
		end += r.index
	}
	str := r.data[r.index:end]
	r.index = end + 1 // デリミタの次の位置にインデックスを更新
	return str, nil // 読み込んだ文字列を返す
}


// 13. io.PipeReaderとio.PipeWriterを使用する例
type MyPipe struct {
	pipeReader *io.PipeReader
	pipeWriter *io.PipeWriter
}


// 14. NewPipe関数を定義
func NewPipe() *MyPipe {
	pipeReader, pipeWriter := io.Pipe()
	return &MyPipe{
		pipeReader: pipeReader,
		pipeWriter: pipeWriter,
	}
}


// 15. Writeメソッドを実装
func (p *MyPipe) Write(data []byte) (n int, err error) {
	n, err = p.pipeWriter.Write(data) // パイプにデータを書き込む
	if err != nil {
		return 0, err // エラーが発生した場合はエラーを返す
	}
	return n, nil // 書き込んだバイト数を返す
}

// 16. Readメソッドを実装
func (p *MyPipe) Read(pData []byte) (n int, err error) {
	n, err = p.pipeReader.Read(pData) // パイプからデータを読み込む
	if err != nil {
		return 0, err // エラーが発生した場合はエラーを返す
	}
	return n, nil // 読み込んだバイト数を返す
}

// 17. Closeメソッドを実装
func (p *MyPipe) Close() error {
	err := p.pipeWriter.Close() // パイプライターをクローズ
	if err != nil {
		return err // エラーが発生した場合はエラーを返す
	}
	err = p.pipeReader.Close() // パイプリーダーをクローズ
	if err != nil {
		return err // エラーが発生した場合はエラーを返す
	}
	return nil // クローズ処理が成功した場合はnilを返す
}


// 18. main関数を定義
func main() {
	// 19. MyWriterを使用してデータを書き込む
	writer := &MyWriter{}
	data := []byte("Hello, World!")
	writer.Write(data)

	// 20. MyReaderを使用してデータを読み込む
	reader := &MyReader{data: "Hello, World!"}
	buf := make([]byte, 5)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Read:", string(buf[:n]))
	}

	// 21. MyCloserを使用してクローズ処理を行う
	closer := &MyCloser{data: "Some data"}
	err = closer.Close()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 22. MySeekerを使用してシーク処理を行う
	seeker := &MySeeker{data: "Hello, World!"}
	offset, err := seeker.Seek(7, io.SeekCurrent)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Seeked to offset:", offset)
	}
	offset, err = seeker.Seek(2, io.SeekCurrent)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Seeked to offset:", offset)
	}

	// 23. MyStringWriterを使用して文字列を書き込む
	stringWriter := &MyStringWriter{}
	n, err = stringWriter.WriteString("Hello, String!")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Written bytes:", n)
	}

	// 24. MyStringReaderを使用して文字列を読み込む
	stringReader := &MyStringReader{data: "Hello, String!"}
	str, err := stringReader.ReadString(',')
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Read string:", str)
	}

	// 25. MyPipeを使用してパイプ処理を行う
	myPipe := NewPipe()
	go func() {
		defer myPipe.Close()
		myPipe.Write([]byte("Hello from Pipe!"))
	}()

	buf = make([]byte, 20)
	n, err = myPipe.Read(buf)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Read from pipe:", string(buf[:n]))
	}
}
