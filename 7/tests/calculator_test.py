import unittest
from calculator import calculator

class TestCalculatorMethods(unittest.TestCase):
    def test_single_expression(self):
        assert(calculator.calculate('1') == '1')
        assert(calculator.calculate('1 + 1') == "2")
    
    def test_multiple_expressions(self):
        assert(calculator.calculate('1 + 1 * 3') == "4")
        assert(calculator.calculate('1 + 2 * 3') == "7")
        assert(calculator.calculate('1 / 3 * 3') == "1")
        assert(calculator.calculate('1 * 3 / 6') == "0.5")
    
    def test_multiple_expressions_with_brackets(self):
        assert(calculator.calculate('( 1 + 1 ) * 5') == "10")
        assert(calculator.calculate('( 1 + 1 ) * ( 5 / 2 )') == "5")
        assert(calculator.calculate('( 1 + 1 ) * ( ( 5 * 2 ) / 2 )') == "10")

if __name__ == '__main__':
    unittest.main()
