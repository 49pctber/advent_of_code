extern crate regex;

use std::collections::HashMap;

fn main() {
    let input = aoc::string_from_file();
    let mut data: HashMap<String, Program> = HashMap::new();
    let mut parents: HashMap<String, String> = HashMap::new();
    parse_tree(&input, &mut data, &mut parents);

    let part1 = find_bottom(&parents);
    println!("Part 1: {}", part1);

    // let part2 = find_imbalanced(&mut data);
    // println!("Part 2: {}", part2);
}

#[derive(Debug, Clone)]
struct Program {
    name: String,
    weight: i32,
    cumulative_weight: i32,
    children: Vec<String>,
}

impl Program {
    fn set_weight(&mut self, weight: i32) {
        self.weight = weight;
    }

    fn get_cumulative_weight(&self) -> i32 {
        self.cumulative_weight
    }

    fn reset_cumulative_weight(&mut self) {
        self.cumulative_weight = self.weight;
    }

    fn add_cumulative_weight(&mut self, weight: i32) {
        self.cumulative_weight += weight;
    }
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

// fn recursively_compute_weight(data: &HashMap<String, Program>, node: String) -> i32 {
//     let p: &Program = data.get(&node).expect("couldn't find node");

//     if p.cumulative_weight != -1 {
//         return p.cumulative_weight;
//     }

//     p.reset_cumulative_weight();
//     for child in p.children.clone() {
//         p.add_cumulative_weight(recursively_compute_weight(data, child.clone()));
//     }

//     p.get_cumulative_weight()
// }

// fn find_imbalanced(data: &mut HashMap<String, Program>) -> i32 {
//     let keys: Vec<String> = data.keys().cloned().collect();
//     for node in keys {
//         recursively_compute_weight(data, node.clone());
//     }

//     return 0;
// }

#[cfg(test)]
mod tests {
    use crate::{find_bottom, parse_tree, Program};
    // use crate::{find_imbalanced};
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
        // assert_eq!(find_imbalanced(&data), 60);
    }
}
