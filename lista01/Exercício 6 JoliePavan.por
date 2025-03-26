programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    inteiro n
    inteiro i
    real x,y, soma
    i=1
    leia (n)
    faca {
      leia (x)
      i=i+1
      y= 5*(x-32)/9
      escreva (mat.arredondar (x,2)," FAHRENHEIT EQUIVALE A ",mat.arredondar(y,2), " CELSIUS \n")
    } enquanto (i <=n)
    
    
  }
}
