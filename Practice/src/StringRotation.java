
public class StringRotation {
	static String leftrotate(String str, int d)
    {
            String ans = str.substring(d) + str.substring(0, d);
            return ans;
    }
	static String rightrotate(String str, int d)
    {
            
            return leftrotate(str, str.length()-d);
    }
	
	public static void main(String ar[])
	{
		String str1 = "abcdabc";
		
		System.out.println(str1.substring(5)+"dsfs"+str1.substring(0,5));
		
		int l=1,r=1,c=1;
		String ans=leftrotate(str1, 2);
		boolean ro=true;
		while(ro)
		{
			c++;
			System.out.println("before ifs"+ans);
			if(!ans.equals(str1)) {
				ans=rightrotate(ans, r);
				System.out.println("insde 1if"+ans);
			}
				else
				{
					ro=false;
					break;
				}
			if(!ans.equals(str1)) {
			ans=leftrotate(ans, 2);
			}
			else
			{
				ro=false;
				break;
			}
			
		}
		
		System.out.println(c);
		
	}

}
