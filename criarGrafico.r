args = commandArgs(trailingOnly = TRUE)
if(length(args) < 1) {
    stop("Nome do arquivo para entrada não encontrado!")
}

require('ggplot2')

# png('resultado.png')
svg('resultado.svg')

f = read.csv(args[1]) # Le o arquivo passado como parametro

f <- within(f, {
    tempo <- as.numeric(as.character(tempo)) # Converte o tempo de string para numerico
    tamanhovetor <- as.numeric(as.character(tamanhovetor)) # Converte tamanho do vetor de string para numérico
})


if(length(args) == 2) {
    f = f[grep(args[2], f$metodo), ]
}

options(scipen=max(f$tempo)) # Tira notação científica

Método = f$metodo

qplot(x=f$tamanhovetor, y=f$tempo, colour = Método, xlab = "Tamanho do vetor", ylab = "Tempo consumido (microsegundos/op)")

dev.off()
