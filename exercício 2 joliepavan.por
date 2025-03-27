programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    inteiro t, i, publico
    real perc_p, perc_g, perc_a, perc_c
    real renda, p, g, a, c

    leia(t)

    para (i = 1; i <= t; i++) 
    {
      leia(publico, perc_p, perc_g, perc_a, perc_c)
      p= (perc_p / 100.0) * publico
      g = (perc_g / 100.0) * publico
      a = (perc_a / 100.0) * publico
      c = (perc_c / 100.0) * publico

      renda = p + (g * 5.00) + (a * 10.00) + (c * 20.00)
      escreva("A RENDA DO JOGO N. ", i, " É = ", mat.arredondar(renda,2), "\n")
        }
    }
}
