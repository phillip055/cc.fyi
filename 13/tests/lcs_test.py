from difff.difff import lowest_common_subsequence, difference

def test_lowest_common_subseq():
    test_cases = [
        ("ABC", "XYZ", ""),
        ("ABCDEF", "ABCDEF", "ABCDEF"),
        ("AABCXY", "XYZ", "XY"),
        ("", "", ""),
        ("ABCD", "AC", "AC")
    ]

    for test_case in test_cases:
        assert "".join(lowest_common_subsequence(list(test_case[0]), list(test_case[1]))) == test_case[2]

def test_lines_diff():
    test_cases = [
        (
            ["This is a test which contains:", "this is the lcs" ],
            ["this is the lcs", "we're testing"],
            ["this is the lcs"]
        ),
        (
            [
                "Coding Challenges helps you become a better software engineer through that build real applications.",
                "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
                "I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
			    "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
            ],
            [
                "Helping you become a better software engineer through coding challenges that build real applications.",
                "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
                "These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
			    "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
            ],
            [
                "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
                "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
            ]
        )
    ]
    for test_case in test_cases:
        line_diff = lowest_common_subsequence(test_case[0], test_case[1])[::-1]
        assert len(line_diff) == len(test_case[2])
        line_diff.sort()
        test_case[2].sort()
        for x, y in zip(line_diff, test_case[2]):
            assert x == y

