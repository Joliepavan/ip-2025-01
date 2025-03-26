programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real n
    leia (n)
    se (n<6)
    {
      escreva ("NOTA =", n, " CONCEITO = D")
    } senao se (n >= 6 e n<7.5)
    {
      escreva ("NOTA =", n, " CONCEITO = C")
    } senao se (n >= 7.5 e n<9)
    {
      escreva ("NOTA =", n, " CONCEITO = B")
    } senao se (n>=9)
    {
      escreva ("NOTA =", n, " CONCEITO = A")
    }
  }
}
