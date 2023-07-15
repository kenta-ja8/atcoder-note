use proconio::input;

fn main() {
    input! {
        mut c:[[i32;3];3],
    }
    for a1 in 0..101 {
        let b1 = c[0][0] - a1;
        let b2 = c[0][1] - a1;
        let b3 = c[0][2] - a1;
        let a2 = c[1][0] - b1;
        let a3 = c[2][0] - b1;
        if a2 + b2 == c[1][1] && a2 + b3 == c[1][2] && a3 + b2 == c[2][1] && a3 + b3 == c[2][2] {
            println!("Yes");
            return;
        }
    }
    println!("No");
}
