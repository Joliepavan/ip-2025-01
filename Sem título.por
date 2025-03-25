programa 
{
  inclua biblioteca Matematica --> mat
  funcao inicio() 
  {
    real salario_minimo, qkw, custokw, custoconsumo, custodesc

    leia (salario_minimo)
    leia (qkw)
    custokw= (salario_minimo * 7) / 1000
    custoconsumo= (custokw * qkw)
    custodesc= custoconsumo - (custoconsumo * 0.1)
    escreva ("Custo por kW: R$", mat.arredondar (custokw,2))
    escreva("\n")
    escreva ("Custo do consumo: R$", mat.arredondar (custoconsumo,2))
    escreva ("\n")
    escreva ("Custo com desconto: R$", mat.arredondar (custodesc,2))
  }
}
