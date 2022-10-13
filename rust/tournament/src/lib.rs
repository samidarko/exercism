use std::cmp::Ordering;
use std::collections::HashMap;
use std::iter::once;

#[derive(Clone)]
struct Stat {
    name: String,
    played: u32,
    won: u32,
    drawn: u32,
    lost: u32,
    points: u32,
}

impl Stat {
    pub fn new() -> Stat {
        Self {
            name: "".to_string(),
            played: 0,
            won: 0,
            drawn: 0,
            lost: 0,
            points: 0,
        }
    }
}

pub fn tally(match_results: &str) -> String {
    let events = match_results
        .split(|c| c == '\n')
        .map(|s| s.split(|c| c == ';').collect())
        .collect::<Vec<Vec<_>>>();

    let mut stats: HashMap<&str, Stat> = HashMap::new();

    for event in events {
        if let [home_team, away_team, result] = event.as_slice() {
            // let home_stat = stats.entry(home_team).or_insert(Stat::new());
            let mut home_stat = stats.get(home_team).unwrap_or(&Stat::new()).clone();
            home_stat.name = home_team.to_string();
            home_stat.played += 1;

            let mut away_stat = stats.get(away_team).unwrap_or(&Stat::new()).clone();
            away_stat.name = away_team.to_string();
            away_stat.played += 1;

            match *result {
                "win" => {
                    home_stat.won += 1;
                    home_stat.points += 3;
                    away_stat.lost += 1;
                }
                "loss" => {
                    away_stat.won += 1;
                    away_stat.points += 3;
                    home_stat.lost += 1;
                }
                "draw" => {
                    away_stat.drawn += 1;
                    away_stat.points += 1;
                    home_stat.drawn += 1;
                    home_stat.points += 1;
                }
                _ => (),
            }

            stats.insert(home_team, home_stat);
            stats.insert(away_team, away_stat);
        }
    }
    let mut stats = stats.values().collect::<Vec<_>>();
    stats.sort_by(|a, b| match a.points.cmp(&b.points) {
        Ordering::Equal => a.name.cmp(&b.name),
        _ => b.points.cmp(&a.points),
    });

    let next_rows = stats
        .iter()
        .map(|stat| {
            format!(
                "{:31}|{:3} |{:3} |{:3} |{:3} |{:3}",
                stat.name, stat.played, stat.won, stat.drawn, stat.lost, stat.points
            )
        })
        .collect::<Vec<String>>();

    once(format!(
        "Team                           | MP |  W |  D |  L |  P"
    ))
    .chain(next_rows.into_iter())
    .collect::<Vec<String>>()
    .join("\n")
}
