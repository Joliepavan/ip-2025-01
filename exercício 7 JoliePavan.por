programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real p,m,f,c
    leia (f)
    leia (p)
      c= 5*(f-32)/9
      m=25.4*p
      escreva ("O VALOR EM CELSIUS= ",mat.arredondar(c,2),"\n")
      escreva ("A QUANTIDADE DE CHUVA É= ", mat.arredondar(m,2),"\n")
  
  }
}

