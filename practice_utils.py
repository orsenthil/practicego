#!/usr/bin/env python3
"""
Shared utilities for practice template generation.
Common functions used by all generate.py scripts in .practice directories.
"""

import os
import json
import re


def get_metadata(practice_dir):
    """
    Load metadata from metadata.json in the .practice directory.
    
    Args:
        practice_dir: Path to the .practice directory
        
    Returns:
        dict: Metadata containing key, display_name, and index
    """
    metadata_path = os.path.join(practice_dir, "metadata.json")
    with open(metadata_path, 'r') as f:
        return json.load(f)


def get_template(practice_dir):
    """
    Load template from template.go in the .practice directory.
    
    Args:
        practice_dir: Path to the .practice directory
        
    Returns:
        str: Go template content
    """
    template_path = os.path.join(practice_dir, "template.go")
    with open(template_path, 'r') as f:
        return f.read()


def topic_to_package_name(topic):
    """
    Convert a topic name to a valid Go package name.
    - Convert to lowercase
    - Replace special characters with underscores
    - Remove consecutive underscores
    - Ensure it starts with a letter
    
    Args:
        topic: Topic key string (e.g., "hello-world")
        
    Returns:
        str: Valid Go package name (e.g., "hello_world")
    """
    name = topic.lower()
    name = re.sub(r'[^a-z0-9]+', '_', name)
    name = re.sub(r'_+', '_', name)
    name = name.strip('_')
    
    if name and name[0].isdigit():
        name = 'go_' + name
    
    return name


def generate_files(practice_dir, target_dir):
    """
    Standard file generation logic used by all templates.
    Creates .go and go.mod files in the target directory.
    
    Args:
        practice_dir: Path to the .practice directory
        target_dir: Directory where .go and go.mod files should be created
        
    Returns:
        bool: True if successful, False otherwise
    """
    try:
        metadata = get_metadata(practice_dir)
        template = get_template(practice_dir)
        package_name = topic_to_package_name(metadata['key'])
        
        # Create .go file
        go_filename = f"{package_name}.go"
        go_filepath = os.path.join(target_dir, go_filename)
        with open(go_filepath, 'w') as f:
            f.write(template)
        
        # Create go.mod file
        go_mod_path = os.path.join(target_dir, "go.mod")
        go_mod_content = f'''module github.com/orsenthil/gobyexample/{package_name}

go 1.25
'''
        with open(go_mod_path, 'w') as f:
            f.write(go_mod_content)
        
        return True
    except Exception as e:
        import sys
        print(f"Error generating files: {e}", file=sys.stderr)
        return False


