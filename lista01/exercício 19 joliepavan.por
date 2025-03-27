programa {
  inclua biblioteca Matematica --> mat
    funcao inicio() {
    inteiro n,i
    real soma = 0.0

    leia(n)
      se (n <= 1) 
      {
      escreva("Numero invalido!\n")
      } 
      senao 
      {
        para (i = 1; i <= n; i++) 
        {
        soma = soma + (1.0 / i)
        }
        escreva(mat.arredondar (soma,6))
      }
    }
}
