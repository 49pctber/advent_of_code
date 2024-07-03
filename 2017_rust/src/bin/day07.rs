extern crate regex;

use std::collections::HashMap;

fn main() {
    let input = aoc::string_from_file();
    let mut data: HashMap<String, Program> = HashMap::new();
    let mut parents: HashMap<String, String> = HashMap::new();
    parse_tree(&input, &mut data, &mut parents);

    let part1 = find_bottom(&parents);
    println!("Part 1: {}", part1);

    let part2 = find_imbalanced(&data);
    find_imbalanced(&data);
    println!("Part 2: {}", part2);
}

#[derive(Debug, Clone)]
struct Program {
    name: String,
    weight: i32,
    cumulative_weight: i32,
    children: Vec<String>,
}

impl std::str::FromStr for Program {
    type Err = std::num::ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let name: String;
        let weight: i32;
        let mut cumulative_weight: i32 = -1;
        let mut children: Vec<String> = vec![];

        let re = regex::Regex::new(r"^(\w+) \((\d+)\)").expect("error parsing regular expression");
        let re2 = regex::Regex::new(r" -> (.*)$").expect("error parsing regular expression");

        let cap = re.captures(s).expect("error parsing string");
        name = cap[1].parse().expect("error parsing name");
        weight = cap[2].parse().expect("error parsing weight");

        if re2.is_match(s) {
            let cap2 = re2.captures(s).expect("error parsing children");
            let list = cap2[1].split(", ");
            for child in list {
                children.push(child.to_string());
            }
        } else {
            cumulative_weight = weight;
        }

        Ok(Program {
            name,
            weight,
            cumulative_weight,
            children,
        })
    }
}

impl std::fmt::Display for Program {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        if self.children.len() > 0 {
            write!(
                f,
                "{} ({}) [{}] -> {:?}",
                self.name, self.weight, self.cumulative_weight, self.children
            )
        } else {
            write!(
                f,
                "{} ({}) [{}]",
                self.name, self.weight, self.cumulative_weight
            )
        }
    }
}

fn parse_tree(
    input: &String,
    data: &mut HashMap<String, Program>,
    parents: &mut HashMap<String, String>,
) {
    for line in input.lines() {
        let program: Program = str::parse(line).expect("error parsing program");
        let name: String = program.name.clone();

        for child in &program.children {
            parents.insert(child.clone(), program.name.clone());
        }

        data.insert(name, program);
    }
}

fn find_bottom(parents: &HashMap<String, String>) -> String {
    let mut name = parents.keys().next().expect("no keys available").clone();
    loop {
        if parents.contains_key(&name) {
            name = parents.get(&name).expect("not found").clone();
        } else {
            break;
        }
    }
    return name;
}

fn get_cumulative_weight(
    data: &HashMap<String, Program>,
    p: &String,
) -> (i32, HashMap<String, i32>) {
    let program = data.get(p).expect("couldn't find program");

    if program.cumulative_weight != -1 {
        return (program.cumulative_weight, HashMap::new());
    }

    let mut weight = program.weight;
    let mut weights: HashMap<String, i32> = HashMap::new();
    for child in program.children.clone() {
        let (new_weight, _) = get_cumulative_weight(&data, &child);
        weight += new_weight;
        weights.insert(child, new_weight);
    }

    (weight, weights)
}

fn find_imbalanced(data: &HashMap<String, Program>) -> i32 {
    let mut program_vector: Vec<Program> = data.values().cloned().collect();

    let mut p: &Program = &Program {
        name: "".to_string(),
        weight: 0,
        cumulative_weight: 0,
        children: vec![],
    };

    let mut curr_weight = 0;
    let mut other_weight = 0;

    for i in 0..program_vector.len() {
        let c = program_vector.get_mut(i).expect("couldn't find program");
        let (weight, weights) = get_cumulative_weight(data, &c.name);
        c.cumulative_weight = weight;
        if weights.len() == 0 {
            continue;
        }

        let mut x: HashMap<i32, i32> = HashMap::new();
        for (_, w) in weights.clone() {
            if x.contains_key(&w) {
                let y = x.get_mut(&w).expect("couldn't find key");
                *y += 1;
            } else {
                x.insert(w, 1);
            }
        }

        if x.len() <= 1 {
            continue;
        }

        println!("{:?}", weights);
        println!("{:?}", x);

        for (key, &value) in x.iter() {
            if value == 1 {
                for (y, z) in weights.iter() {
                    if key == z {
                        p = data.get(y).expect("msg");
                        curr_weight = *z;
                    } else {
                        other_weight = *z;
                    }
                }
            }
        }

        println!("{}", *p);
        return p.weight - curr_weight + other_weight;
    }

    return -1;
}

#[cfg(test)]
mod tests {
    use crate::find_imbalanced;
    use crate::{find_bottom, parse_tree, Program};
    use std::collections::HashMap;

    #[test]
    fn test_part_1() {
        let input: String = "pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)"
            .to_string();
        let mut data: HashMap<String, Program> = HashMap::new();
        let mut parents: HashMap<String, String> = HashMap::new();
        parse_tree(&input, &mut data, &mut parents);
        assert_eq!(find_bottom(&parents), "tknk");
        assert_eq!(find_imbalanced(&mut data), 60);
    }
}
