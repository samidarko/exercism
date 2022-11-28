use std::collections::HashMap;
use std::sync::mpsc;
use std::thread;

pub fn characters_count(input: &[&str]) -> HashMap<char, usize> {
    let mut counts: HashMap<char, usize> = HashMap::new();

    input.iter().for_each(|s| {
        s.to_lowercase()
            .chars()
            .filter(|c| c.is_alphabetic())
            .for_each(|c| {
                *counts.entry(c).or_insert(0) += 1;
            });
    });

    counts
}

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let chunk_size = {
        let chunk_size = (input.len() as f32 / worker_count as f32).floor() as usize;
        if chunk_size == 0 {
            1
        } else {
            chunk_size
        }
    };

    let chunks = input.chunks(chunk_size);
    let worker_count = chunks.len(); // worker count == number of chunks
    let (tx, rx) = mpsc::sync_channel(worker_count);

    thread::scope(move |s| {
        chunks.for_each(|chunk| {
            let tx = tx.clone();
            s.spawn(move || tx.send(characters_count(chunk)));
        });
    });

    let mut result: HashMap<char, usize> = HashMap::new();

    for _ in 0..worker_count {
        if let Ok(counts) = rx.recv() {
            for (c, count) in counts {
                *result.entry(c).or_insert(0) += count;
            }
        }
    }

    result
}
