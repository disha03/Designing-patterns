# Designing-patterns


## Single Responsibilty principle: 
   It says your code only serve purpose that it is suppose to serve. For ex: lets say we are making a journal and in which we are only keeping the entries i.e. it is their sole purpose. Now if you also waant to save the journal in file, you can do this but that is not journal responsibility if you want to persist something if you are saving the jornal by using its object then it will break single respnsibilty principle because jornal is supoose to make entries only so, you can create another field persistance whose responsibilty to save any object i.e book, journal into the file then it makes more sense

## Open - Close principle:
   It says, our code is closed for modification but open for extension. For ex: a product has 3 properties like color, size, name and we are filtering the product based on the property but later the requirement is to filter by city so then need to add new property in struct city but according ton open close principle we are closed for modification so instead of this we need to implements the specification using interface so we can add new specification interface based on requirement

## Liskov Substitution Principle (LSP): 
   It says "functions that use pointers to base classes must be able to use objects of derived classes without knowing it." Thus, the objects of the subclasses should behave in the same way as the objects of the superclass. Simply put, the Liskov Substitution Principle (LSP) states that objects of a superclass should be replaceable with objects of its subclasses without breaking the application. 
