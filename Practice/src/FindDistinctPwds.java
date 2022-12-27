import java.util.HashSet;
import java.util.Iterator;
import java.util.Set;

public class FindDistinctPwds {
	int  MaxChar=26;
	public static void main(String[] args) {
		String[] star= {"abcd", "acbd", "adcb", "cdba","bcda", "badc"};
		Set<String> s=new HashSet<>();
		int dt=0;
		for (int i = 0; i < star.length; i++) {
			if(!s.contains(encode(star[i].toCharArray()))) {
				s.add(encode(star[i].toCharArray()));
				dt++;
			}
			
			
			
		}
		System.out.println(s+"jh"+dt);
		
		//System.out.println(""+(1&1));
		
		
	}
	public static String encode(char[] str) {
		int hashEven[] =new int[26];
		int hashOdd[] =new int[26];
		for(int i=0;i<str.length;i++) {
			char ct=str[i];
			if((i&i)!=0) {
				hashOdd[ct-'a']++;
			}else {
				hashEven[ct-'a']++;
			}
		}
		
		String enc="";
		for(int i=0;i<26;i++) {
			enc+=hashEven[i];
			enc+=('-');
			enc+=hashOdd[i];
			enc+=('-');
		}
		
		return enc;
	}
	

}

/* public class DistinctPasswd
{

private static final int LOWER_ALPHABET_COUNT = 26;


private static String customHash(String str) {

int[] hashEven = new int[LOWER_ALPHABET_COUNT];
int[] hashOdd = new int[LOWER_ALPHABET_COUNT];

for(int i = 0; i < str.length(); i++) {
char c = str.charAt(i);
if((i+1)%2 != 0) {
hashOdd[c-'a']++;
}
else {
hashEven[c-'a']++;
}
}

StringBuilder hashOfInput = new StringBuilder();
for(int i = 0; i < LOWER_ALPHABET_COUNT; i++) {
hashOfInput.append(hashEven[i]).append("#").append(hashOdd[i]).append("|");
}
return hashOfInput.toString();
}

static int countDistinct(String[] input, int n) {
int countOfDistinct = 0;
Set<String> uniqueHashes = new HashSet<>();
for(int i = 0; i < n; i++) {
String hash = customHash(input[i]);
System.out.println("String: " +input[i] + " hash: " +hash) ;
if(!uniqueHashes.contains(hash)) {
uniqueHashes.add(hash);
countOfDistinct++;
}
}
return countOfDistinct;
}

public static void main( String[] args ) throws IOException
{
BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
int n = Integer.parseInt(br.readLine());
String[] input = new String[n];
for(int i = 0; i < n; i++) {
input[i] = br.readLine();
}
System.out.println(countDistinct(input, input.length));

}
} */
