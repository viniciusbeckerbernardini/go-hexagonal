# Arquitetura Hexagonal (Ports and adapters)

## Resumo

Na Arquitetura Hexagonal a aplicação trabalha isolada e independente. Não tem alto acoplamento pois se guia por interfaces então pouco importa como ta fazendo do outro lado desde que tenha o resultado esperado, assim, conseguimos fazer unitários mais facilmente, coneguimos implementar adapters novos (como bancos de dados diferentes, cli, servidores web) sem ferir ou acoplar o core da aplicação pois isso sempre se mantem isolado do resto.

Com isso também manutenções futuras ficam mais facéis de desenvolver bem como novos adapters. O tempo inicial de desenvolvimento do core é maior também por todas as regras que ele precisa carregar e definir.

### Complexidades de um projeto

- Complexidade explicita
    - Complexidade de negócio (problemas reais que o software precisa resolver);
    -  Complexidade técnica (maneira que desenvolvemos o software).

- Misturamos as regras e isso acaba deteriorando a qualidade do software;
- Colocar limites claros para que a complexidade técnica não interfira no negócio;
- Bem implementado permite com que façamos a troca técnica (framework, banco de dados, etc);

## Ciclo de um projeto

### Arquitetura Hexagonal "Ports and Adapters"

#### Pontos importantes sobre arquitetura
- Crescimento sustentável;
- Software precisa se pagar ao longo do tempo;
- Limites que o software precisa ter;
- Software == Qualidade daquilo que está sendo produzido pra gerar resultado;
- Software deve ser desenhado por você e não pelo framework;
- Peças precisam se encaixar e eventualmente serem substituídas

```
Lembre-se! 
Arquitetura diz respeito a futuro do seu software. CRUD qualquer um faz!
```

#### Ciclo de vida de muitos projetos:

Desenhar e arquiteturar é diferente de escrever software:

#### Exemplo

- Fase 1:
    - Banco de dados;
    - Cadastros;
    - Validações;
    - Servidor web;
    - Controllers;
    - Views;
    - Autenticação;
    - Upload de arquivo;

- Fase 2:
    - Regras de negócio;
    - Criação de APIs;
    - Consumo de APIs;
    - Autorização;
    - Relatórios;
    - Logs;

- Fase 3:
    - Mais acessos;
    - Upgrades de hardware;
    - Cache;
    - API parceiros;
    - Regras de parceiros (ex: regras de gateway de pagamento);
    - Mais Relatórios

- Fase 4:
    - Mais acessos;
    - Upgrades de hardware;
    - DB relatórios;
    - Comandos (pra gerar relatórios, executar crons, exportar dados);
    - V2 da API;

- Fase 5:
    - Escala horizontal;
    - Sessões em escala;
    - Uploads em escala;
    - Refatoração;
    - Autoscaling;
    - CI/CD;

- Fase 6:
    - GraphQL;
    - Bugs constantes;
    - Problemas com Logs. Cada log grava em cada máquina. Precisa de um sistema de logs;
    - Integração CRM;
    - Migração para React;

- Fase 7:
    - Inconsistência CRM;
    - Containers;
    - CI/CD;
    - Memória (containers que morrem por funções que usam picos de memórias altos);
    - Logs;
    - Se livrar do legado;

- Fase 8:
    - Microserviços;
    - DB Compartilhado;
    - Problemas com tracing (não saber qual microserviço deu pau);
    - Lentidão em decorrência de dupla latência;
    - Custo elevado;

- Fase 9:
    - Kurbenetes;
    - CI/CD;
    - Mensageria;
    - Perda de mensagens;
    - Contratar Consultorias;

- Fase 10:
    - Use a imaginação porque não tem mais o que fazer, galinho, ta foda já...

### Principais problemas:
- Visão de futuro (inocencia no inicio);
- Limites bem definidos;
- Troca e adição de componentes;
- Escala;
- Otimizações frequentes;
- Preparado pra mudanças;


### Reflexões:
- Está sendo doloroso para o dev atualizar o software?
- Poderia ter sido evitado?
- Software está se pagando?
- Relação com cliente está boa?
- Cliente terá prejuízo com a brusca mudança arquitetural?
- Em qual momento tudo se perdeu?
    - No primeiro dia e no dia após o primeiro, é um dia de cada vez...
- Se você fosse novo na equipe, você julgaria os devs que fizeram tudo isso?
    - O dev se insere no contexto, se ele tem uma má gestão, uma má visão, e ainda sim ele segue pra entregar algo no outro lado

### Arquitetura vs Design
- Arquitetura está em nível acima, design está dentro de arquitetura;
- Uma decisão arquiteturar pode acarretar em uma decisão de design mas não o contrário;
- SOLID por exemplo é design de software;
- Libs, frameworks, DBs, tudo isso é design;
- Fluxos, regras, limites, padrões, responsabilidades, isso é arquitetura;

### Arquitetura Hexagonal "Ports and Adapters":
- Como você pode criar um software que de uma forma ou de outra consiga isolar os processos de acesso e os processos de conexão do software;
- "'Arquitetura' Hexagonal esta muito mais ligado com design que com arquitetura" - Wesley Williams;
- Bem delimitado, não mistura complexidade do negócio com a técnica;

- Definição e limites de proteção nas regras da aplicação;
    - Define a separação da complexidade de négócio e da complexidade técnica;
    - Componentização e desacoplamento;
        - Logs;
        - Cache;
        - Upload;
        - DBs;
        - Comandos;
        - Filas;
        - HTTP / APIs/ Graphql;
- Facilidade na quebra para microserviços;

- Lógica basica:

````
Cliente   =  (Aplicação)  = Servidor
(esquerdo)		centro		(direito)
- REST;					  - DB;
- CLI;					  - Redis;
- RPC;					  - Filesystem;
- GRAPHQL;				  - Lambda;
- UI; 					  - API;
````

- Possui uma camada de interface entre o cliente/servidor e a aplicação, essa interface é o "port";
    - O cliente pode acessar de diversas formas como por exemplo http ou cli mas sempre vai bater primeiro no adapter pra depois ir ou não pra aplicação;

- Dependency Inversion Principle (Principio da inversão de dependencia):
    - Módulos de alto nível não de módulos de baixo nível. Ambos devem depender de abstrações;
    - Abstrações não devem depender de detalhes. Detalhes devem depender de abstrações;

### Arquitetura Hexagonal vs Clean e Onion
- Observações:
    - Quanto mais desacoplado for o seu código, melhor;
    - Em arquitetura hexagonal não há um padrão estabelecido de como o código deve ser organizado;
    - Onion Style e Clean Architecture definem padrões de classes, pastas, organizações, etc;
    - Tem o mesmo principio mas não são iguais.



