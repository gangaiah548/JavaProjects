import java.util.Stack;

public class BalanceString {
	//[()]{​​}​​{​​[()()]()}​​
	public static void main(String[] args) {
		//stack--push,pop lifo
		
	String inp="[()]{​​}​​{​​[()()]()}";
	char[] car=inp.toCharArray();
	Stack<Character> st=new Stack<Character>();
	for (char character : car) {
		char ct=character;
		if(character=='['||character=='{'||character=='(')
		{
			
		st.push(character);
		}else {
		/*
		 * if(st.isEmpty()) { System.out.println("notbalanced878"); break; }
		 */
		
			if(!st.isEmpty()&&st.peek()==ct) {
				st.pop();
			}else if(!st.isEmpty()&&st.peek()!=character) {
				System.out.println("notbalanced");
				break;
			}
		}
		
	}
	if(st.isEmpty())
	{
		System.out.println("balanced");
	}else {
		System.out.println("kjk");
	}
		
	}

}
