trait IsVowel {
    fn is_vowel(&self) -> bool;
}

impl IsVowel for char {
    fn is_vowel(&self) -> bool {
        "aeiouy".contains(|v| v == *self)
    }
}

trait PigLatin {
    fn to_pig_latin(&self) -> String;
}

impl PigLatin for &str {
    fn to_pig_latin(&self) -> String {
        if self.is_empty() {
            return "".to_string();
        }
        let mut output = String::new();
        let first_char = self.chars().nth(0).unwrap();

        match self {
            s if s.len() == 2 && s.chars().nth(1) == Some('y') => {
                output.push('y');
                output.push(s.chars().nth(0).unwrap());
            }
            s if (first_char.is_vowel() && &s[..2] != "ye") || &s[..2] == "xr" => {
                output.push_str(self);
            }
            s if &s[1..=2] == "qu" => {
                output.push_str(&s[3..]);
                output.push(first_char);
                output.push_str("qu");
            }
            s if ["thr", "sch"].contains(&&s[..3]) => {
                output.push_str(&s[3..]);
                output.push_str(&s[..3]);
            }
            s if ["ch", "qu", "th", "rh"].contains(&&s[..2]) => {
                output.push_str(&s[2..]);
                output.push_str(&s[..2]);
            }
            s => {
                output.push_str(&s[1..]);
                output.push(first_char);
            }
        }

        output.push_str("ay");
        output
    }
}

pub fn translate(input: &str) -> String {
    let mut translation: Vec<String> = vec![];
    for word in input.split(" ") {
        translation.push(word.to_pig_latin());
    }
    translation.join(" ")
}

// TODO add interface is_vowel to char
