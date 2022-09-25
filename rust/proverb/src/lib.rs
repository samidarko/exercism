pub fn build_proverb(list: &[&str]) -> String {
    match list.len() {
        0 => String::new(),
        1 => and_all_for_the_want_of_a_thing(list[0]),
        _ => {
            let proverb: Vec<String> = list
                .windows(2)
                .map(|x| for_want_of_a_thing_the_other_thing_was_lost(x[0], x[1]))
                .collect();
            format!(
                "{}\n{}",
                proverb.join("\n"),
                and_all_for_the_want_of_a_thing(list[0])
            )
        }
    }
}

fn for_want_of_a_thing_the_other_thing_was_lost(thing: &str, other_thing: &str) -> String {
    format!("For want of a {thing} the {other_thing} was lost.")
}

fn and_all_for_the_want_of_a_thing(thing: &str) -> String {
    format!("And all for the want of a {thing}.")
}
