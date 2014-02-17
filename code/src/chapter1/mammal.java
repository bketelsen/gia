public class Mammal implements Animal{

   public void eat(){
      System.out.println("eating food");
   }

   public void move(){
      System.out.println("moving");
   } 

  
   public static void main(String args[]){
      Mammal m = new Mammal();
      m.eat();
      m.move();
   }
} 