use std::fs;
use std::path::Path;
use std::process::exit;

fn lines_from_file(filename: impl AsRef<Path>) -> Vec<u32> {
    let file_lines: Vec<u32> = fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(|s| s.parse().unwrap())
        .collect();
    file_lines
}

fn main(){
    let lines = lines_from_file("input.txt");
    let sum = 2020;
    for line in lines.iter() {
        for line2 in lines.iter(){
            if line != line2 {
                if (line + line2) <= sum{
                    let wanted_number = sum - line - line2;
                    if lines.contains(&wanted_number){
                        print!("{}\n", wanted_number * line * line2);
                        exit(0)
                    }
                }
            }
        }
    }
}