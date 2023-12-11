use anyhow::Result;
use std::fs::read_to_string;

#[derive(Debug)]
struct Race {
    time: u64,
    distance: u64,
}

fn main() -> Result<()> {
    let input: String = read_to_string("input.txt")?.parse()?;
    let race1 = parse1(&input)?;
    let race2 = parse2(&input)?;

    let p1 = solve(&race1)?;
    let p2 = solve(&race2)?;

    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);

    Ok(())
}

fn solve(races: &Vec<Race>) -> Result<i32> {
    let mut ans = 1;

    for race in races {
        let mut count = 0;
        for i in 0..race.time {
            if i * (race.time - i) > race.distance {
                count += 1;
            }

            if (race.distance - i) < 1 {
                break;
            }
        }

        ans *= count;
    }

    Ok(ans)
}

fn parse1(input: &String) -> Result<Vec<Race>> {
    let mut races: Vec<Race> = Vec::new();

    let numbers: Vec<u64> = input
        .lines()
        .flat_map(|line| line.split_whitespace())
        .filter_map(|n| n.parse::<u64>().ok())
        .collect();

    let x = numbers.len() / 2;
    for i in 0..x {
        races.push(Race {
            time: numbers[i],
            distance: numbers[x + i],
        })
    }

    Ok(races)
}

fn parse2(input: &String) -> Result<Vec<Race>> {
    let mut races: Vec<Race> = Vec::new();
    let numbers: Vec<u64> = input
        .lines()
        .map(|line| {
            line.split_whitespace()
                .filter_map(|word| word.parse::<u64>().ok())
                .map(|num| num.to_string())
                .collect::<String>()
                .parse::<u64>()
                .unwrap_or(0)
        })
        .collect();

    races.push(Race {
        time: numbers[0],
        distance: numbers[1],
    });

    Ok(races)
}
