import java.util.ArrayList;
import java.util.List;

public class Palindrome {
	public static void main(String[] ar)
	{
		String str1="aba"; //60000 --odd, remaining half of the string
		
		//List<String> ls=new ArrayList<String>();
		
		String longest=str1.substring(0,1);
		
		String str2="";
		for(int i=0; i<str1.length()-1;i++)
		{
			//odd
			String st3=splitpalin(str1,i,i);
			
			String st4=splitpalin(str1,i,i+1); //o(n2)
			//evenif

		}
		
		//if(st)
		
		System.out.println(str2);
	}

	private static String splitpalin(String str1, int i, int j) {
		// TODO Auto-generated method stub
		if(i>j) return null;
		while(i>=0 && j<str1.length() && str1.charAt(i)==str1.charAt(j))
		{
			i--;j++;
		}
		return str1.substring(i+1,j);
	}
	
	

	
}
