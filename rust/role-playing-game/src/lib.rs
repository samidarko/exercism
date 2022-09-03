// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        match self {
            Player { mana: Some(_), health: 0, .. } => Some(Player { health: 100, mana: Some(100), level: self.level }),
            Player { mana: None, health: 0, .. } => Some(Player { health: 100, mana: None, level: self.level }),
            _ => None
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match *self {
            Player { mana: Some(m), .. } if m >= mana_cost => {
                self.mana = Some(m - mana_cost);
                mana_cost * 2
            }
            Player { mana: Some(_), .. } => 0,
            Player { mana: None, .. } => {
                self.health = if self.health >= mana_cost { self.health - mana_cost } else { 0 };
                0
            }
            _ => 0
        }
    }
}
