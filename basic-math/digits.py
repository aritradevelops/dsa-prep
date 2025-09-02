import math
import custom_test


def count_digits_1(n: int) -> int:
    count = 0
    while n > 0:
        n = n // 10
        count += 1
    return count


def count_digits_2(n: int) -> int:
    if n == 0:
        return 1
    return int(math.log10(n)) + 1


def reverse_number(n: int) -> int:
    reversed = 0
    while n > 0:
        digit = n % 10
        n = n // 10
        reversed = reversed * 10 + digit
    return reversed


def check_palindrome(n: int) -> bool:
    return n == reverse_number(n)


def is_armstrong_number(n: int) -> bool:
    num_digits = count_digits_1(n)
    sum = 0
    original = n
    while n > 0:
        digit = n % 10
        sum += digit ** num_digits
        n = n // 10
    return original == sum


if __name__ == '__main__':
    # custom_test.TestCases for count_digits_1 and count_digits_2
    digits_cases = [
        custom_test.TestCase((12345,), 5),
        custom_test.TestCase((0,), 0),        # count_digits_1 returns 0 for 0
        custom_test.TestCase((7,), 1),
        custom_test.TestCase((1000,), 4),
        custom_test.TestCase((99999,), 5),
    ]
    digits2_cases = [
        custom_test.TestCase((12345,), 5),
        custom_test.TestCase((0,), 1),        # count_digits_2 returns 1 for 0
        custom_test.TestCase((7,), 1),
        custom_test.TestCase((1000,), 4),
        custom_test.TestCase((99999,), 5),
    ]
    reverse_cases = [
        custom_test.TestCase((1200,), 21),
        custom_test.TestCase((1234,), 4321),
        custom_test.TestCase((90210,), 1209),
        custom_test.TestCase((1,), 1),
        custom_test.TestCase((100,), 1),
    ]
    palindrome_cases = [
        custom_test.TestCase((121,), True),
        custom_test.TestCase((123,), False),
        custom_test.TestCase((124,), False),
        custom_test.TestCase((1221,), True),
        custom_test.TestCase((132,), False),
        custom_test.TestCase((909,), True),
    ]
    armstrong_cases = [
        custom_test.TestCase((153,), True),      # 1^3 + 5^3 + 3^3 = 153
        custom_test.TestCase((9474,), True),     # 9^4 + 4^4 + 7^4 + 4^4 = 9474
        custom_test.TestCase((1634,), True),     # 1^4 + 6^4 + 3^4 + 4^4 = 1634
        custom_test.TestCase((123,), False),
        custom_test.TestCase((370,), True),
        custom_test.TestCase((407,), True),
    ]

    print("\nTesting count_digits_1")
    custom_test.TestRunner(count_digits_1, digits_cases).run()

    print("\nTesting count_digits_2")
    custom_test.TestRunner(count_digits_2, digits2_cases).run()

    print("\nTesting reverse_number")
    custom_test.TestRunner(reverse_number, reverse_cases).run()

    print("\nTesting check_palindrome")
    custom_test.TestRunner(check_palindrome, palindrome_cases).run()

    print("\nTesting is_armstrong_number")
    custom_test.TestRunner(is_armstrong_number, armstrong_cases).run()
