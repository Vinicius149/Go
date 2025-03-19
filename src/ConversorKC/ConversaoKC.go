package main
import "fmt"

const ebulicaoK int = 373

func main(){
    ebulicaoC := ebulicaoK - 273
    fmt.Printf("A temperatura de ebulição da água em ºC é : %v ºC", ebulicaoC)
}