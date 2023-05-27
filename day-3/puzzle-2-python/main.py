def get_priorities() -> dict:
    map = {}
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for index, letter in enumerate(list(letters)):
        map[letter] = index + 1
    
    return map

def main():
    print("Advent of Code - day 3 puzzle 2, Python")
    
    # Read input from file
    lines = []
    with open("../input.txt", "r") as file:
        lines = file.read().splitlines()
    
    groups = []
    for i in range(0, len(lines), 3):
        group = [lines[i], lines[i+1], lines[i+2]]
        groups.append(group)
    
    priorities = get_priorities()
    
    total_priority = 0    
    for group in groups:
        og_set = set(list(group[0]))
        print(og_set)
        first_inclusive = og_set.intersection(set(list(group[1])))
        print(first_inclusive)
        second_inclusive = first_inclusive.intersection(set(list(group[2])))
        
        priority = priorities[list(second_inclusive)[0]]
        print(f"second_inclusive: {second_inclusive} "
              f"has priority {priority}")

        total_priority += priority
    
    print(f"Total priority: {total_priority}")
        
    
     
if __name__ == "__main__":   
    main()