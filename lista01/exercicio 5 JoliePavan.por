programa {
  funcao inicio() {
    inteiro u
    real f, n
    caracter cons
    leia (u, f, cons)
    escreva ("CONTA:", u)
    escreva("\n")
    se (cons == "R")
    {
      n = 5.00 + (0.05 * f)
    } senao se (cons == "C")
    {se (f>=80)
    {
      n=500.00
    } senao{
      n = 500.00 + (0.25 * (f - 80))
    }
    } senao se (cons == "I")
    {se f<=100{
		    n=800.00
		}senao {
		    n=800+ (0.04 * (f-100))
    }
    }

  }
}
