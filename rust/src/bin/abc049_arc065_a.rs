use std::collections::VecDeque;

use proconio::input;

fn main() {
    input! {
        t: String,
    }

    let str_vec = vec!["dream", "dreamer", "erase", "eraser"];

    let mut queue = VecDeque::new();
    queue.push_back("".to_string());

    while let Some(cur) = queue.pop_front() {
        for s in &str_vec {
            let cstr = format!("{}{}", cur, s);
            if t.starts_with(&cstr) {
                if t == cstr {
                    println!("YES");
                    return;
                }
                queue.push_back(cstr)
            }
        }
    }
    println!("NO");
}
