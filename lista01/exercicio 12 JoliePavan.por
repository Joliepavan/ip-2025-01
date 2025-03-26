programa {
  inclua biblioteca Matematica--> mat
  funcao inicio() {
    inteiro n
    real x
    leia (n)
    se (n % 3 ==0)
    {
   x=n*10/3
    } senao {
      x=((n-n%3)/3) *10 + (n%3)*5
    }
    escreva("O VALOR A PAGAR É = ",x, "\n")
  }
}
