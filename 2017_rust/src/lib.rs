/// Read a file based on the second command-line argument and returns it as a string
pub fn string_from_file() -> String {
    let args: Vec<String> = std::env::args().collect();
    let fname: &String = args.get(1).expect("missing argument");
    let input = std::fs::read_to_string(fname).expect("could not open file");
    return input;
}
