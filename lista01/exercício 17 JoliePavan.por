programa {
    funcao inicio() {
      inteiro x, y,i
      leia(x,y)
        se (x % 2 == 0)
        {
            i = 0
            enquanto (i < y) {
                escreva(x, " ")
                x = x + 2
                i = i + 1
            }
            escreva("\n")
        } senao {
            escreva("O PRIMEIRO NUMERO NÃO É PAR\n")
        }
    }
}
