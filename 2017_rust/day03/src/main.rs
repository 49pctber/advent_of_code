use std::cmp::min;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    let input: i32 = str::parse::<i32>(&args[1]).expect("could not parse input");

    let part1 = distance(&input);
    let part2 = next_largest_allocation(&input);
    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
}

fn distance(input: &i32) -> i32 {
    if *input == 1 {
        return 0;
    }

    let mut i = 0;
    while i32::pow(2 * i + 1, 2) < *input {
        i += 1;
    }

    let mut j = i32::pow(2 * i + 1, 2);
    while j - 2 * i > *input {
        j -= 2 * i;
    }

    return 2 * i - min((j - input).abs(), (j - 2 * i - input).abs());
}

fn next_largest_allocation(input: &i32) -> i32 {
    return *input;
}

#[cfg(test)]
mod tests {
    use crate::distance;
    use crate::next_largest_allocation;

    #[test]
    fn test_distance_1() {
        let input = 1;
        assert_eq!(distance(&input), 0);
    }

    #[test]
    fn test_distance_2() {
        let input = 2;
        assert_eq!(distance(&input), 1);
    }

    #[test]
    fn test_distance_3() {
        let input = 3;
        assert_eq!(distance(&input), 2);
    }

    #[test]
    fn test_distance_4() {
        let input = 4;
        assert_eq!(distance(&input), 1);
    }

    #[test]
    fn test_distance_6() {
        let input = 6;
        assert_eq!(distance(&input), 1);
    }

    #[test]
    fn test_distance_7() {
        let input = 7;
        assert_eq!(distance(&input), 2);
    }

    #[test]
    fn test_distance_12() {
        let input = 12;
        assert_eq!(distance(&input), 3);
    }

    #[test]
    fn test_distance_14() {
        let input = 14;
        assert_eq!(distance(&input), 3);
    }

    #[test]
    fn test_distance_23() {
        let input = 23;
        assert_eq!(distance(&input), 2);
    }

    #[test]
    fn test_distance_26() {
        let input = 26;
        assert_eq!(distance(&input), 5);
    }

    #[test]
    fn test_distance_27() {
        let input = 27;
        assert_eq!(distance(&input), 4);
    }

    #[test]
    fn test_distance_28() {
        let input = 28;
        assert_eq!(distance(&input), 3);
    }

    #[test]
    fn test_distance_29() {
        let input = 29;
        assert_eq!(distance(&input), 4);
    }

    #[test]
    fn test_distance_49() {
        let input = 49;
        assert_eq!(distance(&input), 6);
    }

    #[test]
    fn test_distance_1024() {
        let input = 1024;
        assert_eq!(distance(&input), 31);
    }

    #[test]
    fn test_next_largest_allocation_2() {
        let input = 2;
        assert_eq!(next_largest_allocation(&input), 4);
    }

    #[test]
    fn test_next_largest_allocation_4() {
        let input = 4;
        assert_eq!(next_largest_allocation(&input), 5);
    }

    #[test]
    fn test_next_largest_allocation_300() {
        let input = 300;
        assert_eq!(next_largest_allocation(&input), 304);
    }
}
