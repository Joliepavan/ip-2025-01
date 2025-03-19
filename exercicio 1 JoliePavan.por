programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

    real nota1, nota2, nota3, media_notas
    escreva ("Digite as notas:")
    leia (nota1) 
    leia (nota2)
    leia (nota3)

    media_notas = (nota1 + nota2 + nota3) / 3

    escreva("MÉDIA: ", mat.arredondar(media_notas,2))
    escreva ("\n")

    se (media_notas >= 6)
    {
      escreva ("APROVADO")
    }
    senao 
    {
      escreva ("REPROVADO")
    }
  }
}
