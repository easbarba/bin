#!/usr/bin/env python3

from pathlib import Path
import json


def read__get_file():
    pass


def file_content() -> str:
    f: Path = Path("/tmp").joinpath("opm.json")
    if f.exists():
        with open(f, "r") as fcont:
            return fcont.read()

    return None


def parse_content() -> dict:
    fcont = file_content()
    try:
        return json.loads(fcont)
    except json.JSONDecodeError:
        print("Error decoding api file")

    return None


def main():
    """Go for it."""
    m = parse_content()
    print(m["cover"])


if __name__ == "__main__":
    main()
