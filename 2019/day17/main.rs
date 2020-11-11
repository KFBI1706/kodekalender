
fn digits(mut n: usize) -> usize {
    let mut cnt:usize = 0;
    if n == 0 {
        return 1
    }
    while n > 0 {
        cnt += 1;
        n /= 10;
    }
    return cnt
}

fn rotate(t: usize) -> Vec<usize> {
    let d = digits(t);
    let mut rotates = vec![0usize; d];
    for i in 0..d {
        rotates[i] = t/10usize.pow(i as u32)+t%10usize.pow(i as u32)*10usize.pow((d-i) as u32);
    }
    return rotates
}

fn main() {
    let mut ttall = vec![0usize; 1000000];
    let mut k : usize = 1;
    for i in 1..1000000 {
        ttall[i] = k;
        k += i+1
    }

    let mut cnt = 0;
    for t in ttall {
        for n in rotate(t) {
            if ((n as f64).sqrt() as usize).pow(2) == n {
                cnt += 1;
                break
            }
        }
    }
    println!("{}", cnt);
}
