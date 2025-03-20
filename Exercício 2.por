programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    real n1, n2, n3, X, Y
    escreva ("insira três números entre 1 e 9")
    escreva ("\n")
    leia (n1)
    leia (n2)
    leia (n3)

    X = n1*100 + n2 *10 + n3
    Y = mat.potencia(X, 2.0)

    se (n1<1 ou n1>9 ou n2>9 ou n3>9)
    escreva ("DIGITO INVALIDO")

    senao
    escreva (X, ",",Y)
  }
}
