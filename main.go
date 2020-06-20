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
<<<<<<< HEAD
=======
    //Start from the left most character and find the first
    //character smaller than previous character
>>>>>>> 313b9bd... Go-Lang App deployed using a Docker container
    for i=len(r)-1;i>0;i--{
        if r[i]>r[i-1]{
             break }
     } 
     right:=w[i:]
     r1:=[]rune(right)
<<<<<<< HEAD
=======
     //If no such  character is found, it means string already in descending order
>>>>>>> 313b9bd... Go-Lang App deployed using a Docker container
     if i==0{
         return "no answer"
     }else{
         var min =r1[0]
         var index=0
<<<<<<< HEAD
=======
	 //Now find the smallest character on right part which is also greater than the 
	 //previous character found above
>>>>>>> 313b9bd... Go-Lang App deployed using a Docker container
         for l,v :=range(r1){
             if min>v && v>r[i-1]{ 
                 min=v
                 index=l
             }
         } 
<<<<<<< HEAD
=======
	 //Swap both the charcaters
>>>>>>> 313b9bd... Go-Lang App deployed using a Docker container
        rs:=index+i
        t:=r[i-1]
        r[i-1]=r[rs]
        r[rs]=t
<<<<<<< HEAD
=======
	//Bubble Sort the right part of string now
>>>>>>> 313b9bd... Go-Lang App deployed using a Docker container
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
