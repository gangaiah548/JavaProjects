
public class Escapeminimum {
	public static void main(String[] args) {
		int i = 0, x = 2;
		int[] ar = { 1, 2, 3, 4, 5, 6,7,8 };
		int s = 0, ts = 1, c = 0;
		s += ar[0] + ar[ar.length - 1];
		System.out.println(s);
		int j = 1;
		for (j = 1; j < ar.length - 1; j++) {
			c++;
			if (c <=4) {
				if (c == 1) {
					if (ts > 1) {
						s += ar[ts];
						s = s - x;
						ts++;
					} else {
						s += ar[ts];
						ts++;
					}
					
				} else {
					s += x;
				}
				if (c == 4) {
					c = 0;
				}
			}
			

		}

		System.out.println(""+j+"" + s);
	}

}
