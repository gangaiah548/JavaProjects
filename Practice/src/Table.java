
public class Table {
	public void print(int n) {
		for(int i=1;i<=10;i++)
		System.out.println(n*i);
	}
}
class thread1 extends Thread
{
	Table t;
	thread1(Table t){
		this.t=t;
	}
	
	public void run() {
		t.print(5);
	}
}
class thread2 extends Thread
{
	Table t;
	thread2(Table t){
		this.t=t;
	}
	
	public void run() {
		t.print(7);
	}
}
class Main{
	public static void name() {
		Table tob=new Table();
		thread1 t1=new thread1(tob);
		thread2 t2=new thread2(tob);
		t1.start();t2.start();
	}
}