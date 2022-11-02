use num_bigint::ToBigUint;
use num_traits::cast::ToPrimitive;
use rand::Rng;

pub fn private_key(p: u64) -> u64 {
    let max = p - 2;
    let num = rand::thread_rng().gen_range(0..max);
    num + 2
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    // g.pow(a as u32) % p

    g.to_biguint()
        .unwrap()
        .modpow(&a.to_biguint().unwrap(), &p.to_biguint().unwrap())
        .to_u64()
        .unwrap()
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    // b_pub.pow(a as u32) % p

    b_pub
        .to_biguint()
        .unwrap()
        .modpow(&a.to_biguint().unwrap(), &p.to_biguint().unwrap())
        .to_u64()
        .unwrap()
}
