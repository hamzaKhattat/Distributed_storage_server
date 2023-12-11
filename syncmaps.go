package main
import ("fmt"
	"sync")
func main(){
//	j:=make(int)
//	maps:=make(map[int]int)
	maps:=sync.Map{}
	go func(){
	for i:=0;i<=10;i++{
		maps.Store(i,i)
		}
	}}()
	for i:=0;i<=10;i++{
		k,v:=maps.Load(i)
		
	fmt.Println(k,":",v)
		}
	maps.Delete(5)
	for i:=0;i<=10;i++{
		k,v:=maps.Load(i)
		
	fmt.Println(k,":",v)
		
		}
	//get and put
	value,loader:=maps.LoadOrStore(2,7)
	fmt.Println("The Value is:",value,"Loaded:",loader,"\n")
	maps.Range(func(key,value interface{}) bool{
	fmt.Println(key,value)
	return true
	})
	
}
