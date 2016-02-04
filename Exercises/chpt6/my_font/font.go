package my_font/font

type Font struct{
    Size float32
    Family string
 }

func New(size float32, family string)*Font {
    return &Font{size, family}
}
