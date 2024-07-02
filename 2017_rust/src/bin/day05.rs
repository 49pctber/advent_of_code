fn main() {
    let input = aoc::string_from_file();

    let part1 = count_steps(&input);
    println!("Part 1: {}", part1);

    let part2 = count_steps_2(&input);
    println!("Part 2: {}", part2);
}

fn count_steps(input: &String) -> i32 {
    let mut offsets: Vec<i32> = vec![];

    for line in input.lines() {
        let x = str::parse::<i32>(line).expect("could not parse line");
        offsets.push(x);
    }

    let mut i: i32 = 0;
    let mut count = 0;

    loop {
        count += 1;
        offsets[i as usize] += 1;
        i += offsets[i as usize] - 1;
        if i >= offsets.len() as i32 {
            break;
        }
    }

    return count;
}

fn count_steps_2(input: &String) -> i32 {
    let mut offsets: Vec<i32> = vec![];

    for line in input.lines() {
        let x = str::parse::<i32>(line).expect("could not parse line");
        offsets.push(x);
    }

    let mut i: i32 = 0;
    let mut count = 0;

    loop {
        count += 1;

        let move_amount = offsets[i as usize];

        if offsets[i as usize] >= 3 {
            offsets[i as usize] -= 1;
        } else {
            offsets[i as usize] += 1;
        }

        i += move_amount;

        if i >= offsets.len() as i32 {
            break;
        }
    }

    return count;
}

#[cfg(test)]
mod tests {
    use crate::{count_steps, count_steps_2};

    #[test]
    fn test_part_1() {
        let input: String = "0
3
0
1
-3"
        .to_string();

        assert_eq!(count_steps(&input), 5);
    }

    #[test]
    fn test_part_2() {
        let input: String = "0
3
0
1
-3"
        .to_string();

        assert_eq!(count_steps_2(&input), 10);
    }
}
