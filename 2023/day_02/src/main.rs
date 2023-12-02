use anyhow::Result;
use regex::Regex;
use std::collections::HashMap;
use std::fs::read_to_string;

const RED_LIMIT: i32 = 12;
const GREEN_LIMIT: i32 = 13;
const BLUE_LIMIT: i32 = 14;

fn main() -> Result<()> {
    let input: String = read_to_string("input.txt")?.parse()?;
    let _ = solve(&input);
    Ok(())
}

fn solve(input: &String) -> Result<()> {
    let mut sum1 = 0;
    let mut sum2 = 0;
    let re = Regex::new(r"Game (\d+)").unwrap();
    for line in input.lines() {
        let game_id = re.captures(line).unwrap().get(1).unwrap().as_str();
        let game_id: i32 = game_id.parse()?;
        let parts: Vec<&str> = line.split(":").collect();
        let balls = parts[1].trim();
        if sum_game(balls) {
            sum1 += game_id;
        };
        sum2 += sum_game2(balls)
    }
    println!("Sum 1: {}", sum1);
    println!("Sum 2: {}", sum2);
    Ok(())
}

fn sum_game(games: &str) -> bool {
    let games_array: Vec<&str> = games.split(";").collect();
    for game in games_array {
        let balls: Vec<&str> = game.split(",").collect();
        for ball in balls {
            let parts: Vec<&str> = ball.split_whitespace().collect();
            let number: i32 = parts[0].parse().expect("Not a number");
            let color = parts[1];

            if color == "green" {
                if number > GREEN_LIMIT {
                    return false;
                }
            }
            if color == "red" {
                if number > RED_LIMIT {
                    return false;
                }
            }
            if color == "blue" {
                if number > BLUE_LIMIT {
                    return false;
                }
            }
        }
    }
    true
}

fn sum_game2(games: &str) -> i32 {
    let mut sum = 1;
    let mut colors = HashMap::from([("green", 0), ("red", 0), ("blue", 0)]);
    let games_array: Vec<&str> = games.split(";").collect();
    for game in games_array {
        let balls: Vec<&str> = game.split(",").collect();
        for ball in balls {
            let parts: Vec<&str> = ball.split_whitespace().collect();
            let number: i32 = parts[0].parse().expect("Not a number");
            let color = parts[1];
            if colors[color] < number {
                colors.insert(color, number);
            }
        }
    }

    for (_, v) in colors {
        sum = sum * v;
    }
    sum
}
