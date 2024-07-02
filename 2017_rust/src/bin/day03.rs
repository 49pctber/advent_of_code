use std::{cmp::min, collections::HashMap};

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
    let mut c2v: HashMap<String, i32> = std::collections::HashMap::new();

    let (mut x, mut y) = (0, 0); // coordinates
    let mut v = 1; // the value at a given coordinate
    let mut dir = 0; // incrementing rotates 90 degrees counter-clockwise
    let mut steps_remaining_in_dir = 1;

    // initial conditions
    let c: String = format!("{}.{}", x, y).to_string();
    c2v.insert(c.clone(), 1);

    // "allocate memory"
    while v <= *input {
        // step
        match dir % 4 {
            0 => x += 1,
            1 => y += 1,
            2 => x -= 1,
            3 => y -= 1,
            _ => return -1,
        }
        steps_remaining_in_dir -= 1;

        // turn if necessary
        if steps_remaining_in_dir == 0 {
            dir += 1; // turn
            steps_remaining_in_dir = (dir + 2) / 2;
        }

        // add new value to hashmap
        v = 0;
        for dx in -1..2 {
            for dy in -1..2 {
                if dx == 0 && dy == 0 {
                    continue;
                }
                let c: String = format!("{}.{}", x + dx, y + dy).to_string();
                if c2v.contains_key(&c) {
                    v += c2v.get(&c).expect("coordinate not in hashmap");
                }
            }
        }
        c2v.insert(format!("{}.{}", x, y).to_string(), v);
    }

    return v;
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

    #[test]
    fn test_next_largest_allocation_6591() {
        let input = 6591;
        assert_eq!(next_largest_allocation(&input), 13486);
    }

    #[test]
    fn test_next_largest_allocation_42452() {
        let input = 42452;
        assert_eq!(next_largest_allocation(&input), 45220);
    }
}
