import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.stream.Collectors;

public class NegativePositve {

	public static void main(String[] ar) {
		int[] arr= {1,10,-2,-3,-5,8};
		int[] narr=new int[arr.length];
		int c=0, k=arr.length-1;
		for(int i=0;i<arr.length;i++) {
			if(arr[i]<0) {
			narr[c]=arr[i];
			c++;
			}
			if(arr[i]>0) {
				narr[k]=arr[i];
				k--;
				}
		}
		
		/*
		 * for(int i=0;i<arr.length;i++) { if(arr[i]>0) { narr[c]=arr[i]; c++; } }
		 */
		
		for (int i : narr) {
			System.out.println(i);
			
		}
		
		Map<Integer,String> mp=new HashMap<Integer, String>();
		mp.put(1, null);
		
		List<Integer> ls =mp.keySet().stream().collect(Collectors.toList());
		
		for (Entry<Integer, String> mps : mp.entrySet()) {
			System.out.println(mps);
		}
		
		for (Integer integer : ls) {
			System.out.println(integer);
		}
		
		
	}
	
}
