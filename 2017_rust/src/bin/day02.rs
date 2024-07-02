fn main() {
    let args: Vec<String> = std::env::args().collect();
    let input = std::fs::read_to_string(&args[1]).expect("Could not read file");
    let part1 = checksum(&input);
    let part2 = checksum2(&input);
    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
}

fn checksum(input: &String) -> i32 {
    let rows = input.lines();
    let mut ret = 0;

    for row in rows {
        let values = row.split_whitespace();
        let mut max = std::i32::MIN;
        let mut min = std::i32::MAX;

        for vs in values {
            let v = str::parse::<i32>(vs).expect("could not parse value");
            if v > max {
                max = v;
            }

            if v < min {
                min = v;
            }
        }
        ret += max - min;
    }

    return ret;
}

fn checksum2(input: &String) -> i32 {
    let rows = input.lines();
    let mut ret = 0;

    for row in rows {
        let valuestrings = row.split_whitespace();
        let mut values: Vec<i32> = vec![];

        for s in valuestrings {
            let v: i32 = str::parse::<i32>(s).expect("could not parse value");
            values.insert(0, v);
        }

        values.sort();

        for i in 0..values.len() {
            for j in 0..i {
                if values[i] % values[j] == 0 {
                    ret += values[i] / values[j];
                }
            }
        }
    }

    return ret;
}

#[cfg(test)]
mod tests {
    use crate::checksum;
    use crate::checksum2;

    #[test]
    fn test_checksum_0() {
        let input: String = "5 1 9 5
7 5 3
2 4 6 8"
            .to_string();
        let result = checksum(&input);
        assert_eq!(result, 18);
    }

    #[test]
    fn test_checksum2_0() {
        let input: String = "5 9 2 8
9 4 7 3
3 8 6 5"
            .to_string();
        let result = checksum2(&input);
        assert_eq!(result, 9);
    }
}
