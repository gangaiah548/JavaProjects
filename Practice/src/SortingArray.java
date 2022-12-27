
import java.util.*;
public class SortingArray {
	
	public static void main(String[] args) {
		int[] A= { 6, 2, 1,5,8,9};
		System.out.println(solution(A));
	}


 public static int solution(int[] A) {
     // Implement your solution here [1, 3, 6, 4, 1, 2]
     int rt=0,t=0;
     for(int i=0;i<A.length;i++){
         if(A[i]>0){
             t=1;
             break;
         }
     }
     if(t==0){
         return 1;
     }else{
             int l=0,h=0;
             l=A[0];
             h=A[0];
             
             LinkedHashSet<Integer> s=new LinkedHashSet<>();
             
     for(int i=0;i<A.length;i++){
         s.add(A[i]);
        
         if(l>A[i])
         l=A[i];

         if(h<A[i])
         h=A[i];
     }
     List<Integer> ls=s.stream().sorted((o1,o2)->o1.compareTo(o2)).toList();
     
     Integer[] el=new Integer[ls.size()];
     int mt=0;
     for(Integer lss:ls) {
    	 if((lss-l)!=0) {
    		 break;
    	 }
    	 l++;
    	 
     }
rt=l;
     el=ls.toArray(el);
     int m=0;
		/*
		 * for(int j=0;j<el.length-1;j++){ if((el[j+1]-el[j])==1){ continue; }else{
		 * rt=el[j]; ++rt; break; } } if(rt==0){ rt=el[el.length-1]; ++rt; }
		 */

     
     }


     return rt;
 }
}

