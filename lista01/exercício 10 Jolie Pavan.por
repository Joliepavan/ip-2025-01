programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a,b,c,d, det
    leia(a,b,c,d)
    det= a*d - b*c
    escreva ("O VALOR DO DETERMINANTE É= ", mat.arredondar (det,2),"\n")
  }
}
