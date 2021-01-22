package bmp

import (
    "os"
)

type Bitmap struct {
    data []byte
    size int64
}

func (b Bitmap) Load(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    info, err := file.Stat()
    if err != nil {
        return err
    }
    b.size = info.Size()
    b.data = make([]byte, b.size)
    return nil
}

func (b Bitmap) Size() int64 {
    return b.size
}
