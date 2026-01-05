#!/usr/bin/env python3
"""
Script to set up Go practice modules for learning Go programming concepts.
Creates directories with Go modules using modular template system.
Templates are stored in .practice directories within each module.
"""

import os
import sys
import argparse
import importlib.util
import glob

def discover_templates(base_dir):
    """
    Scan for all .practice directories with generate.py scripts.
    
    Args:
        base_dir: Base directory to search for templates
        
    Returns:
        List of tuples: (module_dir, template_script_path) sorted by directory name
    """
    pattern = os.path.join(base_dir, "*", ".practice", "generate.py")
    template_scripts = glob.glob(pattern)
    
    # Extract module directory and sort by prefix (01, 02, etc.)
    templates = []
    for script_path in template_scripts:
        # Get the module directory (parent of .practice)
        module_dir = os.path.dirname(os.path.dirname(script_path))
        module_name = os.path.basename(module_dir)
        templates.append((module_name, module_dir, script_path))
    
    # Sort by module name (which includes numeric prefix)
    templates.sort(key=lambda x: x[0])
    
    return templates

def load_and_execute_template(script_path, target_dir):
    """
    Load a template's generate.py script and execute its generate() function.
    
    Args:
        script_path: Path to the generate.py script
        target_dir: Directory where files should be generated
        
    Returns:
        bool: True if successful, False otherwise
    """
    try:
        # Load the module dynamically
        spec = importlib.util.spec_from_file_location("template_generator", script_path)
        module = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(module)
        
        # Execute the generate function
        return module.generate(target_dir)
    except Exception as e:
        print(f"  ❌ Error executing {script_path}: {e}")
        return False

def create_practice_modules(base_dir):
    """
    Discover and execute all template scripts to create practice modules.
    
    Args:
        base_dir: Base directory containing template directories
    """
    print(f"🚀 Setting up Go practice modules in: {base_dir}")
    
    templates = discover_templates(base_dir)
    
    if not templates:
        print("❌ No templates found! Run create_template_structure.py first.")
        return
    
    print(f"📚 Found {len(templates)} practice templates\n")
    
    success_count = 0
    for module_name, module_dir, script_path in templates:
        print(f"Creating module: {module_name}")
        
        if load_and_execute_template(script_path, module_dir):
            print(f"  ✅ Generated files")
            success_count += 1
        else:
            print(f"  ❌ Failed")
        
    # Create Go workspace file
    create_go_workspace(base_dir, templates)
    
    print(f"\n✅ Successfully created {success_count}/{len(templates)} Go practice modules!")
    print("\nTo run a specific module:")
    print("  From terminal: cd <module_directory> && go run <module_name>.go")
    print("  From editor: Open any .go file and use the Run/Debug buttons")
    print("\nThe go.work file enables multi-module support in your editor.")
    print("Happy coding! 🚀")

def create_go_workspace(base_dir, templates):
    """
    Create a go.work file to manage all modules in the workspace.
    
    Args:
        base_dir: Base directory where go.work should be created
        templates: List of template tuples (module_name, module_dir, script_path)
    """
    go_work_path = os.path.join(base_dir, "go.work")
    
    print("\nCreating Go workspace file...")
    
    go_work_content = "go 1.25\n\nuse (\n"
    for module_name, _, _ in templates:
        go_work_content += f"    ./{module_name}\n"
    go_work_content += ")\n"
    
    with open(go_work_path, 'w') as f:
        f.write(go_work_content)
    
    print("  ✅ Created go.work (enables multi-module workspace)")

def clean_modules(base_dir):
    """
    Remove generated practice files but preserve .practice directories.
    
    Args:
        base_dir: Base directory containing practice modules
    """
    print(f"🧹 Cleaning up Go practice modules in: {base_dir}")
    
    templates = discover_templates(base_dir)
    
    if not templates:
        print("⏭️  No templates found to clean")
        return
    
    removed_count = 0
    
    for module_name, module_dir, _ in templates:
        # Remove .go files and go.mod files in the module directory
        # but preserve the .practice directory
        removed_files = []
        
        for filename in os.listdir(module_dir):
            if filename == ".practice":
                continue  # Skip the template directory
            
            filepath = os.path.join(module_dir, filename)
            if os.path.isfile(filepath) and (filename.endswith('.go') or filename == 'go.mod'):
                try:
                    os.remove(filepath)
                    removed_files.append(filename)
                except Exception as e:
                    print(f"  ❌ Failed to remove {filepath}: {e}")
        
        if removed_files:
            print(f"  Cleaned {module_name}/ (removed {', '.join(removed_files)})")
            removed_count += 1
    
    # Remove go.work file
    go_work_path = os.path.join(base_dir, "go.work")
    if os.path.exists(go_work_path):
        try:
            os.remove(go_work_path)
            print("  ✅ Removed go.work")
        except Exception as e:
            print(f"  ❌ Failed to remove go.work: {e}")
    
    print(f"\n✅ Cleanup complete! Cleaned {removed_count} modules.")
    print("💡 Template directories (.practice) were preserved.")

def main():
    """Main function to handle command line arguments and operations."""
    parser = argparse.ArgumentParser(
        description="Set up or clean up Go practice modules for learning",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Create all practice modules
  python3 setup_go_practice.py
  
  # Clean up all modules for fresh start
  python3 setup_go_practice.py --clean
  
  # Create modules (same as no arguments)
  python3 setup_go_practice.py --create
        """
    )
    
    group = parser.add_mutually_exclusive_group()
    group.add_argument(
        '--clean', '-c',
        action='store_true',
        help='Remove all practice modules (preserves .practice directories)'
    )
    group.add_argument(
        '--create',
        action='store_true',
        help='Create practice modules (default action)'
    )
    
    args = parser.parse_args()
    script_dir = os.path.dirname(os.path.abspath(__file__))
    
    if args.clean:
        # Clean up modules
        clean_modules(script_dir)
    else:
        # Create modules (default behavior)
        create_practice_modules(script_dir)
        
        print(f"\n💡 Tip: Use '{sys.argv[0]} --clean' to remove all modules for fresh practice")
        print("\n📚 Practice templates are based on Go by Example with implementation removed.")
        print("   Follow the comment instructions to implement each concept from scratch.")
        print("   This hands-on approach will help you learn Go programming effectively!")

if __name__ == "__main__":
    main()
