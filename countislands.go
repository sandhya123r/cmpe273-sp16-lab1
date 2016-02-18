package main

//import "fmt"

/*
func readinput(matrix [][]int,N int,M int){
   
   var element int 
   for i:=0;i<N;i++ {

     matrix[i]=make([]int,M)
     for j:=0;j<M;j++{
     
       fmt.Scan(&element)
       matrix[i][j]=element
     }
   }
   
   
   fmt.Println("You have entered ",matrix)
   
 }
 */

 func dfs(matrix[][] int,x int,y int){
   
   if(x<0 || y<0||x>=len(matrix)|| y>=len(matrix[0])){
    return;
   }
   if(matrix[x][y]==0){
    return;
   }
   matrix[x][y]=0;
   dfs(matrix,x,y+1)
   dfs(matrix,x,y-1)
   dfs(matrix,x+1,y)
   dfs(matrix,x-1,y)
         
   
 }
 
 func CountIslands(matrix[][]int)int{
  
   counter :=0
   for i:=0;i<len(matrix);i++ {
     for j:=0;j<len(matrix[i]);j++ {
     
       if (matrix[i][j]==1) {
         counter++;
         dfs(matrix,i,j)
       }
     }
   }
   return counter 
 }
