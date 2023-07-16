use std::{collections::HashMap, usize};

use proconio::input;

struct UnionFind {
    parent: Vec<usize>,
    rank: Vec<usize>,
}

impl UnionFind {
    fn new(n: usize) -> Self {
        Self {
            parent: (0..n).collect(),
            rank: vec![0; n],
        }
    }
    fn find_root(&mut self, t: usize) -> usize {
        let tp = self.parent[t];
        if tp == t {
            t
        } else {
            self.parent[t] = self.find_root(tp);
            self.parent[t]
        }
    }
    fn unite(&mut self, f: usize, s: usize) {
        let fp = self.find_root(f);
        let sp = self.find_root(s);
        if self.rank[fp] > self.rank[sp] {
            self.parent[sp] = fp;
        } else {
            self.parent[fp] = sp;
            if self.rank[fp] == self.rank[sp] {
                self.rank[sp] += 1;
            }
        }
    }
}
fn main() {
    input! {
        n: usize,
        k: usize,
        l: usize,
        pql: [[usize;2];k],
        rsl: [[usize;2];l],
    }

    let mut pqu = UnionFind::new(n);
    for pq in pql {
        pqu.unite(pq[0] - 1, pq[1] - 1);
    }

    let mut rsu = UnionFind::new(n);
    for rs in rsl {
        rsu.unite(rs[0] - 1, rs[1] - 1);
    }

    let mut map = HashMap::new();
    for i in 0..n {
        *map.entry((pqu.find_root(i), rsu.find_root(i))).or_insert(0) += 1;
    }

    let ans = (0..n)
        .map(|i| map[&(pqu.find_root(i), rsu.find_root(i))].to_string())
        .collect::<Vec<_>>().join(" ");
    println!("{}", ans);
}
