#include <iostream>
#include <string>
#include <cstdlib>
#include <ctime>

// Classe herói atualmente só tem nome e força
class Heroi {
public:
    std::string nome;
    int forca;

    Heroi(std::string n, int str) : nome(n), forca(str) {}
};

class SimuladorBatalha {
public:
    static void simularBatalha(Heroi& heroi1, Heroi& heroi2) {
        std::cout << "Iniciando batalha entre " << heroi1.nome << " e " << heroi2.nome << "!\n";

        // Força final de cada herói será a força inicial mais todos os outros fatores
        int heroi1ForcaFinal = heroi1.forca;
        int heroi2ForcaFinal = heroi2.forca;

        if (std::rand() % 10 < 2) heroi1ForcaFinal += 15;  // Se rand() % 10 for menor que 2 ( 20% de chance ) soma 15 à força dele
        if (std::rand() % 10 < 2) heroi2ForcaFinal += 15;  // ( Talvez no futuro cada herói ter sua própria chance de crítico e o quanto que o crítico aumenta )

        // Probabilidade de vitória será determinada pela razão entre a força do herói 1 e a força dos 2 somadas
        // Exemplo: Se 'heroi1ForcaFinal' for 85 e a 'heroi2ForcaFinal' for 80 a probabilidade será a seguinte:
        //          'forçaTotal' = 165, então a chance do heroi1 vencer será de 85 / 165 = 0,515 ou 51,5%
        // Esse método de calcular as vitórias dá uma chance, nem que seja pequena, de um heroi fraco vencer de um herói forte
        int forcaTotal = heroi1ForcaFinal + heroi2ForcaFinal;
        double heroi1Chance = static_cast<double>(heroi1ForcaFinal) / forcaTotal;
        
        // Para determinar quem será o vitorioso, uma comparação será feita
        // Essa função abaixo irá gerar um número aleatório, usando rand(), de 0 à RAND_MAX (32767)
        // Ao fazer a razão do número aleatório pelo máximo, sempre resulta em um número de 0 a 1
        // Então é só fazer a comparação entre esse valor com a chance do herói1 vencer
        // Exemplo: caso a chance do heroi1 seja 0,515 e o valor aleatório gerado for de 0,6, ele perderá
        //          mas se for 0,5 o heroi1 ganhará
        double random = static_cast<double>(std::rand()) / RAND_MAX;

        std::cout << heroi1.nome << " forca final: " << heroi1ForcaFinal << "\n";
        std::cout << heroi2.nome << " forca final: " << heroi2ForcaFinal << "\n";

        // Comparação
        if (random < heroi1Chance) {
            std::cout << heroi1.nome << " venceu a batalha!\n";
        } else {
            std::cout << heroi2.nome << " venceu a batalha!\n";
        }
    }
};

// Main genérico para testar a função
int main() {
    std::srand(std::time(0));

    Heroi heroi1("Homelander", 100);
    Heroi heroi2("Homelander reverso", 50);

    SimuladorBatalha::simularBatalha(heroi1, heroi2);

    return 0;
}