package collections;

import java.util.LinkedList;
import java.util.Queue;

public class BinarynbrNthPrint {
	public static void main(String[] args) {
		System.out.println((0&1)!=0);
		int i=7, t=0,c=0;
		Queue<String> q=new LinkedList<String>();
		q.add("1");
		while(c<3) {
			String cur=q.poll();
			
			if(!cur.contains("11")) {
				c++;
				System.out.print(cur+" ");
			}
			q.add(cur+"0");
			q.add(cur+"1");
			t++;
		}
		
		System.out.println(c);
		
	}

}
