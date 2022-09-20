use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    time: i32,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // Use `self.number` to refer to each positional data point.
        write!(f, "{}", self.hours_and_minutes())
    }
}

const MINUTES_PER_DAY: i32 = 1440; // minutes per day
const MINUTES_PER_HOUR: i32 = 60; // minutes per hour

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        match hours * MINUTES_PER_HOUR + minutes {
            time if time < 0 => Clock {
                time: MINUTES_PER_DAY + (time - (time / MINUTES_PER_DAY) * MINUTES_PER_DAY),
            },
            time if time > MINUTES_PER_DAY => Clock {
                time: time - (time / MINUTES_PER_DAY) * MINUTES_PER_DAY,
            },
            time if time == MINUTES_PER_DAY => Clock { time: 0 },
            time => Clock { time },
        }
    }

    pub fn hours_and_minutes(&self) -> String {
        format!("{:02}:{:02}", self.hour(), self.minute())
    }

    pub fn hour(&self) -> i32 {
        self.time / MINUTES_PER_HOUR
    }

    pub fn minute(&self) -> i32 {
        self.time % MINUTES_PER_HOUR
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(self.hour(), self.minute() + minutes)
    }
}
