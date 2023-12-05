use anyhow::Result;
use std::fs::read_to_string;

#[derive(Debug)]
struct Game {
    card: Vec<i32>,
    winning: Vec<i32>,
}

fn main() -> Result<()> {
    let input: String = read_to_string("input.txt")?.parse()?;
    let games = parse(&input)?;
    let mut scratchcards: Vec<i32> = vec![0; games.len()];
    let mut sum = 0;
    for (x, g) in games.iter().enumerate() {
        let mut count = 0;
        let mut wins = 0;
        for c in g.card.iter() {
            if g.winning.contains(&c) {
                wins += 1;
                if count == 0 {
                    count += 1;
                    continue;
                }
                // part 1
                count *= 2;
            }
        }
        scratchcards[x] = wins;

        // part 1
        sum += count;
    }
    println!("Part 1: {}", sum);

    //part 2
    let mut accumulated_wins = vec![1; games.len()];
    for (i, wins) in scratchcards.iter().enumerate() {
        for ix in 1..*wins + 1 {
            accumulated_wins[i + ix as usize] += accumulated_wins[i];
        }
    }
    println!("Part 2: {:?}", accumulated_wins.iter().sum::<i32>());

    Ok(())
}

fn parse(input: &str) -> Result<Vec<Game>> {
    let mut games = vec![];
    for line in input.lines() {
        let parts: Vec<&str> = line.split('|').collect();

        let card: Vec<i32> = parts[0]
            .split_whitespace()
            .filter_map(|s| s.parse().ok())
            .collect();
        let winning: Vec<i32> = parts
            .get(1)
            .unwrap_or(&"")
            .split_whitespace()
            .filter_map(|s| s.parse().ok())
            .collect();
        games.push(Game { card, winning });
    }

    Ok(games)
}
