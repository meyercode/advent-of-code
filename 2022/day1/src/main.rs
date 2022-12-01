use std::cmp::Ordering;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

#[derive(Eq)]
struct ElfInventory {
    food_items: Vec<i32>,
    calorie_count: i32,
}

impl PartialOrd for ElfInventory {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        self.calorie_count.partial_cmp(&other.calorie_count)
    }
}

impl Ord for ElfInventory {
    fn cmp(&self, other: &Self) -> Ordering {
        self.calorie_count.cmp(&other.calorie_count)
    }
}

impl PartialEq for ElfInventory {
    fn eq(&self, other: &Self) -> bool {
        self.calorie_count == other.calorie_count
    }
}

fn main() {
    let file_path = "./inputs";

    let mut elves: Vec<ElfInventory> = Vec::new();

    if let Ok(lines) = read_lines(file_path) {
        let mut food_items: Vec<i32> = Vec::new();

        for line in lines {
            if let Ok(food_item) = line {
                if let Ok(food_item) = food_item.parse::<i32>() {
                    food_items.push(food_item);
                } else {
                    elves.push(ElfInventory {
                        food_items: food_items.clone(),
                        calorie_count: food_items.iter().sum(),
                    });
                    food_items = Vec::new();
                }
            }
        }
        elves.push(ElfInventory {
            food_items: food_items.clone(),
            calorie_count: food_items.iter().sum(),
        });

        let max_value = elves.iter().max().unwrap().calorie_count;
        println!("Elf with most calories: {}", max_value);
        elves.sort_by(|a, b| b.calorie_count.cmp(&a.calorie_count));
        let top_three = elves[0].calorie_count + elves[1].calorie_count + elves[2].calorie_count;
        println!("Top three: {}", top_three);
    }
}

fn read_lines<P>(file_name: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(file_name)?;
    Ok(io::BufReader::new(file).lines())
}
