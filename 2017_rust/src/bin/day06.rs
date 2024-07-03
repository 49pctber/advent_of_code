use std::vec;

fn main() {
    let input = aoc::string_from_file();
    let (part1, part2) = count_cycles(&input);
    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
}

fn vec2str(banks: &Vec<i32>) -> String {
    let mut ret: String = "".to_string();
    for x in banks {
        ret = format!("{}.{}", ret, x)
    }
    return ret;
}

fn count_cycles(input: &String) -> (i32, i32) {
    // initialize banks
    let x: Vec<&str> = input.split_whitespace().into_iter().collect();
    let mut banks: Vec<i32> = vec![];
    for y in x {
        banks.push(str::parse::<i32>(y).expect("couldn't parse input value"));
    }

    // redistribute
    let mut count = 0;
    let cycle_length;
    let mut seen: std::collections::HashMap<String, usize> = std::collections::HashMap::new();
    seen.insert(vec2str(&banks), 0);

    loop {
        count += 1;

        // find max
        let mut max = std::i32::MIN;
        let mut maxi = 0;
        for (i, x) in banks.iter().enumerate() {
            if *x > max {
                max = *x;
                maxi = i;
            }
        }

        // redistribute
        let mut remaining = max;
        let mut i = maxi;
        banks[maxi] = 0;
        while remaining > 0 {
            i = (i + 1) % banks.len();
            banks[i] += 1;
            remaining -= 1;
        }

        // check to see if we have been in this state before
        let vstr = vec2str(&banks);
        if seen.contains_key(&vstr) {
            let j = seen.get(&vstr).expect("idk");
            cycle_length = count - *j;
            break;
        }
        seen.insert(vstr, count);
    }

    return (count as i32, cycle_length as i32);
}

#[cfg(test)]
mod tests {
    use crate::count_cycles;

    #[test]
    fn test() {
        let input: String = "0 2 7 0".to_string();
        let (count, cycle_length) = count_cycles(&input);
        assert_eq!(count, 5);
        assert_eq!(cycle_length, 4);
    }
}
