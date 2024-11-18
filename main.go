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

// chanceDeAcertoCritico calcula a chance de acerto crítico com base na popularidade do herói
func (sb SimuladorBatalha) chanceDeAcertoCritico(popularidade int) float64 {
	// A chance de acerto crítico é proporcional à popularidade (máximo de 50%)
	chance := float64(popularidade) / 2.0
	if chance > 50 {
		chance = 50 // Limita a chance de crítico a 50%
	}
	return chance
}

// calculaMoral ajusta a moral com base na diferença de força durante a batalha
func (sb SimuladorBatalha) calculaMoral(forcaHeroi1, forcaHeroi2 int) int {
	// Se o herói 1 está ganhando, sua moral aumenta
	moral := 0
	if forcaHeroi1 > forcaHeroi2 {
		moral = 5
	} else if forcaHeroi2 > forcaHeroi1 {
		moral = -5
	}
	return moral
}

// simularBatalha realiza a simulação da batalha entre dois heróis
func (sb SimuladorBatalha) simularBatalha(heroi1, heroi2 Heroi) {
	fmt.Printf("Iniciando batalha entre %s e %s!\n", heroi1.nome, heroi2.nome)

	// Calcula a força final de cada herói, considerando a popularidade
	forcaFinalHeroi1 := heroi1.forca + (heroi1.popularidade / 2)
	forcaFinalHeroi2 := heroi2.forca + (heroi2.popularidade / 2)

	// Verifica a chance de acerto crítico para cada herói
	chanceCriticoHeroi1 := sb.chanceDeAcertoCritico(heroi1.popularidade)
	chanceCriticoHeroi2 := sb.chanceDeAcertoCritico(heroi2.popularidade)

	// Variáveis para registrar se o herói teve um acerto crítico
	criticoHeroi1 := false
	criticoHeroi2 := false

	// Adiciona o impacto de acerto crítico (10% a mais de força no ataque crítico)
	if rand.Float64()*100 < chanceCriticoHeroi1 {
		criticoHeroi1 = true
		forcaFinalHeroi1 += 20 // Aumenta a força com um bônus de crítico
	}

	if rand.Float64()*100 < chanceCriticoHeroi2 {
		criticoHeroi2 = true
		forcaFinalHeroi2 += 20 // Aumenta a força com um bônus de crítico
	}

	// Inicializa fatores aleatórios (20% de chance)
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

	// Calcula a moral temporária com base nas forças relativas
	moralHeroi1 := sb.calculaMoral(forcaFinalHeroi1, forcaFinalHeroi2)
	moralHeroi2 := sb.calculaMoral(forcaFinalHeroi2, forcaFinalHeroi1)

	// Aplica a moral no cálculo final da força
	forcaFinalHeroi1 += moralHeroi1
	forcaFinalHeroi2 += moralHeroi2

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

	// Relatório do Herói 1
	fmt.Printf("%s:\n", heroi1.nome)
	fmt.Printf("  Força inicial: %d\n", heroi1.forca)
	fmt.Printf("  Popularidade: %d (impacto na força: +%d)\n", heroi1.popularidade, heroi1.popularidade/2)
	fmt.Printf("  Fator aleatório: +%d\n", fatorAleatorioHeroi1)
	if criticoHeroi1 {
		fmt.Printf("  Acerto Crítico: SIM (+15 de força)\n")
	}
	fmt.Printf("  Moral durante a luta: %+d (ajuste de força)\n", moralHeroi1)
	fmt.Printf("  Força final: %d\n", forcaFinalHeroi1)

	// Relatório do Herói 2
	fmt.Printf("%s:\n", heroi2.nome)
	fmt.Printf("  Força inicial: %d\n", heroi2.forca)
	fmt.Printf("  Popularidade: %d (impacto na força: +%d)\n", heroi2.popularidade, heroi2.popularidade/2)
	fmt.Printf("  Fator aleatório: +%d\n", fatorAleatorioHeroi2)
	if criticoHeroi2 {
		fmt.Printf("  Acerto Crítico: SIM (+15 de força)\n")
	}
	fmt.Printf("  Moral durante a luta: %+d (ajuste de força)\n", moralHeroi2)
	fmt.Printf("  Força final: %d\n", forcaFinalHeroi2)
	fmt.Println("\nImpacto da Batalha:")
}

func main() {
	// Criação dos heróis com diferentes popularidades
	heroi1 := Heroi{nome: "Homelander", forca: 100, popularidade: 80}
	heroi2 := Heroi{nome: "Homelander reverso", forca: 50, popularidade: 40}

	// Simulação da batalha
	simulador := SimuladorBatalha{}
	simulador.simularBatalha(heroi1, heroi2)
}
