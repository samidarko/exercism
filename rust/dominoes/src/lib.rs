pub fn is_valid_chain(dominos: &[(u8, u8)]) -> bool {
    if dominos.is_empty() {
        return true;
    }
    dominos[0].0 == dominos[dominos.len() - 1].1
}

pub fn concat(xs: &[(u8, u8)], ys: &[(u8, u8)]) -> Vec<(u8, u8)> {
    let mut v = vec![];

    for x in xs {
        v.push(*x);
    }

    for y in ys {
        v.push(*y);
    }

    v
}

pub fn chain(dominos: &[(u8, u8)]) -> Option<Vec<(u8, u8)>> {
    if is_valid_chain(dominos) {
        return Some(dominos.to_vec());
    }

    let mut dominos = dominos.to_vec();
    dominos.sort_unstable_by(|a, b| {
        if a.0 != b.0 {
            return a.0.cmp(&b.0);
        }
        b.1.cmp(&a.1)
    });

    let mut chain = vec![dominos[0]];
    let mut dominos = dominos[1..].to_vec();
    let mut found;

    while !dominos.is_empty() {
        found = false;
        for i in 0..dominos.len() {
            let last_dots = chain[chain.len() - 1].1;
            let domino = dominos[i].clone();
            if last_dots == domino.0 {
                chain.push(domino);
                dominos.remove(i);
                found = true;
                break;
            }
            if last_dots == domino.1 {
                chain.push((domino.1, domino.0));
                dominos.remove(i);
                found = true;
                break;
            }
            if domino.0 == domino.1 {
                for j in 0..chain.len() - 1 {
                    if chain[j].1 == chain[j + 1].0 && chain[j].1 == domino.0 {
                        chain.insert(j + 1, domino);
                        dominos.remove(i);
                        found = true;
                        break;
                    }
                }
            }
        }

        if !found {
            break;
        }
    }

    if is_valid_chain(&chain) && dominos.len() == 0 {
        return Some(chain);
    }

    None
}
