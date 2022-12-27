import java.time.LocalDate;
import java.time.Month;
import java.time.Period;
import java.util.Date;

public class SimpleDate {
	public static void main(String[] ar){
		//SimpleDate date="";
		LocalDate bdya=LocalDate.of(1955, Month.MAY, 19);
		LocalDate now=LocalDate.now();
		
		Period age=Period.between(bdya, now);
		int years = age.getYears();
        int months = age.getMonths();

        int days = age.getDays();
        System.out.println(""+days+","+years+","+months);
        
       // Date dt=new Date(1991, 08, 05);
        //Date dt2=new Date(1996, 08, 05);
        
	}

}
