# 🐹 Learning Go (Golang)

> **Minha jornada de migração e aprendizado: saindo do Java e explorando o poder do Go.**

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Java](https://img.shields.io/badge/java-%23ED8B00.svg?style=for-the-badge&logo=openjdk&logoColor=white)](https://www.java.com/)

## 📌 Sobre o Repositório

Este repositório documenta meus estudos práticos na linguagem Go (Golang). Como desenvolvedor Backend com background em **Java/Spring**, meu foco aqui é comparar paradigmas e entender a sintaxe e as particularidades do Go.

## 📂 Estrutura dos Estudos

O repositório está organizado em diretórios numerados contendo exercícios práticos:

| Diretório | Tópico | O que foi praticado |
| :--- | :--- | :--- |
| **[EX01](./EX01)** | **Input & Output** | Uso de `fmt.Scan` para entrada de dados e `fmt.Printf` para formatação de strings. |
| **[EX02](./EX02)** | **Funções & Erros** | Criando funções com múltiplos retornos (valor, erro) e validações básicas. |
| **[EX03](./EX03)** | **Lógica & Strings** | Algoritmo para verificar palíndromos, manipulando índices e strings. |
| **[EX04](./EX04)** | **Laços (Loops)** | Uso do `for` clássico para criar uma tabuada dinâmica. |
| **[EX05](./EX05)** | **Interfaces & Structs** | Polimorfismo prático com interface `Geometria` e struct `Quadrado` (Duck Typing). |
| **[EX06](./EX06)** | **Panic & Recover** | Gerenciamento de falhas graves usando `defer`, `panic` e `recover` (simulando try/catch). |
| **[EX07](./EX07)** | **Switch & Condicionais** | Verificação de números pares/ímpares utilizando `switch case` dentro de um loop. |
| **[EX08](./EX08)** | **Conversão & Tipos** | Conversor de Celsius para Fahrenheit explorando tipos `float64` e operações matemáticas. |

## 💡 Principais Diferenças que notei (Java ➡ Go)

Durante esses exercícios, destaquei as maiores mudanças de paradigma:

1. **Sem Classes, só Structs:** Go não tem classes. Dados são separados de comportamentos através de *receivers*.
2. **Herança não existe:** Substituímos o `extends` pela Composição (Embedding).
3. **Interfaces Implícitas:** Se um tipo satisfaz o contrato da interface, ele a implementa automaticamente.
4. **Tratamento de Erros:** Nada de exceções pesadas. Erros são valores de retorno comuns.
5. **Zero Value:** Diferente do Java (onde tudo pode ser `null`), no Go variáveis não inicializadas têm um valor padrão (0, "", false).

---
Developed by [Igor Rooberto](https://github.com/igorRooberto)
