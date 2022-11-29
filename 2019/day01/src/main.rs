use std::{
    fs::File,
    io::{BufRead, BufReader},
    path::Path,
};

fn main() {
    let file = File::open(Path::new("input.txt")).unwrap();
    let reader = BufReader::new(&file);
    let lines: Vec<String> = reader.lines().collect::<Result<_, _>>().unwrap();
    let converted_lines = convert_to_int(lines);
    for n in &converted_lines {
        for n2 in &converted_lines {
            for n3 in &converted_lines {
                if n + n2 + n3 == 2020 {
                    let result = n * n2 * n3;
                    println!("{result}");
                    return;
                }
            }
        }
    }
}

fn convert_to_int(v: Vec<String>) -> Vec<i32> {
    let mut converted: Vec<i32> = Vec::new();
    for l in v {
        converted.push(
            l.parse::<i32>()
                .unwrap_or_else(|_| panic!("Could not convert line to int")),
        );
    }
    converted
}
