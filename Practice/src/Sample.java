import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.function.BiConsumer;
import java.util.function.BiPredicate;
import java.util.function.Function;
import java.util.function.Supplier;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import java.util.stream.Collector;
import java.util.Optional;
import java.util.Set;

class Employee {
	String name;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	int id;

	Employee(String name) {
		this.name = name;
	}

}

class EmploySort implements Comparator<Employee> {
	@Override
	public int compare(Employee o1, Employee o2) {
		// TODO Auto-generated method stub
		return o2.getName().compareTo(o1.getName());
	}
}

class Hosting {

	private int Id;

	public int getId() {
		return Id;
	}

	public void setId(int id) {
		Id = id;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getWebsites() {
		return websites;
	}

	public void setWebsites(int websites) {
		this.websites = websites;
	}

	private String name;
	private int websites;

	public Hosting(int id, String name, int websites) {
		Id = id;
		this.name = name;
		this.websites = websites;
	}
}

public class Sample {
	public static void main(String ar[]) {
		Integer a=7,b=8;
		
		  Function<Integer, Integer> obj=a1->a+a; 
		  
		  System.out.println(obj.apply(a));
		 
		Employee e1=new Employee("abc");
		Employee e2=new Employee("acd");
		Employee e3=new Employee("afd");
		ArrayList<Employee> le=new ArrayList<Employee>();
		le.add(e2);
		le.add(e1);
		le.add(e3);
		
		List<Employee>ls=le.stream().sorted((o1,o2)->o2.getName().compareTo(o1.getName())).collect(Collectors.toList());
		Collections.sort(le,new EmploySort());
		//List<char> l2= new ArrayList<char>();
		char[] car= {'a','b','c'};
		List l2l2=Arrays.asList(car);
		System.out.println(l2l2);
		
		
		
		
		BiConsumer<String, Integer> buc=(x,y)->System.out.print(x+y);
		
		buc.accept("consumerex", b);
		
		BiPredicate<String, Integer> bu=(x,y)->x.length()==y;
		
		System.out.println(bu.test("abc", 3));
		
		//Supplier<String> s="";
		
		for (Employee employee : le) {
			System.out.println(employee.getName());
			
		}
		
		
		List<Hosting> list = new ArrayList<>();
        list.add(new Hosting(1, "liquidweb.com", 80000));
        list.add(new Hosting(2, "linode.com", 90000));
        list.add(new Hosting(3, "digitalocean.com", 120000));
        list.add(new Hosting(4, "aws.amazon.com", 200000));
        list.add(new Hosting(5, "mkyong.com", 1));
		
        list.add(new Hosting(6, "linode.com", 100000));
        
        Map<String,Integer> result1 = list.stream()
                .sorted(Comparator.comparingLong(Hosting::getWebsites).reversed())
                .collect(
                        Collectors.toMap(
                                Hosting::getName, Hosting::getWebsites, // key = name, value = websites
                                (oldValue, newValue) -> oldValue,LinkedHashMap<String,Integer>::new                      
                        ));

        System.out.println("Result 1 : " + result1);
        Map<String, Object> cl=list.stream().collect(Collectors.
        		groupingBy(Hosting::getName,Collectors.collectingAndThen
        				(Collectors.minBy((o1,o2)->o2.getWebsites()-o1.getWebsites()),Optional::get)));
        System.out.println("map with min"+cl);
        
        Set<String> st=list.stream().map(sd->sd.getName()).collect(Collectors.toSet());
        
        System.out.println(st);
        
        List<String> lsn=list.stream()
        		.filter(hs->hs.getWebsites()>100000)
        		.filter(hs->hs.getWebsites()<=200000)
        		.map(Hosting::getName).collect(Collectors.toList());
        
        System.out.println("rang"+lsn);
        
		int[] ab={1,2,3,4};
		int[] ab1={1,2,3,4};
		//int[] mr=Stream.of(ab,ab1).flatMap(ints1->Arrays.stream(ints1)).toArray();
	}

}
