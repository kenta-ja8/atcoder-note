use proconio::input;

fn main() {
    input! {
        n: i32,
        a: i32,
        b: i32,
    }
    let mut total = 0;
    for i in 1..n+1 {
        let mut sum = 0;
        let mut j = i;
        while j > 0 {
            sum += j % 10;
            j /= 10;
        }
        if a <= sum && sum <= b {
            total += i;
        }
    }
    println!("{}", total);
}
