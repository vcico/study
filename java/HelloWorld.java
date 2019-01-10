import com.dao.Student;

public class HelloWorld {
    public static void main(String[] args){
        System.out.println("xxxxxxx");
        for (String str:Student.GetAll())
        {
            System.out.println(str);
        }
        //System.out.println("Hello World!");
    }
}