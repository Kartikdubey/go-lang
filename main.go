package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
// the biggerIsGreater function below.
func biggerIsGreater(w string) string {
    var i int
    r:=[]rune(w)

    //Start from the left most character and find the first
    //character smaller than previous character
    for i=len(r)-1;i>0;i--{
        if r[i]>r[i-1]{
             break }
     } 
     right:=w[i:]
     r1:=[]rune(right)
     //If no such  character is found, it means string already in descending order
     if i==0{
         return "no answer"
     }else{
         var min =r1[0]
         var index=0
	 //Now find the smallest character on right part which is also greater than the 
	 //previous character found above
         for l,v :=range(r1){
             if min>v && v>r[i-1]{ 
                 min=v
                 index=l
             }
         } 
	 //Swap both the charcaters
        rs:=index+i
        t:=r[i-1]
        r[i-1]=r[rs]
        r[rs]=t
	//Bubble Sort the right part of string now
        for t:=i;t<len(r);t++{
            for j:=t+1;j<len(r);j++{
                if r[t]>r[j]{
                    temp:=r[t]
                    r[t]=r[j]
                    r[j]=temp
                }
            }
        }
return string(r)
     }

}

func main() {

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    w := getString()
    result := biggerIsGreater(w)
    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func getString() string {
    data,err := ioutil.ReadFile("./inputfile")
    checkError(err)
    return strings.TrimRight(string(data), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
