
fn main(){
    const SIZE:usize = 217532235;
    let arr = vec![2, 3, 5, 7, 11];
    let mut out = vec![0; SIZE];
    let mut arrcnt = arr[0];
    out[1] = arr[0];
    let mut cnt = 1;
    let mut sum = 0;
    while arrcnt != 217532235 {
        let e = arr[cnt%5];
        let o = out[cnt];
        for i in arrcnt..arrcnt+o {
            out[i] = e;
        }
        if e == 7 {
            sum += o
        }
        arrcnt+=o;
        cnt+=1;
    }
    println!("{}", sum*7);
}
