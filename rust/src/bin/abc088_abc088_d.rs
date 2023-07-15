use proconio::{input, marker::Chars};
use std::collections::VecDeque;

fn main() {
    input! {
        h : i32,
        w : i32,
        s:[Chars;h],
    }

    let dot_count = s.iter().flatten().filter(|&c| *c == '.').count();

    let mut visited_area = vec![vec![false; w as usize]; h as usize];
    let mut queue = VecDeque::new();

    let dxy = [(0, 1), (0, -1), (1, 0), (-1, 0)];
    queue.push_back((0, 0, 2));
    visited_area[0][0] = true;

    // print!("{} ", count);
    while let Some((x, y, counter)) = queue.pop_front() {
        for (dx, dy) in &dxy {
            // print!("{}-{} ", dx, dy);
            let nx = x + dx;
            let ny = y + dy;
            if nx == h - 1 && ny == w - 1 {
                // println!("answer {}-{}", dot_count, count);
                println!("{}", dot_count - counter);
                return;
            }
            if 0 <= nx && nx <= h - 1 && 0 <= ny && ny <= w - 1 {
                if s[nx as usize][ny as usize] == '.'
                    && visited_area[nx as usize][ny as usize] == false
                {
                    visited_area[nx as usize][ny as usize] = true;
                    queue.push_back((nx, ny, counter + 1));
                }
            }
        }
    }

    println!("{}", -1);
}
