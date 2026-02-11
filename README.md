# 🐹 Learning Go (Golang)

> Minha jornada de migração e aprendizado: saindo do **Java** e explorando o poder do **Go**.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Java](https://img.shields.io/badge/java-%23ED8B00.svg?style=for-the-badge&logo=openjdk&logoColor=white)

## 📌 Sobre o Repositório
Sou um Desenvolvedor Backend com foco em Java/Spring e este repositório documenta meus estudos práticos na linguagem Go. O objetivo é comparar paradigmas, entender a sintaxe e dominar a concorrência do Go.

## 📂 Estrutura dos Estudos

Aqui estão os conceitos fundamentais que explorei até agora:

| Pasta | Tópico | Comparativo Java vs Go |
| :--- | :--- | :--- |
| **EX01** | Hello World & Variáveis | Declaração curta (`:=`) vs Tipagem explícita. |
| **EX02** | Controle de Fluxo | O `for` faz papel de `while` e `do-while`. |
| **EX03** | Structs & Ponteiros | Adeus `Class`! Entendendo `*` (Original) vs Valor (Cópia). |
| **EX04** | Composição (Embedding) | Substituindo `extends` (Herança) por Composição de Structs. |
| **EX05** | Interfaces & Polimorfismo | Interfaces implícitas (Duck Typing) vs `implements`. |
| **EX06** |  Defer & Panic | Gerenciamento de recursos e limpeza de pilha. |
| **EX07** | Error Handling | O padrão `if err != nil` substituindo o `try-catch`. |
| **EX08** | Arrays, Slices & Maps | Dinamismo de Slices vs a rigidez dos Arrays no Java. |

## 💡 Principais Diferenças que notei (Java ➡ Go)

Durante esses exercícios, destaquei as maiores mudanças de paradigma:

1. **Sem Classes, só Structs:** Go não tem classes. Dados são separados de comportamentos através de *receivers*.
2. **Herança não existe:** Substituímos o `extends` pela Composição (Embedding).
3. **Interfaces Implícitas:** Se um tipo satisfaz o contrato da interface, ele a implementa automaticamente.
4. **Tratamento de Erros:** Nada de exceções pesadas. Erros são valores de retorno comuns.
5. **Zero Value:** Diferente do Java (onde tudo pode ser `null`), no Go variáveis não inicializadas têm um valor padrão (0, "", false).

---
Developed by [Igor Rooberto](https://github.com/igorRooberto)
