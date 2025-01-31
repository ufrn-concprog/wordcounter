# Trabalho Prático: Contador de Palavras

<sub>Última atualização: 30/01/2025</sub>

## Sumário

- [Objetivo](#objetivo)
- [Descrição](#descrição)
- [Tarefas](#tarefas)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)
- [Dúvidas e informações](#dúvidas-e-informações)

## Objetivo

O objetivo deste trabalho é colocar em prática o projeto e a implementação de programas concorrentes utilizando a linguagem de programação [Go](https://go.dev).

## Descrição

Uma pessoa programadora necessita analisar múltiplos arquivos de texto, de tamanho considerável, para contar a frequência de ocorrência de determinadas palavras. No entanto, para evitar a necessidade de abrir cada arquivo, um a um, para realizar essa contagem para cada uma das palavras, a pessoa programadora decidiu implementar um programa que faça essa análise de forma automatizada. Por outro lado, analisar arquivos de forma sequencial, de forma similar ao que seria feito manualmente, é uma tarefa laboriosa e lenta, de modo que a solução ideal seria implementar um programa que realize essa tarefa de forma concorrente, já que os arquivos são independentes.

## Tarefas

A tarefa central a ser realizada neste trabalho consiste em implementar um programa concorrente na linguagem de programação Go que realize a contagem da ocorrência de um conjunto de palavras em múltiplos arquivos de texto localizados em um diretório. O programa deve esses múltiplos arquivos e, utilizando **gorotinas e canais**, processar cada um deles de forma concorrente para contar a frequência de cada uma das palavras em todos esses arquivos. Ao final, o programa deve agregar as múltiplas contagens feitas e apresentar na saída padrão a contagem para cada arquivo e a soma total para todos os arquivos.

O programa deve receber via linha de comando o caminho do diretório no qual se encontra os arquivos a serem processados e, da entrada padrão, as palavras que se quer contabilizar de forma *case sensitive*. Para este trabalho, devem ser considerados os cinco arquivos de texto disponíveis no diretório [`files`](https://github.com/ufrn-concprog/wordcounter/tree/master/files) deste repositório. O conteúdo desses arquivos consiste em *dummy text* na língua inglesa gerado automaticamente através do *site* [Blind Text Generator](https://www.blindtextgenerator.com/). Um exemplo de execução do programa seria:

```bash
$ wordcounter files
Informe as palavras que deseja contabilizar, separadas por espaço:
me and back like one
```

Para essas palavras, a saída seria:

```bash
Palavra: "me"
- eulanguages.txt: 69
- farfaraway.txt: 0
- kafka.txt: 0
- pangram.txt: 0
- werther.txt: 0
- Total: 69

Palavra: "and"
- eulanguages.txt: 280
- farfaraway.txt: 377
- kafka.txt: 223
- pangram.txt: 150
- werther.txt: 406
- Total: 1436

Palavra: "back"
- eulanguages.txt: 0
- farfaraway.txt: 32
- kafka.txt: 50
- pangram.txt: 0
- werther.txt: 0
- Total: 82

Palavra: "like"
- eulanguages.txt: 70
- farfaraway.txt: 0
- kafka.txt: 29
- pangram.txt: 0
- werther.txt: 95
- Total: 194

Palavra: "one"
- eulanguages.txt: 70
- farfaraway.txt: 0
- kafka.txt: 10
- pangram.txt: 0
- werther.txt: 0
- Total: 80
```

A implementação do programa deve ainda utilizar como ponto de partida o arquivo de código fonte [`wordcounter.go`](https://github.com/ufrn-concprog/wordcounter/tree/master/wordcounter.go) disponibilizado neste repositório. Esse arquivo contém apenas a implementação da função `processFile`, que recebe como parâmetro o caminho do arquivo de texto a ser processado, além do trecho da função principal (`main`) que implementa a entrada do programa. A contagem de palavras deve ser armazenada em uma estrutura de dados do tipo *map* em que as chaves referem-se às palavras e os valores referem-se à frequência delas no arquivo em questão, ou seja, `map[string]int`. **Faz parte da tarefa deste trabalho modificar essa função e/ou criar outras, se necessário, para operar de forma concorrente mediante o uso de gorotinas e canais.**

## Autoria e política de colaboração

Este trabalho deverá ser realizado **individualmente**. O arquivo [`author.md`](https://github.com/ufrn-concprog/wordcounter/tree/master/author.md) presente no repositório deverá ser editado preenchendo o nome da pessoa autora do programa, na seção [Informações de Autoria](https://github.com/ufrn-concprog/wordcounter/tree/master/author.md#identificação-de-autoria). **Trabalhos copiados no todo ou em parte de outros colegas ou da Internet ou ainda gerados por ferramentas de Inteligência Artificial serão anulados e receberão nota zero.**

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://github.com) serão utilizados para possibilitar a entrega da implementação realizada. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (*assignment*) no GitHub Classroom. Cada estudante deverá acessar este [*link*](https://classroom.github.com/a/X6rl7L5H), aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelas pessoas integrantes da equipe e pelo docente, sob a organização [`ufrn-concprog`](https://github.com/ufrn-concprog). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo dos arquivos executáveis gerado a partir da compilação do código fonte.

A entrega deste trabalho deverá ser realizada até as **23:59 do dia 31 de janeiro de 2025** no respectivo repositório Git. Não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.

## Avaliação

A avaliação do trabalho contabilizará nota de até 10,0 pontos. O trabalho será avaliado de acordo com os seguintes critérios:

- utilização correta de gorotinas, canais e dos mecanismos de sincronização providos pela linguagem de programação Go;
- a corretude da execução do programa implementado, que deve apresentar saída em conformidade com a especificação, e com relação a concorrência, e;
- a aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte.

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível enviar mensagem privada diretamente ao docente, utilizando o servidor Discord criado para a turma.
