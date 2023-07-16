use proconio::input;

fn main() {
    input! {
        n: i32,
        mut ds: [i32;n],
    }
    ds.sort_unstable();

    let mut total = 1;
    for i in 1..n {
        if ds[i as usize] != ds[(i - 1) as usize] {
            total += 1;
        }
    }
    println!("{}", total);
}
