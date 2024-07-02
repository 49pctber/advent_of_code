use std::collections::HashSet;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    let fname: &String = args.get(1).expect("missing argument");
    let input = std::fs::read_to_string(fname).expect("could not open file");

    let mut count1 = 0;
    let mut count2 = 0;

    for passphrase in input.lines() {
        if valid_passphrase(passphrase) {
            count1 += 1;
        }

        if valid_passphrase_2(passphrase) {
            count2 += 1;
        }
    }

    println!("Part 1: {}", count1);
    println!("Part 2: {}", count2);
}

fn valid_passphrase(passphrase: &str) -> bool {
    let mut words: HashSet<&str> = HashSet::new();

    for word in (*passphrase).split_whitespace() {
        if words.contains(word) {
            return false;
        }

        words.insert(word);
    }

    return true;
}

fn valid_passphrase_2(passphrase: &str) -> bool {
    let mut words: HashSet<String> = HashSet::new();

    for word in (*passphrase).split_whitespace() {
        // sort word
        let mut chars: Vec<char> = word.chars().collect();
        chars.sort_by(|a, b| b.cmp(a));
        let sorted = chars.into_iter().collect();

        // check if set contains sorted word
        if words.contains(&sorted) {
            return false;
        }

        // insert sorted word
        words.insert(sorted);
    }

    return true;
}

#[cfg(test)]
mod tests {
    use crate::valid_passphrase;
    use crate::valid_passphrase_2;

    #[test]
    fn test_valid_passphrase_0() {
        assert!(valid_passphrase("aa bb cc dd ee"));
    }

    #[test]
    fn test_valid_passphrase_1() {
        assert!(!valid_passphrase("aa bb cc dd aa"));
    }

    #[test]
    fn test_valid_passphrase_2() {
        assert!(valid_passphrase("aa bb cc dd aaa"));
    }

    #[test]
    fn test_valid_passphrase_2_0() {
        assert!(valid_passphrase_2("abcde fghij"));
    }

    #[test]
    fn test_valid_passphrase_2_1() {
        assert!(!valid_passphrase_2("abcde xyz ecdab"));
    }

    #[test]
    fn test_valid_passphrase_2_2() {
        assert!(valid_passphrase_2("a ab abc abd abf abj"));
    }

    #[test]
    fn test_valid_passphrase_2_3() {
        assert!(valid_passphrase_2("iiii oiii ooii oooi oooo"));
    }

    #[test]
    fn test_valid_passphrase_2_4() {
        assert!(!valid_passphrase_2("oiii ioii iioi iiio"));
    }
}
