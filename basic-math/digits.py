"""
  Link: https://www.youtube.com/watch?v=1xNbjMdbjug&list=LL&index=7
"""

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

# All the divisor of n lies within 1 - sqrt(n)
# Time Complexity: O(sqrt(n)) + O(number of factor)


def get_factors(n: int) -> list[int]:
    i = 1
    out = []
    # instead of checking i < sqrt(n) which may take more computation
    # we can check it like this
    while i * i <= n:
        if n % i == 0:
            # divisible by i
            out.append(i)
            other = n // i
            if other != i:
                out.append(other)
        i = i + 1
    return sorted(out)


# if a number only have 2 factors that is 1 and the number itself


def is_prime(n: int) -> bool:
    i = 1
    count = 0
    while i * i <= n:
        if n % i == 0:
            count = count + 1
            other = n // i
            if other != i:
                count = count + 1
        i = i + 1
    return count == 2


# GCD = Greatest Common Divisor | HCF = Highest Common Factor
# Time Complexity: O(min(n1, n2))
def calculate_gcd(n1: int, n2: int) -> int:
    for i in range(min(n1, n2), 0, -1):
        if n1 % i == 0 and n2 % i == 0:
            return i
# Equilateral Algorithm gcd(a,b) = gcd(a-b, b) where a> b
# Time Complexity: O(log of min(a,b) to the base Fi)


def gcd_with_equilateral_algo(n1: int, n2: int) -> int:
    while n1 > 0 and n2 > 0:
        if n1 > n2:
            n1 = n1 % n2
        else:
            n2 = n2 % n1
    return n1 if n2 == 0 else n2


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

    factors_cases = [
        custom_test.TestCase((1,), [1]),
        custom_test.TestCase((2,), [1, 2]),
        custom_test.TestCase((4,), [1, 2, 4]),
        custom_test.TestCase((36,), [1, 2, 3, 4, 6, 9, 12, 18, 36]),
    ]

    prime_cases = [
        custom_test.TestCase((1,), False),
        custom_test.TestCase((2,), True),
        custom_test.TestCase((4,), False),
        custom_test.TestCase((7,), True),
        custom_test.TestCase((9,), False),
    ]
    gcd_cases = [
        custom_test.TestCase((11, 13), 1),
        custom_test.TestCase((4, 2), 2),
        custom_test.TestCase((27, 18), 9),
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

    print("\nTesting get_factors")
    custom_test.TestRunner(get_factors, factors_cases).run()

    print("\nTesting is_prime")
    custom_test.TestRunner(is_prime, prime_cases).run()

    print("\nTesting calculate_gcd")
    custom_test.TestRunner(calculate_gcd, gcd_cases).run()

    print("\nTesting calculate_gcd with equilateral algorithm")
    custom_test.TestRunner(gcd_with_equilateral_algo, gcd_cases).run()
