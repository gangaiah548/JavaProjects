
public class Threads implements Runnable {
	public static void main(String[] args) throws InterruptedException {
		Threads ts=new Threads();
		Thread t1= new Thread(ts);
		Thread t2= new Thread(ts);
		Thread t3= new Thread(ts);
		
		t1.start();
		t1.sleep(1200);
		t2.start();
		t3.start();
		Thread.sleep(200);
		
	}

	@Override
	public void run() {
		// TODO Auto-generated method stub
		String sname=Thread.currentThread().getName();
		try {
			 for(int i = 5; i > 0; i--) {
			     System.out.println(sname + ": " + i);
			      Thread.sleep(1000);
			 }
			
		}catch(InterruptedException e) {
			System.out.println(e);
		}
		System.out.println("Exiting"+sname);
		
	}

}
