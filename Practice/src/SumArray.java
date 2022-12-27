import java.util.Arrays;
import java.util.stream.Stream;
import java.util.stream.Collectors;

public class SumArray {
public static void main(String[] args) {
	int[] ar={1,2,3,4};
	int sum=Arrays.stream(ar).sum();
	int sum1=Arrays.stream(ar).reduce((x,y)->x+y).getAsInt();
	Arrays.stream(ar).forEach(k->System.out.print("-"+k));
	System.out.println(sum+""+sum1);
	
	//Arrays.stream(ar).map( String::valueOf ).collect(Collectors.joining( " " ) );
	
	System.out.println(Stream.of(ar).map( String::valueOf ).collect(Collectors.joining( " " ) ));
}

}
