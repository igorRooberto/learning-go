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

## 💡 Principais Diferenças que notei (Java ➡ Go)

Durante esses exercícios, destaquei as maiores mudanças de paradigma:

1.  **Sem Classes, só Structs:** Go não tem classes. Dados são separados de comportamentos.
2.  **Herança não existe:** Não existe `Funcionario extends Pessoa`. Existe uma `Pessoa` dentro de `Funcionario`.
3.  **Interfaces Implícitas:** Não preciso declarar que implemento uma interface. Se eu tenho o método, eu implemento.
4.  **Tratamento de Erros:** Nada de `try-catch`. O erro é retornado como valor.

---
Developed by [Igor Rooberto](https://github.com/igorRooberto)
