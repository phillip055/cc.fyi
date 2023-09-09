import os
import json_parser

for root, dirs, files in os.walk('json_parser/cases'):
    for filename in files:
        with open('json_parser/cases/' + filename) as file:
            try:
                t = file.read()
                o = json_parser.JsonParser(t)
                o.parse()
                if filename.startswith('fail'):
                    print(filename, "Expected to fail but passed")
                else:
                    continue
            except (Exception) as e:
                if filename.startswith('fail'):
                    continue
                else:
                    print(filename, "Expected to pass but failed ->", e)
