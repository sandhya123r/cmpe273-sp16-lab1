package main

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3


type LRUCache struct{

  Hashtable map[int]*node
  size int
  count int
  ll *list

}

type node struct{
  key int
  value int
  next *node
  prev *node
}

type list struct{
  head *node
  tail *node
}

func addnode(l *list,n *node){
  /*
if(l==nil){
    newlist:=new(list)
    l:=newlist
    l.head=n
    l.head.prev=nil
    l.tail=l.head
    l.head.next=nil
    return
  }*/
  n.prev=l.head
  n.next=l.head.next
  l.head.next.prev=n
  l.head.next=n
}

func removenode(l *list,n *node){

  
  prev_node:=n.prev
  next_node:=n.next
  //fmt.Println("prev= ",prev_node,"next =",next_node)
  if(next_node==nil){
    prev_node.next=nil
  }else{
    prev_node.next=next_node
    next_node.prev=prev_node
  }
  
}

func movetohead(l *list, n *node){
  /*
if(l==nil){
    return
  }*/
  
  removenode(l,n)
  addnode(l,n)
}

func pop(l *list)*node{ 
  popped_node:=l.tail.prev
  //fmt.Println("value=",popped_node.key)
  //if(l.tail!=nil && l.tail.prev!=nil){
    //l.tail=l.tail.prev
  //
  //fmt.Println("popping")
  removenode(l,popped_node)
  //fmt.Println("popped ")
  return popped_node
}

func CreateList(mylist *list){
  
  /*
  head:=new(node)
  head.prev=nil
  */
  
  mylist.head =new(node)
  mylist.head.prev=nil

  mylist.tail=new(node)
  mylist.head.next=mylist.tail
  mylist.tail.next=nil
  mylist.tail.prev=mylist.head
  //fmt.Println("Creating list",mylist.head.value,mylist.tail.value)
  //return mylist
}

func Cache()*LRUCache{

  cache := new(LRUCache)
  cache.count=0
  cache.size=CACHE_SIZE
  cache.ll=new(list)
  CreateList(cache.ll)
  
 
  return cache
}


var cache = Cache()

func Set(key int, value int) {
  
  hash:=cache.Hashtable
  targetnode,flag:=hash[key]
  if(!flag){
    newnode:=new (node)
    newnode.key=key
    newnode.value=value
    //fmt.Println("Created hashtable")
    if(len(cache.Hashtable)==0){
      Hashtable:=make(map[int]*node)
      cache.Hashtable=Hashtable
    }
    cache.Hashtable[key]=newnode
   // fmt.Println("Printing hash table",cache.Hashtable[key])
    addnode(cache.ll,newnode)
    cache.count++
    if(cache.count>cache.size){
      //fmt.Println("Count > size ")
      
      tailnode:=pop(cache.ll)
      //fmt.Println(tailnode)
      _, ok :=cache.Hashtable[tailnode.key]
      //fmt.Println("value=",cache.Hashtable[tailnode.key])
      if(ok){
        delete(cache.Hashtable,tailnode.key)
        //fmt.Println("HASHTABLE : ",cache.Hashtable)
        //fmt.Println("Deleted key")
      }
      cache.count--
    }
  } else {
    targetnode.value=value
    movetohead(cache.ll,targetnode)
  }

}

func Get(key int) int {
	
  hash:=cache.Hashtable
  //fmt.Println("Get: ",cache.Hashtable)
  targetnode,flag:=hash[key]
  //fmt.Println(key , "Get function : ",hash[key])
  if(!flag){
    return -1
  }
  
  linkedlist:=cache.ll
  movetohead(linkedlist,targetnode)
  
  return targetnode.value
}



