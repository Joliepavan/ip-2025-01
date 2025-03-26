programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a,b,c
    inteiro delta
    leia (a)
    leia (b)
    leia (c)
    delta= (mat.potencia (b,2.0))-4*a*c
    escreva("O VALOR DE DELTA É = ", (mat.arredondar(delta,2)))
    
  }
}
