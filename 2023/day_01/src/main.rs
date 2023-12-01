use anyhow::Result;
use std::fs::read_to_string;

fn main() -> Result<()> {
    let input: String = read_to_string("input.txt")?.parse()?;
    let _ = solve_one(&input);
    let _ = solve_two(&input);
    Ok(())
}

fn solve_one(input: &String) -> Result<()> {
    let mut sum = 0;
    for line in input.lines() {
        let number = get_ints(line).unwrap();
        sum += number;
    }
    println!("Sum 1: {}", sum);
    Ok(())
}

fn solve_two(input: &String) -> Result<()> {
    let mut sum = 0;
    for line in input.lines() {
        let number = get_ints_and_strings(line).unwrap();
        sum += number;
    }
    println!("Sum 2: {}", sum);

    Ok(())
}

fn get_ints_and_strings(line: &str) -> Option<u32> {
    let mut string_numbers = Vec::new();
    string_numbers.push("one");
    string_numbers.push("two");
    string_numbers.push("three");
    string_numbers.push("four");
    string_numbers.push("five");
    string_numbers.push("six");
    string_numbers.push("seven");
    string_numbers.push("eight");
    string_numbers.push("nine");

    let mut int_vector = Vec::new();

    for (i, x) in line.chars().enumerate() {
        if x.is_digit(10) {
            int_vector.push(x.to_string());
            continue;
        }
        for (ix, yx) in string_numbers.iter().enumerate() {
            let lx: String = line.chars().skip(i).collect();
            if lx.starts_with(yx) {
                int_vector.push((ix + 1).to_string());
            }
        }
    }
    let mut parsed_int: String = int_vector.first()?.to_string();
    parsed_int = parsed_int + &int_vector.last()?;

    parsed_int.parse::<u32>().ok()
}

fn get_ints(data: &str) -> Option<u32> {
    let mut int_string = String::new();
    for x in data.chars() {
        if x.is_digit(10) {
            int_string += &x.to_string();
            break;
        }
    }

    for x in data.chars().rev() {
        if x.is_digit(10) {
            int_string += &x.to_string();
            break;
        }
    }

    int_string.parse::<u32>().ok()
}
