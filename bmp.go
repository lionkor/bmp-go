package bmp

type BitmapHeader struct {
    data [2 + 4 + 2 + 2 + 4]byte
}

type Bitmap struct {
    header BitmapHeader
}
