package main

import (
	"fmt"
	"math/rand"
)

// Heroi representa um herói com nome, força e popularidade
type Heroi struct {
	nome         string
	forca        int
	popularidade int
}

// SimuladorBatalha simula uma batalha entre dois heróis
type SimuladorBatalha struct{}

// simularBatalha realiza a simulação da batalha entre dois heróis
func (sb SimuladorBatalha) simularBatalha(heroi1, heroi2 Heroi) {
	fmt.Printf("Iniciando batalha entre %s e %s!\n", heroi1.nome, heroi2.nome)

	// Calcula a força final de cada herói, considerando a popularidade
	forcaFinalHeroi1 := heroi1.forca + (heroi1.popularidade / 2)
	forcaFinalHeroi2 := heroi2.forca + (heroi2.popularidade / 2)

	// Inicializa fatores aleatórios
	fatorAleatorioHeroi1 := 0
	fatorAleatorioHeroi2 := 0

	// Adiciona um fator aleatório para cada herói (20% de chance)
	if rand.Intn(10) < 2 {
		fatorAleatorioHeroi1 = 15
		forcaFinalHeroi1 += fatorAleatorioHeroi1
	}
	if rand.Intn(10) < 2 {
		fatorAleatorioHeroi2 = 15
		forcaFinalHeroi2 += fatorAleatorioHeroi2
	}

	// Calcula a probabilidade de vitória do herói 1
	forcaTotal := forcaFinalHeroi1 + forcaFinalHeroi2
	chanceHeroi1 := float64(forcaFinalHeroi1) / float64(forcaTotal)

	fmt.Printf("%s força final: %d (força: %d, popularidade: %d)\n", heroi1.nome, forcaFinalHeroi1, heroi1.forca, heroi1.popularidade)
	fmt.Printf("%s força final: %d (força: %d, popularidade: %d)\n", heroi2.nome, forcaFinalHeroi2, heroi2.forca, heroi2.popularidade)

	// Determina o vencedor com base na probabilidade calculada
	vencedor := heroi2
	perdedor := heroi1
	if rand.Float64() < chanceHeroi1 {
		vencedor = heroi1
		perdedor = heroi2
	}
	fmt.Printf("%s venceu a batalha!\n", vencedor.nome)

	// Gera e exibe o relatório da batalha
	fmt.Println("\nRelatório da Batalha:")
	fmt.Println("----------------------")
	fmt.Printf("Vencedor: %s\n", vencedor.nome)
	fmt.Printf("Perdedor: %s\n", perdedor.nome)
	fmt.Println("\nDesempenho dos Heróis:")
	fmt.Printf("%s:\n", heroi1.nome)
	fmt.Printf("  Força inicial: %d\n", heroi1.forca)
	fmt.Printf("  Popularidade: %d (impacto na força: +%d)\n", heroi1.popularidade, heroi1.popularidade/2)
	fmt.Printf("  Fator aleatório: +%d\n", fatorAleatorioHeroi1)
	fmt.Printf("  Força final: %d\n", forcaFinalHeroi1)
	fmt.Printf("%s:\n", heroi2.nome)
	fmt.Printf("  Força inicial: %d\n", heroi2.forca)
	fmt.Printf("  Popularidade: %d (impacto na força: +%d)\n", heroi2.popularidade, heroi2.popularidade/2)
	fmt.Printf("  Fator aleatório: +%d\n", fatorAleatorioHeroi2)
	fmt.Printf("  Força final: %d\n", forcaFinalHeroi2)
	fmt.Println("\nImpacto da Batalha:")
	fmt.Printf("%s: Popularidade %+d, Força %+d\n", vencedor.nome, 5, 2)
	fmt.Printf("%s: Popularidade %+d, Força %+d\n", perdedor.nome, -3, -1)
}

func main() {
	heroi1 := Heroi{nome: "Homelander", forca: 100, popularidade: 80}
	heroi2 := Heroi{nome: "Homelander reverso", forca: 50, popularidade: 40}

	simulador := SimuladorBatalha{}
	simulador.simularBatalha(heroi1, heroi2)
}
