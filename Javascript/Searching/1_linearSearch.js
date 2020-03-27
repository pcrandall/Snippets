function linearSearch(arr,val){
    for(let i = 0; i < arr.length; i++){
        if(arr[i] === val) return i;
    }
    return -1;
}

linearSearch([43,512,63,14,52,66,12], 14)
