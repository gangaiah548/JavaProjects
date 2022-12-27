
public class Stringcompress {
	public static void main(String[] args) {
		String s = "abbbaa";
		int r = 1;
		String rs = "";
		char ts = 0;
		for (int i = 0; i < s.length() - 1; i++) {
			ts = s.charAt(i);
			if (s.charAt(i) == s.charAt(i + 1)) {
				r += 1;
			} else {
				if (r > 1) {
					rs += ts;
					rs += r;
				}
				rs += ts;
				r = 1;
			}
		}
		
		if(r>1) {
			rs+=ts;
			rs+=r;
		}
		
		System.out.println(rs);
	}

}
