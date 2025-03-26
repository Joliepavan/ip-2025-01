programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    inteiro i,n
    i=2
    leia (n)
    se (5<n e n<2000)
    {
      enquanto (i <= n)
      {
        escreva (i, "^", i, " = ", i * i, "\n")
        i=i+2
      }
    }
  }
}
