
public class StringRightrotation {
	public static void main(String ar[]) {
		String str1 = "abcdabc";
		int c = 0;
		String ans = rightRatate(str1, str1.length() - 1);
		System.out.println(c+" "+ans);
		boolean ro = true;
		while (ro) {
			c++;
			if (!ans.equals(str1)) {
				ans=rightRatate(ans, ans.length() - 2);
				System.out.println(c+" "+ans);
			} else {
				ro = false;
				break;
			}
			if (!ans.equals(str1)) {
					ans=rightRatate(ans, ans.length() - 1);
					System.out.println(c+" "+ans);
			} else {
				ro = false;
				break;
			}

		}
		System.out.println(c+""+ans);
	}

	public static String rightRatate(String str, int d) {

		String rs = str.substring(d) + str.substring(0, d);
		return rs;
	}
}
