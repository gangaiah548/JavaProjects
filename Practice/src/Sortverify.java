import java.util.ArrayList;

class GoodStanding{
	private String crud;
	private String name;

	public GoodStanding(String crud, String name) {
		this.crud=crud;
		this.name=name;
	}
	
	public String getCrud() {
		return crud;
	}

	public void setCrud(String crud) {
		this.crud = crud;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}
}
class RegFileStatus{
	private String crud;
	private String name;
	
 public RegFileStatus(String crud, String name) {
	this.crud=crud;
	this.name=name;
}

	public String getCrud() {
		return crud;
	}

	public void setCrud(String crud) {
		this.crud = crud;
	}
	
	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}
	
}
public class Sortverify {
	public static void sortgds(ArrayList<GoodStanding> list)
	{
		list.sort((o1,o2)->o2.getCrud().compareTo(o1.getCrud()));
	}
	
	public static void sortfds(ArrayList<RegFileStatus> list)
	{
		list.sort((o1,o2)->o2.getCrud().compareTo(o1.getCrud()));
	}
	
	public static void main(String[] ar)
	{
		ArrayList<GoodStanding> gdslist=new ArrayList<GoodStanding>();
		ArrayList<RegFileStatus>rgslist=new ArrayList<RegFileStatus>();
		gdslist.add(new GoodStanding("C","gs"));
		gdslist.add(new GoodStanding("D","gs1"));
		gdslist.add(new GoodStanding("C","gs2"));
		gdslist.add(new GoodStanding("D","gs3"));
		
		rgslist.add(new RegFileStatus("C","fs"));
		rgslist.add(new RegFileStatus("D","fs1"));
		rgslist.add(new RegFileStatus("C","fs5"));
		rgslist.add(new RegFileStatus("D","fs1"));
		
		sortfds(rgslist);
		sortgds(gdslist);
		
		System.out.println(gdslist);
		
		for(GoodStanding obj:gdslist) {
			System.out.println(obj.getCrud()+" "+obj.getName());
		}
		
for(RegFileStatus obj:rgslist) {
	System.out.println(obj.getCrud()+" "+obj.getName());
		}
		
	}
}
