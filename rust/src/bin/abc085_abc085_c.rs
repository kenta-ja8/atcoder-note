use proconio::input;

fn main() {
    input! {
        n: i32,
        amount: i32,
    }
    for i10 in 0..n + 1 {
        for i5 in 0..n + 1 {
            let i1 = n - i10 - i5;
            if i1 < 0 {
                break;
            }
            let total = 10000 * i10 + 5000 * i5 + 1000 * i1;
            if total > amount {
                break;
            }
            if total < amount {
                continue;
            }
            println!("{} {} {}", i10, i5, i1);
            return;
        }
    }

    println!("-1 -1 -1");
}
