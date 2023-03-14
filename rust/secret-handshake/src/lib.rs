const SIGNS: [&str; 4] = ["wink", "double blink", "close your eyes", "jump"];
const REVERSE_SIGNS: u8 = 16;

pub fn actions(code: u8) -> Vec<&'static str> {
    let mut action: i32 = 0;
    let mut action_incr: i32 = 1;
    let mut end: i32 = SIGNS.len() as i32;

    if (code & REVERSE_SIGNS) != 0 {
        action = 3;
        action_incr = -1;
        end = -1;
    }

    let mut output = vec![];

    while action != end {
        if (code & (1 << action)) != 0 {
            output.push(SIGNS[action as usize]);
        }

        action += action_incr;
    }

    output
}
