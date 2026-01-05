# Practice Go

This repository is for me to help practice Go programming concepts progressively.

Based on: https://gobyexample.com/

## Architecture

This project uses a **modular template system** with **shared utilities** where each Go concept has its own self-contained directory:

```
01HelloWorld/
├── .practice/                       # Template source (preserved across cleanups)
│   ├── generate.py                  # Lightweight generator (uses shared utils)
│   ├── metadata.json                # Topic metadata (key, display_name, index)
│   └── template.go                  # Go code template
├── hello_world.go                   # Generated practice file (your workspace)
└── go.mod                           # Generated module file

practice_utils.py                    # Shared utilities for all templates
```

### Key Design Principles

- **DRY (Don't Repeat Yourself)**: Common code extracted to `practice_utils.py`
- **Single Source of Truth**: One place to update shared logic  
- **Easy to Maintain**: Each template only contains what's unique to it
- **Easy to Extend**: Add new templates by creating a directory with `.practice/`
- **Easy to Experiment**: Test individual templates without affecting others
- **Version Control Friendly**: Track changes to specific templates easily

## 🚀 Quick Start

### 1. Generate All Practice Modules

```sh
python3 setup_go_practice.py
```

This will:
- Discover all `.practice` directories
- Execute each template's `generate.py` script
- Create `.go` and `go.mod` files for practice
- Generate a `go.work` file for multi-module support

### 2. Practice a Concept

Navigate to any module and start coding:

```sh
cd 01HelloWorld/
# Read hello_world.go and complete the TODOs
```

### 3. Run Your Code

**From terminal:**
```sh
go run hello_world.go
```

**From VS Code/Cursor:**
- Open the `.go` file
- Use Command+Shift+B (task runner) or Control+F5 (run)

### 4. Clean Up for Fresh Practice

```sh
python3 setup_go_practice.py --clean
```

This removes:
- Generated `.go` and `go.mod` files from all practice modules
- `go.work` workspace file

But **preserves**:
- All `.practice` directories with templates
- Python scripts and utilities

**Note**: `__pycache__` directories (Python bytecode cache) are hidden in VS Code via `.vscode/settings.json` and excluded from version control via `.gitignore`.

## 📚 Available Modules

The repository includes **84 progressive Go practice modules** covering:

1. **Basics**: Hello World, Values, Variables, Constants
2. **Control Flow**: For, If/Else, Switch
3. **Data Structures**: Arrays, Slices, Maps
4. **Functions**: Functions, Multiple Return Values, Variadic Functions, Closures
5. **Advanced Types**: Structs, Methods, Interfaces, Generics
6. **Concurrency**: Goroutines, Channels, Select, WaitGroups, Mutexes
7. **Error Handling**: Errors, Custom Errors, Panic, Defer, Recover
8. **Standard Library**: String Functions, JSON, XML, Time, HTTP
9. **System**: File I/O, Command-Line Args, Environment Variables, Signals
10. **Testing**: Testing and Benchmarking

## 🛠️ Advanced Usage

### Adding a New Practice Template

1. Create a numbered directory (e.g., `85NewConcept/`)
2. Create `.practice/` subdirectory
3. Add three files:

**metadata.json:**
```json
{
  "key": "new-concept",
  "display_name": "New Concept",
  "index": 85
}
```

**template.go:**
```go
// Your Go template code with TODO comments
package main

import "fmt"

func main() {
    // TODO: Implement new concept
}
```

**generate.py:**
```python
#!/usr/bin/env python3
# Use the template from create_template_structure.py
# or copy from an existing module
```

4. Run `python3 setup_go_practice.py` - it will auto-discover your new template!

### Modifying an Existing Template

1. Navigate to the template directory:
   ```sh
   cd 01HelloWorld/.practice/
   ```

2. Edit `template.go` with your changes

3. Test the template:
   ```sh
   python3 generate.py
   ```

4. Regenerate all modules:
   ```sh
   cd ../..
   python3 setup_go_practice.py --clean
   python3 setup_go_practice.py
   ```

### Testing a Single Template

```sh
cd 01HelloWorld/.practice/
python3 generate.py
cd ..
go run hello_world.go
```

## 📖 Learning Approach

1. **Read**: Open a practice file and read the instructions
2. **Complete**: Implement the TODOs following the comments
3. **Test**: Run the program and verify it works
4. **Understand**: Ask questions, experiment with variations
5. **Enhance**: Add your own improvements or related features
6. **Repeat**: Move to the next module

## 🔧 Scripts Reference

### `practice_utils.py`

Shared utilities module used by all `generate.py` scripts. Contains:

- `get_metadata(practice_dir)` - Load metadata.json
- `get_template(practice_dir)` - Load template.go
- `topic_to_package_name(topic)` - Convert topic to Go package name
- `generate_files(practice_dir, target_dir)` - Generate .go and go.mod files

**Benefits**: 
- Single source of truth for common logic
- Update once, affects all 84 templates
- Reduces each generate.py from 80 to 34 lines (56% reduction)

### `setup_go_practice.py`

Main script for managing practice modules.

```sh
# Create all modules (default)
python3 setup_go_practice.py

# Clean all modules (preserves templates)
python3 setup_go_practice.py --clean

# Show help
python3 setup_go_practice.py --help
```

### `create_template_structure.py`

One-time migration script (already run). Used to convert the original monolithic template system to the modular architecture.

```sh
# Run migration (creates .practice directories)
python3 create_template_structure.py
```

## 📂 Project Structure

```
practicego/
├── README.md                          # This file
├── setup_go_practice.py               # Main script (simplified, ~150 lines)
├── create_template_structure.py      # Migration utility
├── go.work                            # Go workspace file (generated)
├── 01HelloWorld/
│   ├── .practice/                     # Template source
│   └── hello_world.go                 # Generated practice file
├── 02Values/
│   ├── .practice/
│   └── values.go
...
└── 84Exit/
    ├── .practice/
    └── exit.go
```

## 🎯 References

- **Language Docs**: https://go.dev/doc/
- **Go by Example**: https://gobyexample.com/
- **Effective Go**: https://go.dev/doc/effective_go
- **Go Spec**: https://go.dev/ref/spec

## 🎨 Keyboard Shortcuts (VS Code/Cursor)

- **Command+Shift+B**: Build/Run task
- **Control+F5**: Run without debugging
- **F5**: Debug

## 💡 Tips

- Practice concepts in order (01 → 84) for progressive learning
- Use `--clean` frequently to practice from scratch
- Modify templates to create your own variations
- Each module is independent - jump to what interests you
- The `go.work` file enables your editor to understand all modules

## 🤝 Contributing

This is a personal practice repository, but feel free to:
- Fork and adapt for your own learning
- Suggest improvements via issues
- Share your practice approach

Happy Go learning! 🚀