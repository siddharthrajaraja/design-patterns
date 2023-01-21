# Factory pattern
* Probably second-best known and used design pattern.
* It's purpose is to abstract the user from knowledge of the struct he needs to achieve for a specific purpose.
* Used to gain extra layer of encapsulation so that our program can grow in controlled environment.

## Objectives
* Delegating the creation of new instances of structures to a different part of program
* Working at the interface level instead of with concrete implementations
* Grouping families of objects to obtain family object creator
* Finally we have seen that tests must be written with cae if you don't want to tie yourself to certain implementations that do not have anything to do with the tests directly.

## Example :
We will take example of Payment methods where we want to pay either by using Cash or Debit Card.
We can further enhance our implementations with Credit Card and can replace implementation of Debit Card with Credit Card without deleting DebitCard struct.
Or we can simply add a new payment method whose object can be created in switch case inside `GetPaymentMehtod` func.
## Acceptance Criteria
* To have common method for every payment method we create a Pay method
* To be able to delegate the creation of payments methods we have need to make Factory
* To be able to add more payments methods to the library by just adding it to the factory method