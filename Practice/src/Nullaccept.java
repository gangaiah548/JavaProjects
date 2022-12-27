
public class Nullaccept {
	
	public static void Strcall(Object str)
{
	System.out.println("Object method"+str);
}
	
	public static void Strcall(String str)
	{
		System.out.println("string method"+str);
	}
	
	
	
	public static void main(String ar[]) {
		
		Strcall(null);
		//Objcall(null);
	}

}
