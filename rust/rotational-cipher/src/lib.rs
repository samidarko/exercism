// use std::collections::HashMap;

const ALPHABET: [char; 26] = [
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
    't', 'u', 'v', 'w', 'x', 'y', 'z',
];

pub fn rotate(input: &str, key: i8) -> String {
    input
        .chars()
        .map(|c| match c.to_ascii_lowercase() {
            _ if c.is_alphabetic() => ALPHABET
                .into_iter()
                .position(|letter| letter == c.to_ascii_lowercase())
                .map(|i| {
                    let rotated = match i + key as usize {
                        position if position >= 26 => ALPHABET[position - 26],
                        position => ALPHABET[position],
                    };
                    if c.is_uppercase() {rotated.to_ascii_uppercase()} else { rotated }
                })
                .unwrap(),
            _ => c,
        })
        .collect::<String>()
}

// func RotationalCipher(s string, shift int) string {
// 	var output strings.Builder
// 	for _, r := range s {
// 		if unicode.IsLetter(r) {
// 			output.WriteRune(rotate(r, shift))
// 		} else {
// 			output.WriteRune(r)
// 		}
// 	}
// 	return output.String()
// }
//
// func rotate(r rune, shift int) rune {
// 	isUpper := unicode.IsUpper(r)
// 	r = unicode.ToLower(r)
//
// 	r += int32(shift)
//
// 	if r > 'z' {
// 		// if passed 'z' returns to 'a'
// 		r = 'a' + (r - ('z' + 1))
// 	}
//
// 	if isUpper {
// 		return unicode.ToUpper(r)
// 	}
// 	return r
// }
