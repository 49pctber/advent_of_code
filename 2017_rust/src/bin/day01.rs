fn main() {
    let args: Vec<String> = std::env::args().collect();
    let input = std::fs::read_to_string(&args[1]).expect("Could not read file");
    let part1 = captcha(&input);
    let part2 = captcha2(&input);

    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
}

fn captcha(s: &String) -> i32 {
    let b = s.as_bytes();
    let offset: u8 = b'0';
    let mut ret: i32 = 0;

    for i in 0..b.len() {
        let curr = b[i];
        let next = b[(i + 1) % b.len()];
        if curr == next {
            ret += (b[i] - offset) as i32;
        }
    }

    return ret;
}

fn captcha2(s: &String) -> i32 {
    let b = s.as_bytes();
    let offset: u8 = b'0';
    let mut ret: i32 = 0;

    for i in 0..b.len() {
        let curr = b[i];
        let next = b[(i + b.len() / 2) % b.len()];
        if curr == next {
            ret += (b[i] - offset) as i32;
        }
    }

    return ret;
}

#[cfg(test)]
mod tests {
    use crate::captcha;
    use crate::captcha2;

    #[test]
    fn test_captcha_0() {
        let input: String = "1122".to_string();
        let result = captcha(&input);
        assert_eq!(result, 3);
    }

    #[test]
    fn test_captcha_1() {
        let input: String = "1111".to_string();
        let result = captcha(&input);
        assert_eq!(result, 4);
    }

    #[test]
    fn test_captcha_2() {
        let input: String = "1234".to_string();
        let result = captcha(&input);
        assert_eq!(result, 0);
    }

    #[test]
    fn test_captcha_3() {
        let input: String = "91212129".to_string();
        let result = captcha(&input);
        assert_eq!(result, 9);
    }

    #[test]
    fn test_captcha2_0() {
        let input: String = "1212".to_string();
        let result = captcha2(&input);
        assert_eq!(result, 6);
    }

    #[test]
    fn test_captcha2_1() {
        let input: String = "1221".to_string();
        let result = captcha2(&input);
        assert_eq!(result, 0);
    }

    #[test]
    fn test_captcha2_2() {
        let input: String = "123425".to_string();
        let result = captcha2(&input);
        assert_eq!(result, 4);
    }

    #[test]
    fn test_captcha2_3() {
        let input: String = "123123".to_string();
        let result = captcha2(&input);
        assert_eq!(result, 12);
    }

    #[test]
    fn test_captcha2_4() {
        let input: String = "12131415".to_string();
        let result = captcha2(&input);
        assert_eq!(result, 4);
    }
}
