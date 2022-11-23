public class Bird {
   public void fly() {
      System.out.println("I'm flying");
  }
  public void walk() {
    System.out.println("I'm walking");
 }
 public class Dove extends Bird{
 }

}
##Here, the Dove can fly because it is a Bird.

public class Penguin extends Bird{
}

##In this above inheritance, much as technically a penguin is a bird, penguins do not fly. 
##Thus, the "fly" method is not applicable to all types of birds; violating the LSV principle


###############################################################################################################cORRECT CODE

public class Bird {
   public void walk() {
      System.out.println("I'm walking");
   }

  public class FlyingBird extends Bird{
   public void fly() {
      System.out.println("I'm flying");
   }
  }

  public class Dove extends FlyingBird{}
  public class Penguin extends Bird{ }
}

##In this snippet, the Dove can both walk(from the Bird class) and fly (from FlyingBird class) and the Penguin can only walk (from the Bird class).
