programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real r,h, at, ac, al, x
    leia(r)
    leia (h)
    ac=(mat.potencia(r,2.0))*3.14159
    al= 2*r*h*3.14159
    at= 2*ac+al
    x=at*100
    escreva ("O VALOR DO CUSTO É= ", (mat.arredondar (x,2)))
  }
}
