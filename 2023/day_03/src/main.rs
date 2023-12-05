use anyhow::Result;
use std::collections::HashMap;
use std::fs::read_to_string;

struct Numbers {
    number: i32,
    y: i32,
    x: (i32, i32),
}

fn main() -> Result<()> {
    let input: String = read_to_string("input.txt")?.parse()?;
    let (numbers, symbols) = parse(&input)?;

    let mut sum = 0;
    for n in numbers.iter() {
        for (s, _) in symbols.iter() {
            if (n.y - s.0).abs() < 2 {
                if (n.x.0 - s.1).abs() < 2 || (n.x.1 - s.1).abs() < 2 {
                    sum += n.number;
                }
            }
        }
    }
    println!("Part 1: {}", sum);

    let mut sum2: i32 = 0;
    for (s, k) in symbols {
        if k == "*" {
            let mut adjacent = Vec::new();
            for n in numbers.iter() {
                if (s.0 - n.y).abs() < 2 {
                    if (n.x.0 - s.1).abs() < 2 || (n.x.1 - s.1).abs() < 2 {
                        adjacent.push(n.number);
                    }
                }
            }
            if adjacent.len() >= 2 {
                let mut added_nums = 1;
                for x in adjacent.iter() {
                    added_nums *= x;
                }
                sum2 += added_nums;
            }
        }
    }
    println!("Part 2: {}", sum2);

    Ok(())
}

fn parse(input: &String) -> Result<(Vec<Numbers>, HashMap<(i32, i32), String>)> {
    let mut numbers = Vec::new();
    let mut symbols = HashMap::new();
    let mut start = 0;

    for (y, line) in input.lines().enumerate() {
        let mut number = Vec::new();
        for (x, c) in line.chars().enumerate() {
            if c.is_digit(10) {
                if number.is_empty() {
                    start = x as i32;
                    number.push(c);
                } else {
                    number.push(c);
                }
            }

            if !c.is_digit(10) && c != '.' {
                symbols.insert((y as i32, x as i32), c.to_string());
            }

            if (!c.is_digit(10) && !number.is_empty())
                || (!number.is_empty() && x == line.len() - 1)
            {
                let int = number
                    .iter()
                    .collect::<String>()
                    .parse::<i32>()
                    .expect("could not parse vector to int");
                numbers.push(Numbers {
                    number: int,
                    y: y as i32,
                    x: (start, x as i32 - 1),
                });
                start = 0;
                number = Vec::new();
            }
        }
    }
    Ok((numbers, symbols))
}
