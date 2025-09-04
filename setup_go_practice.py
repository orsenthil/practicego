#!/usr/bin/env python3
"""
Script to set up Go practice modules for learning Go programming concepts.
Creates directories with Go modules, each containing a template .go file and go.mod file.
"""

import os
import re
import sys

# List of Go concepts to create practice modules for
TOPICS = [
    "Hello World",
    "Values",
    "Variables", 
    "Constants",
    "For",
    "If/Else",
    "Switch",
    "Arrays",
    "Slices",
    "Maps",
    "Functions",
    "Multiple Return Values",
    "Variadic Functions",
    "Closures",
    "Recursion",
    "Range over Built-in Types",
    "Pointers",
    "Strings and Runes",
    "Structs",
    "Methods",
    "Interfaces",
    "Enums",
    "Struct Embedding",
    "Generics",
    "Range over Iterators",
    "Errors",
    "Custom Errors",
    "Goroutines",
    "Channels",
    "Channel Buffering",
    "Channel Synchronization",
    "Channel Directions",
    "Select",
    "Timeouts",
    "Non-Blocking Channel Operations",
    "Closing Channels",
    "Range over Channels",
    "Timers",
    "Tickers",
    "Worker Pools",
    "WaitGroups",
    "Rate Limiting",
    "Atomic Counters",
    "Mutexes",
    "Stateful Goroutines",
    "Sorting",
    "Sorting by Functions",
    "Panic",
    "Defer",
    "Recover",
    "String Functions",
    "String Formatting",
    "Text Templates",
    "Regular Expressions",
    "JSON",
    "XML",
    "Time",
    "Epoch",
    "Time Formatting / Parsing",
    "Random Numbers",
    "Number Parsing",
    "URL Parsing",
    "SHA256 Hashes",
    "Base64 Encoding",
    "Reading Files",
    "Writing Files",
    "Line Filters",
    "File Paths",
    "Directories",
    "Temporary Files and Directories",
    "Embed Directive",
    "Testing and Benchmarking",
    "Command-Line Arguments",
    "Command-Line Flags",
    "Command-Line Subcommands",
    "Environment Variables",
    "Logging",
    "HTTP Client",
    "HTTP Server",
    "Context",
    "Spawning Processes",
    "Exec'ing Processes",
    "Signals",
    "Exit"
]

def topic_to_package_name(topic):
    """
    Convert a topic name to a valid Go package name.
    - Convert to lowercase
    - Replace spaces and special characters with underscores
    - Remove consecutive underscores
    - Ensure it starts with a letter
    """
    # Convert to lowercase and replace problematic characters
    name = topic.lower()
    name = re.sub(r'[^a-z0-9]+', '_', name)
    name = re.sub(r'_+', '_', name)  # Remove consecutive underscores
    name = name.strip('_')  # Remove leading/trailing underscores
    
    # Ensure it starts with a letter (prepend 'go_' if it starts with a number)
    if name and name[0].isdigit():
        name = 'go_' + name
    
    return name

def create_go_file(directory, package_name, topic):
    """Create a template .go file with package main and empty main function."""
    go_filename = f"{package_name}.go"
    go_filepath = os.path.join(directory, go_filename)
    
    go_content = f'''package main

// {topic}
// Practice exercises for learning Go programming concepts

import "fmt"

func main() {{
    fmt.Println("Learning: {topic}")
    // TODO: Add your practice code here
}}
'''
    
    with open(go_filepath, 'w') as f:
        f.write(go_content)
    
    print(f"  Created {go_filename}")

def create_go_mod(directory, package_name):
    """Create a go.mod file for the module."""
    go_mod_path = os.path.join(directory, "go.mod")
    
    go_mod_content = f'''module github.com/orsenthil/gobyexample/{package_name}

go 1.25
'''
    
    with open(go_mod_path, 'w') as f:
        f.write(go_mod_content)
    
    print(f"  Created go.mod")

def create_practice_module(topic, base_dir):
    """Create a complete practice module for a given topic."""
    package_name = topic_to_package_name(topic)
    directory = os.path.join(base_dir, package_name)
    
    print(f"\nCreating module for '{topic}' -> {package_name}")
    
    # Create directory
    os.makedirs(directory, exist_ok=True)
    
    # Create .go file
    create_go_file(directory, package_name, topic)
    
    # Create go.mod file
    create_go_mod(directory, package_name)

def create_go_workspace(base_dir, topics):
    """Create a go.work file to manage all modules in the workspace."""
    go_work_path = os.path.join(base_dir, "go.work")
    
    print("\nCreating Go workspace file...")
    
    go_work_content = "go 1.25\n\nuse (\n"
    for topic in topics:
        package_name = topic_to_package_name(topic)
        go_work_content += f"    ./{package_name}\n"
    go_work_content += ")\n"
    
    with open(go_work_path, 'w') as f:
        f.write(go_work_content)
    
    print("  Created go.work (enables multi-module workspace)")

def main():
    """Main function to set up all Go practice modules."""
    script_dir = os.path.dirname(os.path.abspath(__file__))
    print(f"Setting up Go practice modules in: {script_dir}")
    print(f"Total modules to create: {len(TOPICS)}")
    
    # Create all modules
    for topic in TOPICS:
        create_practice_module(topic, script_dir)
    
    # Create Go workspace file
    create_go_workspace(script_dir, TOPICS)
    
    print(f"\n✅ Successfully created {len(TOPICS)} Go practice modules!")
    print("\nTo run a specific module:")
    print("  From terminal: cd <module_directory> && go run <module_name>.go")
    print("  From editor: Open any .go file and use the Run/Debug buttons")
    print("\nThe go.work file enables multi-module support in your editor.")
    print("Happy coding! 🚀")

if __name__ == "__main__":
    main()
