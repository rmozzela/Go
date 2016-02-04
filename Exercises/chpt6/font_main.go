package main

import (
    "fmt"
//    "my_font/font"
)

type Font struct{
    size float32
    Family string
 }

func New(size float32, Family string) *Font {
    return &Font{size, Family}
}

func (font *Font) SetSize(size float32) (err error) {

  if size < 5 || size > 144 {
    return fmt.Errorf("Font Size needs to be between 5 and 144 %f",size)
  } else { font.size = size
    return nil
    }
    return nil
}

func (font *Font) SetFamily(Family string)(err error){
  if Family == ""{
    return fmt.Errorf("Family can not be blank %s",Family)
  } else {font.Family = Family
  return nil
  }
  return nil
}
//add getters
func (font Font) GetFamily() string{
  return font.Family
}

func (font Font) GetSize() float32{
  return font.size
}

func (font *Font) String() string {
  return fmt.Sprintf("%f, %s", font.size, font.Family)
}

func main() {
    titleFont := New(12, "Times Roman")
    if err := titleFont.SetSize(145); err != nil {
       fmt.Println(err)
     } else { titleFont.GetSize() //set to current size;
     }

    if err := titleFont.SetFamily(""); err != nil {
       fmt.Println(err)
      }else { titleFont.GetFamily() //set to current size;
    }
    fmt.Println(titleFont)
}
