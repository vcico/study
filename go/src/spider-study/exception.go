package main

func mian(){
    outerFunc()
}

func outerFunc(){
    innerFunc()
}

func innerFunc(){
    panic(errors.New("An intended fatal error!"))
}