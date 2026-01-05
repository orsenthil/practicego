#!/usr/bin/env python3


import os
import sys

# Add parent directory to path to import practice_utils
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))))

from practice_utils import generate_files


def generate(target_dir):
    practice_dir = os.path.dirname(os.path.abspath(__file__))
    return generate_files(practice_dir, target_dir)


if __name__ == "__main__":
    # Allow running directly for testing
    target_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    success = generate(target_dir)
    sys.exit(0 if success else 1)
