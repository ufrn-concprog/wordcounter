# Trabalho Prático: Contador de Palavras

<sub>Última atualização: 30/01/2025</sub>

## Sumário

- [Objetivo](#objetivo)
- [Descrição](#descrição)
- [Tarefas](#tarefas)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)

## Objetivo

O objetivo deste trabalho é colocar em prática o projeto e a implementação de programas concorrentes utilizando a linguagem de programação [Go](https://go.dev).

## Descrição

Uma pessoa programadora necessita analisar múltiplos arquivos de texto, de tamanho considerável, para contar a frequência de ocorrência de determinadas palavras. No entanto, para evitar a necessidade de abrir cada arquivo, um a um, para realizar essa contagem, a pessoa programadora decidiu implementar um programa que faça essa análise de forma automatizada. Por outro lado, analisar arquivos de forma sequencial, de forma similar ao que seria feito manualmente, é uma tarefa lenta, de modo que a solução ideal seria implementar um programa concorrente para isso.

## Tarefas

A tarefa central a ser realizada neste trabalho consiste em implementar um programa concorrente na linguagem de programação Go que realize a contagem da ocorrência de um conjunto de palavras em múltiplos arquivos de texto localizados em um diretório. O programa deve esses múltiplos arquivos e, utilizando **gorotinas e canais**, processar cada um deles de forma concorrente para contar a frequência de cada uma das palavras em todos esses arquivos. Ao final, o programa deve agregar as múltiplas contagens feitas e apresentar na saída padrão a contagem para cada arquivo e a soma total para todos os arquivos. Um exemplo de saída para o programa seria:

```bash
Palavra "
Aeronave #1: Solicitando permissão para pouso | Aeronaves aguardando pouso: 1
Aeronave #2: Solicitando permissão para decolagem | Aeronaves aguardando decolagem: 1
Aeronave #3: Solicitando permissão para pouso | Aeronaves aguardando pouso: 2
Aeronave #4: Solicitando permissão para decolagem | Aeronaves aguardando decolagem: 2
Aeronave #5: Solicitando permissão para pouso | Aeronaves aguardando pouso: 3
Aeronave #6: Solicitando permissão para decolagem | Aeronaves aguardando decolagem: 3

Aeronave #1: Pouso em andamento...
Aeronave #1: Pouso finalizado.
Aeronave #3: Pouso em andamento...
Aeronave #3: Pouso finalizado.
Aeronave #5: Pouso em andamento...
Aeronave #5: Pouso finalizado.
Aeronave #2: Decolagem em andamento...
Aeronave #2: Decolagem finalizada.
Aeronave #4: Decolagem em andamento...
Aeronave #4: Decolagem finalizada.
Aeronave #6: Decolagem em andamento...
Aeronave #6: Decolagem finalizada.

Todas as aeronaves completaram suas operações.
```

uma nova versão da emulação de um sistema de gestão da pista do aeroporto em questão feita anteriormente, desta vez na linguagem de programação Go. As múltiplas aeronaves devem ser implementadas como gorotinas realizando operações de solicitação de pouso ou de decolagem. Conforme descrito para o problema, as requisições de pouso têm preferência sobre as requisições de decolagem. A sincronização das gorotinas para a utilização da pista pode ser feito uso de canais de comunicação (com ou sem *buffer*) e/ou de mecanismos de sincronização avançada providos pelo pacote [`sync`](https://pkg.go.dev/sync) da linguagem Go.

Além da implementação do núcleo da solução, deverá ser implementado, para fins de demonstração, um código fonte principal que realize a instanciação de um determinado número de aeronaves realizando as operações de decolagem e pouso sobre a pista do aeroporto. Esse número pode tanto ser fornecido pela pessoa usuária quando da execução do programa, via entrada padrão, quanto ser fixado no próprio código fonte.

A execução do programa deverá exibir, na saída padrão, o comportamento das aeronaves. Para tornar a execução menos simplista, sugere-se colocar cada gorotina em um estado de espera cronometrada (*sleep*) para simular a utilização efetiva da pista do aeroporto pela aeronave. O programa só deve ser encerrado quando **todas** as aeronaves tiverem decolado ou pousado. Um exemplo de saída para um programa com seis aeronaves, três delas solicitando pouso e três solicitando decolagem, seria:

```bash
Aeronave #1: Solicitando permissão para pouso | Aeronaves aguardando pouso: 1
Aeronave #2: Solicitando permissão para decolagem | Aeronaves aguardando decolagem: 1
Aeronave #3: Solicitando permissão para pouso | Aeronaves aguardando pouso: 2
Aeronave #4: Solicitando permissão para decolagem | Aeronaves aguardando decolagem: 2
Aeronave #5: Solicitando permissão para pouso | Aeronaves aguardando pouso: 3
Aeronave #6: Solicitando permissão para decolagem | Aeronaves aguardando decolagem: 3

Aeronave #1: Pouso em andamento...
Aeronave #1: Pouso finalizado.
Aeronave #3: Pouso em andamento...
Aeronave #3: Pouso finalizado.
Aeronave #5: Pouso em andamento...
Aeronave #5: Pouso finalizado.
Aeronave #2: Decolagem em andamento...
Aeronave #2: Decolagem finalizada.
Aeronave #4: Decolagem em andamento...
Aeronave #4: Decolagem finalizada.
Aeronave #6: Decolagem em andamento...
Aeronave #6: Decolagem finalizada.

Todas as aeronaves completaram suas operações.
```

### Relato

Uma vez realizadas a tarefa de implementação, deverá ser elaborado um documento simples contendo (i) uma descrição da evolução da implementação feita anteriormente para a versão atual na linguagem de programação Go e (ii) uma análise comparativa no que se refere a esforço de programação entre a solução atual e a anteriormente implementada, ressaltando inclusive pontos positivos e negativos. O documento deverá estar preferenciamente em formato PDF.

## Autoria e política de colaboração

Este trabalho deverá necessariamente ser realizado em equipe composta de **até dois estudantes**, sendo importante, dentro do possível, dividir as tarefas igualmente entre as pessoas integrantes da equipe. O arquivo [`author.md`](https://github.com/ufrn-concprog/arms-golang/tree/master/author.md) presente no repositório deverá ser editado preenchendo as informações de identificação das pessoas integrantes da equipe, na seção [Informações de Autoria](https://github.com/ufrn-concprog/arms-golang/tree/master/author.md#identificação-de-autoria). Por sua vez, instruções sobre como compilar e executar o programa devem ser adicionadas à seção [Instruções](https://github.com/ufrn-concprog/arms-golang/tree/master/author.md#instrucoes) do mesmo arquivo [`author.md`](https://github.com/ufrn-concprog/arms-golang/tree/master/author.md) do repositório.

O trabalho em cooperação entre estudantes da turma é estimulado, sendo admissível a discussão de ideias e estratégias. Contudo, tal interação não deve ser entendida como permissão para utilização de (parte de) código fonte de colegas, o que pode caracterizar situação de plágio. **Trabalhos copiados no todo ou em parte de outros colegas ou da Internet ou ainda gerados por ferramentas de Inteligência Artificial serão anulados e receberão nota zero.**

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://github.com) serão utilizados para possibilitar a entrega da implementação realizada. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (*assignment*) no GitHub Classroom. Cada integrante de equipe deverá acessar este [*link*](https://classroom.github.com/a/IY8_AGpL), aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelas pessoas integrantes da equipe e pelo docente, sob a organização [`ufrn-concprog`](https://github.com/ufrn-concprog). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo dos arquivos executáveis gerado a partir da compilação do código fonte.

A entrega deste trabalho deverá ser realizada até as **23:59 do dia 26 de janeiro de 2025** no respectivo repositório Git da equipe. Além disso, o documento produzido deverá ser submetido através da opção *Tarefas* da Turma Virtual do SIGAA e o endereço do repositório da equipe deverá ser informado no campo *Comentários*. **Apenas uma pessoa integrante da equipe deve realizar esse envio** e não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.

## Avaliação

A avaliação do trabalho contabilizará nota de até 10,0 pontos na 3ª Unidade do módulo. O trabalho será avaliado de acordo com os seguintes critérios:

- utilização correta de gorotinas e dos mecanismos de sincronização providos pela linguagem de programação Go;
- a corretude da execução do programa implementado, que deve apresentar saída em conformidade com a especificação, e com relação a concorrência;
- a aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte, e;
- correta utilização do repositório Git, no qual deverá ser registrado todo o histórico da implementação por meio de *commits*.

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível:

- enviar *e-mail* para o endereço <everton.cavalcante@ufrn.br>;
- enviar mensagem privada diretamente ao docente, utilizando o servidor Discord;
- enviar mensagem no canal de texto `#duvidas` no servidor Discord, ou;
- agendar encontros síncronos pelo canal de áudio `Fale com Prof. Everton` no servidor Discord.
