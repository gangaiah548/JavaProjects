import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.stream.*;

//import com.sun.xml.internal.stream.Entity;

public class Test {
	public static void main(String[] args) {
		Base obj = new Base2();
		((Base) obj).hu();
		List<String> ls = new ArrayList<>();
	List<Double> ls2=new ArrayList<Double>();

		ls.add("abc");
		ls.add("bbc");
		ls.add("cbc");
		ls.add("fbc");
		ls.add("dbc");
		ls.add("ebc");
		Map<Character, List<String>> ms = ls.stream().sorted()
				.collect(Collectors.groupingBy(s -> s.charAt(0), LinkedHashMap::new, Collectors.toList()));

		for (Map.Entry<Character, List<String>> ss : ms.entrySet()) {
			System.out.println(ss);

		}
		List<String> tok = Arrays.asList("1", "2");
		String joi = tok.stream().collect(Collectors.joining(",", "[", "]"));
		System.out.println(joi);
		LinkedList<String> l1s = new LinkedList<String>();
		l1s.add("sf");
		// l1s.add(3,"sfd");

		System.out.println("str" == new String("str")); //false
		System.out.println("str" == "str");//true
		System.out.println("str".equals("str")); //true
		System.out.println("str".equals(new String("str")));
		System.out.println("str".substring(0).equals(new String("str")));
		System.out.println(new String("str")==(new String("str")));
		System.out.println("str".substring(0,3));
		System.out.println("str".length());
		

		Object x[] = new Integer[10];
		// x[0]=new String("j");
		System.out.println(x[0]);
	}

}