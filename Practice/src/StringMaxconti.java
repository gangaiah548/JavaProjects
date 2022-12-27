import java.util.Iterator;

public class StringMaxconti {
	public static void main(String[] args) {
		String inp="aabccc";
		int m=1,t=0;
		String rs="",ts="";
		for(int i=0;i<inp.length();i++) {
			rs=rs+inp.charAt(i);
			//System.out.println(rs);
			if(inp.length()-1>i && inp.charAt(i)==inp.charAt(i+1)) {
				//System.out.print(inp.charAt(i));
				m++;
			}else {
				//System.out.println(rs);
				if(t<m) {
					ts="";
					ts=rs;
					t=m;
				}
				rs="";
				m=1;
			}
		}
		
		if(t<m) {
			ts="";
			ts=rs;
			t=m;
		}
		System.out.println(ts+t);
	}

}
