const SIGNS: [&str; 4] = ["wink", "double blink", "close your eyes", "jump"];
const REVERSE_SIGNS: u8 = 16;

pub fn actions(code: u8) -> Vec<&'static str> {
    let (mut action, action_incr, end) = match code & REVERSE_SIGNS {
        0 => (0, 1, 4),
        _ => (3, -1, -1),
    };

    let mut output = vec![];

    while action != end {
        if (code & (1 << action)) != 0 {
            output.push(SIGNS[action as usize]);
        }

        action += action_incr;
    }

    output
}
