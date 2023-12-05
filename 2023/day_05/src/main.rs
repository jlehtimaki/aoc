use anyhow::Result;
use std::fs::read_to_string;

#[derive(Clone)]
struct MapLine {
    destination_range_start: u64,
    source_range_start: u64,
    range_length: u64,
}

#[derive(Clone)]
struct CategoryMap(Vec<MapLine>);

struct Almanac {
    seeds: Vec<u64>,
    maps: Vec<CategoryMap>,
}

fn main() -> Result<()> {
    let input: String = read_to_string("input.txt")?.parse()?;
    let parse = parse_input(input.as_str());
    let p1 = part1(&parse);
    let p2 = part2(&parse);

    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);

    Ok(())
}

fn parse_input(almanac: &str) -> Almanac {
    use aoc_parse::{parser, prelude::*};

    let almanac_parser = parser!(
        rule map_line: MapLine = destination_range_start:u64 " " source_range_start:u64 " " range_length:u64 =>
            MapLine {
                destination_range_start,
                source_range_start,
                range_length,
            };

        rule category_map: CategoryMap =
            line(string(any_char+))
            map_lines:lines(map_line) =>
                CategoryMap(map_lines);

        seeds:section(line("seeds: " repeat_sep(u64, " ")))
        maps:sections(category_map) =>
            Almanac {
                seeds,
                maps,
            }
    );

    almanac_parser.parse(almanac).unwrap()
}

fn location(seed: &u64, maps: &[CategoryMap]) -> u64 {
    let mut id = *seed;

    for category_map in maps.iter() {
        for map_line in category_map.0.iter() {
            if id >= map_line.source_range_start
                && id < map_line.source_range_start + map_line.range_length
            {
                id = map_line.destination_range_start + (id - map_line.source_range_start);
                break;
            }
        }
    }

    id
}

fn part1(almanac: &Almanac) -> u64 {
    almanac
        .seeds
        .iter()
        .map(|seed| location(seed, &almanac.maps))
        .min()
        .unwrap()
}

fn part2(almanac: &Almanac) -> u64 {
    let mut source_ranges = Vec::new();

    for seed_range in almanac.seeds.chunks_exact(2) {
        source_ranges.push((seed_range[0], seed_range[0] + seed_range[1] - 1));
    }

    for category_map in almanac.maps.iter() {
        let mut destination_ranges = Vec::new();

        'source_ranges: while let Some(source_range) = source_ranges.pop() {
            for map_line in category_map.0.iter() {
                let line_source_start = map_line.source_range_start;
                let line_source_end = map_line.source_range_start + map_line.range_length - 1;

                let line_destination_start = map_line.destination_range_start;

                if line_source_start <= source_range.1 && line_source_end >= source_range.0 {
                    if source_range.0 < line_source_start {
                        source_ranges.push((source_range.0, line_source_start - 1));
                    }

                    if source_range.1 > line_source_end {
                        source_ranges.push((line_source_end + 1, source_range.1));
                    }

                    destination_ranges.push((
                        u64::max(line_source_start, source_range.0) - line_source_start
                            + line_destination_start,
                        u64::min(line_source_end, source_range.1) - line_source_start
                            + line_destination_start,
                    ));

                    continue 'source_ranges;
                }
            }

            destination_ranges.push((source_range.0, source_range.1));
        }

        source_ranges = destination_ranges;
    }

    source_ranges
        .iter()
        .map(|(range_min, _)| *range_min)
        .min()
        .unwrap()
}
