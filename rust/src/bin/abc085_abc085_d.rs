use proconio::input;

fn main() {
    input! {
        n: i32,
        mut h: i32,
        sl: [[i32;2];n],
    }

    let mut sll = vec![];
    for s in sl {
        sll.push((s[0], false));
        sll.push((s[1], true));
    }
    sll.sort_unstable_by(|a, b| b.0.cmp(&a.0));

    let mut counter = 0;

    let mut index = 0;
    while h > 0 {
        if !sll[index].1 {
            let add = h / sll[index].0;
            counter += add;
            if (h % sll[index].0) != 0 {
                counter += 1;
            }
            break;
        }
        counter += 1;
        h -= sll[index].0;
        index += 1;
    }
    println!("{}", counter);
}
