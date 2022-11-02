use rand::Rng;

pub fn private_key(p: u64) -> u64 {
    let max = p - 2;
    let num = rand::thread_rng().gen_range(0..max);
    num + 2
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    // g.pow(a as u32) % p
    modular_pow(g, a, p)
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    // b_pub.pow(a as u32) % p
    modular_pow(b_pub, a, p)
}

pub fn modular_pow(base: u64, mut exponent: u64, modulus: u64) -> u64 {
    let mut base = base as u128;
    let modulus = modulus as u128;
    if modulus == 1 {
        return 0;
    }

    let mut r: u128 = 1;
    base = base % modulus;
    while exponent > 0 {
        if exponent % 2 == 1 {
            r = (r * base) % modulus;
        }
        base = (base * base) % modulus;
        exponent = exponent >> 1;
    }
    r as u64
}
