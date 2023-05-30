 ## Desafio E-mail
  
 ### Crie um cliente de email em Go que receberá os seguintes provedores:

- SendGrid
- AWS SES    
- Provider "Null" que retorna sempre verdadeiro simulando um envio    
  
### *O cliente aceita os providers em qualquer ordem*
### *O cliente utiliza cada provider em ordem até que o e-mail seja disparado por um deles.*
### *Tenha em mente que no futuro novos providers podem ser adicionados ao cliente.*

## Links Importantes

-  [Sendgrind](https://github.com/sendgrid/sendgrid-go)
-  [AWS SES](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-ses-with-go-sdk.html)

## Design Patters

- [Factory](https://refactoring.guru/pt-br/design-patterns/factory-method)
- [Strategy](https://refactoring.guru/pt-br/design-patterns/strategy)


## Ideia do desafio 

- [DevGyn](https://app.devgym.com.br/challenges/8d84979a-e8b2-4f87-a6c8-0e0d2f380a2a)