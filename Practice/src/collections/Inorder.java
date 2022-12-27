package collections;

public class Inorder {
	public static void main(String[] args) {
		// 0,1 0 1 01 ==0011

		String s = "0101110";
		
		int[] cr= {1,0,0,1,1,0,1};
		int p=0;
		int t=cr.length-1;
		while(t>p) {
			if(cr[p]==1) {
				if(cr[t]!=1) {
					cr[t]=cr[t]+cr[p];
					cr[p]=cr[t]-cr[p];
					cr[t]=cr[t]-cr[p];
				}
				t--;
			}else {
				p++;
			}
		}
		
		for(int a:cr) {
			System.out.print(a);
		}
		String str = "Techie Delight";
		char ch = '_';
		int pos=6;
str = str.substring(0, pos) + ch + str.substring(pos + 1);
System.out.println(str);

	}
	

}
