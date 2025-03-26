programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real h,a,v,ab
    leia (h)
    leia (a)
    ab= 3*(mat.potencia (a,2.0))* (mat.raiz (3,2.0))/2
    v=ab*h/3
    escreva("O VOLUME DA PIRAMIDE É = ",(mat.arredondar(v,2))," METROS CUBICOS")
  }
}
